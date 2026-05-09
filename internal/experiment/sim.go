package experiment

import (
    "math/rand"
    "gdlb-simulator/internal/algorithms"
    "gdlb-simulator/internal/core"
    "gdlb-simulator/internal/workload"
)

type Config struct {
    Algorithm core.Algorithm
    Proxies int
    Backends int
    Duration float64
    Rate float64
    BurstRate float64
    BurstProb float64
    Workload string
    Seed int64
    ServiceRate float64
    TelemetryDelay float64
    PenaltyW float64
    DecayAlpha float64
    FailureAt float64
    FailureDuration float64
}

func Run(c Config) *core.Metrics {
    rng := rand.New(rand.NewSource(c.Seed))
    backs := make([]*core.Backend,c.Backends)
    for i:=range backs { backs[i] = core.NewBackend(i, c.ServiceRate*(0.8+0.4*rng.Float64())) }
    routers := make([]algorithms.Router,c.Proxies)
    for i:=range routers { routers[i]=algorithms.NewRouter(c.Algorithm,c.Backends,rand.New(rand.NewSource(c.Seed+int64(i)+101)),c.PenaltyW,c.DecayAlpha) }
    gen := workload.New(c.Workload,c.Rate,c.BurstRate,c.BurstProb,rng)
    m := core.NewMetrics(c.Algorithm,c.Backends)
    now:=0.0; id:=0
    for now < c.Duration {
        now += gen.NextInterarrival()
        if c.FailureDuration>0 && now>=c.FailureAt && now<c.FailureAt+c.FailureDuration { backs[0].Failed=true } else { backs[0].Failed=false }
        p := id % c.Proxies
        b := routers[p].Pick(now, backs)
        req := core.Request{ID:id, ProxyID:p, Arrival:now, BackendID:b}
        routers[p].OnDispatch(b,now)
        tel := backs[b].Handle(&req,now,rng)
        routers[p].OnTelemetry(tel, req.Finish+c.TelemetryDelay)
        m.Record(req); id++
    }
    for _, r := range routers { m.ControlBytes += r.ControlBytes() }
    return m
}
