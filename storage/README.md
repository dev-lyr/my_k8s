# 一 概述:
## (1)背景:
- 容器中磁盘上文件是临时的, 当容器crash时, kubelet会重启它, 但是文件内容会丢失(新启动的容器是clean状态).
- Pod内的容器需要在容器间共享文件.
- Docker也有卷的概念, 但Docker中的卷管理松散, 只是磁盘上或其他容器内的一个目录, Docker虽然提供了卷driver, 但目前功能有限(目前官方只有local volume driver).
- K8s使用卷(volume)来解决上面的问题.

## (2)k8s的卷:
- 有显式的生命周期, 被定义为Pod的一部分, 和Pod有同样的生命周期, 当Pod启动是创建卷, 在Pod销毁时销毁.
- 当容器重启时, 卷内容不变, 新启动容器可以识别前容器写入卷的所有文件.
- 若Pod包含多个容器, 则这个卷可以被Pod内所有容器**共享使用**.
- K8s支持多种类型的卷, 且一个pod可以同时使用任意数量卷.
- 通过spec.volumes属性来指定Pod内容器可以使用的卷; 使用spec.containers.volumeMounts属性来指定容器的挂载点.

## (3)卷的类型:
- **emptyDir**: 用于存储临时数据的简单空目录.
- **hostPath**: 挂载host node上的一个文件或目录到Pod内.
- **configMap**: 提供一种将配置数据注入Pods的方式.
- **dowmwardAPI**: 使downward API数据对应用可见.
- **secret**: 用于向Pod传送敏感信息, 例如:密码.
- **persistentVolumeClaim**: 单独介绍.
- **gitRepo**: 已废弃.
- **CSI**: 重要, 后续单独介绍.
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

# 二 emptyDir
## (1)概述:
- emptyDir当Pod被分配到一个Node时创建, 初始为空.
- 和Pod有同样的生命周期, 当Pod被从node上删除时, emptyDir中的数据也被删除.
- emptyDir卷可以存储在后端node上的任何媒介上, 例如: disk,SSD或网络存储, 也可以设置emptyDir.midium为Memory告诉k8s挂载一个tmpfs使用.

## (2)EmptyDirVolumeSource:
- medium: 目录后端存储介质的类型, 可选: 空字符串(默认)或Memory, 空字符串表示使用node默认的介质.
- sizeLimit: 默认为nil, 无限制.

# 三 hostPath:
## (1)概述:
- hostPath卷挂载host node文件系统上的一个文件或者目录到pod中.

## (2)HostPathVolumeSource属性:
- path: 必选.
- type(可选): DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice, BlockDevice, 默认为空(表示在挂载卷前不会执行任何check).

## (3)使用场景

## (4)谨慎使用

# 四 downward API:
## (1)使用场景:
- 使用configMap和secret卷应用传递Pod调度, 运行前的数据是可行的, 但对于不能预先知道的数据, 比如: Pod的Ip,主机名或Pod自身的名字名称等, 则需要使用downward API来解决.
- downward API允许用户通过**环境变量**或**文件(downward API卷)**来将Pod和容器的元数据传递给它们内部运行的进程.
- 详细: https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/

## (2)downward API可传暴露递的数据:
- pod的名称,ip,所在namespace,运行节点的名称,所归属服务账号的名称.
- 每个容器请求的CPU和内存使用量, CPU和内存的限制.
- Pod标签和注解.
- 备注: Pod标签和注解只能通过downward API**卷**来暴露, 其它的可以使用环境变量或downward API卷来暴露.

## (3)DownwardAPIVolumeSource:
- defaultMode: 创建文件使用的mode bits, 默认为0644, 可选范围:0-0777.
- items(DownwardAPIVolumeFile数组): downward API volume文件的数组.

## (4)DownwardAPIVolumeFile:
- fieldRef: 选择pod的一个field, 当前支持: 注解,labels, name和namespace.
- mode: 该文件上使用的mode bits, 若不指定则使用volume上defaultMode.
- path: 需要创建文件的相对路径, 不能是绝对路径也不能包含..路径.
- resourceFiledRef(ResourceFiledSelector): 选择容器的一个资源, 当前只支持resource limit和requests.

## (5)备注:
- https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/
