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

import (
	"cmp"
	"math/rand/v2"
	"slices"
)

// FloatWeightedCaster is a weighted caster for float64 values
type FloatWeightedCaster = WeightedCaster[float64]

// IntWeightedCaster is a weighted caster for int values
type IntWeightedCaster = WeightedCaster[int]

// ticket holds a value and its probability weight for deterministic custom rolling
type ticket[T constraint] struct {
	value  T
	weight float64
}

// ticketsFromWeights converts a weight map to a sorted slice of tickets
// Sorting by value ensures deterministic iteration order
func ticketsFromWeights[T constraint](weights Weights[T]) []ticket[T] {
	tickets := make([]ticket[T], 0, len(weights))
	for v, w := range weights {
		tickets = append(tickets, ticket[T]{value: v, weight: w})
	}
	slices.SortFunc(tickets, func(a, b ticket[T]) int {
		return cmp.Compare(a.value, b.value)
	})
	return tickets
}

// WeightedCaster holds a derived RNG and config for custom weight-based rolling
type WeightedCaster[T constraint] struct {
	rng *rand.Rand
	cfg weightedConfig[T]
}

// Custom sets the custom weights
// Does nothing if w is nil
func (c *WeightedCaster[T]) Custom(w Weights[T]) *WeightedCaster[T] {
	if len(w) > 0 {
		c.cfg.tickets = ticketsFromWeights(w)
	}
	return c
}

// RerollBelow sets the lower reroll threshold
func (c *WeightedCaster[T]) RerollBelow(v T) *WeightedCaster[T] {
	c.cfg.RerollBelow = v
	return c
}

// LowerExplosions sets the max lower explosions
func (c *WeightedCaster[T]) LowerExplosions(n int) *WeightedCaster[T] {
	c.cfg.MaxLowerExplosions = n
	return c
}

// RerollAbove sets the upper reroll threshold
func (c *WeightedCaster[T]) RerollAbove(v T) *WeightedCaster[T] {
	c.cfg.RerollAbove = v
	return c
}

// UpperExplosions sets the max upper explosions
func (c *WeightedCaster[T]) UpperExplosions(n int) *WeightedCaster[T] {
	c.cfg.MaxUpperExplosions = n
	return c
}

// With applies options to the caster config
// Preserves existing tickets if opts.Custom is nil
func (c *WeightedCaster[T]) With(opts Options[T]) *WeightedCaster[T] {
	existingTickets := c.cfg.tickets
	c.cfg = weightedConfigFromOptions(opts)
	if len(opts.Custom) == 0 {
		c.cfg.tickets = existingTickets
	}
	return c
}

// Fork creates a deep copy of the WeightedCaster
func (c *WeightedCaster[T]) Fork() WeightedCaster[T] {
	return WeightedCaster[T]{
		rng: c.rng,
		cfg: c.cfg.fork(),
	}
}

// One rolls a single value based on custom weights
func (c *WeightedCaster[T]) One(_ ...Range[T]) Result[T] {
	value := c.rollWeighted()

	return Result[T]{
		First:           value,
		Last:            value,
		Sum:             value,
		LowerExplosions: 0,
		UpperExplosions: 0,
		Rolls:           []T{value},
	}
}

// Multiple rolls multiple values and returns individual results
func (c *WeightedCaster[T]) Multiple(count int, _ ...Range[T]) []Result[T] {
	results := make([]Result[T], count)
	for i := 0; i < count; i++ {
		results[i] = c.One()
	}
	return results
}

// Odds calculates the probability distribution from custom weights
func (c *WeightedCaster[T]) Odds(_ ...Range[T]) Odds {
	result := Odds{
		Probabilities: make(map[int]float64),
	}

	var totalWeight float64
	for _, t := range c.cfg.tickets {
		totalWeight += t.weight
	}

	if totalWeight == 0 {
		return result
	}

	for _, t := range c.cfg.tickets {
		result.Probabilities[int(t.value)] = (t.weight / totalWeight) * 100
	}

	return result
}

// rollWeighted selects a value based on custom probability weights (deterministic order)
func (c *WeightedCaster[T]) rollWeighted() T {
	var totalWeight float64
	for _, t := range c.cfg.tickets {
		totalWeight += t.weight
	}

	rv := c.rng.Float64() * totalWeight

	var cumulative float64
	for _, t := range c.cfg.tickets {
		cumulative += t.weight
		if rv < cumulative {
			return t.value
		}
	}

	if len(c.cfg.tickets) > 0 {
		return c.cfg.tickets[len(c.cfg.tickets)-1].value
	}

	return 0
}
