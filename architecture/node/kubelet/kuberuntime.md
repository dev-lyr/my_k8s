# 一 /pkg/kubelet/container
## (1)相关类型:
- Runtime接口: 定义容器运行时必须实现的接口.
- ImageService
- StreamingRuntime: 运行时实现用来提供streaming调用(exec,attach和port-forward).

## (2)Runtime接口:
- ImageService接口(内嵌)
- Type: 容器运行时的类型.
- SyncPod
- KillPod
- GetPodStatus
- GetPods
- UpdatePodCIDR
- GetContainerLog
- DeleteContainer
- GarbageCollect

## (3)ImageService:
- PullImage
- RemoveImage
- ImageStats
- GetImageRef
- ListImage
- 等等.

## (4)StreamingRuntime:
- GetAttach
- GetExec
- GetPortForward

# 二 pkg/kubelet/kuberuntime
## (1)相关类型:
- kubeGenericRuntimeManager: 实现Runtime等接口.

## (2)kubeGenericRuntimeManager属性:
- runtimeName
- runtimeService: GRPC service client.
- imageService: GRPC service client.
- 等等.

## (3)kubeGenericRuntimeManager方法:
- SyncPod

# 三 SyncPod

