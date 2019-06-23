# 一 概述:
## (1)概述:
- 管理pods集合的部署和扩容, 保证这些pod**有序**和**唯一**, 用来管理**有状态**应用.
- 与Deployment一样, 基于相同的container spec来管理Pods; 不同的是, StatefulSet会为每个Pod维护一个sticky identity, Pod不是可互换的(interchangeable), pod有一个持久化identifier.

## (2)使用场景:
- Stable, 唯一网络identifier.
- Stable, 持久化存储.
- Ordered, 平滑部署和扩容.
- Ordered, 自动rolling updates.

## (3)备注:
- https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/
- https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/
