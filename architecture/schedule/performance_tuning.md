# 一 概述:
## (1)概述:
- 在大规模集群中, 需要调整调度行为来平衡调度**延迟**(pod更快找到位置)和**准确性**(更准确的调度).

## (2)参考:
- https://kubernetes.io/docs/concepts/scheduling/scheduler-perf-tuning/

# 二 percentageOfNodesToScore:
## (1)概述:
- 控制调度器找到多少feasible节点就不再继续查找了,开始进入打分阶段.
- 范围: 0-100, 100表示找到所有feasible节点, 0表示使用kube-schedule默认的.

## (2)计算逻辑:
- kubernetes/pkg/scheduler/core/generic_scheduler.go: numFeasibleNodesToFind

## (3)经验:
- 集群几百节点或者更少不需要配置该值,没有太大效果.
- 不要将值配置的太小,避免小于10%,除非应用对调度器的throughput更看重,节点的分数不重要.
