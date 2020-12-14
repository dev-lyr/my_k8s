# 一 概述:
## (1)概述:
- local卷表示一个挂载的本地存储设备, 例如:disk,partition或目录.
- local卷只能通过静态创建pv卷来使用, 暂还不支持dynamic provisioning.

## (2)与hostPath比较:
- local卷无需手动调度pod到node, 系统会通过PersistentVolume上的node affinity来感知卷的node约束.

## (3)使用场景:
- local卷依赖node的稳定性,当node不健康时,local卷也是不可访问的,使用该卷的pod也不能运行.
- 使用本地卷的应用必须能够容忍潜在的数据丢失.

## (4)备注:
- https://kubernetes.io/blog/2019/04/04/kubernetes-1.14-local-persistent-volumes-ga/
