# 一 概述:
## (1)概述:
- Pod可以有优先级, 优先级表示该Pod相对其它pod的重要性.
- PriorityClass是一个non-namespaced对象, 定义了一个优先级类name和优先级整数值的映射, 越高优先级值, 级别越高.

## (2)功能:
- 若Pod不能被调度, 调度器会尝试抢占(preemption)低优先级的Pod来使得高优先级得到调度.
- 优先级也会影响pod的调度顺序以及node上out-of-resources时的eviction顺序.

## (3)使用:
- 使用优先级和抢占: 添加一个或多个**PriorityClass**资源; 创建Pod时候PodSpec的priorityClassName指定为特定的PriorityClass对象.
- 创建Pod在spec指定了priorityClassName, 则**priority admission controller**使用priorityClassName属性并填入(populate)priority的整数值.
- 关闭抢占: 将PodSpec的preemptionPolicy设置为Nerver; 将kube-schedule的flag disablePreemption设置为true.

## (4)优先级和QoS:
- 调度器的抢占逻辑在选择被抢占目标时不考虑QoS,只考虑优先级.
- 唯一同时考虑优先级和QoS的是kubelet的out-of-resource eviction.

## (5)备注:
- 在kube 1.12+, 当集群资源有压力时, 重要的Pod依赖调度器抢占来调度, 因此不推荐关闭抢占.
- https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/

# 二 PriorityClass:
## (1)概述:
- PriorityClass是个global资源, 定义一个优先级类名到优先级值的映射, 值越大优先级越高.

## (2)PriorityClass资源:
- globlaDefault: 表示该PriorityClass的值作为没有指定priorityClassName的pod的默认值, 系统只能存在一个globalDefault的priorityClass对象, 若不指定, 则没有PriorityClassName的Pod的优先级为0.
- value: 优先级类的值.

# 三 抢占(preemption):
## (1)概述:
- 当pod被创建后放入队列并且等调度, 若没有node满足pod需求, 则**抢占逻辑**被触发.
- 当Pod P抢占Node N上一个或多个Pods时,Pod P的status的nominatedNodeName属性被设置为Node N的名字.

