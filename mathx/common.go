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

package mathx

// Clamp returns the value clamped to the range [min, max]
func Clamp(value, min, max float64) float64 {
	if value <= min {
		return min
	}
	if value >= max {
		return max
	}
	return value
}

// Normalize converts x from [lower, upper] to [0, 1]
func Normalize(x, lower, upper float64) float64 {
	return (x - lower) / (upper - lower)
}

// Scale scales a normalized Beta value [0,1] to [lower, upper]
func Scale(normalized, lower, upper float64) float64 {
	return lower + normalized*(upper-lower)
}
