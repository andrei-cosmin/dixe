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
	"math/rand/v2"
)

// FloatParams bundles parameters passed to distribution methods (float64)
type FloatParams = DistParams[float64]

// IntParams bundles parameters passed to distribution methods (int)
type IntParams = DistParams[int]

// Uniform returns a uniform distribution (equal probability for all values)
func Uniform() Distribution { return uniformDist }

// Normal returns a bell-curve distribution centered in the range
func Normal() Distribution { return normalDist }

// Skewed returns a U-shaped distribution favoring extremes
func Skewed() Distribution { return skewedDist }

// WeightedLow returns a distribution biased toward lower values
func WeightedLow() Distribution { return weightedLowDist }

// WeightedHigh returns a distribution biased toward higher values
func WeightedHigh() Distribution { return weightedHighDist }

// WeightedMin returns a distribution with strong bias toward minimum
func WeightedMin() Distribution { return weightedMinDist }

// WeightedMax returns a distribution with strong bias toward maximum
func WeightedMax() Distribution { return weightedMaxDist }

// DistParams bundles parameters passed to distribution methods
type DistParams[T constraint] struct {
	Range[T]
	Weight float64
	Rng    *rand.Rand
}

// Distribution is a probability distribution with both sampling and CDF capabilities
type Distribution interface {
	// Rand generates a random value in [lower, upper) using this distribution
	Rand(p FloatParams) float64

	// CDF computes P(X < x) for X drawn from distribution on [lower, upper)
	CDF(x float64, p FloatParams) float64
}
