# 一 概述:
## (1)概述:
- 管理pods集合的部署和扩容, 保证这些pod**有序**和**唯一**, 用来管理**有状态**应用.
- 与Deployment一样, 基于相同的container spec来管理Pods; 不同的是, StatefulSet会为每个Pod维护一个sticky identity, Pod不是可互换的(interchangeable), pod有一个持久化identifier.
- 与其它controller操作方式一致, 在StatefulSet对象中定义期望的状态, Statefule控制器来负责调整当前状态到期望状态.

## (2)使用场景:
- Stable, 唯一网络identifier.
- Stable, 持久化存储.
- Ordered, 平滑部署和扩容.
- Ordered, 自动rolling updates.

## (3)限制:
- StatefuleSet当前需要一个Headless Service来作为Pods的网络身份(identity), 需要你创建该服务.
- 给定Pod的存储必须事先分配好(provisioned).
- 删除或缩小statefuleSet不会删除statefulset关联的volume.
- 当删除statefuleset被删除时, 不保证pod的终止顺序, 如要获得有序和平滑的pod终止方式, 可通过scale statefulset到0来替换删除.

## (4)备注:
- https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/
- https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/