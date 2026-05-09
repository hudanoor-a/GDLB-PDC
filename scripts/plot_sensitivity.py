import csv
import sys
import os
import matplotlib.pyplot as plt

if len(sys.argv) < 3:
    print("Usage: python scripts\\plot_sensitivity.py <output_folder> <csv1:label1> <csv2:label2> ...")
    sys.exit(1)

output_folder = sys.argv[1]
items = sys.argv[2:]

os.makedirs(output_folder, exist_ok=True)

def get_value(row, possible_names):
    for name in possible_names:
        if name in row:
            return float(row[name])
    raise KeyError(f"None of these columns found: {possible_names}. Available columns: {list(row.keys())}")

labels = []
p99 = []
p999 = []
goodput = []
fairness = []
control = []

for item in items:
    csv_path, label = item.split(":", 1)

    with open(csv_path, newline="") as f:
        reader = csv.DictReader(f)
        rows = list(reader)

    if not rows:
        print("Skipping empty CSV:", csv_path)
        continue

    row = rows[0]

    labels.append(label)

    p99.append(get_value(row, [
        "p99_latency_ms",
        "p99_ms",
        "p99",
        "latency_p99_ms"
    ]))

    p999.append(get_value(row, [
        "p999_latency_ms",
        "p999_ms",
        "p99.9_ms",
        "p999",
        "latency_p999_ms"
    ]))

    goodput.append(get_value(row, [
        "goodput_rps",
        "goodput",
        "throughput_rps",
        "rps"
    ]))

    fairness.append(get_value(row, [
        "jain_fairness",
        "fairness",
        "jain_index"
    ]))

    control.append(get_value(row, [
        "control_bytes",
        "control_overhead_bytes",
        "message_bytes",
        "overhead_bytes"
    ]))

def make_bar(values, title, ylabel, filename):
    plt.figure(figsize=(10, 6))
    plt.bar(labels, values)
    plt.title(title)
    plt.xlabel("Sensitivity Setting")
    plt.ylabel(ylabel)
    plt.xticks(rotation=20)
    plt.tight_layout()
    plt.savefig(os.path.join(output_folder, filename))
    plt.close()

make_bar(p99, "P99 Latency Sensitivity Comparison", "P99 Latency (ms)", "sensitivity_p99.png")
make_bar(p999, "P99.9 Latency Sensitivity Comparison", "P99.9 Latency (ms)", "sensitivity_p999.png")
make_bar(goodput, "Goodput Sensitivity Comparison", "Goodput (RPS)", "sensitivity_goodput.png")
make_bar(fairness, "Fairness Sensitivity Comparison", "Jain Fairness Index", "sensitivity_fairness.png")
make_bar(control, "Control Overhead Sensitivity Comparison", "Control Bytes", "sensitivity_control.png")

print("Sensitivity comparison plots saved in:", output_folder)