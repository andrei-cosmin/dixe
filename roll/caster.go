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

// FloatCaster is a caster interface for float64 values
type FloatCaster = Caster[float64]

// IntCaster is a caster interface for int values
type IntCaster = Caster[int]

// Caster is the interface for rolling dice
// Implemented by both DistCaster (distribution-based) and WeightedCaster (custom weights)
type Caster[T constraint] interface {
	// One rolls a single value and returns the result
	One(r ...Range[T]) Result[T]

	// Multiple rolls multiple values and returns individual results
	Multiple(count int, r ...Range[T]) []Result[T]

	// Odds calculates the probability distribution for a roll
	Odds(r ...Range[T]) Odds
}

// defaultRange returns the first range if provided, or a default range otherwise
func defaultRange[T constraint](r ...Range[T]) Range[T] {
	if len(r) > 0 {
		return r[0]
	}
	return Range[T]{Lower: 1, Upper: 100}
}
