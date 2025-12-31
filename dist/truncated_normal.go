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

package dist

import (
	"math"
	"math/rand/v2"
)

// TruncatedNormal represents a truncated normal distribution
// Values are guaranteed to be within [Lower, Upper]
type TruncatedNormal struct {
	Mu    float64    // Mean of the underlying normal
	Sigma float64    // Standard deviation of the underlying normal
	Lower float64    // Lower bound (inclusive)
	Upper float64    // Upper bound (inclusive)
	Rng   *rand.Rand // Random generator (required)
}

// CDF computes the value of the cumulative distribution function at x.
// For truncated normal: CDF(x) = (Φ((x-μ)/σ) - Φ((a-μ)/σ)) / (Φ((b-μ)/σ) - Φ((a-μ)/σ))
func (t TruncatedNormal) CDF(x float64) float64 {
	if x <= t.Lower {
		return 0
	}
	if x >= t.Upper {
		return 1
	}

	// Standardize
	alphaLower := (t.Lower - t.Mu) / t.Sigma
	alphaUpper := (t.Upper - t.Mu) / t.Sigma
	alphaX := (x - t.Mu) / t.Sigma

	// CDF values
	cdfLower := normCDF(alphaLower)
	cdfUpper := normCDF(alphaUpper)
	cdfX := normCDF(alphaX)

	return (cdfX - cdfLower) / (cdfUpper - cdfLower)
}

// Rand returns a random sample from the truncated normal distribution
// Uses the inverse CDF method for efficiency
func (t TruncatedNormal) Rand() float64 {
	if t.Sigma <= 0 {
		panic("truncated normal: sigma <= 0")
	}
	if t.Lower >= t.Upper {
		panic("truncated normal: lower >= upper")
	}

	// Standardize bounds
	alphaLower := (t.Lower - t.Mu) / t.Sigma
	alphaUpper := (t.Upper - t.Mu) / t.Sigma

	// CDF values at bounds
	cdfLower := normCDF(alphaLower)
	cdfUpper := normCDF(alphaUpper)

	// Generate uniform in [cdfLower, cdfUpper)
	uniform := cdfLower + t.Rng.Float64()*(cdfUpper-cdfLower)

	// Inverse CDF to get truncated normal value
	return t.Mu + t.Sigma*normQuantile(uniform)
}

// normCDF computes the standard normal cumulative distribution function
func normCDF(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}

// normQuantile computes the inverse standard normal CDF (quantile function)
func normQuantile(p float64) float64 {
	// Handle edge cases
	if p <= 0 {
		return math.Inf(-1)
	}
	if p >= 1 {
		return math.Inf(1)
	}
	return math.Sqrt2 * math.Erfinv(2*p-1)
}
