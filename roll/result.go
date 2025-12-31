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

// FloatResult alias for a roll result (float64)
type FloatResult = Result[float64]

// IntResult alias for a roll result (int)
type IntResult = Result[int]

// Result contains the outcome of a roll and metadata
type Result[T constraint] struct {
	// Value is the first roll result
	First T

	// Last is the last roll resul
	Last T

	// Sum is the sum of all rolls before multiplier/bonus
	Sum T

	// LowerExplosions is the number of lower explosions that occurred
	LowerExplosions int

	// UpperExplosions is the number of upper explosions that occurred
	UpperExplosions int

	// Rolls contains all individual rolls (for explosive dice)
	Rolls []T
}
