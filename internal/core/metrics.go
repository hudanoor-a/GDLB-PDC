package core

import (
    "encoding/csv"
    "fmt"
    "math"
    "os"
    "sort"
)

type Metrics struct {
    Algorithm Algorithm
    Latencies []float64
    Failed int
    Completed int
    ControlBytes int
    BackendCounts []int
    TotalRequests int
}

func NewMetrics(a Algorithm, backends int) *Metrics { return &Metrics{Algorithm:a, BackendCounts:make([]int, backends)} }

func (m *Metrics) Record(req Request) {
    m.TotalRequests++
    if req.Failed { m.Failed++; return }
    m.Completed++
    m.Latencies = append(m.Latencies, req.Finish-req.Arrival)
    if req.BackendID >= 0 && req.BackendID < len(m.BackendCounts) { m.BackendCounts[req.BackendID]++ }
}

func percentile(vals []float64, p float64) float64 {
    if len(vals)==0 { return 0 }
    cp := append([]float64{}, vals...); sort.Float64s(cp)
    idx := int(math.Ceil((p/100.0)*float64(len(cp))))-1
    if idx < 0 { idx=0 }; if idx>=len(cp) { idx=len(cp)-1 }
    return cp[idx]
}

func Jain(counts []int) float64 {
    var sum, sumSq float64
    for _, c := range counts { x:=float64(c); sum += x; sumSq += x*x }
    if sumSq == 0 { return 0 }
    return (sum*sum)/(float64(len(counts))*sumSq)
}

func (m *Metrics) Summary(duration float64) []string {
    mean := 0.0
    for _, l := range m.Latencies { mean += l }
    if len(m.Latencies)>0 { mean /= float64(len(m.Latencies)) }
    return []string{
        string(m.Algorithm),
        fmt.Sprintf("%d", m.TotalRequests), fmt.Sprintf("%d", m.Completed), fmt.Sprintf("%d", m.Failed),
        fmt.Sprintf("%.6f", mean), fmt.Sprintf("%.6f", percentile(m.Latencies,50)), fmt.Sprintf("%.6f", percentile(m.Latencies,95)), fmt.Sprintf("%.6f", percentile(m.Latencies,99)), fmt.Sprintf("%.6f", percentile(m.Latencies,99.9)),
        fmt.Sprintf("%.2f", float64(m.Completed)/duration), fmt.Sprintf("%d", m.ControlBytes), fmt.Sprintf("%.6f", Jain(m.BackendCounts)),
    }
}

func WriteCSV(path string, rows [][]string) error {
    f, err := os.Create(path); if err != nil { return err }
    defer f.Close()
    w := csv.NewWriter(f); defer w.Flush()
    return w.WriteAll(rows)
}
