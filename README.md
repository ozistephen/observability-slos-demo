# observability-slos-demo
# 📊 observability-slos-demo

SLO-driven observability stack for a fictional SaaS platform (**PulseDesk**) using Prometheus, Grafana, Loki, and Tempo.

## 🎯 Goal
Show how to measure and alert on **availability** and **latency** using real SLOs — not just CPU graphs.

## 🧰 Stack
- OpenTelemetry (metrics, traces)
- Prometheus (metrics & alerts)
- Grafana (dashboards)
- Loki (logs)
- Tempo (traces)
- k6 (load testing)

## ⚙️ Metrics & Alerts
**Availability SLI:**  
`success_ratio = successful_requests / total_requests`

**Latency SLI:**  
`histogram_quantile(0.95, sum(rate(request_duration_seconds_bucket[5m])) by (le))`

**Alerts:**  
- Fast burn: >2% of error budget in 2h  
- Slow burn: >10% of budget in 24h  

## 📐 Diagram
```mermaid
graph TD
A[App] -->|OTel Export| B[Prometheus]
A -->|Logs| C[Loki]
A -->|Traces| D[Tempo]
B --> E[Grafana Dashboards]
C --> E
D --> E