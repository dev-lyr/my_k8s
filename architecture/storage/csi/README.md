# 一 概述:
## (1)概述:
- 容器存储接口(CSI)为容器编排系统(例如:k8s)定义一个标准的接口, 用来将任意存储系统暴露给它们的容器workload.

## (2)相关名词:
- Volume
- Block Volume
- Mounted Volume
- SP: storage provider.

## (3)备注:
- https://kubernetes.io/docs/concepts/storage/volumes/#csi
- https://kubernetes-csi.github.io/docs/introduction.html
- https://github.com/container-storage-interface/spec/blob/master/spec.md
- kubernetes/pkg/volume/csi
