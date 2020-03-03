# 一 命令和配置:
## (1)概述:
- kubernetes/cmd/kubelet

## (2)kubeletServer(struct):
- KubeletFlag: NewKubeletFlags方法创建默认值.
- kubeletconfig.KubeletConfiguration
- 创建方法: NewKubeletServer.
- 备注: options.go, 封装启动kubelet所需的所有参数.

## (3)kubelet.Dependencies:
- Auth
- CAdvisorInterface
- Cloud
- ContainerManager
- DockerClientConfig
- KubeClient
- HeartbeatClient
- EventClent
- HostUtil
- 等等.
- 创建: UnsecuredDependencies构造默认依赖.


# 二 Kubelet(struct):
## (1)概述:
- 是kubelet的主要实现, 创建方法: NewMainKubelet.
- kubernetes/pkg/kubelet/kubelet.go

## (2)属性:
- kubeClient
- heartbeatClient
- podWorkers
- podManager
- evictionManager
- cadvisor
- dnsConfigurer
- serviceLister
- nodeLister
- volumePluginMgr
- probeManager
- livenessManager
- imageManager
- containerLogManager
- 等等.


## (3)方法: 
- Run
- syncLoop: 是处理变化的主要循环.
- syncLoopIteration: 从不同channel读取数据然后将Pod分发给不同的handler, 被syncLoop调用.
- 等等.

## (4)syncLoopIteration:
- configCh(kubetypes.PodUpdate): 根据事件类型将配置变更的pod分发给合适的handler.
- plegCh(pleg.PodLifecycleEvent)
- syncCh(time.Time): 同步所有等待同步的pod.
- housekeepingCh(time.Time): 触发pod的清理.
- liveness Manager.

# 三 kubelet_pod.go

# 四 pod_workers.go
