package workload

import "math/rand"

type Generator struct {
    Mode string
    BaseRate float64
    BurstRate float64
    BurstProb float64
    StateBurst bool
    rng *rand.Rand
}

func New(mode string, base, burst, burstProb float64, rng *rand.Rand) *Generator {
    return &Generator{Mode:mode, BaseRate:base, BurstRate:burst, BurstProb:burstProb, rng:rng}
}

func (g *Generator) NextInterarrival() float64 {
    rate := g.BaseRate
    switch g.Mode {
    case "mmpp":
        if g.rng.Float64() < g.BurstProb { g.StateBurst = !g.StateBurst }
        if g.StateBurst { rate = g.BurstRate }
    case "incast":
        if g.rng.Float64() < g.BurstProb { rate = g.BurstRate * 2 }
    }
    if rate <= 0 { rate = 1 }
    return g.rng.ExpFloat64()/rate
}
