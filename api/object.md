# 一 概述:
## (1)概述:
- K8s对象是K8s系统中的持久化实体(persistent entities), k8s用这些实体来表示集群的状态. 例如: 描述在运行中的容器化应用是什么, 这些应用使用的资源, 以及容器的行为策略(重启策略, 升级和容错)等.
- 对象的增删改查通过**k8s API**完成, kubectl底层也是通过调用API.
- 可通过**kubectl api-resources**查询所有资源类型.

## (2)基本对象:
- Pod
- Service
- Volume
- Namespace
- 等等

## (3)Controllers:
- ReplicaSet
- ReplicationController
- Deployment
- StatefulSet
- DaemonSet
- Job
- 备注: Controller基于基本对象创建, 提供额外的功能性和遍历的特性.

## (4)yaml文件属性:
- **apiVersion(必选)**: 用来创建对象的k8s API版本.
- **kind(必选)**: 需要创建的对象类型, 例如: Pod, Service等.
- **metadata(必选)**: 用来唯一标识一个对象, 包括**name**, **UID**和**可选的namespace**.
- **spec**: 每个k8s对象的格式不同, 包含特定对象的嵌套属性, 参考: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/.

## (5)备注:
- https://kubernetes.io/docs/concepts/overview/working-with-objects/kubernetes-objects/

# 二 对象spec和status:
## (1)概述:
- 每个K8s对象包含两个内置对象属性: **对象spec**和**对象status**.
- **对象spec**: 创建对象时必须提供spec来描述期望的对象状态.
- **对象status**: 描述对象的实际状态, 由k8s系统提供和更新.
- 在任意时间, k8s控制面板主动操作对象的实际状态使其到达你提供的期望状态.

## (2)描述对象:
- 创建对象(通过API或kubectl)需提供对象spec来描述期望的状态, 就像其他基本信息一样(例如:name).
- 使用API是通过在JSON请求体来包含spec信息, 更常用是为kubectl提供一个**yaml**文件来创建, kubectl将信息转换为JSON来调用API.

# 三 Names和UID:
## (1)概述:
- Kubernetes REST API中的所有对象都通过一个**Name**和**UID**唯一标识.
- 针对不唯一的用户提供的属性, k8s可通过**label**和**annotation**.

## (2)Names:
- **用户侧**提供一个引用到资源URL中对象的字符串.
- 同一时刻, 一个给定类型资源的对象的Name需唯一, 若删除该对象, 则可以使用同样名字创建新的对象.
- **格式**: 按照惯例, name最大大小为253字符, 并且由小写字母,-和.组成, 特定资源有自己制定的限制.

## (3)UID:
- **Kubernetes系统**生成的唯一标识对象的字符串.
- Kubernetes集群生命周期创建的每个对象都有一个不同的UID.

# 四 Namespaces:
## (1)概述:
- Kubernetes支持把一个后端物理集群分为多个虚拟集群, 这些虚拟集群成为namespaces.
- Namespace提供一个names的范围, 在同一namespace内, 资源的name需是唯一的.
- Namespace是一种在多个用户间划分集群资源的方式(通过资源quota).
- 在未来版本中, 同一namespace的对象默认会有同样的访问控制策略.

## (2)三个初始namespace:
- **default**: 没指定namespace的对象的默认namespace.
- **kube-system**: 所有被kubernetes系统创建的对象的namespace.
- **kube-public**: 该namespace是自动创建的, 可以被所有用户访问(包括未授权), 该namespace通常为集群使用预留.

## (3)不是所有对象都在namespace:
- 大部分kubernetes资源(例如:pod,services,replication controllers和其它)都在namespace内, 但是namespace资源本身不在namespace内, 还有一些低层次资源也不再namespace内(例如:nodes, persistentVolumes和其它).
- 查询方法: `kubectl api-resources --namespaced=true/false`.

## (4)使用场景:
- 当需要跨多个团队或项目的用户使用环境时, 少量用户时不需要创建和考虑namespace.

## (5)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/architecture/namespaces.md

# 五 Labels
## (1)概述:
- **标签**是attach到对象的键值对, 用来给对象指定一些对用户有意义和相关的属性.
- 标签用来组织和选择对象的子集, 标签可以在创建时添加, 随后添加并且可以在任意时间修改.
- 通过**标签选择器**, 用户可以指定对象的集合, 标签选择器是k8s中核心的分组原语.
- 可使用,将多个lable串联起来进行选择, 逗号代表and关系.
- 格式: `metadata:{labels:{"k1":"v1", "k2":"v2"}}`, 给定对象的key是唯一的.

## (2)选择器类型:
- **equality-based**: 支持操作=,==和!=.
- **set-based**: 支持操作in,notin和exists.

## (3)常用labels:
- "release" : "stable", "release" : "canary"
- "environment" : "dev", "environment" : "qa", "environment" : "production"
- "tier" : "frontend", "tier" : "backend", "tier" : "cache"
- "partition" : "customerA", "partition" : "customerB"
- "track" : "daily", "track" : "weekly"
- 备注: 可以自定义.

## (4)使用:
- 命令: `kubectl get pods -l environment=production,tier=frontend`.
- API: `?labelSelector=environment%3Dproduction,tier%3Dfrontend`.

# 六 Annotations:
## (1)概述:
- 格式: `metadata:{annotations:{"k1":"v1", "k2":"v2"}}`

## (2)注解和标签对比:
- 可通过标签(label)和注解(annotation)attach metadata到k8s对象, 标签用来选择对象和找到满足特定条件的对象集合.
- 注解与标签不同, 不是用来标记和选择资源, 注解中的metadata可以大或小, 结构化或非结构化, 注解中可以包含在标签中不允许的字符.

# 七 Field Selectors:
## (1)概述:
- **属性选择器**可以通过一个或多个资源属性来选择k8s资源.
- 不同资源类型支持的属性不同, 所有资源支持metadata.name和metadata.namespace, 不支持的属性会报错.
- 支持操作: =,==和!=, 其中=和==一样.
- 可通过,分隔将多个属性选择器串联起来.
- 例如: `kubectl get pods --field-selector status.phase=Running`

## (2)备注:
- https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors/
