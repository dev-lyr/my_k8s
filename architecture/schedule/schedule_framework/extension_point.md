# 一 概述:
## (1)概述:
- 调度框架定义一些**扩展点**, 在扩展点上**注册的调度器插件**会被调用.
- 一些插件可以改变调度决定, 一些可能仅仅是信息性(informational)插件.
- 一个插件可以注册在多个扩展点.

## (2)queue-sort扩展点:
- 这些插件提供一个排序函数,对调度队列中pending的pod进行排序, 同一时刻只能开启一个排序插件.

## (3)调度cycle扩展点:
- pre-filter
- filter
- pre-scoring
- scoring
- normalize scoring
- reserve
- permit

## (4)绑定cycle扩展点:
- prebind
- bind
- postbind

## (5)调度插件:
- pkg/scheduler/framework/plugins

# 二 pre-filter:
## (1)概述:
- 这些插件用来预处理pod的信息或者检查集群或pod必须满足的条件.
- 可将pod标记为不可调度.

## (2)PreFilterPlugin接口:
- Name()
- PreFilter()
- PreFilterExtensions()

# 三 filter:
## (1)概述:
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

## (8)permit:
- 在每个pod的调度cycle结束后被调用, 用来prevent或delay绑定到候选nodes.
- approve: 一旦所有permit插件approve一个node, 则它会被发送到binding.
- deny: 任一permit插件拒绝一个pod, 它会被退回到调度queue, 同时触发Unreserve插件.
- wait: 若一个permit插件返回wait, 则pod会被放入一个内部的waiting pod列表, 直至pod被approved, 若timeout发生, 则wait会变为deny, 则pod会被返回到调度queue, 同时触发Unreserve插件.

# 五 绑定cycle扩展点:
## (1)pre-bind:
- 用来执行pod bind前需要做的任何事情,例如:一个pre-bind插件会在pod bind前提供一个network volume并将它挂载到目标node上.

## (2)bind:
- 这些插件用于将一个pod bind到一个node.
- bind插件在所有pre-bind插件执行完成后才会被调用.
- 根据配置的order来调用每个bind插件, bind插件可以选择是否处理给定pod, 若选择处理, 则其余的bind插件会被跳过.

## (3)post-bind:
- 一个informational扩展点, post-bind插件会在pod被成功bind后调用, 是binding阶段的end, 可以用来清理一些相关资源.

## (4)Unreserve:
- 一个informational扩展点, 若pod被reserved且在后续阶段被拒绝, 则unreserve插件会被通知.
- Unreserve插件用来清理一些reserved pod相关的state.
