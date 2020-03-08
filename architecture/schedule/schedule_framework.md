# 一 概述:
## (1)概述:
- 调度框架是针对kube调度器的新的可插拔(plugable)框架, 使得更加容易定制化调度.
- 调度框架定义一些扩展点, 在扩展点上注册的调度器插件会被调用.
- 一些插件可以改变调度决定, 一些可能仅仅是信息性插件.

## (2)阶段:
- 一个Pod的调度可分为两个阶段: **scheduling cycle**和**binding cycle**, 两个阶段组成了**schedule context**.
- **调度cycle**: 为pod选择一个node; 是串行执行.
- **绑定cycle**: 将decision apply到集群; 可能会并发执行.
- 当pod被判断是不可调度或者有内部错误时, 调度或绑定cycle可以被中断, 该pod会被返回到队列和重试.

## (3)备注:
- https://kubernetes.io/docs/concepts/configuration/scheduling-framework/

# 二 扩展点:
## (1)概述:
- 一个插件可以注册在多个扩展点.

## (2)调度cycle扩展点:
- queue-sort
- pre-filter
- filter
- post-filter
- scoring
- normalize scoring
- reserve.

## (3)binding cycle扩展点:
- permit
- pre-bind
- bind
- post-bind
