# 一 概述:
## (1)概述:
- CRI(容器运行时接口): 一个plugin接口, 可以让kubelet无需重新编译的情况下使用多种容器运行时.

## (2)组成:
- protocol buffer和gRPC API.
- specifications/requirements
- 容器运行时集成与kubelet所需的库(https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/server/streaming).

## (3)相关cri实现:
- docker-shim: https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/dockershim
- containerd-cri: https://github.com/containerd/cri
- 备注: docker-shim->docker->containerd->oci runtime
- 备注: containerd-cri->containerd->oci runtime
- 备注: containerd-cri性能更好, 占用资源更少, docker生态更丰富点.

## (4)备注:
- https://kubernetes.io/blog/2016/12/container-runtime-interface-cri-in-kubernetes/
- https://kubernetes.io/docs/setup/cri/
- https://godoc.org/k8s.io/cri-api

# 二 cri接口:
## (1)概述:
- RuntimeService: 需被一个容器运行时实现, 方法需是线程安全的.
- ImageManagerService: 需被一个容器镜像manager实现, 方法需是线程安全的.
- 备注: docker-shim实现的是v1alpha2

## (2)RuntimeService

## (3)ImageManagerService
