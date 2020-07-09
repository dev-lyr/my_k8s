# 一 概述:
## (1)概述:
- 当可用计算资源较少时, kubelet需要保持node的稳定性, 尤其是不可压缩资源(内存,磁盘等),当资源耗尽时,node就会不稳定.

## (2)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource

# 二 驱逐策略:
## (1)概述:
- kubelet主动的监控和避免资源饥饿,当出现饥饿时,kubelet通过failing一个或多个pod来回收资源.
- 当pod被kubelet fail后, pod会终止所有容器且PodPhase转换为Failed,若Pod被Deployment等管理时, Deployment会创建一个新的Pod.
- 当到达驱逐阀门时, kubelet优先尝试回收node级别资源, 然后才是驱逐pod.

## (2)驱逐信号:
- memory.available
- nodefs.available
- nodefs.inodesFree
- imagefs.available
- imagefs.inodesFree

## (3)驱逐阀门:
- kubelet支持指定触发kubelet回收资源的驱逐阀门.
- 格式: [eviction-signal] [operator] [quantity]
- eviction-signal是(2)指定的信号.
- operator: 关系操作符, 例如小于.
- quantity: 驱逐阀门的数量, 例如:1Gi,也可以是百分比.

## (4)硬驱逐阀门:
- 没有宽限期, 当出现资源饥饿时kubelet立即kill pod.
- kubelet参数eviction-hard: 指定一个驱逐阀门集合.
- kubelet有默认的驱逐阀门.

## (5)软驱逐阀门:
- 指定一个宽限期, 当资源出现饥饿时, kubelet会在宽限期之后再回收资源.
- eviction-soft
- eviction-soft-grace-period
- eviction-max-pod-grace-period
- 相关: pod.Spec.TerminationGracePeriodSeconds.

## (6)Node conditions:
- kubelet将一个或多个驱逐信号映射到对应的node condition上, 当到达驱逐阀门时, kubelet会报告一个condition来反应node正处于压力下.
- MemoryPressure
- DiskPressure
- 配置node-status-update-frequency
- 配置eviction-pressure-transition-period: 避免condition频繁变化.

## (7)回收node级别资源:
- 指定imagefs情况: 若nodefs达到驱逐阀门, 则kubelet通过删除dead pod和它们容器来释放磁盘空间; 当imagefs达到驱逐阀门, 通过删除所有未使用镜像来释放磁盘空间.
- 未指定imagefs情况: 先删除dead pod和它们的容器;删除所有未使用镜像.

## (8)调度器:
- 当node处于压力情况下,调取器避免将其他pod调度到该node.
- MemoryPressure: 不会有新想BestEffort pod被调度到该node.
- DiskPressure: 不会有新的pod调度到该节点.

# 三 驱逐pod:

# 四 Node OOM行为:
## (1)概述:
- 在kubelet能够回收内存之前, node遇到系统OOM事件, 此时node根据**oom-killer**来响应.
- 与Pod驱逐不同, 若Pod的容器被oom杀死, 根据RestartPolicy容器可能被kubelet重启.

## (2)设置oom_score_adj:
- Guaranteed: -998
- BestEffort: 1000
- Burstable: min(max(2, 1000 - (1000 * memoryRequestBytes) / machineMemoryCapacityBytes), 999)

## (3)oom-killer行为:
- 根据容器使用内存的百分比计算出oom_score, 并加入oom_score_adj计算出一个有效的oom_score, 接着杀死得分最高的容器.

# 五 最佳实践
