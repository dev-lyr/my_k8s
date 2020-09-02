# 一 概述:
## (1)概述:
- CRI(容器运行时接口): 一个plugin接口, 可以让kubelet无需重新编译的情况下使用多种容器运行时, 通过GRPC进行通信.

## (2)组成:
- protocol buffer和gRPC API.
- specifications/requirements
- 容器运行时集成与kubelet所需的库(https://github.com/kubernetes/kubernetes/tree/master/pkg/kubelet/server/streaming).

## (3)CRI接口:
- RuntimeService: 需被一个容器运行时实现, 方法需是线程安全的.
- ImageManagerService: 需被一个容器镜像manager实现, 方法需是线程安全的.
- 备注: docker-shim实现的是v1alpha2

## (4)相关CRI实现:
- docker-shim: docker-shim->docker->containerd->oci runtime.
- containerd-cri: containerd-cri->containerd->oci runtime.
- 备注: containerd-cri性能更好, 占用资源更少, docker生态更丰富点.

## (5)备注:
- https://github.com/kubernetes/cri-api
- https://kubernetes.io/docs/setup/cri/
- https://godoc.org/k8s.io/cri-api
- https://github.com/containerd/cri

# 二 RuntimeService
## (1)sandbox相关:
- RunPodSandbox: 创建和启动一个pod级别的sandbo(pause容器),返回成功就表示sandbox处于ready状态.
- StopPodSandbox
- RemovePodSandbox
- PodSandboxStatus
- ListPodSandbox

## (2)container相关:
- CreateContainer
- StartContainer
- StopContainer
- RemoveContainer
- ListContainers
- ContainerStatus
- UpdateContainerResources
- ReopenContainerLogs
- Exec
- ExecSync
- Attach
- PortForward

# 三 ImageManagerService:
## (1)方法:
- ListImages
- ImageStatus
- ImageFsInfo
- PullImage
- RemoveImage
