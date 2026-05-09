package core

import "math/rand"

type Backend struct {
    ID int
    ServiceRate float64
    NextFree float64
    InFlight int
    QueueDepth int
    Failed bool
    SlowFactor float64
}

func NewBackend(id int, serviceRate float64) *Backend {
    return &Backend{ID: id, ServiceRate: serviceRate, SlowFactor: 1.0}
}

func (b *Backend) Handle(req *Request, now float64, rng *rand.Rand) BackendTelemetry {
    if b.Failed {
        req.Failed = true
        req.Start = now
        req.Finish = now
        return b.Telemetry(now)
    }
    if b.NextFree < now { b.NextFree = now }
    b.QueueDepth = int((b.NextFree - now) * b.ServiceRate)
    b.InFlight++
    serviceMean := 1.0 / (b.ServiceRate / b.SlowFactor)
    service := rng.ExpFloat64() * serviceMean
    req.Start = b.NextFree
    req.Finish = b.NextFree + service
    b.NextFree = req.Finish
    if b.InFlight > 0 { b.InFlight-- }
    b.QueueDepth = int((b.NextFree - now) * b.ServiceRate)
    return b.Telemetry(req.Finish)
}

func (b *Backend) Telemetry(t float64) BackendTelemetry {
    return BackendTelemetry{BackendID:b.ID, QueueDepth:b.QueueDepth, RequestsInFlight:b.InFlight, ServiceRate:b.ServiceRate, Timestamp:t, Failed:b.Failed}
}
