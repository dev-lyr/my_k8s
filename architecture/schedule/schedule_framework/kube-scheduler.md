# 一 概述:
## (1)概述:
- kube-scheduler是默认的调度器, 并且作为控制面板的一部分运行,若需要可以实现自己的调度组件来替换.
- 已存在nodes需要根据调度需求进行过滤(filter), 满足调度需求的node被称作**feasible** nodes.
- scheduler找一个pod的feasible nodes, 随后执行一些函数对node进行打分, 在这些feasible nodes中选择得分最高的node来运行pod.
- **binding**: scheduler将决定通知API server的过程.

## (2)2阶段:
- Filtering
- Scoring

## (3)备注:
- kubernetes/pkg/scheduler
- https://kubernetes.io/docs/reference/scheduling/policies

# 二 Filtering:
## (1)概述:
- 找出可用于调度pod的node的集合, filtering后, 剩下的都是合适的Nodes, 通常超过一个, 若为空, 则pod当前是不可调度的.

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

## (3)备注:
- kubernetes/pkg/scheduler/algorithm

# 三 Scoring:
## (1)概述:
- 调度器对filtering后剩余的Node进行排序, 选择最合适的Node.
- 调度器会给每个node分配一个score, 根据score规则进行打分.

## (2)score策略(priority):
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

## (3)备注:
- kubernetes/pkg/scheduler/algorithm
- 每个Priority函数的打分范围0-10,10表示优先级最高.
