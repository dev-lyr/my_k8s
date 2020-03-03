# 一 概述:
## (1)概述
- kube-scheduler: watch新创建的且未分配node的Pod, 为这些pod选择一个pod运行.
- 调度影响因素包括: 个别和集体资源需要, 硬件/软件/策略约束, affinity/anti-affinity, 数据locality, inter-workload和deadlines.

## (2)将pod assign给Node的方案:
- 可以通过一些方法来限制pod只能够或优先运行在特定的node上. 
- 方案: node selectors, Affinity/Anti-Affinity, nodeName(最简单,但限制多很少使用).
- 备注: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/

## (3)Taints和Tolerations:
- 与Node Affinity相反, Taints和Tolerations一起使用保证pods不会被调度到不合适的node上.
- 备注: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/

## (4)pod priority和preeption:
- 备注: https://kubernetes.io/docs/concepts/configuration/pod-priority-preemption/

# 二 nodeSelector:
## (1)概述:
- **nodeSelector**是最简单推荐的node选择约束, nodeSelector是**PodSpec**的一个属性,它指定一个kv对map.
- 为了让pod能够运行在node上, node必须有**所有**kv对的labels, 通常是一个kv对.
- 给node添加label: `kubectl label nodes <node-name> <label-key>=<label-value>`

## (2)内置的node labels:
- kubernetes.io/hostname
- failure-domain.beta.kubernetes.io/zone
- failure-domain.beta.kubernetes.io/region
- beta.kubernetes.io/instance-type
- kubernetes.io/os
- kubernetes.io/arch

## (3)Node isolation/restriction:
- 功能: 保证node的label不会被kubelet进程修改.
- 实现方式: 通过NodeRestriction admission插件来阻止kubelet设置和修改**node-restriction.kubernetes.io**为前缀的labels.

# 三 Affinity和Anti-Affinity:
## (1)概述:
- affinity: 亲和性; Anti-Affinity: 反亲和性.
- nodeSelector提供使用指定label来约束pod调度到node的方式, affinity/anti-affinity特性则扩展了可以表达的约束类型.
- nodeSelector最终会被废弃, 因为affinity可以表达它的约束.
- affinity只在调度pod时候有用, 删除或修改node上lable对已被调度的pod无影响.

## (2)类型:
- node的affinity: 与nodeSelector类似, 但有(3)说的前2个优势.
- pod间affinity/anti-affinity: 通过已经运行在node上的pod的label来调度, 而不是node label来约束.

## (3)Affinity/anti-affinity的优势:
- 更强的表达式(不仅仅是AND).
- 可以设置某rule是soft/preference而不是硬性要求, 当scheduler不能满足时, 该pod仍会被调度.
- you can constrain against labels **on other pods running on the node** (or other topological domain), rather than against labels on the node itself, which allows rules about which pods can and cannot be co-located.

## (4)Node的affinity:
- 通过PodSpec的affinity属性的nodeAffinity属性指定. 
- 当前支持两类Node Affinity: **requiredDuringSchedulingIgnoredDuringExecution(硬限制)**和**preferredDuringSchedulingIgnoredDuringExecution(软限制)**.
- 支持operator: In, NotIn, Exists, DoesNotExist, Gt, Lt; 可通过NotIn和DostNotExist来实现Node的anti-affinity.
- nodeSelector和NodeAffinity同时使用, 则Pod需要都满足时才能被调度到该节点.
- 若多个nodeSelectorTerms, 则pod满足任一即可被调度到该节点; 若nodeSelectorTerms包含多个matchExpression, 则pod需满足所有matchExpression.

## (5)Pod的Affinity和Ant-Affinity:
- 通过PodSpec的affinity属性的**podAffinity**和**podAntiAffinity**属性来配置.
- 当前支持两类Pod Affinity和Anti-Affinity: requiredDuringSchedulingIgnoredDuringExecution和preferredDuringSchedulingIgnoredDuringExecution.

# 四 Taints和Tolerations:
## (1)概述:
- 将一个或多个**taints**应用到一个node, 表示node不接受任何不能tolerate这些taints的pod. 备注: NodeSpec的taints数组.
- **tolerations**被应用到pods, 允许(但是不是必须)pod调度到匹配taints的node上. 备注: PodSpec的tolerations数组.
- toleration和taint匹配的条件: key和effect是一样的, operator是exist或者equal且value相等.

## (2)Taint
- effect: 可选值:NoSchedule, PreferNoSchedule和NoExecute.
- key
- timeAdded
- value
- 备注: 通常,若一个effect为NoExecute的taint添加到node上, 则不能忍受该taint的pod会被立即清除; pod可以通过使用effect为NoExecute且指定tolerationTimeout属性来设置在tained被添加到node多久后再清除该pod.

## (3)Toleration:
- effect: 空表示match所有taint effects, 若指定, 可选值:NoSchedule, PreferNoSchedule(soft版NoSchedule, 系统会尝试避免将一个pod放在不能忍受taint的node上)和NoExecute.
- key: toleration应用的taint key, 空的key表示可以容忍所有.
- operator: 合理是Exists和Equal, 默认为Equal.
- tolerationSeconds: 表示toleration容忍的taint的时间周期, 默认不设置, 表示永久容忍该taint(不会被evict); 该值只在effect为NoExecute时有效.
- value: toleration需match的taint的值; 若operator是Exists, 该值需为空; 其它operator时, 为普通字符串.

## (4)备注:
- kubectl taint

# 五 Pod优先级和抢占:
## (1)概述:
- Pod可以有优先级, 优先级表示该Pod相对其它pod的重要性.
- 若Pod不能被调度, 调度器会尝试抢占(驱逐)低优先级的Pod来使得高优先级得到调度.
- 在kube1.9后, 优先级也会影响pod的调度顺序以及node上out-of-resources时的eviction顺序.
- PriorityClass是一个non-namespaced对象, 定义了一个优先级类name和优先级整数值的映射, 越高优先级值, 级别越高.

## (2)使用:
- 使用优先级和抢占: 添加一个或多个**PriorityClass**资源; 创建Pod时候PodSpec的priorityClassName指定为特定的PriorityClass对象.
- 关闭抢占: 将PodSpec的preemptionPolicy设置为Nerver; 将kube-schedule的flag disablePreemption设置为true.
- 备注: 在kube 1.12+, 当集群资源有压力时, 重要的Pod依赖调取器抢占来调度, 因此不推荐关闭抢占.

## (3)PriorityClass资源:
- globlaDefault: 表示该PriorityClass的值作为没有指定priorityClassName的pod的默认值, 系统只能存在一个globalDefault的priorityClass对象, 若不指定, 则没有PriorityClassName的Pod的优先级为0.

## (4)Pod优先级:
- 创建Pod在spec指定了priorityClassName, 则**priority admission controller**使用priorityClassName属性并填入(populate)priority的整数值.

## (5)抢占.
