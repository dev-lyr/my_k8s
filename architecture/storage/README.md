# 一 概述:
## (1)背景:
- 容器中磁盘上文件是临时的, 当容器crash时, kubelet会重启它(新建一个), 但是文件内容会丢失(新启动的容器是clean状态).
- Pod内的容器需要在容器间共享文件.
- Docker也有卷的概念, 但Docker中的卷管理松散, 只是磁盘上或其他容器内的一个目录, Docker虽然提供了卷driver, 但目前功能有限(目前官方只有local volume driver).
- K8s使用卷(volume)来解决上面的问题.

## (2)k8s的卷:
- 当容器重启时, 卷内容不变, 新启动容器可以识别前容器写入卷的所有文件.
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
- **csi**: csi-ephemeral-volume.
- **flexVolume**: csi出现之前的out-of-tree插件机制.
- **nfs**: 挂载到Pod中的NFS共享卷.
- **其它类型网络存储**: glusterfs, iscsi, cephfs等.
- **云厂商提供**: awsElasticBlockStore, gcePersistentDisk, azureDisk.

## (4)使用:
- podSpec.volumes(**Volume**数组): 属于pod的可以被挂载到容器里的volumes.
- 挂载: 在container对象的volumeMounts属性表示.
- 备注: Volume对应(3)的类型.

## (5)相关模块:
- pkg/controller/volume: 各种controller.
- pkg/volume
- pkg/kubelet/volumemanager

## (6)out-of-tree volume插件:
- out-of-truee插件包括:**CSI**和**FlexVolume**, 允许存储供应商创建自定义存储插件, 无需将它们添加到kubernetes仓库.

## (7)备注
- https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md

# 二 重要接口:
## (1)概述:
- pkg/volume/plugins.go
- 每类插件都有个ProbeVolumePlugins方法, 在kubelet启动时候调用进行注册.

## (2)类型:
- **VolumePluginMgr**: tracks registered plugins.
- **VolumePlugin**: 被kubelet用来实例化和管理volume的volume plugin接口, 被各类volume实现.
- **PersistentVolumePlugin**: VolumePlugin接口的扩展, 增加了GetAccessModes方法.
- **DeviceMountableVolumePlugin** 
- **AttachableVolumePlugin**
- **Volume**: represents a directory used by pods or hosts on a node.
- **Attacher/Detacher**: attach/detach一个volume到node.
- **Mounter/Unmounter**: mount/unmount volume.
- **PersistentVolumePlugin**
- **VolumeHost**: an interface that plugins can use to access the kubelet.

## (3)VolumePluginMgr:
- 创建kubelet时通过NewInitializedVolumePluginMgr函数创建, 并传递给kubelet的volumeManager使用.
- InitPlugins: 调用VolumePlugin列表的Init初始化插件.
- Run
- 等等.

## (4)VolumePlugin
- Init: 初始化该插件, 
- GetPluginName: 获取插件名字.
- GetVolumeName
- CanSupport: 判断插件是否支持指定的volume specification.
- RequiresRemount
- ConstructVolumeSpec:
- NewMounter
- NewUnmounter
- SupportsMountOption: 是否支持mount选项.
- SupportsBulkVolumeVerification

## (5)PersistentVolumePlugin:
- VolumePlugin
- GetAccessModes

# 三 临时卷:
## (1)类型:
- emptyDir
- configMap
- downwardAPI
- secret
- CSI ephemeral volumes
- generic ephemeral volumes

## (2)备注:
- https://kubernetes.io/docs/concepts/storage/ephemeral-volumes
