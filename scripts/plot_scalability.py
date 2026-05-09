import csv
import os
import sys
import matplotlib.pyplot as plt

if len(sys.argv) != 5:
    print("Usage:")
    print("python scripts\\plot_scalability.py <small_csv> <medium_csv> <large_csv> <output_folder>")
    sys.exit(1)

small_csv = sys.argv[1]
medium_csv = sys.argv[2]
large_csv = sys.argv[3]
output_folder = sys.argv[4]

os.makedirs(output_folder, exist_ok=True)

scale_labels = ["Small", "Medium", "Large"]
csv_files = [small_csv, medium_csv, large_csv]

def read_gdlb_row(csv_path):
    with open(csv_path, newline="") as f:
        reader = csv.DictReader(f)
        rows = list(reader)

    if not rows:
        raise ValueError(f"No rows found in {csv_path}")

    # Prefer GDLB row if CSV contains multiple algorithms
    for row in rows:
        if row.get("algorithm", "").lower() == "gdlb":
            return row

    # If only one row exists, use that row
    return rows[0]

def get_float(row, column_names):
    for col in column_names:
        if col in row:
            return float(row[col])
    raise KeyError(f"None of these columns found: {column_names}. Available columns: {list(row.keys())}")

p99_values = []
goodput_values = []
fairness_values = []

for csv_path in csv_files:
    row = read_gdlb_row(csv_path)

    p99_values.append(get_float(row, ["p99", "p99_ms", "p99_latency_ms"]))
    goodput_values.append(get_float(row, ["goodput_rps", "goodput", "throughput_rps"]))
    fairness_values.append(get_float(row, ["jain_fairness", "fairness", "jain_index"]))

# -----------------------------
# Graph 1: Actual P99 latency
# -----------------------------
plt.figure(figsize=(9, 6))
plt.plot(scale_labels, p99_values, marker="o")
plt.title("Scalability Trend: P99 Latency")
plt.xlabel("Scale Level")
plt.ylabel("P99 Latency (ms)")
plt.grid(True, alpha=0.3)
plt.tight_layout()
plt.savefig(os.path.join(output_folder, "scalability_p99.png"))
plt.close()

# -----------------------------
# Graph 2: Actual goodput
# -----------------------------
plt.figure(figsize=(9, 6))
plt.plot(scale_labels, goodput_values, marker="o")
plt.title("Scalability Trend: Goodput")
plt.xlabel("Scale Level")
plt.ylabel("Goodput (RPS)")
plt.grid(True, alpha=0.3)
plt.tight_layout()
plt.savefig(os.path.join(output_folder, "scalability_goodput.png"))
plt.close()

# -----------------------------
# Graph 3: Actual fairness
# -----------------------------
plt.figure(figsize=(9, 6))
plt.plot(scale_labels, fairness_values, marker="o")
plt.title("Scalability Trend: Jain Fairness Index")
plt.xlabel("Scale Level")
plt.ylabel("Jain Fairness Index")
plt.grid(True, alpha=0.3)
plt.tight_layout()
plt.savefig(os.path.join(output_folder, "scalability_fairness.png"))
plt.close()

# -----------------------------
# Graph 4: Combined normalized graph
# This puts P99, goodput, and fairness on one graph.
# -----------------------------
def normalize(values):
    min_v = min(values)
    max_v = max(values)

    if max_v == min_v:
        return [1.0 for _ in values]

    return [(v - min_v) / (max_v - min_v) for v in values]

p99_norm = normalize(p99_values)
goodput_norm = normalize(goodput_values)
fairness_norm = normalize(fairness_values)

plt.figure(figsize=(10, 6))
plt.plot(scale_labels, p99_norm, marker="o", label="P99 latency normalized")
plt.plot(scale_labels, goodput_norm, marker="s", label="Goodput normalized")
plt.plot(scale_labels, fairness_norm, marker="^", label="Fairness normalized")
plt.title("Scalability Trend Across Increasing Cluster Sizes")
plt.xlabel("Scale Level")
plt.ylabel("Normalized Metric Value")
plt.grid(True, alpha=0.3)
plt.legend()
plt.tight_layout()
plt.savefig(os.path.join(output_folder, "scalability_combined_normalized.png"))
plt.close()

# -----------------------------
# Save extracted values as CSV
# -----------------------------
summary_path = os.path.join(output_folder, "scalability_summary.csv")

with open(summary_path, "w", newline="") as f:
    writer = csv.writer(f)
    writer.writerow(["scale_level", "p99", "goodput_rps", "jain_fairness"])
    for i in range(len(scale_labels)):
        writer.writerow([
            scale_labels[i],
            p99_values[i],
            goodput_values[i],
            fairness_values[i]
        ])

print("Scalability plots saved in:", output_folder)
print("Caption: Scalability trend across increasing cluster sizes.")