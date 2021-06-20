# 一 概述:
## (1)概述:
- 管理node的生命周期.
- Node控制器会根据Node的condition自动添加一些taint到node上, 包括: NoExecute和NoSchedule.
- 相关: kubelet的evictionManager实现的是pod级别的驱逐.

## (2)功能:
- 监控node的健康并驱逐不健康节点上的pods.
- Keeping the node controller's internal list of nodes up to date with the cloud provider's list of available machines.

## (3)watch资源:
- lease
- node
- pod 
- daemonSet

## (4)配置:
- --node-monitor-grace-period: Amount of time which we allow running Node to be unresponsive before marking it unhealthy, 默认是40s, 该值必须是kubelet的nodeStatusUpdateFrequency的N倍.
- --node-monitor-period: The period for syncing NodeStatus in NodeController, 默认5s.
- --node-startup-grace-period: Amount of time which we allow starting Node to be unresponsive before marking it unhealthy, 默认1m.
- --enable-taint-manager: 默认为true, 为true表示开启NoExecute Taints并且会驱逐带有该类型Taint的node上不能tolerate该taint的pod.
- --pod-eviction-timeout: The grace period for deleting pods on failed nodes, 默认5m.
- --node-eviction-rate
- --secondary-node-eviction-rate
- --unhealthy-zone-threshold

## (5)备注:
- pkg/controller/nodelifecycle
- https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/
- https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/

# 二 实现:
## (1)概述:
- 控制器通过给node添加**NoExecute** taints(unreachable或notReady)来驱逐pods.
- 同时也会加上表示node问题的taints来让调度器不会调度pod到该node.

## (2)Controller属性:
- taintManager
- nodeHealthMap
- nodeEvictionMap
- podUpdateQueue
- nodeUpdateQueue

## (3)Run:
- 若开启taintManager(默认开启), 则启动taintManager.
- 启动doNodeProcessingPassWorker
- 启动doPodProcessingWorker
- 若开启taintManager启动doNoExecuteTaintingPass(给node打上taint, 由taint manager执行驱逐); 否则启动doEvictionPass(直接驱逐).
- 启动monitorNodeHealth.

# 三 taint类型:
## (1)v1.TaintEffectNoExecute:
- UnreachableTaintTemplate: v1.TaintNodeUnreachable
- NotReadyTaintTemplate: v1.TaintNodeNotReady

## (2)v1.TaintEffectNoSchedule:
- nodeConditionToTaintKeyStatusMap: node条件到taint的映射.

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

## (2)processPod:
- 取消NodeReady为true节点上pod的taint驱逐.
- 不运行taintManager时, 对于NodeReady为false或unknow节点pod的直接驱逐.
- 标记pod为not ready.

# 六 doNoExecuteTaintingPass:
## (1)概述:
- 根据zoneNoExecuteTainter(monitorNodeHealth维护), 判断并对node打taint.

# 七 monitorNodeHealth:
## (1)概述:
- 验证node监控情况是否不断的被kubelet更新, 若不是则post "NodeReady==ConditionUnknown".
- taint nodes who are not ready or not reachable for a long period of time.
- 当nodeReady为true时, 会将node设置为healthy.
