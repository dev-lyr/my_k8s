# 一 概述:
## (1)概述:
- Kubernetes garbage collector: delete certain objects that once had an owner, but no longer have an owner.
- DeleteOptions资源: 可在删除对象时候提供.
- 一些对象是其它对象的owner, **被owner的对象是owner对象的依赖**, 每个依赖对象有一个**metadata.ownerReferences**属性指向owner对象, 在一些情况下, 该属性会被kube自动设置, 也可以手动设定.

## (2)删除类型:
- **级联删除(cascading deletion)**: 删除对象时自动删除它的依赖, 级联删除分为2种: backgroup和foregroup.
- **不删除依赖**: 若删除对象时不删除它的依赖, 则依赖会变成孤儿(orphaned).

## (3)设置删除策略:
- 当删除对象时, 在删除选项(deleteOptions)参数中设置**propagationPolicy**, 可能值为**Orphan**, **Foreground**或**Background**.
- kubectl删除时通过--cascade来控制, 默认为true.

## (4)备注:
- https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/
- 代码: kubernetes/pkg/controller/garbagecollector

# 二 级联删除:
## (1)类型:
- background
- foreground

## (2)foreground
- root对象先进入"deletion in progress"阶段: 对象通过REST API依旧可见; 对象的deletionTimestamp被设置; metadata.finalizers包含"foregroundDeletions".
- gc开始删除对象依赖, 一旦gc删除所有blocking依赖(对象的ownerReference.blockOwnerDeletion=true), 此后gc删除root对象.
- 若对象的ownerReferences被controller设置, 则blockOwnerDeletion会被自动设置.

## (3)background:
- k8s直接删除owner对象, gc稍后在后台删除它的依赖.
