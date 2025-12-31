// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the THIRD_PARTY_LICENSES file.
//
// Derived from gonum.org/v1/gonum/stat/distuv

package dist

import "math/rand/v2"

// Beta represents a Beta distribution
type Beta struct {
	Alpha float64    // Left shape parameter (must be > 0)
	Beta  float64    // Right shape parameter (must be > 0)
	Rng   *rand.Rand // Random generator (required)
}

// CDF computes the value of the cumulative distribution function at x.
func (b Beta) CDF(x float64) float64 {
	if x <= 0 {
		return 0
	}
	if x >= 1 {
		return 1
	}
	return regIncBeta(b.Alpha, b.Beta, x)
}

// Rand returns a random sample drawn from the distribution
// Uses ratio of two Gamma variates
func (b Beta) Rand() float64 {
	ga := Gamma{Alpha: b.Alpha, Beta: 1, Rng: b.Rng}.Rand()
	gb := Gamma{Alpha: b.Beta, Beta: 1, Rng: b.Rng}.Rand()
	return ga / (ga + gb)
}
