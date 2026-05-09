#!/usr/bin/env python3
import csv, sys, os
import matplotlib.pyplot as plt

inp = sys.argv[1] if len(sys.argv)>1 else 'results/summary.csv'
outdir = sys.argv[2] if len(sys.argv)>2 else 'results'
os.makedirs(outdir, exist_ok=True)
rows = list(csv.DictReader(open(inp)))
for metric, ylabel in [('p99','P99 latency (s)'),('goodput_rps','Goodput (RPS)'),('control_bytes','Control bytes'),('jain_fairness','Jain fairness')]:
    labels=[r['algorithm'] for r in rows]
    vals=[float(r[metric]) for r in rows]
    plt.figure(figsize=(8,4.8))
    plt.bar(labels, vals)
    plt.ylabel(ylabel)
    plt.xlabel('Algorithm')
    plt.title(ylabel + ' by algorithm')
    plt.tight_layout()
    plt.savefig(os.path.join(outdir, metric+'.png'), dpi=160)
    plt.close()
print('plots written to', outdir)
