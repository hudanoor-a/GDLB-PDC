package algorithms

import (
    "math"
    "math/rand"
    "gdlb-simulator/internal/core"
)

type Router interface {
    Pick(now float64, backends []*core.Backend) int
    OnDispatch(backend int, now float64)
    OnTelemetry(t core.BackendTelemetry, now float64)
    ControlBytes() int
}

type State struct { Auth float64; LocalSent int; LastUpdate float64; EWMA float64; Failed bool }

type Base struct { States []State; rr int; rng *rand.Rand; ctrl int }
func (b *Base) ControlBytes() int { return b.ctrl }
func (b *Base) OnDispatch(id int, now float64) {}
func (b *Base) OnTelemetry(t core.BackendTelemetry, now float64) { b.States[t.BackendID].Auth = float64(t.QueueDepth + t.RequestsInFlight); b.States[t.BackendID].LastUpdate=now; b.States[t.BackendID].Failed=t.Failed }

func NewRouter(name core.Algorithm, n int, rng *rand.Rand, w, alpha float64) Router {
    base := Base{States:make([]State,n), rng:rng}
    switch name {
    case core.RoundRobin: return &RR{Base:base}
    case core.LeastConnections: return &Least{Base:base}
    case core.PowerOfTwo: return &P2C{Base:base}
    case core.Probing: return &Probe{Base:base}
    case core.MLPredictive: return &ML{Base:base}
    case core.GDLB: return &GDLB{Base:base, W:w, Alpha:alpha}
    default: return &GDLB{Base:base, W:w, Alpha:alpha}
    }
}

type RR struct { Base }
func (r *RR) Pick(now float64, bs []*core.Backend) int { id:=r.rr%len(bs); r.rr++; return id }

type Least struct { Base }
func (r *Least) Pick(now float64, bs []*core.Backend) int { best:=0; bestScore:=math.Inf(1); for i:=range bs{ s:=r.States[i].Auth; if r.States[i].Failed {s=math.Inf(1)}; if s<bestScore{best=i; bestScore=s}}; return best }

type P2C struct { Base }
func (r *P2C) Pick(now float64, bs []*core.Backend) int { a:=r.rng.Intn(len(bs)); b:=r.rng.Intn(len(bs)); if r.States[a].Auth <= r.States[b].Auth {return a}; return b }

type Probe struct { Base }
func (r *Probe) Pick(now float64, bs []*core.Backend) int { a:=r.rng.Intn(len(bs)); b:=r.rng.Intn(len(bs)); r.ctrl += 64*2; if bs[a].QueueDepth <= bs[b].QueueDepth {return a}; return b }

type ML struct { Base }
func (r *ML) Pick(now float64, bs []*core.Backend) int { best:=0; bestScore:=math.Inf(1); for i:=range bs{ st:=r.States[i]; age:=now-st.LastUpdate; score:=0.7*st.EWMA + 0.3*st.Auth + 0.05*age; if st.Failed {score=math.Inf(1)}; if score<bestScore{best=i; bestScore=score}}; return best }
func (r *ML) OnTelemetry(t core.BackendTelemetry, now float64) { old:=r.States[t.BackendID].EWMA; obs:=float64(t.QueueDepth+t.RequestsInFlight); if old==0{old=obs}; r.States[t.BackendID].EWMA=0.8*old+0.2*obs; r.Base.OnTelemetry(t,now) }

type GDLB struct { Base; W float64; Alpha float64 }
func (r *GDLB) score(i int, now float64) float64 { st:=r.States[i]; if st.Failed { return math.Inf(1) }; age:=now-st.LastUpdate; penalty:=r.W*float64(st.LocalSent)*math.Exp(-r.Alpha*age); return st.Auth + penalty }
func (r *GDLB) Pick(now float64, bs []*core.Backend) int { best:=0; bestScore:=math.Inf(1); for i:=range bs{ s:=r.score(i,now); if s<bestScore{best=i; bestScore=s}}; return best }
func (r *GDLB) OnDispatch(id int, now float64) { r.States[id].LocalSent++ }
func (r *GDLB) OnTelemetry(t core.BackendTelemetry, now float64) { r.Base.OnTelemetry(t,now); r.States[t.BackendID].LocalSent=0 }
