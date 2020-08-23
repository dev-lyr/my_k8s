# 一 概述:
## (1)概述:
- node控制器是一个管理node许多方面的控制面板组件, 在node的生命周期内有很多作用.

## (2)功能:
- 当node注册时给node**分配CIDR块**.
- Keeping the node controller's internal list of nodes up to date with the cloud provider's list of available machines.
- 监控node的监控: 当node不可达时将NodeStatus的**NodeReady condition**更新为ConditionUnknow; 随后若未恢复, 将会驱逐node上所有pods.

## (3)备注:
- https://kubernetes.io/docs/concepts/architecture/nodes/#node-controller
- pkg/controller/nodelifecycle
