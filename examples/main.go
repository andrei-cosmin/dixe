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

package main

import (
	"fmt"

	"github.com/andrei-cosmin/dixe/roll"
)

func main() {
	// Create a seeded source for reproducible results
	src := roll.NewIntSource("my-game-seed").
		Dist(roll.Normal()).
		Weight(0.5)

	// Create a caster with a salt (e.g., player ID or action type)
	caster := src.SaltDist("player-123-attack")

	// Roll a D20
	result := caster.One(roll.D20())
	fmt.Printf("D20 roll: %d\n", result.First)

	// Roll multiple dice
	results := caster.Multiple(3, roll.D6())
	fmt.Printf("3D6 rolls: ")
	for _, r := range results {
		fmt.Printf("%d ", r.First)
	}
	fmt.Println()

	// Get probability distribution
	odds := caster.Odds(roll.D6())
	fmt.Println("\nD6 probabilities:")
	for i := 1; i <= 6; i++ {
		fmt.Printf("  %d: %.2f%%\n", i, odds.Probabilities[i])
	}

	// Different distributions
	fmt.Println("\nDifferent distributions (D10):")
	distributions := []struct {
		name string
		dist roll.Distribution
	}{
		{"Uniform", roll.Uniform()},
		{"Normal", roll.Normal()},
		{"WeightedLow", roll.WeightedLow()},
		{"WeightedHigh", roll.WeightedHigh()},
	}

	for _, d := range distributions {
		c := roll.NewIntSource("demo").Dist(d.dist).Weight(0.5).SaltDist("test")
		r := c.One(roll.D10())
		fmt.Printf("  %s: %d\n", d.name, r.First)
	}
}
