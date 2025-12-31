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

import "slices"

// config holds the base configuration shared by all casters (explosion settings)
type config[T constraint] struct {
	RerollBelow        T
	MaxLowerExplosions int
	RerollAbove        T
	MaxUpperExplosions int
}

// shouldExplodeLower returns true if the roll should trigger a lower explosion
func (c *config[T]) shouldExplodeLower(rollValue T, rollCount int) bool {
	return rollCount < c.MaxLowerExplosions && rollValue < c.RerollBelow
}

// shouldExplodeUpper returns true if the roll should trigger an upper explosion
func (c *config[T]) shouldExplodeUpper(rollValue T, rollCount int) bool {
	return rollCount < c.MaxUpperExplosions && rollValue > c.RerollAbove
}

// distConfig holds configuration for distribution-based rolling
type distConfig[T constraint] struct {
	config[T]
	Dist   Distribution
	Weight float64
}

// distConfigFromOptions extracts a distConfig from Options
func distConfigFromOptions[T constraint](opts Options[T]) distConfig[T] {
	return distConfig[T]{
		config: config[T]{
			RerollBelow:        opts.RerollBelow,
			MaxLowerExplosions: opts.MaxLowerExplosions,
			RerollAbove:        opts.RerollAbove,
			MaxUpperExplosions: opts.MaxUpperExplosions,
		},
		Dist:   opts.Dist,
		Weight: opts.Weight,
	}
}

// weightedConfig holds configuration for custom weight-based rolling
type weightedConfig[T constraint] struct {
	config[T]
	tickets []ticket[T]
}

// weightedConfigFromOptions extracts a weightedConfig from Options
func weightedConfigFromOptions[T constraint](opts Options[T]) weightedConfig[T] {
	return weightedConfig[T]{
		config: config[T]{
			RerollBelow:        opts.RerollBelow,
			MaxLowerExplosions: opts.MaxLowerExplosions,
			RerollAbove:        opts.RerollAbove,
			MaxUpperExplosions: opts.MaxUpperExplosions,
		},
		tickets: ticketsFromWeights(opts.Custom),
	}
}

// fork creates a deep copy of the weightedConfig
func (c *weightedConfig[T]) fork() weightedConfig[T] {
	return weightedConfig[T]{
		config:  c.config,
		tickets: slices.Clone(c.tickets),
	}
}
