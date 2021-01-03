# 一 概述:
## (1)概述:
- k8s系统组件的metrics可以让用户看清楚内部发生了什么, 通常用于构建数据大盘和告警.
- k8s控制面板中的metrics是prometheus格式且人可读.
- 在生产环境可通过配置一个Prometheus服务器或其它metrics抓取器来收集这些metrics并将他们存储到一些时序数据库.

## (2)相关组件:
- apiserver
- kube-proxy
- controller-manager
- kube-schedule
- kubelet
- 备注: 大多数metrics通过http服务器的/metrics路径暴露出来.

## (3)备注:
- https://kubernetes.io/docs/concepts/cluster-administration/monitoring/
- k8s.io/component-base/metrics

