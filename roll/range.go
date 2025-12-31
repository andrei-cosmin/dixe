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

// FloatRange alias for roll range (float64)
type FloatRange = Range[float64]

// IntRange alias for roll range (int)
type IntRange = Range[int]

// D4 returns a range of 1-4
func D4() IntRange {
	return IntRange{Lower: 1, Upper: 4}
}

// D6 returns a range of 1-6
func D6() IntRange {
	return IntRange{Lower: 1, Upper: 6}
}

// D8 returns a range of 1-8
func D8() IntRange {
	return IntRange{Lower: 1, Upper: 8}
}

// D10 returns a range of 1-10
func D10() IntRange {
	return IntRange{Lower: 1, Upper: 10}
}

// D20 returns a range of 1-20
func D20() IntRange {
	return IntRange{Lower: 1, Upper: 20}
}

// D100 returns a range of 1-100
func D100() IntRange {
	return IntRange{Lower: 1, Upper: 100}
}

// Dice returns a range of 1-x
func Dice(x int) IntRange {
	return IntRange{Lower: 1, Upper: x}
}

// Range represents a range of values
type Range[T constraint] struct {
	Lower T
	Upper T
}

// Width returns the range width
func (r Range[T]) Width() T {
	return r.Upper - r.Lower
}

// Sum returns the sum of the range
func (r Range[T]) Sum() T {
	return r.Lower + r.Upper
}

// Midpoint returns the midpoint of the range
func (r Range[T]) Midpoint() T {
	return (r.Lower + r.Upper) / 2
}
