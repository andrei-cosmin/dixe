# dixe

Deterministic dice rolling with statistical distributions for Go.

## Features

- **Seeded & Salted RNG** - Reproducible rolls using BLAKE3 key derivation
- **Multiple Distributions** - Uniform, Normal, Skewed, WeightedLow/High/Min/Max
- **Exploding Dice** - Configurable upper/lower explosion thresholds
- **Probability Calculation** - Get exact odds for any roll configuration

## Install

```bash
go get github.com/andrei-cosmin/dixe
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/andrei-cosmin/dixe/roll"
)

func main() {
    // Create a seeded source
    src := roll.NewIntSource("game-seed").
        Dist(roll.Normal()).
        Weight(0.5)

    // Create a caster with a salt
    caster := src.SaltDist("player-123")

    // Roll dice
    result := caster.One(roll.D20())
    fmt.Printf("D20: %d\n", result.First)

    // Get probabilities
    odds := caster.Odds(roll.D6())
    for i := 1; i <= 6; i++ {
        fmt.Printf("%d: %.2f%%\n", i, odds.Probabilities[i])
    }
}
```

## Distributions

| Distribution     | Behavior                   |
|------------------|----------------------------|
| `Uniform()`      | Equal probability          |
| `Normal()`       | Bell curve (center-biased) |
| `Skewed()`       | U-shaped (extreme-biased)  |
| `WeightedLow()`  | Favors lower values        |
| `WeightedHigh()` | Favors higher values       |
| `WeightedMin()`  | Strong bias toward minimum |
| `WeightedMax()`  | Strong bias toward maximum |

The `Weight` parameter (0.0-1.0) controls distribution intensity.

## License

MIT License - see [LICENSE](LICENSE)

Portions of the `dist` package are derived from [Gonum](https://github.com/gonum/gonum) (BSD-3-Clause)
and [Cephes](https://www.netlib.org/cephes/). See [THIRD_PARTY_LICENSES](THIRD_PARTY_LICENSES).