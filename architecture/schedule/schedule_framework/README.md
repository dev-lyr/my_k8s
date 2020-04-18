# 一 概述:
## (1)概述:
- 调度框架是针对kube调度器的新的可插拔(plugable)框架, 使得更加容易定制化调度.
- 调度框架定义一些扩展点, 在扩展点上注册的调度器插件会被调用.
- 一些插件可以改变调度决定, 一些可能仅仅是信息性插件.
- 一个插件可以注册在多个扩展点.

## (2)阶段:
- 一个Pod的调度可分为两个阶段: **scheduling cycle**和**binding cycle**, 两个阶段组成了**schedule context**.
- **调度cycle**: 为pod选择一个node; 是串行执行.
- **绑定cycle**: 将decision apply到集群; 可能会并发执行.
- 当pod被判断是不可调度或者有内部错误时, 调度或绑定cycle可以被中断, 该pod会被返回到队列和重试.

## (3)备注:
- https://kubernetes.io/docs/concepts/configuration/scheduling-framework/
- https://kubernetes.io/docs/reference/scheduling/profiles/

# 二 代码:
## (1)概述:
- kubernetes/pkg/scheduler
- kubernetes/cmd/kube-scheduler

## (2)framework(struct):
- registry: 一个map, 表示可用的插件的集合.
- snapshotSharedLister
- waitingPods
- pluginNameToWeightMap
- queueSortPlugins: 只能有一个工作.
- preFilterPlugins
- filterPlugins
- preScorePlugins
- scorePlugins
- reservePlugins
- preBindPlugins
- bindPlugins
- postBindPlugins
- unreservePlugins
- permitPlugins
- clientSet
- informerFactory
- volumeBinder
- metricsRecorder
- runAllFilters

## (3)Scheduler(struct):
- SchedulingQueue: 存放等待调度的pod.
- Algorithm: 一个接口.
- VolumeBinder: 处理pod的pv/pvc绑定.
- podPreemptor: 用于驱逐(evit)pod和更新被驱逐pod的NominatedNode属性.
- Profiles
- NextPod
- 等等.

# 三 调度cycle扩展点:
## (1)queue-sort: 
- 用于将调度队列中的pod排序, 队列排序查询需提供一个Less函数, 同一时刻只能有一个排序插件开启.
## (2)pre-filter
## (3)filter
## (4)post-filter
## (5)scoring
## (6)normalize scoring
## (7)reserve.

# 四 绑定cycle扩展点:
## (1)permit
## (2)pre-bind
## (3)bind
## (4)post-bind
