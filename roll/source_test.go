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
	"fmt"
	"math"
	"testing"
)

const (
	sides     = 100
	samples   = 100_000
	tolerance = 0.02 // 2% tolerance
)

func TestDistributionProbabilities(t *testing.T) {
	dists := []struct {
		name string
		dist Distribution
	}{
		{"Uniform", Uniform()},
		{"Normal", Normal()},
		{"Skewed", Skewed()},
		{"WeightedLow", WeightedLow()},
		{"WeightedMin", WeightedMin()},
		{"WeightedHigh", WeightedHigh()},
		{"WeightedMax", WeightedMax()},
	}
	weights := []float64{0.0, 0.25, 0.50, 0.75, 1.0}

	for _, d := range dists {
		for _, w := range weights {
			name := fmt.Sprintf("%s/weight=%.0f%%", d.name, w*100)
			t.Run(name, func(t *testing.T) {
				testDistribution(t, d.dist, w)
			})
		}
	}
}

func testDistribution(t *testing.T, dist Distribution, weight float64) {
	t.Helper()

	caster := NewIntSource("test-seed").Dist(dist).Weight(weight).SaltDist("test-salt")
	r := Dice(sides)
	counts := make(map[int]int)

	for i := 0; i < samples; i++ {
		counts[caster.One(r).First]++
	}

	expected := caster.Odds(r).Probabilities
	for bucket := r.Lower; bucket <= r.Upper; bucket++ {
		empirical := float64(counts[bucket]) / float64(samples) * 100.0
		diff := math.Abs(empirical - expected[bucket])
		if diff > tolerance*100 {
			t.Errorf("bucket %d: got %.2f%%, want %.2f%% (±%.0f%%)",
				bucket, empirical, expected[bucket], tolerance*100)
		}
	}
}

func TestFloatDistributionBounds(t *testing.T) {
	dists := []struct {
		name string
		dist Distribution
	}{
		{"Uniform", Uniform()},
		{"Normal", Normal()},
		{"Skewed", Skewed()},
		{"WeightedLow", WeightedLow()},
		{"WeightedMin", WeightedMin()},
		{"WeightedHigh", WeightedHigh()},
		{"WeightedMax", WeightedMax()},
	}

	r := FloatRange{Lower: 1, Upper: 100}

	for _, d := range dists {
		t.Run(d.name, func(t *testing.T) {
			src := NewFloatSource("test-seed").Dist(d.dist).Weight(0.50)
			caster := src.SaltDist("test-salt")

			for i := 0; i < samples; i++ {
				result := caster.One(r)
				if result.First < r.Lower || result.First > r.Upper {
					t.Errorf("value %.4f out of bounds [%.4f, %.4f]", result.First, r.Lower, r.Upper)
				}
			}
		})
	}
}
