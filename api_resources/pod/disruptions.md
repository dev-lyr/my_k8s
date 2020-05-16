# 一 概述:
## (1)类型:
- involuntary disruptions: 不可避免.
- voluntary disruptions

## (2)involuntary disruptions:
- node的物理设备硬件故障.
- kernel panic
- 集群管理员误删虚拟机.
- 云provider或者hypervisor故障导致虚拟机消失
- 由于集群网络partition导致的集群的node消失.
- 由于out of resources导致的pod驱逐.

## (3)voluntary disruptions:
- 删除pod对应的controller.
- 更新deployment的pod模板导致的重启.
- 直接删除pod.
- drain一个node维修或升级.
- drain一个node来缩容集群.

## (4)备注:
- https://kubernetes.io/docs/concepts/workloads/pods/disruptions/

# 二 非自愿中断:
## (1)减轻非自愿中断的方法:
- 确保pod请求它需要的资源.
- 使应用多副本.
- 使应用多跨racks(使用anti-affinity)或者跨zones(使用multi-zone cluster)来运行多副本.

# 三 自愿中断:
## (1)概述:
- 默认情况下, 基本的k8s集群没有自愿中断发生.
- 但是集群管理员或者云提供商可能会导致自愿中断发生, 例如:滚动升级node上的软件. 一些集群(node)自动伸缩会导致自愿中断来进行碎片整理和压缩(compact)node.
- k8s提供一些features来帮助在频繁发生自愿中断的情况下运行高可用的应用, 将这些特性称作**Disruption Budgets**.

## (2)PodDisruptionBudget(PDB)


# 四 Drain Node:
- https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/
