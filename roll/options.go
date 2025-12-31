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

// FloatOptions alias for roll options (float64)
type FloatOptions = Options[float64]

// IntOptions alias for roll options (int)
type IntOptions = Options[int]

// Weights is a type alias for custom probability weights
type Weights[T constraint] = map[T]float64

// FloatWeights alias for float64 weights
type FloatWeights = Weights[float64]

// IntWeights alias for int weights
type IntWeights = Weights[int]

// Options configures the behavior of the roll function
type Options[T constraint] struct {
	// Dist is the probability distribution to use
	Dist Distribution

	// Weight controls distribution behavior [0.0, 1.0]:
	//   uniform:      ignored
	//   normal:       sigma = range * weight / 2 (lower = tighter bell curve)
	//   skewed:       higher = more extreme U-shape (values cluster at edges)
	//   weightedHigh: higher = stronger bias toward upper values
	//   weightedLow:  higher = stronger bias toward lower values
	//   weightedMin:  probability of returning the minimum value directly
	//   weightedMax:  probability of returning the maximum value directly
	Weight float64

	// Custom holds custom probability weights for discrete value selection
	// Used by WeightedCaster, ignored by DistCaster
	Custom Weights[T]

	// RerollBelow: reroll if result is strictly below this value (ints: 30 means 29 and down)
	RerollBelow T

	// MaxLowerExplosions limits reroll attempts (default 0 - none)
	MaxLowerExplosions int

	// RerollAbove: reroll if result is strictly above this value (ints: 75 means 76 and up)
	RerollAbove T

	// MaxUpperExplosions limits recursive explosions (default 0 - none)
	MaxUpperExplosions int
}

// DefaultOptions returns sensible defaults for RollOptions
func DefaultOptions[T constraint]() Options[T] {
	return Options[T]{
		Dist:               Uniform(),
		Weight:             0.5,
		RerollBelow:        0,
		MaxLowerExplosions: 0,
		RerollAbove:        0,
		MaxUpperExplosions: 0,
	}
}

// MergeWith merges the provided options into this one
func (o *Options[T]) MergeWith(override Options[T]) {
	if override.Dist != nil {
		o.Dist = override.Dist
	}
	if override.Weight != 0 {
		o.Weight = override.Weight
	}
	if override.Custom != nil {
		o.Custom = override.Custom
	}
	if override.RerollBelow != 0 {
		o.RerollBelow = override.RerollBelow
	}
	if override.MaxLowerExplosions != 0 {
		o.MaxLowerExplosions = override.MaxLowerExplosions
	}
	if override.RerollAbove != 0 {
		o.RerollAbove = override.RerollAbove
	}
	if override.MaxUpperExplosions != 0 {
		o.MaxUpperExplosions = override.MaxUpperExplosions
	}
}

// MergeOptions merges multiple RollOptions structs into a single one
func MergeOptions[T constraint](opts ...Options[T]) Options[T] {
	// Return default options if none provided
	if len(opts) == 0 {
		return DefaultOptions[T]()
	}

	// Merge options
	mergedOpts := opts[0]
	for index := range opts[1:] {
		mergedOpts.MergeWith(opts[index])
	}

	// Return merged options
	return mergedOpts
}
