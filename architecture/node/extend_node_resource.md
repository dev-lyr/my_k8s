# 一 概述:
## (1)概述:
- 扩展资源允许集群管理员向k8s告知其它节点级别资源.

## (2)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/extended-resource-node/
- https://kubernetes.io/docs/tasks/configure-pod-container/extended-resource/
- pod使用扩展资源时, 需指定requests和limits, 否则报错: Limit must be set for non overcommitable resources.
