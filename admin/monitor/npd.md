# 一 概述:
## (1)概述:
- Node problem detector是一个用来监控node健康度的daemonset.
- 从不同daemon处收集node相关问题并以**NodeCondition**和**Event**上报给apiserver.
- 当前k8s不会针对npd生成的NodeConditon和Event采取任何操作, 未来会接入一个remedy system来处理node相关问题.

## (2)备注:
- https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/
- https://github.com/kubernetes/node-problem-detector
