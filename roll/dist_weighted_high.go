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

// weightedHighDist biases toward the upper half using beta distribution with alpha > beta
var weightedHighDist = NewBetaDist(func(w float64) (float64, float64) {
	return 1.0 + w*4, 1.0
})

// weightedMaxDist has strong bias toward maximum with direct probability check
var weightedMaxDist = weightedMax{
	BetaDist: NewBetaDist(func(w float64) (float64, float64) {
		return 2.0, 1.0
	}),
}

type weightedMax struct {
	BetaDist
}

func (w weightedMax) Rand(p FloatParams) float64 {
	if p.Rng.Float64() < p.Weight {
		return p.Upper
	}
	return w.BetaDist.Rand(p)
}

func (w weightedMax) CDF(x float64, p FloatParams) float64 {
	if x <= p.Lower {
		return 0
	}
	if x >= p.Upper {
		return 1
	}
	betaCDF := w.BetaDist.Beta(p).CDF(mathx.Normalize(x, p.Lower, p.Upper))
	return (1 - p.Weight) * betaCDF
}
