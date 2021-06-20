# 一 概述:
## (1)概述:
- Metrics Server is a scalable, efficient source of container resource metrics for Kubernetes built-in autoscaling pipelines.
- Metrics Server从kubelet收集资源metrics并在apiserver中通过**Metrics API**暴露出来.
- Metrics Server不要用于non-autoscaling目的.

## (2)使用方:
- HPA
- VPA
- kubectl top.

## (3)不适用场景:
- 非k8s集群.
- 准确的资源用量metrics源.
- 基于其它资源(非CPU和内存)的水平扩缩.

## (4)参数:
- --kubelet-preferred-address-types
- --kubelet-insecure-tls
- --requestheader-client-ca-file

## (5)备注:
- https://github.com/kubernetes-sigs/metrics-server

# 二 设计:
## (1)概述:
- Resource Metrics API is an effort to provide a first-class Kubernetes API (stable, versioned, discoverable, available through apiserver and with client support) that serves resource usage metrics for **pods and nodes**.
- It will be a cluster level component which periodically scrapes metrics from all Kubernetes nodes served by Kubelet through Summary API.
- Then metrics will be aggregated, stored in **memory** and served in **Metrics API format**.
