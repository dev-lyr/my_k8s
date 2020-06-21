# 一 概述:
## (1)概述:
- 管理pods集合的部署和扩容, 保证这些pod**有序**和**唯一**, 用来管理**有状态**应用.
- 与Deployment一样, 基于相同的container spec来管理Pods; 不同的是, StatefulSet会为每个Pod维护一个sticky identity, Pod不是可互换的(interchangeable), pod有一个持久化identifier.

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

## (4)创建步骤:
- 创建存储数据的PVC.
- 创建一个headless服务.
- 创建statefulset本身.

## (5)备注:
- https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/
- https://kubernetes.io/docs/tutorials/stateful-application/basic-stateful-set/

# 二 Pod Identity:
## (1)概述:
- 组成: 一个ordinal, 一个stable网络identity, 固定存储.

## (2)ordinal index:
- 假设statefuleset N副本, 则ordinal为0到N-1.

## (3)稳定网络ID:
- statefulset中每个pod的hostname格式为: statefulsetname-ordinal, pod的name一样.
- statefulset使用headless服务来控制pods的域名, 当pod被创建后, pod的域名为podname.servicename.namespace.svc.cluster.local, cluster.local是集群domains.

## (4)稳定存储:
- 使用pv.
- 当pod被删除或statefuleset被删除时, pv不会被删除, 需要手动删除.

## (5)pod name lable:
- statefulset创建的pod会加上一个statefulset.kubernetes.io/pod-name标签, 值是pod的name,格式和hostname一样.

## (6)备注:
- 参考newStatefulSetPod函数.

# 三 statefulSetSpec
## (1)属性:
- podManagementPolicy
- replicas
- selector
- template
- serviceName: 管理statefuleset的服务的名字, 服务需在statefuleset之前创建并且负责set的网络唯一标识.
- volumeClaimTemplates
- updateStratery

