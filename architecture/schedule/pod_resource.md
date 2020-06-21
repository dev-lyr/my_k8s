# 一 概述:
## (1)概述:
- 默认情况下,k8s集群中容器有unbounded计算资源(cpu和mem)可使用.
- 使用**ResourceQuota**, 集群管理员可以限制一个namespace的资源消费和创建.
- 使用**LimitRange**可以限制namespace内每个pod或容器的资源.
- Pod的内存和cpu资源的request和limit是Pod内所有container的总和.

## (2)资源类型:
- 可压缩资源(compressible resources): 当资源不足时, 只会饥饿, 不会退出, 例如:CPU.
- 不可压缩资源(incompressible resources): 例如:内存, 当资源不住时, 会OOM被杀死.

## (3)requests和limits：
- Pod的调度是基于资源的**request**.
- limits不受可用资源量的约束, 即limits总和可以超过节点资源总量, 即**超卖**.

## (4)备注:
- https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container
- k8s允许用户为节点添加自定义资源并支持pod的request来申请.

# 二 LimitRange对象:
## (1)功能:
- 限制一个namespace内每个Pod或容器的最小和最大计算资源使用量.
- 限制一个namespace内每个PersistentVolumeClaim的最小和最大存储请求.
- 限制一个namespace内request和limit的ratio.
- 设置一个namespace内默认的计算资源的request/limit, 并在运行时自动将它们inject到容器.

## (2)备注:
- LimitRange只能应用于单个pod或容器, 创建大量的Pod仍然会吃点集群资源, 此时需要ResourceQuota.
- https://kubernetes.io/docs/concepts/policy/limit-range/

# 三 ResourceQuota对象:
## (1)功能:
- 资源quota使用**ResourceQuota**对象来定义, 约束每个**namespace**可使用的资源量.
- 可以限制namespace内每种对象类型可以创建的对象数量.
- Resource Quota默认是开启的, 当apiserver --enable-admission-plugins flag有ResourceQuota参数时开启.

## (2)使用场景:
- 不同团队在不同namespace内工作.
- 管理员为每个namespace创建ResourceQuota.
- 用户在namespace内创建资源(pod,服务等), quota系统跟踪使用量, 并确保它不会超过ResourceQuota定义的硬资源限制.
- 若创建或更新一个资源超过quota限制, 则请求会以HTTP 403 FORBIDDEN返回失败.
- 若在namespace针对计算资源(cpu和mem)开启ResourceQuota, 用户必须指定这些值的request或limit, 否则quota系统将拒绝pod创建.

## (3)备注:
- https://kubernetes.io/docs/concepts/policy/resource-quotas/

# 四 内存资源:
## (1)概述:
- resource:requests:memory: 一个容器的内存资源请求.
- resource:limits:memory: 一个容器的内存资源限制.
- 备注: 内存资源计量单位为字节.
- 参考: https://kubernetes.io/docs/tasks/configure-pod-container/assign-memory-resource/

## (2)超出内存限制:
- 当Node有内存可用时, 容器允许使用超过request的内存数量, 但是不允许使用超过内存limit的数量.
- 若容器分配超过limit的内存数量, 该容器会成为被停止的候选者; 若容器持续消费超过它limit的内存, 该容器会被终止; 若被终止的容器可以被重启, kubelet会重启它, 与其他类型运行时失败一样.

## (3)不指定内存限制:
- case1: 容器可以使用Node所有可用的内存, 没有资源限制的容器更可能被OOM.
- case2: 容器运行在一个有默认内存限制的namespace(LimitRange)中, 则容器继承该内存限制.

## (4)使用:
- 通过配置容器的内存请求和限制, 可以高效利用集群中节点的内存资源.
- 配置一个小的内存请求, 可以让Pod更多机会被调度.
- 通过配置大于请求的内存限制, 可以让Pod在burst时有内存可用; 同时又将内存限制在一个合理的数量.

# 五 cpu资源:
## (1)概述:
- resource:request:cpu: 设置一个容器的cpu请求.
- resource:limit:cpu: 设置一个容器的cpu限制.

## (2)cpu单位:
- cpu资源通过cpu unit来计量, 在k8s中, 一个cpu等于: 1个aws vcpu, 一个gcp core, 一个Azure vCore, 带超线程的bare-metal intel处理的一个超线程.
- 允许使用小数, 例如:请求0.5cpu则保证使用一个cpu的一半, 可以使用m来表示千分之一, 例如:500m和0.5cpu是一样的.
- cpu请求是绝对值, 例如: 0.1cpu在一个单核或者多核机器上都一样.

## (3)不指定cpu限制:
- case1: 容器可以使用Node所有可用的cpu资源.
- case2: 容器运行在一个有默认cpu限制的namespace中, 则容器继承该cpu限制.

## (4)使用:
- 通过配置容器的cpu请求和限制, 可以高效利用集群中节点的cpu资源.
- 配置一个小的cpu请求, 可以让Pod更多机会被调度.
- 通过配置大于请求的cpu限制, 可以让Pod在burst时有cpu可用; 同时又将cpu限制在一个合理的数量.

# 六 Pod的QoS:
## (1)QoS类型:
- Guaranteed: Pod内的每个容器都必须有内存和cpu请求和限制, 且请求和限制值相等.
- BestEffort: Pod内的容器不能有内存或cpu的请求或限制.
- Burstable: 不满足Guaranteed, Pod内至少有一个容器有内存或cpu请求.

## (2)使用场景:
- 当宿主机资源紧张时, kubelet会对Pod进行Eviction, 此时参考QoS类型.

# 七 扩展资源:
## (1)概述:
- k8s允许定义自定义资源,同时支持pod资源request中申请该类资源.
- 扩展的资源的名称不能以kubernetes.io域名开头.

## (2)node级别扩展资源:
- 方式: 通过HTTP PATCH请求将可用资源数量更新的node的**status.capacity**属性.
- 相应实现: device插件等.

## (3)集群级别扩展资源

## (4)备注:
- https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/#extended-resources
