# 一 概述:
## (1)概述

## (2)Restart Policy:
- Always: 默认.
- OnFailure
- Never
- 备注: PodSpect的restartPolicy属性, 针对pod内所有容器.

## (3)Pod lifetime:
- 通常, pod只有在被人工或controller销毁时才会消失.
- 例外, 当Pods处于Successed或Failed超过一段时间(由master的**terminated-pod-gc-threshold**控制), Pod会被自动销毁.

## (4)备注:
- https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/

# 二 Pod 阶段(phase)
## (1)概述:
- Pod的**status**属性是一个**PodStatus**对象, 该对象有一个**phase**属性.

## (2)phase的值:
- **Pending**: k8s系统已经接受Pod, 但是容器镜像还没有创建, 包含被调度前的时间和通过网络下载镜像花费的时间.
- **Running**: Pod已经绑定到Node, 所有容器已经被创建, 至少一个容器在运行, 或处于启动或重启过程中.
- **Successed**: Pod内所有容器已经成功结束, 不会被重启.
- **Failed**: Pod内所有容器已经终止, 且至少一个容器终止失败, 容器以非0退出或被系统终止.
- **Unknown**: 由于一些原因不能获得Pod的状态, 通常是和Pod的Node通信失败.
- **Completed**: Pod由于没有事情做而运行完成, 例如: Completerd Jobs.
- **CrashLoopBackOff**: Pod内的一个容器非预期退出, 并且可能基于重启策略重启也报非0错误.

## (3)备注:
- 当node失败时, node上的所有pod的phase会被设置为Failed.

# 三 Pod Condition:
## (1)概述:
- 表示pod当前的服务状态, PodStatus的conditions属性, 是一个**PodCondition**对象数组.

## (2)PodCondition对象属性:
- lastProbeTIme: 表示pod最近一次被探测时间戳.
- lastTransitionTime: 表示Pod最近一个状态变化的时间戳.
- message: transition的细节.
- reason: transition的原因.
- status: 字符串, 可能值"True", "False"和"Unknown".
- type: PodScheduled(Pod已被调度到一个node); Ready(Pod可以接收请求并且可以被添加到负载均衡pool); Initialized(所有init容器已经成功启动); UnSchedulable(调度器当前不能调度Pod, 例如:缺少资源或其它限制); ContainersReady(Pod内所有容器都已经ready).

# 四 容器probes:
## (1)概述:
- kubelet定期在容器上执行诊断性探测(Probe), 为了执行诊断, kubelet调用容器实现的handler.
- 每次探测的结果: Success(容器通过诊断), Failure(容器诊断失败), Unknown(诊断失败, 不会执行action).

## (2)探测类型:
- livenessProbe: 表示容器是否在运行中, 若探测失败, kubelet会kill容器, 并根据restart policy来决定如何操作, 若容器不提供livenessProbe探测则默认状态为Success.
- readinessProbe: 表示容器是否准备好接收请求, 若探测失败, 则**endpoint控制器**会从所有服务的endpoints中删除pod的ip地址.readiness在init delay之前的默认状态为Failure, 若不提供readiness探测则默认state为Success.

## (3)handler类型:
- ExecAction: 在容器内执行一个指定命令, 若退出码为0则表明诊断执行成功.
- TCPSocketAction: 在容器IP地址上指定Port执行一个TCP检查, 若port是打开则表示诊断成功.
- HTTPGetAction: 在容器IP地址上指定port和路径上执行一个Get请求, 若返回码>=200且<400则表示诊断成功.

## (4)ReadinessGate

## (5)使用场景:
- 若希望容器在测探失败时被kill并且重启, 则提供一个liveness probe, 且指定restartPolicy为Always或OnFailure.
- 若希望在探测成功后再发生请求给Pod, 则提供一个readiness probe.

## (6)备注:
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/

# 五 容器状态:
## (1)概述:
- Pod的status属性包含一个containerStatuses属性表示Pod每个容器的状态, initContainerStatuses表示每个init容器的状态.
- containerStatuses是一个**ContainerStatus**对象的数组, ContainerStatus对象的state表示容器的状态.

## (2)ContainerState:
- Running: 表示容器正在没问题的执行中, 当容器进入Running状态时, postStart hook(若有)会被执行.
- Terminated: 表示容器已经完成执行操作且已停止运行, 当容器成功执行完成或由于一些原因失败时候会进入该状态, 在进入Terminated状态之前, PreStop hook(若有)会被执行.
- Waiting: 容器的默认状态, 若容器不在Running或Terminated状态, 则处于Waiting状态, Waiting状态的容器依旧执行它要求的操作, 例如拉镜像, 应用secrets等.

