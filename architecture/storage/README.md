# 一 概述:
## (1)背景:
- 容器中磁盘上文件是临时的, 当容器crash时, kubelet会重启它(新建一个), 但是文件内容会丢失(新启动的容器是clean状态).
- Pod内的容器需要在容器间共享文件.
- Docker也有卷的概念, 但Docker中的卷管理松散, 只是磁盘上或其他容器内的一个目录, Docker虽然提供了卷driver, 但目前功能有限(目前官方只有local volume driver).
- K8s使用卷(volume)来解决上面的问题.

## (2)k8s的卷:
- 当容器重启时, 卷内容不变, 新启动容器可以识别前容器写入卷的所有文件.
- 卷的声明周期与pod一样, 当pod不存在时卷也将不存在(pod被自动重建后呢?).
- 若Pod包含多个容器, 则这个卷可以被Pod内所有容器**共享使用**.
- K8s支持多种类型的卷, 且一个pod可以同时使用任意数量卷.
- 通过spec.volumes属性来指定Pod内容器可以使用的卷; 使用spec.containers.volumeMounts属性来指定容器的挂载点.

## (3)卷的类型(podSpec.volumes):
- **emptyDir**: 用于存储临时数据的简单空目录.
- **hostPath**: 挂载host node上的一个文件或目录到Pod内.
- **configMap**: 提供一种将配置数据注入Pods的方式.
- **dowmwardAPI**: 使downward API数据对应用可见.
- **secret**: 用于向Pod传送敏感信息, 例如:密码.
- **persistentVolumeClaim(PVC)**: 单独介绍.
- **gitRepo**: 已废弃.
- **csi**: 重要, 后续单独介绍.
- **flexVolume**: csi出现之前的out-of-tree插件机制.
- **nfs**: 挂载到Pod中的NFS共享卷.
- **其它类型网络存储**: glusterfs, iscsi, cephfs等.
- **云厂商提供**: awsElasticBlockStore, gcePersistentDisk, azureDisk.

## (4)使用:
- podSpec.volumes(**Volume**数组): 属于pod的可以被挂载到容器里的volumes.
- 挂载: 在container对象的volumeMounts属性表示.
- 备注: Volume对应(3)的类型.

## (5)源码:
- pkg/controller/volume: 各种controller.
- pkg/volume

## (6)out-of-tree volume插件:
- out-of-truee插件包括CSI和FlexVolume, 允许存储供应商创建自定义存储插件, 无需将它们添加到kubernetes仓库.
- 备注: https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md

## (7)PV使用步骤:
- Attach
- Mount
- 将主机上volume通过Mounts传递给docker.

## (8)备注:
- pv的在kubelet的主控制循坏之外, 不阻塞主控流程.

# 二 重要接口:
## (1)概述:
- Volume
- Attacher/Detacher: attach/detach一个volume到node.
- Mounter/Unmounter: mount/unmount volume.
- VolumePlugin
- PersistentVolumePlugin

## (2)Volume:
- GetPath
- MetricsProvider

## (3)Attcher:
- Attach
- VolumesAreAttached
- WaitforAttach

## (4)Detacher:
- Detach
- UnmountDevice

## (5)Mounter:
- CanMount
- SetUp
- SetUpAt

## (6)UnMounter:
- TearDown
- TearDownAt

