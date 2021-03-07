# 一 概述: 
## (1)概述:
- kubeGenericRuntimeManager: 实现container目录下的Runtime等接口(支持pod相关操作), 同时包含cri的调用.
- kubelet调用NewKubeGenericRuntimeManager方法创建.

## (2)kubeGenericRuntimeManager属性:
- runtimeName
- runtimeService: GRPC service client.
- imageService: GRPC service client.
- 等等.

## (3)kubeGenericRuntimeManager方法:
- SyncPod

## (4)目录下:
- instrumented_services.go: 对RuntimeService进行了封装, 记录了操作和错误metrics.
- log

## (5)相关:
- pkg/kubelet/container: 支持一些pod级别的操作, kubeGenericRuntimeManager会实现相关接口.
- pkg/kubelet/cri/remote: cri的runtimeService和imageService的grpc实现, 作为kubeGenericRuntimeManager的属性.

## (6)备注:
- pkg/kubelet/kuberuntime

# 二 cri/remote目录:
## (1)概述:
- 包含cri-api的RuntimeService和ImageManagerService的实现, 用来创建到cri实现的client.

# 三 container目录:
## (1)概述:
- runtime.go: 定义一些容器运行时需实现的方法, kubeGenericRuntimeManager实现.
