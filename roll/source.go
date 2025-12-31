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
	crand "crypto/rand"
	"math/rand/v2"

	"lukechampine.com/blake3"
)

// FloatSource is a source for float64 casters
type FloatSource = Source[float64]

// IntSource is a source for int casters
type IntSource = Source[int]

// Source holds the seed and default options, acts as a template for casters
type Source[T constraint] struct {
	seed       string
	opts       Options[T]
	floatRange func(Range[T]) FloatRange
	convert    func(float64) T
}

// NewFloatSource creates a new float source
// If no seed is provided, a random seed is generated
func NewFloatSource(seed ...string) *FloatSource {
	return &Source[float64]{
		seed:       initSeed(seed),
		opts:       DefaultOptions[float64](),
		floatRange: floatFloatRange,
		convert:    floatConvert,
	}
}

// NewIntSource creates a new int source
// If no seed is provided, a random seed is generated
func NewIntSource(seed ...string) *IntSource {
	return &Source[int]{
		seed:       initSeed(seed),
		opts:       DefaultOptions[int](),
		floatRange: intFloatRange,
		convert:    intConvert,
	}
}

// initSeed returns the seed or generates a random one if not provided
func initSeed(seed []string) string {
	if len(seed) > 0 && seed[0] != "" {
		return seed[0]
	}
	var buf [32]byte
	crand.Read(buf[:])
	var out [32]byte
	blake3.DeriveKey(out[:], "random-seed", buf[:])
	return string(out[:])
}

// SaltDist creates a DistCaster with a derived RNG from seed+salt
func (s *Source[T]) SaltDist(salt string) *DistCaster[T] {
	var chachaSeed [32]byte
	blake3.DeriveKey(chachaSeed[:], s.seed, []byte(salt))

	return &DistCaster[T]{
		rng:        rand.New(rand.NewChaCha8(chachaSeed)),
		cfg:        distConfigFromOptions(s.opts),
		floatRange: s.floatRange,
		convert:    s.convert,
	}
}

// SaltWeighted creates a WeightedCaster with a derived RNG from seed+salt
func (s *Source[T]) SaltWeighted(salt string) *WeightedCaster[T] {
	var chachaSeed [32]byte
	blake3.DeriveKey(chachaSeed[:], s.seed, []byte(salt))

	return &WeightedCaster[T]{
		rng: rand.New(rand.NewChaCha8(chachaSeed)),
		cfg: weightedConfigFromOptions(s.opts),
	}
}

// SaltCustomWeighted creates a WeightedCaster with provided Weights and a derived RNG from seed+salt
func (s *Source[T]) SaltCustomWeighted(salt string, weights Weights[T]) *WeightedCaster[T] {
	var chachaSeed [32]byte
	blake3.DeriveKey(chachaSeed[:], s.seed, []byte(salt))

	cfg := weightedConfigFromOptions(s.opts)
	if len(weights) > 0 {
		cfg.tickets = ticketsFromWeights(weights)
	}

	return &WeightedCaster[T]{
		rng: rand.New(rand.NewChaCha8(chachaSeed)),
		cfg: cfg,
	}
}

// Dist sets the distribution
func (s *Source[T]) Dist(d Distribution) *Source[T] {
	s.opts.Dist = d
	return s
}

// Weight sets the weight option
func (s *Source[T]) Weight(w float64) *Source[T] {
	s.opts.Weight = w
	return s
}

// Custom sets the custom weights (used by SaltWeighted)
func (s *Source[T]) Custom(w Weights[T]) *Source[T] {
	s.opts.Custom = w
	return s
}

// RerollBelow sets the lower reroll threshold
func (s *Source[T]) RerollBelow(v T) *Source[T] {
	s.opts.RerollBelow = v
	return s
}

// LowerExplosions sets the max lower explosions
func (s *Source[T]) LowerExplosions(n int) *Source[T] {
	s.opts.MaxLowerExplosions = n
	return s
}

// RerollAbove sets the upper reroll threshold
func (s *Source[T]) RerollAbove(v T) *Source[T] {
	s.opts.RerollAbove = v
	return s
}

// UpperExplosions sets the max upper explosions
func (s *Source[T]) UpperExplosions(n int) *Source[T] {
	s.opts.MaxUpperExplosions = n
	return s
}

// With merges the provided options
func (s *Source[T]) With(opts Options[T]) *Source[T] {
	s.opts.MergeWith(opts)
	return s
}

// Fork creates a copy of the Source
func (s *Source[T]) Fork() Source[T] {
	return *s
}
