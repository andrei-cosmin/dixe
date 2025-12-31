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

import "math"

// constraint is a type that can be used as a roll range
type constraint interface {
	int | float64
}

// intFloatRange converts an IntRange to FloatRange with upper+1 for discrete rolling
func intFloatRange(r IntRange) FloatRange {
	return FloatRange{
		Lower: float64(r.Lower),
		Upper: float64(r.Upper) + 1,
	}
}

// floatFloatRange converts a FloatRange to FloatRange with nextafter upper
func floatFloatRange(r FloatRange) FloatRange {
	return FloatRange{
		Lower: r.Lower,
		Upper: math.Nextafter(r.Upper, r.Upper+1),
	}
}

// intConvert converts a float to int via floor
func intConvert(value float64) int {
	return int(math.Floor(value))
}

// floatConvert returns the float as-is
func floatConvert(value float64) float64 {
	return value
}

// rollFloat generates a random float in the range using the distribution
// The result is clamped to [lower, upper) to handle edge cases
func rollFloat(distribution Distribution, p FloatParams) float64 {
	value := distribution.Rand(p)
	// Clamp to ensure we never exceed upper bound
	if value >= p.Upper {
		return math.Nextafter(p.Upper, p.Lower)
	}
	return value
}
