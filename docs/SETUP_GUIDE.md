# Setup Guide - GDLB Deliverable 3

## 1. Software requirements

Install the following:

- Go 1.21+
- Python 3.9+
- matplotlib
- Optional: OMNeT++ 6.x

## 2. Folder placement

Place the folder anywhere, for example:

```bash
~/projects/GDLB_Deliverable3
```

Then open a terminal inside the folder.

## 3. Python setup

```bash
python3 -m pip install matplotlib
```

## 4. Go setup

No external Go packages are required. The simulator uses only the Go standard library.

Check Go installation:

```bash
go version
```

## 5. OMNeT++ setup, optional

Install OMNeT++ 6.x and import the `omnetpp` folder as an OMNeT++ project.

Compatible language: C++.

OMNeT++ is appropriate because it natively supports discrete-event simulation, message passing, queueing, network delay, and modular network components.

## 6. Expected files

Make sure these files exist:

```text
cmd/simulator/main.go
internal/algorithms/router.go
internal/core/backend.go
internal/core/metrics.go
internal/workload/workload.go
internal/experiment/sim.go
scripts/run_all.sh
scripts/plot_results.py
configs/default.yaml
```
