# 一 概述:
## (1)分类:
- Resource metrics pipeline
- Full metrics pipeline
- CronJob monitor
- Node problem detector(npd): https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health
- kube-state-metrics: https://github.com/kubernetes/kube-state-metrics
- 备注: k8s中, 应用监控不依赖一个单独的监控方案, 可以使用多个.

## (2)相关方案:
- 监控kubernetes节点: node_exporter和npd.
- 监控kuberbetes对象(deployment等): kube-state-metrics
- 监控容器: cadvisor
- 监控应用: 使用sidecar模式部署各种exporter(prometheus相关)或者让应用向外暴露metrics.

## (3)备注:
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/monitoring_architecture.md
- https://kubernetes.io/docs/tasks/debug-application-cluster/resource-usage-monitoring/
- https://kubernetes.io/docs/tasks/debug-application-cluster/resource-metrics-pipeline/
- https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/
- 与prometheus相关: https://github.com/coreos/kube-prometheus

# 二 resource metrics pipeline:
## (1)概述:
- resource metrics pipeline为集群组件(例如:HorizontalPodAutoscaler控制器)提供有限的metrics集; 用户也可以通过kubectl top命令来直接访问.
- metrics通过**metrics-server**来收集并通过**metrics.k8s.io** API暴露出来.
- metrics server发现集群中的所有node并向kubelet查询CPU和内存使用量; kubelet通过**cAdvisor**来获取数据.

## (2)cAdvisor:
- 是一个开源的容器资源使用量和性能分析**agent**, cAdvisor集成与kubelet的二进制文件中.
- 用于分析资源使用和运行中容器的性能特征.

## (3)备注:
- https://github.com/kubernetes/metrics/blob/master/pkg/apis/metrics/v1beta1/types.go
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/metrics-server.md
- https://github.com/google/cadvisor
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/instrumentation/resource-metrics-api.md

# 三 full metrics pipeline:
## (1)概述:
- 提供更丰富的metrics, k8s可以针对这些metrics进行自动扩容或者调整集群, 例如: Horizontal Pod Autoscaler.
- 通过kubelet来fetch metrics, 并通过一个实现**custom.metrics.k8s.api**或**external.metrics.k8s.io** API的**adapter**来暴露出来.

## (2)实现方案:
- prometheus
- sysdig

## (3)备注:
- https://github.com/DirectXMan12/k8s-prometheus-adapter

# 四 kube-state-metrics:
## (1)概述:
- https://github.com/kubernetes/kube-state-metrics

# 五 CronJob monitoring:
## (1)概述:
- 管理员可以用来查看那些job是运行中以及完成的job的状态.

## (2)备注:
- https://github.com/pietervogelaar/kubernetes-job-monitor
