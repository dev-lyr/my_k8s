# 一 概述:
## (1)概述:
- kube-scheduler是**默认**的调度器, 并且作为控制面板的一部分运行,若需要可以实现自己的调度组件来替换.
- 已存在nodes需要根据调度需求进行过滤(filter), 满足调度需求的node被称作**feasible** nodes.
- scheduler找一个pod的feasible nodes, 随后执行一些函数对node进行打分, 在这些feasible nodes中选择得分最高的node来运行pod.
- **binding**: scheduler将决定通知API server的过程.

## (2)选择node的2阶段:
- Filtering
- Scoring

## (3)配置filtering和scoring行为的方式:
- 调度policies
- 调度Profiles

## (4)备注:
- kubernetes/pkg/scheduler
- https://kubernetes.io/docs/reference/scheduling/policies

# 二 Scheduling Policies:
## (1)概述:
- 允许配置kube-scheduler用来进行filtering和scoring 节点的**predicates**和**priorities**.

## (2)Filter策略(predicate):
- PodFitsHostPorts
- PodFitsHost
- PodFitsResources
- PodMatchNodeSelector
- NoVolumeZoneConflict
- NoDiskConflict
- MaxCSIVolumeCount
- CheckNodeMemoryPressure
- CheckNodePIDPressure
- CheckNodeDiskPressure
- CheckNodeCondition
- PodToleratesNodeTaints
- CheckVolumeBinding

## (3)score策略(priority):
- SelectorSpreadPriority
- InterPodAffinityPriority
- LeastRequestedPriority
- MostRequestedPriority
- RequestedToCapacityRatioPriority
- BalancedResourceAllocation
- NodePreferAvoidPodsPriority
- NodeAffinityPriority
- TaintTolerationPriority
- ImageLocalityPriority
- ServiceSpreadingPriority
- CalculateAntiAffinityPriorityMap
- EqualPriorityMap

## (4)备注:
- kubernetes/pkg/scheduler/algorithm/provider
- 每个Priority函数的打分范围0-10,10表示优先级最高.
- https://kubernetes.io/docs/reference/scheduling/policies/

# 三 profiles
