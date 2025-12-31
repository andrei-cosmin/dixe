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

import "math/rand/v2"

// Uniform represents a uniform distribution
type Uniform struct {
	Min float64
	Max float64
	Rng *rand.Rand
}

// CDF computes the value of the cumulative distribution function at x.
func (u Uniform) CDF(x float64) float64 {
	if x < u.Min {
		return 0
	}
	if x > u.Max {
		return 1
	}
	return (x - u.Min) / (u.Max - u.Min)
}

// Rand returns a random number in the range [Min, Max]
func (u Uniform) Rand() float64 {
	return u.Rng.Float64()*(u.Max-u.Min) + u.Min
}
