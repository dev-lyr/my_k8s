# 一 概述:
## (1)概述:
- CRI(容器运行时接口): 一个plugin接口, 可以让kubelet无需重新编译的情况下使用多种容器运行时.

## (2)组成:
- protocol buffer和gRPC API.
- specifications/requirements
- 容器运行时集成与kubelet所需的库(https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/server/streaming).

## (3)备注:
- https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/
- https://github.com/kubernetes/kubernetes/blob/242a97307b34076d5d8f5bbeb154fa4d97c9ef1d/docs/devel/container-runtime-interface.md
- https://kubernetes.io/docs/setup/cri/
