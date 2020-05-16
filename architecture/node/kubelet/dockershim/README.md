# 一 概述:
## (1)概述:
- cri的docker实现.

## (2)备注:
- https://github.com/kubernetes/cri-api/blob/master/pkg/apis/runtime/v1alpha2/api.proto
- pkg/kubelet/dockershim

# 二 docker_service:
## (1)dockerService结构:
- client libdocker.Interface
- os kubecontainer.OSInterface
- podSandboxImage
- network 
- networkReady
- containerManager
- checkpointManager
- containerCleanUpInfos
- streamingRuntime
- streamingServer

## (2)NewDockerService:


# 三 docker_sandbox:
## (1)概述:
- docker_sandbox.go

## (2)RunPodSandbox:
- 1.拉取sandbox镜像, 若没配置则默认使用k8s.gcr.io/pause:3.2.
- 2.创建sandbox容器.
- 3.创建sandbox容器的checkpoint.
- 4.启动sandbox容器.
- 5.配置sandbox容器的网络,通过CNI插件完成,若使用host网络则跳过这步.

## (3)StopPodSandbox:

## (4)RemovePodSandbox:

## (5)PodSandboxStatus:

## (6)ListPodSandbox

# 四 docker_container:
## (1)概述:
- docker_container.go
- docker_stat.go

## (2)CreateContainer

# 五 docker_image:
## (1)概述:
- docker_image.go

## (2)PullImage

## (3)ListImage

## (4)ImageStatus

## (5)RemoveImage

# 六 docker_streaming:
## (1)概述:
- docker_streaming.go
- exec.go

## (2)Attach

## (3)Exec

## (4)PortForward

# 七 metrics
