# 一 概述:
## (1)概述:
- Pod是k8s中最基本的构建单元(最小和最简单的创建和部署单元).
- 一个Pod封装: 一个应用容器(或一些情况下多个容器),存储资源和一个唯一的网络IP和其它容器运行需要的选项.
- 一个Pod表示一个部署单元: k8s中一个应用的一个实例, 可能由一个容器或多个需共享资源的紧耦合的容器组成.
- Docker是k8s中pod最常用的容器运行时, 也支持其他容器运行时.
- 通常使用controller来管理pods, 很少直接操作pod.
- Pod使用一个执行pause命令的基础容器来保存所有命名空间, 其它用户自定义容器使用该基础容器的命名空间.

## (2)Controllers:
- Replication Controller: 不建议使用.
- ReplicaSet: 下一代Replication Controller.
- Deployments
- 等等.
- 备注: 推荐使用合适的controller来创建Pods, 而不是直接创建Pod, 因为当Node失败时, 单独创建的Pod不会被恢复.

## (3)隔离性:
- 同一Pod内容器默认共享network,UTS和IPC namespace, 可通过配置podspec.shareprocessnamespace来共享Pid命名空间.
- 由于共享相同的网络命名空间, 因此同一pod内容器共享网络接口和端口空间, 各个容器需要注意使用不同的端口.
- 容器的文件系统来自容器镜像, 因此默认情况下, 容器的文件系统与其它容器完全隔离, 但可以使用volume来共享文件目录.

## (4)PodTemplate:
- PodTemplate是被include到其它对象的pod规范, controller根据Pod Templates来制造真正的pod.

## (5)备注:
- https://kubernetes.io/docs/concepts/workloads/pods/pod-overview

# 二 PodSpec
## (1)概述:
- 描述期望的Pod的状态.

## (2)容器相关:
- containers(Container数组): 属于该Pod的容器.
- ephemeralContainers(EphemeralContainer数组).
- initContainers
- imagePullSecrets
- runtimeClassName: 引用一个RuntimeObject对象, 表示运行该Pod的运行时.
- restartPolicy: pod内所有容器的重启策略, 可选:Always(默认),Never和OnFailure.

## (3)host相关:
- hostAliases.
- hostIPC: 使用host的ipc namespace, 默认为false.
- hostNetwork: 使用host的network namespace, 默认为false.
- hostPID: 使用host的pid namespace, 默认为false.

## (4)调度相关:
- affinity
- nodeName
- nodeSelector
- tolerations
- topologySpreadConstraints
- schedulerName
- priority: Pod的优先级值, 许多系统组件依赖该值来发现pod的优先级.
- priorityClassName: Pod的优先级.
- preemptionPolicy: 抢占低优先级pod的策略, 可选: Never或PreemptionLowerPriority, 默认PreemptionLowerPriority.

## (5)存储相关:
- volumes(Volume数组): 属于pod的可以被挂载到容器里的volumes.
- 挂载是在container对象的volumeMounts属性表示.

## (6)网络相关:
- dnsConfig
- dnsPolicy
- enableServiceLinks: 表示服务的信息是否以环境变量的形式注入到Pod, 默认为true.
- readinessGates

## (7)安全相关:
- serviceAccountName: 运行该Pod所使用的ServiceAccount的名字.
- serviceAccount: 已废弃, 使用serviceAccountName.
- automountServiceAccountToken: 表示serviceAccount token是否被自动挂载.
- securityContext: pod级别的安全属性和通用的容器配置,容器也有securityContext,若存在相同的属性则容器内的会覆盖pod级别的.

# 三 podStatus
## (1)状态相关:
- conditions: pod当前服务状态.
- phase: pod生命周期所处阶段的高度简单的总结, 可选值: Pending,Running,Successed,Failed,Unkown.
- message
- reason: 解释pod为什么处于某种state, 例如:"Evicted".
- startTime: pod被kubelet认可的时间, 在kubelet pull pod中容器镜像之前.
- qosClass: Pod的QoS类型.

## (2)容器状态相关(ContainerStatu数组):
- containerStatuses:每个容器一项,每项是docker inspect的当前输出.
- ephemeralContainerStatuses: pod中运行的任意临时容器的状态.
- initContainerStatuses: 每个临时容器一项.

## (3)网络相关:
- podIP: 分配给Pod的IP地址, 至少在集群内路由,若未分配则为空.
- podIPs: 包含分配给Pod的IP地址, 若指定则第0项必须是PodIP的值,若未分配则为空.
- hostIP: pod分配的host的IP地址, 若未被调度则为空.

# 四 容器管理:
## (1)概述:
- Pod为其内部容器提供两种共享资源: **网络**和**存储**.
- **网络**: 每个Pod被分配一个唯一的IP地址, Pod内的容器共享网络空间, 包括IP地址和网络端口; Pod内的容器可以通过localhost来通信, 当Pod内容器与外界通信时, 需协调怎么使用共享的网络资源(例如:端口).
- **存储**: Pod可以指定一个共享存储volumes的集合, 所有Pod内的容器可以访问共享volumes, 从而实现容器共享数据.

## (2)两种主要用法:
- **一个pod一个容器**: 最常用的情况, 可认为pod是一个单独容器的封装, k8s管理pod而不是容器.
- **一个pod和多个需要一起工作的容器**: 一个pod可能封装一个由多个容器组成的应用, 这些容器紧密耦合并且需要共享一些资源, 例如: **sidecar**容器.

## (3)混合(composite)容器模式:
- **sidecar容器**: 用于扩展和加强主容器的功能,使用场景: 日志轮转和收集器,数据处理器,通信适配器等.
- **ambassador容器**: 是一种特定的sidecar, 隐藏了访问pod外部服务的复杂性并提供统一的接口, 用于proxy本地连接到外界. 
- **adapter容器**: 用于标准化和统一输出. 例如: 为了监控N个不通应用, 监控系统期望它监控的数据是一致的数据模型, 但是不同应用有自己export监控数据的方式, 因此可以使用adapter容器将不通应用的监控数据转换为一致统一的表示方式.

## (4)备注:
- https://www.usenix.org/system/files/conference/hotcloud16/hotcloud16_burns.pdf
- https://kubernetes.io/blog/2015/06/the-distributed-system-toolkit-patterns
- https://kubernetes.io/blog/2016/06/container-design-patterns.

# 五 Init容器:
## (1)概述:
- Init容器是一种在app容器之前运行的, 包含一些在app镜像中不包含的工具(utilities)或启动(setup)脚本.
- 一个Pod可以有一个或多个Init容器, 当Pod内的一个Init容器失败, 则k8s会重复重启该Pod直至Init容器成功, 除非Pod的restartPolicy设置为Nerver.
- 若在Pod内指定多个Init容器, 则这些容器在顺序执行, 同一时刻只能运行一个, 且必须运行成功才会运行下一个. 当所有Init容器都执行完成, K8s初始化Pod并运行应用容器.

## (2)Init容器和普通容器区别:
- Init容器通常运行至完成.
- 每个Init容器必须成功结束下一个才会开始.

## (3)资源相关:
- 所有init容器中最大的资源request和limit是**effective init request/limit**.
- pod的effective init request/limit是所有业务容器request/limit和init的effective request/limit的两者间的最大者.

## (4)使用场景:
- Init容器可以包含和运行由于安全原因不期望包含在应用容器镜像的工具.
- 可以包含用于setup的工具或自定义代码, 例如:当仅仅想使用sed, awk等工具时, 不需要FROM一个镜像.
- 应用镜像的builder和deployer可以独立工作, 不需要仅仅构建一个单独的app镜像.
- 使用Linux Namespace, 所以可以有不同于应用容器的文件系统, 可以用于访问一些app容器不能访问的私有数据.
- 在app容器启动前执行, app容器是并行运行的, 所以Init容器提供一种block或delay应用容器启动的方法, 直至满足指定preconditions.

# 六 Pause/Infra容器:
## (1)概述:
- 扮演pod内容器的"父容器", 两个主要功能: 共享namespace; 当pid空间共享开启时, 作为每个pod的PID 1并且回收僵尸进程.
- 又叫infra容器.

## (2)备注:
- https://www.ianlewis.org/en/almighty-pause-container

# 七 Pod Preset:
## (1)概述:
- Pro Preset是一种用来在Pod创建时inject特定信息的API资源.
- 使用Pro Preset允许Pod模板的作者不需要为每个Pod显式提供所有信息.
- K8s提供一个admission控制器(PodPreset), 当开启时,应用Pod Presets到进来的Pod创建请求.

## (2)工作方式:

## (3)备注:
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/service-catalog/pod-preset.md

# 八 静态pod:
## (1)概述:
- 静态pod被运行node上的kubelet直接管理,与其它workload(deployment等)管理的pod不同, kubelet watch每个静态pod并在crash时候重启.
- 静态容器通常与给定node上的kubelet绑定, kubelet会自动为每个静态pod创建一个在k8s ApiServer上的mirror pod, 因此通过API server可见但不能通过它修改.

## (2)创建方式:
- 基于文件系统的配置文件: kubelet的staticPodPath指定的目录, kubelet会定时轮询该目录,创建和删除文件中指定的静态Pods.
- 基于web的配置文件: kubelet的--manifest-url配置指定, kubelet定时refetch配置文件来执行apply.

## (3)相关行为:
- kubelet启动时会自动启动所有静态Pod.
- 删除API server中的mirror pod不会删除静态Pod.
- 判断是不是静态pod: pod存在kubernetes.io/config.source注解且注解的值不为api.

## (4)使用场景:
- kubeadm启动API Server和controller-manager等几个管控组件都是使用静态Pod.

## (5)备注:
- https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/

# 九 临时容器:
## (1)概述:
- 一种特殊类型的容器, 在一个**已存在**Pod中临时运行的容器, 用来完成一些用户触发的操作, 比如: 问题定位.

## (2)备注:
- https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers
