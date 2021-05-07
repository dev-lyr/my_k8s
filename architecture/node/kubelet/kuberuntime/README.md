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
- computePodActions
- createPodSandbox
- startContainer

## (4)相关:
- pkg/kubelet/container: 支持一些pod级别的操作, kubeGenericRuntimeManager会实现相关接口.
- pkg/kubelet/cri/remote: cri的runtimeService和imageService的grpc实现, 作为kubeGenericRuntimeManager的属性.

## (5)备注:
- pkg/kubelet/kuberuntime

# 二 SyncPod:
## (1)概述:
- 被kubelet的syncPod调用.
- 计算sandbox和容器的改动(新建,杀死等), 然后根据计算结果执行对应操作.

## (2)pod创建步骤:
- 创建sandbox
- 创建临时容器
- 创建init容器
- 创建普通容器

