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

import "github.com/andrei-cosmin/dixe/mathx"

// uniformDist is a uniform distribution
var uniformDist = uniform{}

// uniform is a uniform distribution
type uniform struct{}

// Rand generates a random value in [lower, upper]
func (u uniform) Rand(p FloatParams) float64 {
	return p.Lower + p.Rng.Float64()*(p.Upper-p.Lower)
}

// CDF returns the cumulative distribution function at x
func (u uniform) CDF(x float64, p FloatParams) float64 {
	if x <= p.Lower {
		return 0
	}
	if x >= p.Upper {
		return 1
	}
	return mathx.Normalize(x, p.Lower, p.Upper)
}
