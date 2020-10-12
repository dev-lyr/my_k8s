# 一 概述:
## (1)概述:
- 当可用计算资源较少时,kubelet需要保持node的稳定性, 尤其是不可压缩资源(内存,磁盘等),当资源耗尽时,node就会不稳定.
- kubelet通过evictionManager来实现node资源的监控和响应.
- 相关: node_lifecycle_controller实现node级别的驱逐.

## (2)驱逐策略:
- kubelet会主动监控和避免计算资源的starvation.
- 当出现饥饿时, kubelet会通过终止一个或多个pods来reclaim饥饿资源.
- kubelet终止pod的所有容器且pod将PodPhase转换为Failed; 若Pod由controller管理则会另外创建一个pod.
- 当到达驱逐阀门时, kubelet优先尝试回收node级别资源, 然后才是驱逐pod.

## (3)调度器:
- 当node处于压力情况下,调度器避免将其他pod调度到该node.
- MemoryPressure: 不会有新BestEffort pod被调度到该node.
- DiskPressure: 不会有新的pod调度到该节点.

## (4)备注:
- pkg/kubelet/eviction
- https://kubernetes.io/docs/tasks/administer-cluster/out-of-resource
- https://help.aliyun.com/document_detail/155464.html, 可基于psi提供一些主动驱逐的能力.

# 二 驱逐信号:
## (1)概述:
- kubelet基于(2)描述的signal做驱逐决定.
- 每个signal支持常量或者百分比.

## (2)驱逐信号:
- memory.available
- nodefs.available
- nodefs.inodesFree
- imagefs.available
- imagefs.inodesFree
- 备注: nodefs文件系统kubelet用于volumes,daemon日志等; imagefs文件系统被容器运行时用来存储镜像和容器写入层.

# 三 驱逐阀门:
## (1)概述:
- kubelet支持指定触发kubelet回收资源的驱逐阀门.
- 阈值的格式: [eviction-signal] [operator] [quantity]
- eviction-signal是(2)指定的信号.
- operator: 关系操作符, 例如小于.
- quantity: 驱逐阀门的数量, 例如:1Gi,也可以是百分比.

## (2)软驱逐阀门:
- 指定一个宽限期, 当资源出现饥饿时, kubelet会在宽限期之后再回收资源.
- eviction-soft: 指定一个驱逐阈值, 超过一段时间后会触发pod驱逐.
- eviction-soft-grace-period: 但达到软驱逐时等待多久再驱逐.
- eviction-max-pod-grace-period: 由于软驱逐导致pod被终止的最长grace period.
- 若指定最长grace period, 则kubelet使用它和pod.Spec.TerminationGracePeriodSeconds间的最小值.

## (3)硬驱逐阀门:
- 没有宽限期, 当出现资源饥饿时kubelet立即kill pod; 终止pod时也没有grace period.
- kubelet参数eviction-hard: 指定一个驱逐阀门集合.
- kubelet有默认的驱逐阀门.

## (4)驱逐监控间隔:
- housekeeping-interval
- global-housekeeping-interval 

# 四 Node conditions:
## (1)概述:
- kubelet将一个或多个驱逐信号映射到对应的node condition上, 当到达驱逐阀门(硬或软)时, kubelet会上报一个condition来反应node正处于压力下.
- 配置node-status-update-frequency: 默认10s,控制node状态的更新频率.

## (2)conditions:
- MemoryPressure: memory.available
- DiskPressure: nodefs.available,nodefs.inodesFree,imagesfs.available或imagefs.inodesFree.

## (3)conditions震荡(oscillation):
- 若node在软驱逐阀门间震荡, 但有没有超过设置的grace period, 则会导致node的condition在true和false之间经常震荡, 影响调度.
- 配置eviction-pressure-transition-period: 避免condition频繁变化.

# 五 node资源回收:
## (1)概述:
- 当到达驱逐阀门时, kubelet优先尝试回收node级别资源, 然后才是驱逐pod.

## (2)回收node级别资源:
- 指定imagefs情况: 若nodefs达到驱逐阀门, 则kubelet通过删除dead pod和它们容器来释放磁盘空间; 当imagefs达到驱逐阀门, 通过删除所有未使用镜像来释放磁盘空间.
- 未指定imagefs情况: 先删除dead pod和它们的容器;删除所有未使用镜像.

# 六 驱逐终端用户pod:
## (1)概述:
- 当kubelet不能从node上回收足够资源时, kubelet会开始驱逐pods.
- kubelet首先判断pods使用的饥饿资源是否超过请求量,然后根据**priority**,再根据饥饿资源相对于请求量的使用情况对pods进行排序.

## (2)排序和驱逐pod的顺序:
- BestEffort和Burstable中Pods饥饿资源使用量超过request的, 这些pod先根据优先级然后超过请求的使用量进行排序.
- Guaranteed和Burstable Pods中使用量小于请求量的最后被驱逐.
- 当系统daemon(kubelet,docker和journald)使用资源超过通过system-reserverd或kube-reserverd分配的量, 并且node只有Guaranteed和Burstable pods使用量少于请求量时, 根据Pod的Priority进行驱逐.

## (3)最小资源回收:
- 特定场景下, 驱逐pods会导致回收少量资源, 会导致kubelet重复的hit驱逐阀门.
- 为了解决上述问题, kubelet可指定每个资源的**mininum-reclaim**, kubelet观察资源压力, 使用总量减去阈值超过mininum-reclaims指定的量时再回收.
- 默认情况下eviction-minimum-reclaim为0.

# 七 Node OOM行为:
## (1)概述:
- 在kubelet能够回收内存之前, node遇到系统OOM事件, 此时node根据**oom-killer**来响应.
- 与Pod驱逐不同, 若Pod的容器被oom杀死, 根据RestartPolicy容器可能被kubelet重启.

## (2)设置oom_score_adj:
- Guaranteed: -998
- BestEffort: 1000
- Burstable: min(max(2, 1000 - (1000 * memoryRequestBytes) / machineMemoryCapacityBytes), 999)

## (3)oom-killer行为:
- 根据容器使用内存的百分比计算出oom_score, 并加入oom_score_adj计算出一个有效的oom_score, 接着杀死得分最高的容器.

# 八 最佳实践
## (1)DaemonSet:
- 若不希望ds中的pod被驱逐, 则指定一个足够高的优先级.
- 若希望ds中pod只有在资源丰富时运行, 则指定一个低的或默认优先级.
