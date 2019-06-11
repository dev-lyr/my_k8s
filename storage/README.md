# 一 概述:
## (1)背景:
- 容器中磁盘上文件是临时的, 当容器crash时, kublet会重启它, 但是文件内容会丢失(新启动的容器是clean状态).
- Pod内的容器需要在容器间共享文件.
- Docker也有卷的概念, 但Docker中的卷管理松散, 只是磁盘上或其他容器内的一个目录, Docker虽然提供了卷driver, 但目前功能有限.
- K8s使用卷(volume)来解决上面的问题.

## (2)k8s的卷:
- 有显式的生命周期, 被定义为Pod的一部分, 和Pod有同样的生命周期, 当Pod启动是创建卷, 在Pod销毁时销毁.
- 当容器重启时, 卷内容不变, 新启动容器可以识别前容器写入卷的所有文件.
- 若Pod包含多个容器, 则这个卷可以被Pod内所有容器共享使用.
- K8s支持多种类型的卷, 且一个pod可以同时使用任意数量卷.
- 通过spec.volumes属性来指定Pod内容器可以使用的卷; 使用spec.containers.volumeMounts属性来指定容器的挂载点.

## (3)卷的类型:
- **emptyDir**: 用于存储临时数据的简单空目录.
- **hostPath**: 挂载host node上的一个文件或目录到Pod内.
- **configMap**: 提供一种将配置数据注入Pods的方式.
- **dowmwardAPI**: 使downward API数据对应用可见.
- **secret**: 用于向Pod传送敏感信息, 例如:密码.
- **gitRepo**: 已废弃.
- **CSI**: 重要, 后续单独介绍.
- **nfs**: 挂载到Pod中的NFS共享卷.
- **其它类型网络存储**: glusterfs, iscsi, cephfs等.
- **云厂商提供**: awsElasticBlockStore, gcePersistentDisk, azureDisk.

# 二 emptyDir
## (1)概述:
- emptyDir当Pod被分配到一个Node时创建, 初始为空.
- 和Pod有同样的生命周期, 当Pod被从node上删除时, emptyDir中的数据也被删除.
- emptyDir卷可以存储在后端node上的任何媒介上, 例如: disk,SSD或网络存储, 也可以设置emptyDir.midium为Memory告诉k8s挂载一个tmpfs使用.

# 三 configMap:
## (1)概述:
- configMap资源提供一种向Pod内注入配置数据的方式.
- 数据存储在ConfigMap对象中, 可通过configMap类型卷来引用, 被运行在Pod内的容器化应用消费.

# 四 downward API:
## (1)使用场景:
- 使用configMap和secret卷应用传递Pod调度,运行前的数据是可行的, 但对于不能预先知道的数据, 比如: Pod的Ip,主机名或Pod自身的名字名称等, 则需要使用downward API来解决.
- downward API允许用户通过环境变量或文件(downward API卷)来将Pod和容器的元数据传递给它们内部运行的进程.
- 详细: https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/

## (2)downward API可传暴露递的数据:
- pod的名称,ip,所在namespace,运行节点的名称,所归属服务账号的名称.
- 每个容器请求的CPU和内存使用量, CPU和内存的限制.
- Pod标签和注解
- 备注: Pod标签和注解只能通过downward API卷来暴露, 其它的可以使用环境变量或downward API卷来暴露.

# 五 secret:
## (1)概述:
- 用于向Pod传送敏感数据, 例如:密码等.
- secret后端是tmpfs, 不会被写入non-volatile存储.
- 详细: https://kubernetes.io/docs/concepts/configuration/secret/
