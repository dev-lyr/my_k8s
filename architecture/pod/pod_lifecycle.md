# 一 概述:
## (1)概述:
- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#podstatus-v1-core

## (2)Restart Policy:
- Always: 默认.
- OnFailure
- Never
- 备注: PodSpec的restartPolicy属性, 针对pod内所有容器.
- 备注: After containers in a Pod exit, the kubelet restarts them with an exponential back-off delay (10s, 20s, 40s, …), that is capped at five minutes.

## (3)Pod lifetime:
- 通常, pod只有在被人工或controller销毁时才会消失.
- 例外, 当Pods处于Successed或Failed超过一段时间(由master的**terminated-pod-gc-threshold**控制), Pod会被自动销毁.

## (4)备注:
- https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/

# 二 Pod 阶段(phase)
## (1)概述:
- Pod的**status**属性是一个**PodStatus**对象, 该对象有一个**phase**属性.
- The phase of a Pod is a simple, high-level summary of where the Pod is in its lifecycle. The conditions array, the reason and message fields, and the individual container status arrays contain more detail about the pod's status.

## (2)phase的值:
- **Pending**: k8s系统已经接受Pod, 但是容器镜像还没有创建, 包含被调度前的时间和通过网络下载镜像花费的时间.
- **Running**: Pod已经绑定到Node, 所有容器已经被创建, 至少一个容器在运行或处于启动或重启过程中.
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
- type: 对于(3)的type.

## (3)type:
- **PodScheduled**: Pod已被调度到一个node. 
- **Ready**: Pod可以接收请求并且可以被添加到负载均衡pool.
- **Initialized**: 所有init容器已经成功启动. 
- **UnSchedulable**: 调度器当前不能调度Pod, 例如:缺少资源或其它限制. 
- **ContainersReady**: Pod内所有容器都已经ready. 
- **其它类型**: **ReadinessGate**指定的额外的类型.

# 四 Pod readiness gate
## (1)概述:
- 可使用ReadinessGate来指定计算pod readiness的额外的condition条件, 若status.conditions没有某个condition, 则表示该condition状态为false.
- 是一个PodReadinessGate数组, 若指定, 则Pod只有在**所有容器都是Ready**且**所有ReadinessGate中的condition等于true**时候才算是ready(即PodCondition中type=Ready的值为true).
- PodReadinessGate: 只有一个conditionType字符串属性, 表示pod的condition list需要match的类型.

## (2)备注:
- https://github.com/kubernetes/enhancements/blob/master/keps/sig-network/0007-pod-ready++.md
