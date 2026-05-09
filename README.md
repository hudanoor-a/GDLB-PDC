# GDLB Deliverable 3 - Working Software

This package implements the Final Demonstration software for the Parallel and Distributed Computing project: Gossip-Driven Predictive Decay Load Balancer (GDLB).

## Project structure

```text
GDLB_Deliverable3/
  cmd/simulator/main.go              # CLI entry point
  internal/core/                     # backend, telemetry, metrics, request model
  internal/algorithms/               # RR, Least, P2C, Probing, ML, GDLB routers
  internal/workload/                 # Poisson, MMPP, incast workload generators
  internal/experiment/               # simulation driver and config
  configs/default.yaml               # reproducible default configuration
  scripts/run_all.sh                 # one-command local experiment runner
  scripts/plot_results.py            # result plot generator
  omnetpp/                           # OMNeT++ compatible C++/NED skeleton
  docs/                              # final report and setup guide
  results/                           # generated CSVs and plots
```

## Setup guide

### Required tools

1. Go 1.21 or newer
2. Python 3.9 or newer
3. Python package: matplotlib
4. Optional: OMNeT++ 6.x for network-level simulation

### Install dependencies

```bash
cd GDLB_Deliverable3
python3 -m pip install matplotlib
```

Go uses only the standard library, so no external Go dependency is required.

## Algorithms implemented

- Round Robin
- Least Connections
- Power-of-Two Choices
- Prequal-inspired probing
- ML-inspired EWMA predictive routing
- GDLB novel predictive decay load balancer

## Deliverable 3 requirements covered

- Experimental setup
- Workload modeling
- Scalability parameters
- Failure scenario evaluation
- Comparative baseline analysis
- Sensitivity parameters
- Measurable improvement metrics
- C-6 novelty implementation
- Reproducibility support

Execution commands are intentionally separated from this setup guide and can be followed after setup is complete.
