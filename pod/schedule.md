# 一 概述:
## (1)概述
- kube-scheduler: watch新创建的且未分配node的Pod, 为这些pod选择一个pod运行.
- 调度影响因素包括: 个别和集体资源需要, 硬件/软件/策略约束, affinity/anti-affinity, 数据locality, inter-workload和deadlines.

## (2)将pod assign给Node的方案:
- 可以通过一些方法来限制pod只能够或优先运行在特定的node上, 推荐使用**lable selectors**.
- 其它方案: Affinity/Anti-Affinity, nodeName(最简单,但限制多很少使用).
- 备注: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/

## (3)Taints和Tolerations:
- 与Node Affinity相反, Taints和Tolerations一起使用保证pods不会被调度到不合适的node上.
- 备注: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/

# 二 nodeSelector:
## (1)概述:
- **nodeSelector**是最简单推荐的node选择约束, nodeSelector是**PodSpec**的一个属性,它指定一个kv对map.
- 为了让pod能够运行在node上, node必须有所有kv对的labels, 通常是一个kv对.
- 给node添加label: `kubectl label nodes <node-name> <label-key>=<label-value>`

## (2)内置的node labels:
- kubernetes.io/hostname
- failure-domain.beta.kubernetes.io/zone
- failure-domain.beta.kubernetes.io/region
- beta.kubernetes.io/instance-type
- kubernetes.io/os
- kubernetes.io/arch

## (3)Node isolation/restriction:

# 三 Affinity/Anti-Affinity:
## (1)概述:
- nodeSelector提供使用指定label来约束pod调度到node的方式, affinity/anti-affinity特性则扩展了可以表达的约束类型.
- nodeSelector最终会被废弃, 因为affinity可以表达它的约束.

## (2)Affinity/anti-affinity的扩展点:
- 更具有表达式(不仅仅是AND).
- 可以设置某rule是soft/preference而不是硬性要求, 当scheduler不能满足时, 该pod仍会被调度.
- you can constrain against labels on other pods running on the node (or other topological domain), rather than against labels on the node itself, which allows rules about which pods can and cannot be co-located.

# 四 Taints和Tolerations:
## (1)原理:
- 将一个或多个taints应用到一个node, 表示node不接受任何不能tolerate这些taints的pod.
- tolerations被应用到pods, 允许(但是不是必须)pod调度到匹配taints的node上.
