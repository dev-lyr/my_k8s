# 一 概述:
## (1)概述:
- 调度框架是针对kube调度器的新的可插拔(plugable)框架, 使得更加容易定制化调度.
- 调度框架增加一些插件API到现有的调度器, 插件被编译进调度器; 将调度特性以插件形式实现, 可以保持调度核心的简单性和可维护性.
- 调度框架定义一些**扩展点**, 在扩展点上**注册的调度器插件**会被调用.
- 一些插件可以改变调度决定, 一些可能仅仅是信息性(informational)插件.
- 一个插件可以注册在多个扩展点.
- 只有一个worker来执行scheduleOne.

## (2)阶段:
- 一个Pod的调度可分为两个阶段: **scheduling cycle**和**binding cycle**, 两个阶段组成了**schedule context**.
- **调度cycle**: 为pod选择一个node, 顺序执行.
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

# 二 Scheduler结构:
## (1)概述:
- Scheduler watch未被调度的pods,尝试找到匹配的node并将binding写回给apiserver.

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

