# 一 概述:
## (1)概述:
- k8s提供一个device插件框架, 运行将系统硬件资源报告给kubelet.
- ventor实现device插件并部署(手动或作为dameon), 目标资源包括: GPUs,高性能网卡(NICS), FPGAs或其它类型计算资源(需要vendor特定初始化和部署).

## (2)备注:
- https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/
- https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#extended-resources
