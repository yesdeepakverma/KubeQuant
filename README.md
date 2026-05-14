# 🚀 KubeQuant
<a href="https://github.com/kubequant/KubeQuant">
  <img src="https://img.shields.io/badge/GitHub-KubeQuant-181717?style=for-the-badge&logo=github" />
</a>

**Percentile-Driven Kubernetes Resource Optimization Engine**

> KubeQuant is an open-source system that analyzes container resource usage using percentile-based metrics (p50, p95, p99) to generate intelligent CPU right-sizing recommendations—improving efficiency while preserving performance SLOs.

---

## 🧠 Why KubeQuant?

Modern Kubernetes workloads are **bursty and unpredictable**. Most existing tools rely on *average CPU usage*, which leads to:

* ❌ Over-provisioning → wasted cloud cost
* ❌ Under-provisioning → throttling & latency spikes

KubeQuant takes a different approach:

✅ Uses **tail-aware metrics (p95/p99)**

✅ Captures real workload behavior

✅ Produces **SLO-safe optimization recommendations**

---

## ⚙️ Key Features

* 📊 **Percentile-Based Analysis**
  Analyze CPU usage using p50, p95, and p99 distributions

* 🎯 **Right-Sizing Recommendations**
  Suggest optimal CPU requests & limits per container

* 🔍 **Over/Under-Provision Detection**
  Identify inefficient workloads automatically

* 📈 **Workload Profiling**
  Classify workloads (steady, bursty, periodic)

* 🔌 **Kubernetes Integration**
  Works seamlessly with Kubernetes APIs

* 📡 **Metrics Pipeline Support**
  Integrates with Prometheus

---

## 🏗️ Architecture

```
                +----------------------+
                |   Prometheus Metrics |
                +----------+-----------+
                           |
                           v
                +----------------------+
                |   Metrics Collector  |
                +----------+-----------+
                           |
                           v
                +----------------------+
                |  Percentile Analyzer |
                | (p50 / p95 / p99)   |
                +----------+-----------+
                           |
                           v
                +----------------------+
                | Recommendation Engine|
                +----------+-----------+
                           |
                           v
                +----------------------+
                | Kubernetes API Layer |
                +----------------------+
```

---

## 🚀 Getting Started

### Prerequisites

* Kubernetes cluster (v1.24+ recommended)
* Prometheus installed and scraping metrics
* Go / Python runtime (depending on build)

---

### Installation

```bash
git clone git@github.com:kubequant/KubeQuant.git
cd KubeQuant
make build
```

---

### Run Analysis

```bash
kubequant analyze \
  --namespace default \
  --window 7d \
  --output report.json
```

---

### Sample Output

```json
{
  "pod": "api-service",
  "current_cpu_request": "500m",
  "recommended_cpu_request": "320m",
  "p95_usage": "300m",
  "status": "over-provisioned"
}
```

---

## 📊 Benchmark Results (Example)

| Metric                   | Before | After | Improvement |
| ------------------------ | ------ | ----- | ----------- |
| Avg CPU Utilization      | 42%    | 68%   | +26%        |
| Idle CPU Allocation      | High   | Low   | -35%        |
| Throttling Events        | Low    | Low   | No impact   |
| Estimated Cost Reduction | —      | —     | ~25%        |

---

## 🔬 Research Motivation

KubeQuant is built on the hypothesis that:

> **Percentile-based resource allocation (p95/p99) provides a more reliable foundation for Kubernetes optimization than mean-based approaches.**

This project supports:

* Reproducible experiments
* Dataset export for research
* Integration with academic publications

---

## 📁 Project Structure

```
kubequant/
 ├── cmd/                # CLI entrypoints
 ├── pkg/                # Core logic
 ├── api/                # Kubernetes integration
 ├── analyzer/           # Percentile engine
 ├── recommender/        # Optimization logic
 ├── benchmarks/         # Performance experiments
 ├── docs/               # Documentation
 ├── examples/           # Sample configs
 └── research/           # Paper & datasets
```

---

## 🧪 Roadmap

* [x] Percentile-based CPU analysis
* [x] Recommendation engine (p95 baseline)
* [ ] Workload classification (bursty vs steady)
* [ ] Predictive optimization (ML-based)
* [ ] Auto-apply recommendations
* [ ] Multi-cluster support

---

## 🤝 Contributing

We welcome contributions!

* Open issues for bugs or feature requests
* Submit PRs with clear descriptions
* Follow coding + documentation standards

---

## 📢 Use Cases

* Cloud cost optimization
* SRE performance tuning
* Capacity planning
* Multi-tenant cluster efficiency

---

## 🧩 Comparison

| Feature                 | KubeQuant | VPA | HPA |
| ----------------------- | --------- | --- | --- |
| Percentile-aware        | ✅         | ❌   | ❌   |
| Tail (p95/p99) analysis | ✅         | ❌   | ❌   |
| SLO-focused tuning      | ✅         | ⚠️  | ❌   |
| Explainable output      | ✅         | ❌   | ❌   |

---

## 📄 License

MIT License

---

## 💡 Citation

If you use KubeQuant in research, please cite:

```
KubeQuant: Percentile-Driven Kubernetes Resource Optimization Engine
(Author: Your Name)
```

---

## 🌍 Vision

KubeQuant aims to redefine how Kubernetes resources are managed:

> From static provisioning → to **adaptive, data-driven optimization**
