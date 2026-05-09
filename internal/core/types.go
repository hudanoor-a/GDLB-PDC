package core

type Algorithm string

const (
    RoundRobin Algorithm = "rr"
    LeastConnections Algorithm = "least"
    PowerOfTwo Algorithm = "p2c"
    Probing Algorithm = "probing"
    MLPredictive Algorithm = "ml"
    GDLB Algorithm = "gdlb"
)

type Request struct {
    ID int
    ProxyID int
    Arrival float64
    BackendID int
    Start float64
    Finish float64
    Failed bool
}

type BackendTelemetry struct {
    BackendID int
    QueueDepth int
    RequestsInFlight int
    ServiceRate float64
    Timestamp float64
    Failed bool
}
