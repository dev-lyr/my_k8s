# 一 概述:
## (1)概述:
- node控制器是一个管理node许多方面的控制面板组件, 在node的生命周期内有很多作用.

## (2)功能:
- 当node注册时给它**分配CIDR块**.
- Keeping the node controller's internal list of nodes up to date with the cloud provider's list of available machines.
- 监控node的健康

## (3)node lease

## (4)备注:
- https://kubernetes.io/docs/concepts/architecture/nodes/#node-controller
- pkg/controller/nodelifecycle
- 配置: https://kubernetes.io/docs/reference/command-line-tools-reference/kube-controller-manager/

# 二 CIDR管理:

# 三 node管理:
## (1)概述:
- The node controller is responsible for updating the **NodeReady condition** of NodeStatus to **ConditionUnknown** when a node becomes unreachable (i.e. the node controller stops receiving heartbeats for some reason, for example due to the node being down), and then later evicting all the pods from the node (using graceful termination) if the node continues to be unreachable.

## (2)相关配置:
- --node-eviction-rate
- --pod-eviction-timeout: The grace period for deleting pods on failed nodes, 默认5m.
- --node-startup-grace-period: Amount of time which we allow starting Node to be unresponsive before marking it unhealthy, 默认1m.
- --node-monitor-period: The period for syncing NodeStatus in NodeController, 默认5s.
- --node-monitor-grace-period: Amount of time which we allow running Node to be unresponsive before marking it unhealthy. Must be N times more than kubelet's nodeStatusUpdateFrequency, where N means number of retries allowed for kubelet to post node status.
