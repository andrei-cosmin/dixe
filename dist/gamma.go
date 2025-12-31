// Copyright ©2016 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the THIRD_PARTY_LICENSES file.
//
// Derived from gonum.org/v1/gonum/stat/distuv
//
// Original algorithm references:
// - Liu, Martin, Syring. "Simulating from a gamma distribution with small
//   shape parameter" https://arxiv.org/abs/1302.1884
// - Marsaglia, Tsang. "A simple method for generating gamma variables."
//   ACM TOMS 26.3 (2000): 363-372.

package dist

import (
	"math"
	"math/rand/v2"
)

const (
	// The 0.2 threshold is from https://www4.stat.ncsu.edu/~rmartin/Codes/rgamss.R
	// described in detail in https://arxiv.org/abs/1302.1884.
	smallAlphaThresh = 0.2
)

// Gamma represents a Gamma distribution
type Gamma struct {
	Alpha float64    // Shape parameter (must be > 0)
	Beta  float64    // Rate parameter (must be > 0)
	Rng   *rand.Rand // Random generator (required)
}

// CDF computes the value of the cumulative distribution function at x.
func (g Gamma) CDF(x float64) float64 {
	if x < 0 {
		return 0
	}
	return gammaIncReg(g.Alpha, g.Beta*x)
}

// Rand returns a random sample drawn from the distribution
func (g Gamma) Rand() float64 {
	if g.Beta <= 0 {
		panic("gamma: beta <= 0")
	}

	a := g.Alpha
	b := g.Beta
	switch {
	case a <= 0:
		panic("gamma: alpha <= 0")
	case a == 1:
		return g.Rng.ExpFloat64() / b
	case a < smallAlphaThresh:
		return gammaSmallAlpha(a, b, g.Rng.ExpFloat64)
	case a >= smallAlphaThresh:
		return gammaMarsagliaTsang(a, b, g.Rng.Float64, g.Rng.NormFloat64)
	}
	panic("unreachable")
}

// gammaSmallAlpha generates a gamma variate for small alpha (< 0.2)
// using Liu, Chuanhai, Martin, Ryan and Syring, Nick. "Simulating from a
// gamma distribution with small shape parameter"
// https://arxiv.org/abs/1302.1884
// Reference: http://link.springer.com/article/10.1007/s00180-016-0692-0
// Algorithm adjusted to work in log space as much as possible.
func gammaSmallAlpha(a, b float64, exprnd func() float64) float64 {
	lambda := 1/a - 1
	lr := -math.Log1p(1 / lambda / math.E)
	for {
		e := exprnd()
		z := computeZ(e, lr, lambda, exprnd)
		eza := math.Exp(-z / a)
		lh := -z - eza
		lEta := computeLEta(z, lambda)
		if lh-lEta > -exprnd() {
			return eza / b
		}
	}
}

// computeZ computes the z value for the Liu-Martin-Syring algorithm
func computeZ(e, lr, lambda float64, exprnd func() float64) float64 {
	if e >= -lr {
		return e + lr
	}
	return -exprnd() / lambda
}

// computeLEta computes the lEta value for the Liu-Martin-Syring algorithm
func computeLEta(z, lambda float64) float64 {
	if z >= 0 {
		return -z
	}
	return -1 + lambda*z
}

// gammaMarsagliaTsang generates a gamma variate using
// Marsaglia, George, and Wai Wan Tsang. "A simple method for generating
// gamma variables." ACM Transactions on Mathematical Software (TOMS)
// 26.3 (2000): 363-372.
func gammaMarsagliaTsang(a, b float64, unifrnd, normrnd func() float64) float64 {
	d, m := computeDM(a, unifrnd)
	c := 1 / (3 * math.Sqrt(d))
	for {
		x := normrnd()
		v := 1 + x*c
		if v <= 0.0 {
			continue
		}
		v = v * v * v
		u := unifrnd()
		if acceptMarsagliaTsang(u, x, d, v) {
			return m * d * v / b
		}
	}
}

// computeDM computes the d and m values for the Marsaglia-Tsang algorithm
func computeDM(a float64, unifrnd func() float64) (d, m float64) {
	d = a - 1.0/3
	m = 1.0
	if a < 1 {
		d += 1.0
		m = math.Pow(unifrnd(), 1/a)
	}
	return d, m
}

// acceptMarsagliaTsang checks if the sample should be accepted
func acceptMarsagliaTsang(u, x, d, v float64) bool {
	xPow2 := x * x
	return u < 1.0-0.0331*(xPow2)*(xPow2) || math.Log(u) < 0.5*xPow2+d*(1-v+math.Log(v))
}
