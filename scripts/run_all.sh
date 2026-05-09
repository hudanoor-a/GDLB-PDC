#!/usr/bin/env bash
set -euo pipefail
mkdir -p results
go run ./cmd/simulator -out results/summary.csv -workload mmpp -duration 60 -proxies 20 -backends 50 -rate 900 -burst 5000 -burstprob 0.02 -telemetrydelay 0.005 -seed 42
python3 scripts/plot_results.py results/summary.csv results
