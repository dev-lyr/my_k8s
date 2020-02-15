# 一 概述:
## (1)概述:
- kube-scheduler是默认的调度器, 并且作为控制面板的一部分运行,若需要可以实现自己的调度组件来替换.
- 已存在nodes需要根据调度需求进行过滤(filter), 满足调度需求的node被称作**feasible** nodes.
- scheduler找一个pod的feasible nodes, 随后执行一些函数对node进行打分, 在这些feasible nodes中选择得分最高的node来运行pod.
- **binding**: scheduler将决定通知API server的过程.

## (2)2阶段:
- Filtering: 选择pod调度的feasible的nodes.
- Scoring: 对Filtering nodes进行排序从而选择最合适的pod placement, 若存在相同分数的则随机选择一个.

# 二 Filtering:

# 三 Scoring
