# 一 概述:
## (1)概述:
- Replicaset用来确保在任意时间指定数量的pod在运行.
- Deployment是更高层次的概念,它管理一个rs并提供pod的声明式更新,所以推荐deployment,除非你需要自定更新策略或不需要更新.

## (2)工作方式:
- rs定义一个selector指定如何标记它管理的pod, 一个replicass数量表示应该维护多少pods; 一个pod模板指定新创建pod的数据.
- rs通过创建和删除pod来达到期望的数量.
- rs通过pod的ownerReferences属性来连接到它的pods, 若一个pod没有ownerrefence或它的ownerreference不是一个controller且它满足rs的selector, 则该pod会被rs获得.

## (3)备注:
- rs修改镜像不会自动更新已有pod, 此时需使用deploy来支持; 可以手动删除pod, 此时新创建出来的就是新的镜像.

# 二 ReplicaSet资源:
## (1)ReplicaSetSpec:
- minReadySeconds
- replicas: 期望副本数, 默认为1.
- selector: 查询哪些pod match使用的label, 需match pod模板的labels.
- template: podTemplateSpec.

## (2)ReplicaSetStatus:
- availableReplicas: 当用副本的数量(至少ready minReadySeconds).
- conditions
- fullyLabledReplicas
- readyReplicas: ready的副本数量.
- replicas


# 三 ReplicaSetController:
## (1)功能:

## (2)属性:
- kubeclient
- syncHandler
- queue
- rsLister
- podLister
- expectations
- podControl
- burstReplicas
