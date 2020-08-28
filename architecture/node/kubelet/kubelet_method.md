# 一 概述:
## (1)主要:
- Run: 启动.
- syncLoop: 处理各种变化的main loop, for循环中调用syncLoopIteration.
- syncLoopIteration: 从不同channel中读取事件并dispatch pod到指定的handler.

## (2)pod相关:
- canRunPod
- canAdmitPod
- HandlerPod(Additions,Reconcile,Updates,Syncs,Removes): 底层都调用dispatchWork.
- dispatchWork: 开始**异步**sync pod, 使用pod worker(调用podWorkers的UpdatePod).
- **syncPod**: podWorker的UpdatePod会异步调用syncPod方法.

# 二 Kubelet启动(Run):
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

## (2)initializeModules: - metrics
- 调用setupDataDirs创建相关数据目录: root目录, pods目录, plugin目录和pod资源目录.
- 创建容器日志目录(若不存在),目录为/var/log/containers.
- 启动imageManager.
- 启动ServerCertificateManager.
- 启动oom watcher.
- 启动resourceAnalyzer.

# 三 syncLoopIteration:
## (1)概述:
- 从不同channel中读取事件并dispatch pod到指定的handler, handler就是kubelet自己, 它实现了很多方法.

## (2)事件类型:
- configCh: pod config event.
- plegCh: pleg event.
- syncCh: periodic sync event.
- housekeepingCh: trigger clean up of pods.
-liveness manager update channel
