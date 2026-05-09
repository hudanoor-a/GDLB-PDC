package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "gdlb-simulator/internal/core"
    "gdlb-simulator/internal/experiment"
)

func main() {
    algs := flag.String("algorithms","rr,least,p2c,probing,ml,gdlb","comma-separated algorithms")
    workload := flag.String("workload","mmpp","poisson|mmpp|incast")
    out := flag.String("out","results/summary.csv","CSV output path")
    proxies := flag.Int("proxies",20,"number of distributed proxies")
    backends := flag.Int("backends",50,"number of backend nodes")
    duration := flag.Float64("duration",60,"simulation duration in seconds")
    rate := flag.Float64("rate",900,"base request arrival rate")
    burst := flag.Float64("burst",5000,"burst arrival rate")
    burstProb := flag.Float64("burstprob",0.02,"MMPP state-change probability")
    service := flag.Float64("service",30,"average service rate per backend")
    delay := flag.Float64("telemetrydelay",0.005,"telemetry delay in seconds")
    w := flag.Float64("w",1.8,"GDLB penalty weight")
    alpha := flag.Float64("alpha",0.7,"GDLB decay alpha")
    failAt := flag.Float64("failat",20,"backend-0 failure start time")
    failDur := flag.Float64("faildur",5,"backend-0 failure duration")
    seed := flag.Int64("seed",42,"random seed")
    flag.Parse()

    header := []string{"algorithm","total","completed","failed","mean_latency","p50","p95","p99","p999","goodput_rps","control_bytes","jain_fairness"}
    rows := [][]string{header}
    for _, a := range strings.Split(*algs, ",") {
        cfg := experiment.Config{Algorithm:core.Algorithm(strings.TrimSpace(a)), Proxies:*proxies, Backends:*backends, Duration:*duration, Rate:*rate, BurstRate:*burst, BurstProb:*burstProb, Workload:*workload, Seed:*seed, ServiceRate:*service, TelemetryDelay:*delay, PenaltyW:*w, DecayAlpha:*alpha, FailureAt:*failAt, FailureDuration:*failDur}
        m := experiment.Run(cfg)
        rows = append(rows, m.Summary(*duration))
    }
    if err := os.MkdirAll(filepath.Dir(*out),0755); err != nil { panic(err) }
    if err := core.WriteCSV(*out, rows); err != nil { panic(err) }
    fmt.Println("wrote", *out)
}
