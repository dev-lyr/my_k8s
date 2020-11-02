# 一 概述:
## (1)概述:
- 调度框架是针对kube调度器的新的可插拔(plugable)框架, 使得更加容易定制化调度.
- 调度框架增加一些插件API到现有的调度器, 插件被编译进调度器; 将调度特性以插件形式实现, 可以保持调度核心的简单性和可维护性.
- 调度框架定义一些扩展点, 在扩展点上注册的调度器插件会被调用.
- 一些插件可以改变调度决定, 一些可能仅仅是信息性(informational)插件.
- 一个插件可以注册在多个扩展点.
- 只有一个worker来执行scheduleOne.

## (2)阶段:
- 一个Pod的调度可分为两个阶段: **scheduling cycle**和**binding cycle**, 两个阶段组成了**schedule context**.
- **调度cycle**: 为pod选择一个node, 顺序执行(无锁化).
- **绑定cycle**: 将decision apply到集群; 可能会并发执行(代码中异步goroutine实现), 乐观绑定.
- 当pod被判断是不可调度或者有内部错误时, 调度或绑定cycle可以被中断, 该pod会被**返回到队列和重试**.

## (3)特点:
- 无锁化
- cache化
- 乐观绑定
- 备注: kubelet会调用HandlePodAdditions检查pod是否应该运行在该节点.

## (4)备注:
- https://kubernetes.io/docs/concepts/configuration/scheduling-framework/
- https://kubernetes.io/docs/reference/scheduling/profiles/
- https://github.com/kubernetes-sigs/scheduler-plugins
- kubernetes/pkg/scheduler
- kubernetes/cmd/kube-scheduler
- https://github.com/kubernetes-sigs/scheduler-plugins

# 二 Scheduler结构:
## (1)概述:
- Scheduler watch未被调度的pods,尝试找到匹配的node并将binding写回给apiserver.
- 相关文件: scheduler.go, core和factory.go

## (2)属性:
- SchedulingQueue: 存放等待调度的pod.
- Algorithm: 对应一个genericScheduler.
- VolumeBinder: 处理pod的pv/pvc绑定.
- podPreemptor: 用于驱逐(evit)pod和更新被驱逐pod的NominatedNode属性.
- Profiles
- 等等.

## (3)方法:
- NextPod: 一个函数, block直至下个pod可用.
- **scheduleOne**: 执行一个pod全部调度流程.

## (4)方法

# 三 调度cycle扩展点:
## (1)queue-sort: 
- 用于将调度队列中的pod排序, 队列排序查询需提供一个Less函数, 同一时刻只能**开启一个排序插件**.

## (2)pre-filter:
- 这些插件用来预处理pod的信息或者检查集群或pod必须满足的条件.

## (3)filter:
- 用来过滤掉不能运行pod的node, 针对每个node, 调度器会根据配置的顺序调用filter插件.
- 若任一filter插件将某node标记为infeasible, 则剩余的filter插件不会被调用.
- Node可以被并发的评估.

## (4)pre-scoring:
- 这些插件用来执行pre-scoring工作, 会生成一个由score插件共享的state.

## (5)scoring:
- 用来对通过filting阶段的pod进行排序, 会对每个node调用scoring插件.

## (6)normalize scoring:
- 用来在scheduler计算node的最终排名前修改scores.

## (7)reserve:
- 一个informational扩展点.
- Plugins which maintain runtime state (aka "stateful plugins") should use this extension point to be notified by the scheduler when resources on a node are being reserved for a given Pod.

# 四 绑定cycle扩展点:
## (1)permit:
- 在每个pod的调度cycle结束后被调用, 用来prevent或delay绑定到候选nodes.
- approve: 一旦所有permit插件approve一个node, 则它会被发送到binding.
- deny: 任一permit插件拒绝一个pod, 它会被退回到调度queue, 同时触发Unreserve插件.
- wait: 若一个permit插件返回wait, 则pod会被放入一个内部的waiting pod列表, 直至pod被approved, 若timeout发生, 则wait会变为deny, 则pod会被返回到调度queue, 同时触发Unreserve插件.

## (2)pre-bind:
- 用来执行pod bind前需要做的任何事情,例如:一个pre-bind插件会在pod bind前提供一个network volume并将它挂载到目标node上.

## (3)bind:
- 这些插件用于将一个pod bind到一个node.
- bind插件在所有pre-bind插件执行完成后才会被调用.
- 根据配置的order来调用每个bind插件, bind插件可以选择是否处理给定pod, 若选择处理, 则其余的bind插件会被跳过.

## (4)post-bind:
- 一个informational扩展点, post-bind插件会在pod被成功bind后调用, 是binding阶段的end, 可以用来清理一些相关资源.

## (5)Unreserve:
- 一个informational扩展点, 若pod被reserved且在后续阶段被拒绝, 则unreserve插件会被通知.
- Unreserve插件用来清理一些reserved pod相关的state.
