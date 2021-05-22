# 一 概述: 
## (1)概述:
- kubeGenericRuntimeManager: 实现container目录下的Runtime等接口(支持pod相关操作), 同时包含cri的调用.
- kubelet调用NewKubeGenericRuntimeManager方法创建.

## (2)kubeGenericRuntimeManager属性:
- runtimeName
- runtimeService: GRPC service client.
- imageService: GRPC service client.
- imagePuller
- runtimeHelper(kubecontainer.RuntimeHelper): kubelet实现该接口的方法.
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
- 计算sandbox和容器的改动判断需要执行什么操作(新建,杀死等), 然后根据计算结果执行对应操作.

## (2)pod创建步骤:
- 创建sandbox(createPodSandbox)
- 创建临时容器(startContainer)
- 创建init容器(startContainer)
- 创建普通容器(startContainer)

# 三 computePodActions:

# 四 createPodSandbox:
## (1)功能:
- 创建一个pod sandbox并返回ID,失败信息(若有)等.

## (2)步骤:
- 生成podSandboxConfig配置
- 创建pod日志目录
- 调用runtimeService.RunPodSandbox

## (3)generatePodSandboxConfig

# 五 startContainer:
## (1)步骤:
- pull镜像
- 创建容器
- 启动容器
- 执行postStart Hook(若有)
