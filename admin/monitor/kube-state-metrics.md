# 一 概述:
## (1)概述:
- kube-state-metrics是一个服务, 它监听apiserver并生成对象状态的metrics, 例如: deployments, nodes, pods等.
- metrics在HTTP endpoint /metrics和端口8080(默认)暴露出来.

## (2)kube-state-metrics和metrics-server:
- Metrics-Server is a cluster level component which periodically scrapes metrics from all Kubernetes nodes served by Kubelet through Summary API. The metrics are aggregated, stored in memory and served in Metrics API format. The metrics-server stores the latest values only and is not responsible for forwarding metrics to third-party destinations.
- kube-state-metrics is focused on generating completely new metrics from Kubernetes' object state (e.g. metrics based on deployments, replica sets, etc.). It holds an entire snapshot of Kubernetes state in memory and continuously generates new metrics based off of it. And just like the metrics-server it too is not responsible for exporting its metrics anywhere.
- Having kube-state-metrics as a separate project also enables access to these metrics from monitoring systems such as Prometheus.

## (3)备注:
- https://github.com/kubernetes/kube-state-metrics
