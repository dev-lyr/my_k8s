# 一概述:
## (1)功能:
- DaemonSet确保**一些或者全部**Node上运行指定pod,当node添加到集群时,该pod会被自动创建,当node从集群中移除时,pod会被垃圾回收.
- DaemonSet被删除时,它创建的Pod会被清除.
- 最简单情况, daemonSet在所有节点运行每种类型daemon; 复杂情况是一种类型使用多个daemonSet来运行, 但是不同daemonSet使用不同flags或者对不同硬件类型使用不同的cpu或内存request.

## (2)典型使用场景:
- 在所有node上运行一个node monitoring daemon, 例如: Prometheus Node Exporter.
- 在所有Node上运行一个日志收集daemon, 例如:fluentd或logstash.
- 在每个node上允许一个集群存储daemon, 例如: glusterd或cepth.

## (3)labels:
- controller-revision-hash: 加到到已存在ds的pod上,用于在ds的template更新期间区分旧的和新的pod.
- pod-template-generation: 作用同上,用上替代,已废弃.


# 二 DaemonSet资源:
## (1)DaemonSetSpec
- minReadySeconds
- revisionHistoryLimit: 需保留的历史版本的数量, 用于回滚, 默认为10.
- selector: 查询哪些pod match使用的label, 需match pod模板的labels.
- template: podTemplateSpec.
- updateStrategy: 更新策略(RollingUpdate或OnDelete).

## (2)DaemonSetStatus

# 三 DaemonSet调度:
## (1)概述:
- DaemonSet pods的创建和调度是通过daemonSet控制器(普通pod调度是通过调度器).
- DaemonSet遵守taints和tolerations, 根据相关特性一些tolerations会被自动添加到ds的pod上.

## (2)自动添加的tolerations:
- node.kubernetes.io/not-ready
- node.kubernetes.io/unreachable
- node.kubernetes.io/disk-pressure
- node.kubernetes.io/memory-pressure
- node.kubernetes.io/unschedulable
- node.kubernetes.io/network-unavailable

## (3)相关问题:
- 普通pod等待调度是pending状态,但daemonset pod pending状态表示没被创建.
- pod抢占是由默认调度器处理, daemonset控制器不考虑pod优先级和抢占.
