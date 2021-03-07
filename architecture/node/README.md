# 一 概述:
## (1)概述:
- node是k8s中的工作节点, node可以是虚拟机或物理机.

## (2)node上组件:
- kubelet
- 容器运行时
- kube-proxy

## (3)和node交互的组件:
- node controller
- kubelet
- kubectl

## (4)影响pod情况:
- kubectl drain
- 资源紧张触发驱逐(kuebelt的evictionManager)
- node不健康(心跳丢失或者宕机等)触发的驱逐(node controller)
- node平滑shutdown: https://kubernetes.io/docs/concepts/architecture/nodes/#graceful-node-shutdown
- 探针

## (5)备注:
- https://kubernetes.io/docs/concepts/architecture/nodes/
- API: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.15/#node-v1-core
- 监控: https://kubernetes.io/docs/tasks/debug-application-cluster/monitor-node-health/
 
# 二 Node管理:
## (1)概述:
- 与pod和service等不同, node不是由kubernetes内部创建的, node是通过外部云服务商或者存在于你自己的物理机或虚拟机池.
- kubernetes创建一个node对象来表示node, 创建后会检查node是否合法.
- 管理员可以修改node属性, 设置label或者标记不可调度等, 标记为不可调度不会影响已存在的pod.
- **kubectl cordon nodename**将node标记为不可用, **kubectl uncordon**解除限制.

## (2)node注册方式:
- 自注册(常用): 通过kubelet --register-node设置为true来实现, kubelet会尝试向api server注册它自己.
- 手动注册: 将kubelet --register-node设置为false, 然后自己创建node对象. 

# 三 nodeSpec:
## (1)概述:
- configService
- podCIDR
- podCIDRS
- providerID: cloud provider分配的node ID.
- taints
- unschedulable: 控制node为不可调度, 默认是可调度.

# 四 node.status:
## (1)addresses:
- HostName: node的内核上报的hostname, 可被kubelet --hostname-override参数覆盖.
- ExternalIP: node的ip地址, 可被外部路由到(集群外).
- InternalIP: node的内部ip地址, 只能集群内访问.

## (2)conditions: 
- **OutOfDisk**: True表示没有足够的磁盘空间添加新的pod了, 否则为False.
- **Ready**: True表示node是健康的且准备好接收pod; False表示node是不健康且不能接收pod; Unknown表示node controller在过去node-monitor-grace-period(默认40s)没有接收到node的心跳.
- **MemoryPressure**: True表示内存紧张(pressure); False表示内存使用量低.
- **PIDPressure**: True表示进程紧张, 表示node上有太多进程.
- **DiskPressure**: True表示磁盘紧张, False表示磁盘容量低.
- **NetworkUnavailable**: True表示node的网络配置不正确; 反之False.

## (3)capacity:
- 描述node上可用的资源, 例如:CPU,memory和可被调度的最大pod数量.
- 容量(cpus数量和内存总量等)是node对象的一部分, 通常node注册自己并在创建node对象时候报告自己的容量; 若是手动注册, 则需要在添加node时候设置node的容量.
- kubernete调度器保证node上的pod有足够的资源, 它会检查容器请求的资源不超过node的容量, 但不包括不是kubelet启动的容器.
- 默认情况下, Pod可使用node上的所有资源, 但也需要为一些管理os和kubernete自己的系统daemon预留一些资源, kubelet通过**Node Allocatable**特性来为系统daemon预留资源.
- kubernetes建议集群管理员基于每个node的workload density来配置Node Allocatable.
- 备注: https://kubernetes.io/docs/tasks/administer-cluster/reserve-compute-resources/

## (4)allocatable:
- Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
- 这里是可分配单资源数量, 不是剩余资源数量, 剩余的可通过kubectl describe node查询.

## (5)config:
- Status of the config assigned to the node via the dynamic Kubelet config feature.

## (6)nodeInfo:
- 包含node的常规信息, 例如:内核版本,kubernerate版本,docker版本等.

## (7)存储相关:
- node上容器的image列表.
- images: 
- volumesAttached: attach到node的卷的列表.
- volumesInUse: attached列表中已在使用(被mount)的volume列表.

## (8)daemonEndpoints:
- Endpoints of daemons running on the Node.

## (9)images:
- 该node上容器镜像列表.

# 五 资源预留:
## (1)概述:
- 资源划分: Allocatable, kube-reserved, system-reserved, eviction-threshold.
 
## (2)Allocatable: 
- 可用于Pod的计算资源的总量, 调度器不能超过该总量, 当前支持cpu,mem和ephemeral-storage, 通过kubectl descirbe node nodename来查询.
- 当全部pod使用量超过Allocatable时, kubelet会evicting pods, 具体策略参考五的eviction policy.

## (3)kube-reserved:
- 用与kubernets系统daemon的资源预留, 例如:kubelet, container runtime和node problem detector等, 不会为以pod形式运行的系统daemon预留.
- 支持的资源类型: cpu, mem, ephemeral-storage和pid.
- kubelet flag: --kube-reserved和--kube-reserved-cgroup

## (4)system-reserved: 
- 为OS系统daemon预留资源, 例如:sshd等.
- kubelet flag: --system-reserved和--system-reserved-cgroup

## (5)eviction thresholds:
- node上的内存压力会导致OOM.
- 为了避免系统oom, kubelet提供了**out of resource**管理.
- eviction只支持mem和ephemeral-storage, 通过kubelet的--eviction-hard来设置.

# 六 心跳:
## (1)概述:
- node通过发送心跳来帮助判断node的可用性.

## (2)两种形式:
- 更新**NodeStatus**
- 更新Lease Object: Node在**kube-node-lease** namespace内有一个关联的Lease对象, Lease是一个轻量资源, 可以提高node的心跳上报性能.

## (3)更新NodeStatus:
- kubelet在**status改变时**或者**超过配置的时间间隔**后更新NodeStatus.

## (4)更新Lease对象:
- kubelet会间隔10s(默认更新间隔)更新它关联的lease对象, Lease对象的更新不依赖NodeStatus的更新.
- kubelet的nodeLeaseController.

## (5)备注:
- pkg/kubelet/nodelease
- https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/0009-node-heartbeat.md
