// MIT License
//
// Copyright (c) 2025 Andrei Casu-Pop
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package roll

import "math/rand/v2"

// FloatDistCaster is a distribution caster for float64 values
type FloatDistCaster = DistCaster[float64]

// IntDistCaster is a distribution caster for int values
type IntDistCaster = DistCaster[int]

// DistCaster holds a derived RNG and config for distribution-based rolling
type DistCaster[T constraint] struct {
	rng        *rand.Rand
	cfg        distConfig[T]
	floatRange func(Range[T]) FloatRange
	convert    func(float64) T
}

// Dist sets the distribution
func (c *DistCaster[T]) Dist(d Distribution) *DistCaster[T] {
	if c.cfg.Dist != nil {
		c.cfg.Dist = d
	}
	return c
}

// Weight sets the weight option
func (c *DistCaster[T]) Weight(w float64) *DistCaster[T] {
	c.cfg.Weight = w
	return c
}

// RerollBelow sets the lower reroll threshold
func (c *DistCaster[T]) RerollBelow(v T) *DistCaster[T] {
	c.cfg.RerollBelow = v
	return c
}

// LowerExplosions sets the max lower explosions
func (c *DistCaster[T]) LowerExplosions(n int) *DistCaster[T] {
	c.cfg.MaxLowerExplosions = n
	return c
}

// RerollAbove sets the upper reroll threshold
func (c *DistCaster[T]) RerollAbove(v T) *DistCaster[T] {
	c.cfg.RerollAbove = v
	return c
}

// UpperExplosions sets the max upper explosions
func (c *DistCaster[T]) UpperExplosions(n int) *DistCaster[T] {
	c.cfg.MaxUpperExplosions = n
	return c
}

// With applies options to the caster config
// Preserves existing Dist if opts.Dist is nil
func (c *DistCaster[T]) With(opts Options[T]) *DistCaster[T] {
	existingDist := c.cfg.Dist
	c.cfg = distConfigFromOptions(opts)
	if opts.Dist == nil {
		c.cfg.Dist = existingDist
	}
	return c
}

// Fork creates a copy of the DistCaster
func (c *DistCaster[T]) Fork() DistCaster[T] {
	return *c
}

// One rolls a single value and returns the result
func (c *DistCaster[T]) One(r ...Range[T]) Result[T] {
	distRange := defaultRange(r...)

	// Build the float parameters for rolling
	p := FloatParams{
		Range:  c.floatRange(distRange),
		Weight: c.cfg.Weight,
		Rng:    c.rng,
	}

	// Roll the first value
	firstRoll := c.convert(rollFloat(c.cfg.Dist, p))

	// Generate explosions
	lowerRolls := c.processExplosion(firstRoll, p, c.cfg.shouldExplodeLower)
	upperRolls := c.processExplosion(firstRoll, p, c.cfg.shouldExplodeUpper)

	// Combine all rolls
	total := len(lowerRolls) + len(upperRolls) + 1
	allRolls := make([]T, total)
	allRolls[0] = firstRoll
	copy(allRolls[1:], lowerRolls)
	copy(allRolls[1+len(lowerRolls):], upperRolls)

	// Calculate the sum
	var sum T
	for _, v := range allRolls {
		sum += v
	}

	// Return the result
	return Result[T]{
		First:           allRolls[0],
		Last:            allRolls[total-1],
		Sum:             sum,
		LowerExplosions: len(lowerRolls),
		UpperExplosions: len(upperRolls),
		Rolls:           allRolls,
	}
}

// Multiple rolls multiple values and returns individual results
func (c *DistCaster[T]) Multiple(count int, r ...Range[T]) []Result[T] {
	distRange := defaultRange(r...)
	results := make([]Result[T], count)
	for i := 0; i < count; i++ {
		results[i] = c.One(distRange)
	}
	return results
}

// Odds calculates the probability distribution for a roll
func (c *DistCaster[T]) Odds(r ...Range[T]) Odds {
	distRange := defaultRange(r...)
	result := Odds{
		Probabilities: make(map[int]float64),
	}

	p := FloatParams{
		Range:  c.floatRange(distRange),
		Weight: c.cfg.Weight,
		Rng:    nil,
	}

	for v := int(distRange.Lower); v <= int(distRange.Upper); v++ {
		prob := c.cfg.Dist.CDF(float64(v+1), p) - c.cfg.Dist.CDF(float64(v), p)
		result.Probabilities[v] = prob * 100
	}

	explosionRange := c.floatRange(Range[T]{Lower: c.cfg.RerollBelow, Upper: c.cfg.RerollAbove})

	if c.cfg.MaxLowerExplosions > 0 {
		result.LowerExplosionChance = c.cfg.Dist.CDF(explosionRange.Lower, p) * 100
	}

	if c.cfg.MaxUpperExplosions > 0 {
		result.UpperExplosionChance = (1 - c.cfg.Dist.CDF(explosionRange.Upper, p)) * 100
	}

	return result
}

// processExplosion generates additional rolls while the condition is met
func (c *DistCaster[T]) processExplosion(currentRoll T, p FloatParams, shouldExplode func(T, int) bool) []T {
	var rolls []T
	for shouldExplode(currentRoll, len(rolls)) {
		currentRoll = c.convert(rollFloat(c.cfg.Dist, p))
		rolls = append(rolls, currentRoll)
	}
	return rolls
}
