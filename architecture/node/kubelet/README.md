# 一 概述:
## (1)功能:
- kubelet是每个node上运行的"node agent", kubelet将node注册到apiserver(使用hostname或一个覆盖hostname的flag或一个cloud provider的特定逻辑).
- 持续watchAPI服务器是否把Node分配给Pod, 然后通过运行时启动Pod容器, 随后持续健康运行中的容器向API服务器汇报它们状态,事件和资源消耗.
- kubelet运行容器的探针, 当探针失败时会重启容器; 当Pod删除时, kubelet会终止容器.
- Kubelet和PodSpec一起工作, 保证PodSpec描述的容器健康运行, kubelet不管理非kebernetes创建的容器.

## (2)PodSpec的来源:
- apiserver(主要).
- File: 通过命令行flag传递文件路径, 该路径下的文件会被周期性监控更新, 默认监控周期为20s.
- HTTP endpoint
- HTTP server
- 备注: 代码makePodSourceConfig.

## (3)源码:
- kubernetes/cmd/kubelet
- kubernetes/pkg/kubelet

## (4) kubelet结构体:
- 主要的kubelet实现
- pkg/kubelet/kubelet.go

# 二 Kubelet结构属性:
## (1)pod相关:
- podWorkers: 根据event进行pods的sync.
- podManager: 存储和管理对pod的访问,以及static pods和mirror pods间的映射.
- serviceLister
- admitHandlers
- softAdmitHandlers

## (2)容器相关
- probeManager
- livenessManager
- startupManager
- containerGC: dead容器的gc.
- containerLogManager: 容器日志管理.

## (3)Node相关:
- victionManager: 观察和响应影响node稳定性的情况.
- nodeLister
- imageManager: 管理镜像的gc.
- evictionManager 

# 三 方法:
## (1)主要:
- Run: 启动.
- syncLoop: 处理各种变化的main loop, for循环中调用syncLoopIteration.
- syncLoopIteration: 从不同channel中读取事件并dispatch pod到指定的handler.

## (2)pod相关:
- canRunPod
- canAdmitPod
- HandlerPod(Additions,Reconcile,Updates,Syncs,Removes): 底层都调用dispatchWork.
- dispatchWork: 开始**异步**sync pod, 使用pod worker(调用podWorkers的UpdatePod).
- **syncPod**: podWorker的UpdatePod会**异步**调用syncPod方法.

# 四 Run:
## (1)流程:
- 若cloudResourceSyncmManager不为nil, 则启动.
- 调用initializeModules初始化不需要容器运行时启动的内部模块.
- 启动volumeManager.
- syncNodeStatus.
- fastStatusUpdateOnce.
- 启动nodeLeaseController.
- updateRuntimeUp.
- 调用initNetworkUtil创建iptables rule.
- 启动一个goroutine用于killing pod(podKiller).
- 启动statusManager
- 启动probeManager
- 启动runtimeClassManager
- 启动pod lifecycle event generator.
- 调用syncLoop: 处理变化的主要loop.

## (2)initializeModules:
- 调用setupDataDirs创建相关数据目录: root目录, pods目录, plugin目录和pod资源目录.
- 创建容器日志目录(若不存在),目录为/var/log/containers.
- 启动imageManager.
- 启动ServerCertificateManager.
- 启动oom watcher.
- 启动resourceAnalyzer.

# 五 syncLoopIteration:
## (1)概述:
- 从不同channel中读取事件并dispatch pod到指定的handler, handler就是kubelet自己, 它实现了很多方法.

## (2)事件类型:
- configCh: pod config event.
- plegCh: pleg event.
- syncCh: periodic sync event.
- housekeepingCh: trigger clean up of pods.
- liveness manager update channel

## (3)configCh
- ADD(HandlePodAdditions): pods that are new to this source.
- DELETE(HandlePodUpdates): pods that are gracefully deleted from this source.
- REMOVE(HandlePodRemoves): pods that have been removed from this source.
- UPDATE(HandlePodUpdates): pods have been updated in this source.
- RECONCILE(HandlePodReconcile): pods that have unexpected status in this source,kubelet should reconcile status with this source.
- 备注: kubelet/type/pod_update.go, kubelet/config/config.go
- 备注: merges many sources of pod configuration.

## (4)plegCh:
- ContainerStarted
- ContainerDied: 调用cleanUpContainersInPod.
- ContainerRemoved
- ContainerChanged
- PodSync
- 备注: kubelet/pleg/pleg.go

## (5)syncCh:
- 写死1s触发一次事件, 叫醒kubelet来检查是否有pod workers(调用getPodsToSync查询)需要sync(调用HandlePodSyncs).

## (6)housekeepingCh:
- 触发pod的cleanup, 写死2s触发一次.
- 调用HandlePodCleanups方法.

## (7)liveness manager:
- liveness探测失败则调用HandlePodSyncs来sync pod.
