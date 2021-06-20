# 一 概述:
## (1)概述:
- 由于(2)的场景, 造成一些pod运行的node已经不是最期望的node了, 因此descheduler可以根据它的policy, 将运行中的pod移动到其它节点.
- descheduler可以以Job或CronJob的形式运行在k8s集群内, 也可以指定descheduling-interval来定时循环执行, descheduler pod作为重要pod运行在kube-system命名空间中, 避免被自己或kubelet驱逐.
- Descheduler的policy是可配置的, 同时policy里的strategies也是可以enabled或disabled的.

## (2)使用场景:
- 一些node过载或者利用率低.
- 一些最初的调度desicion已经改变, 例如node上的taints或labels被添加或删除,pod或node affinity要求已不再满足.
- 一些node失败, 它上面的pod被移动到其它node.
- 新node加入到集群.

## (3)Strageties:
- RemoveDuplicates
- LowNodeUtilization
- RemovePodsViolatingInterPodAntiAffinity
- RemovePodsViolatingNodeAffinity
- RemovePodsViolatingNodeTaints
- RemovePodsViolatingTopologySpreadConstraint
- RemovePodsHavingTooManyRestarts
- PodLifeTime

## (4)备注:
- https://github.com/kubernetes-sigs/descheduler

# 二 descheduler命令:
## (1)概述:
- 运行descheduler, 镜像中执行的也是该命令.

## (2)参数:
- --kubeconfig
- --policy-config-file
- --dry-run 
- --descheduling-interval: 指定descheduler执行的间隔时间, 若未指定则只会运行一次.




