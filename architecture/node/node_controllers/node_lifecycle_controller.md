# 一 概述:
## (1)概述:
- 管理node的lifecycle.

## (2)功能:
- Keeping the node controller's internal list of nodes up to date with the cloud provider's list of available machines.
- 监控node的健康

## (3)配置:
- --node-monitor-grace-period: Amount of time which we allow running Node to be unresponsive before marking it unhealthy, 默认是40s, 该值必须是kubelet的nodeStatusUpdateFrequency的N倍.
- --node-monitor-period: The period for syncing NodeStatus in NodeController, 默认5s.
- --node-startup-grace-period: Amount of time which we allow starting Node to be unresponsive before marking it unhealthy, 默认1m.
- --enable-taint-manager: 默认为true, 为true表示开启NoExecute Taints并且会驱逐带有该类型Taint的node上不能tolerate该taint的pod.
- --pod-eviction-timeout: The grace period for deleting pods on failed nodes, 默认5m.
- --node-eviction-rate
- --secondary-node-eviction-rate
- --unhealthy-zone-threshold

## (4)备注:
- pkg/controller/nodelifecycle
- https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/

# 二 实现:
## (1)Controller属性:
- taintManager
- nodeHealthMap
- nodeEvictionMap
- podUpdateQueue
- nodeUpdateQueue

## (2)Run:
- 若开启taintManager(默认开启), 则启动taintManager.
- 启动doNodeProcessingPassWorker
- 启动doPodProcessingWorker
- 若开启taintManager启动doNoExecuteTaintingPass(给node打上taint, 由taint manager执行驱逐); 否则启动doEvictionPass(直接驱逐).
- 启动monitorNodeHealth.

# 三 taintManager:
## (1)功能:
- 监听taint和toleration的变化, 并且负责删除有NoExecute Taints的Node上的pods. 

## (2)NoExecuteTaintManager

# 四 doNodeProcessingPassWorker:
## (1)功能:
- reconcile **labels** and/or update **NoSchedule taint** for nodes

## (2)doNoScheduleTaintingPass:
- 根据node.Status.Conditions来给node添加taint, 只处理effect=NoSchedule的taint.
- 映射关系见nodeConditionToTaintKeyStatusMap和taintKeyToNodeConditionMap变量.

## (3)reconcileNodeLabels

# 五 doPodProcessingWorker:
## (1)概述:
- 从podUpdateQueue取变化pod来处理.

## (2)processPod

# 六 doExecuteTaintingPass:
## (1)概述:
- 根据zoneNoExecuteTainter(monitorNodeHealth维护), 判断并对node打taint.

# 七 monitorNodeHealth:
## (1)概述:
- verifies node health are constantly updated by kubelet, and if not, post "NodeReady==ConditionUnknown".
- taint nodes who are not ready or not reachable for a long period of time.
