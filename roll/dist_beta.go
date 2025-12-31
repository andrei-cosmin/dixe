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
	"github.com/andrei-cosmin/dixe/dist"
	"github.com/andrei-cosmin/dixe/mathx"
)

// BetaParams computes alpha and beta from weight
type BetaParams func(weight float64) (alpha, beta float64)

// BetaDist is a configurable beta-based distribution
type BetaDist struct {
	params BetaParams
}

// NewBetaDist creates a new beta distribution with the given params function
func NewBetaDist(params BetaParams) BetaDist {
	return BetaDist{params: params}
}

// Rand generates a random value in [lower, upper]
func (b BetaDist) Rand(p FloatParams) float64 {
	return mathx.Scale(b.Beta(p).Rand(), p.Lower, p.Upper)
}

// CDF returns the cumulative distribution function at x
func (b BetaDist) CDF(x float64, p FloatParams) float64 {
	if x <= p.Lower {
		return 0
	}
	if x >= p.Upper {
		return 1
	}
	return b.Beta(p).CDF(mathx.Normalize(x, p.Lower, p.Upper))
}

// Beta returns the underlying beta distribution
func (b BetaDist) Beta(p FloatParams) dist.Beta {
	alpha, beta := b.params(p.Weight)
	return dist.Beta{Alpha: alpha, Beta: beta, Rng: p.Rng}
}
