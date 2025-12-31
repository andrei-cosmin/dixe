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

// Normal represents a normal (Gaussian) distribution
type Normal struct {
	Mu    float64    // Mean of the normal distribution
	Sigma float64    // Standard deviation of the normal distribution
	Rng   *rand.Rand // Random generator (required)
}

// CDF computes the value of the cumulative distribution function at x.
func (n Normal) CDF(x float64) float64 {
	return 0.5 * math.Erfc(-(x-n.Mu)/(n.Sigma*math.Sqrt2))
}

// Rand returns a random sample drawn from the distribution
func (n Normal) Rand() float64 {
	return n.Rng.NormFloat64()*n.Sigma + n.Mu
}
