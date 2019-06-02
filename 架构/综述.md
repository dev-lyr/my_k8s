# 一 概述:
## (1)组件分类:
- Master Components
- Node Components
- Addons

## (2)运行:
- 控制组件和kube-proxy可直接部署在系统或者作为Pod运行, kubelet是唯一需要作为常规组件在系统上运行的组件(即不能容器化).

## (3)HA:
- 为了高可用, master组件可以有多个实例, 其中etcd和api服务器可以并行运行, 其它的则需要是一个实例起作用, 其它处于待命状态.
- 高可用: https://kubernetes.io/docs/setup/independent/high-availability

## (4)备注: 
- https://kubernetes.io/docs/concepts/overview/components/

# 二 Master组件:
## (1)概述:
- Master组件提供集群的控制面板(control plane).
- Master组件在集群中做全局决定(比如:调度), 并检测和响应集群事件.
- Master组件可运行在集群内的任何机器上, 然而为了简单, 启动脚本通常在同一个非运行用户容器的机器上启动master组件.

## (2)kube-apiserver:
- 暴露Kubernetes API, 是Kubernetes的**控制面板**的前端.
- k8s系统组件间只能通过API服务器通信, 它们之间不会直接通信.
- API服务器是和etcd通信的唯一组件, 其它组件都通过api服务器来修改集群状态.
- 支持水平扩展(scale horizontally).

## (3)etcd:
- 一致和高可用的键值存储, kubernetes用来作为所有集群数据的辅助存储(backing store).
- https://github.com/etcd-io/etcd/blob/master/Documentation/docs.md

## (4)kube-scheduler:
- kubernetes的**调度器**, 观察新创建的但未分配node的pods, 并选择一个node去执行它们.
- 调度决定考虑的因素包括: individual and collective resource requirements, hardware/software/policy constraints, affinity and anti-affinity specifications, data locality, inter-workload interference and deadlines.

## (5)kube-controller-manager:
- 运行在master上负责运行各种控制器的组件, 控制器负责确保系统中资源状态向API服务器定义的期望状态收敛.
- **Node Controller**: Responsible for noticing and responding when nodes go down.
- **Replication Controller**: Responsible for maintaining the correct number of pods for every replication controller object in the system.
- **Endpoints Controller**: Populates the Endpoints object (that is, joins Services & Pods).
- **Service Account & Token Controllers**: Create default accounts and API access tokens for new namespaces.

## (6)cloud-controller-manager:
- cloud-controller-manager runs controllers that interact with the underlying cloud providers.

# 三 Node组件:
## (1)Node介绍:
- Node组件运行在每个Node上, 维护运行中的pods和提供kebernetes运行时环境.
- Node是kubernetes中的worker机器, node可以是vm或物理机.
- 每个node包含运行pod和被master组件管理的必须的服务.

## (2)kubelet:
- An agent that runs on each node in the cluster. It makes sure that containers are running in a pod.
- The kubelet takes a set of PodSpecs that are provided through various mechanisms and ensures that the containers described in those PodSpecs are running and healthy. The kubelet doesn’t manage containers which were not created by Kubernetes.

## (3)kube-proxy:
- kube-proxy enables the Kubernetes service abstraction by maintaining network rules on the host and performing connection forwarding.

## (4)Container Runtime:
- The container runtime is the software that is responsible for running containers. 
- Kubernetes supports several runtimes: **Docker**, **rkt**, **runc** and any OCI runtime-spec implementation.

# 四 扩展组件(Addons):
## (1)概述:
- Addons是一些实现集群特性的**Pods**和**Services**, 这些Pod被Deployments, ReplicationControllers等管理.
- Addons用来扩展K8s的功能.
- 命名空间化的Addon对象在**kube-system**空间中创建.
- 备注: https://kubernetes.io/docs/concepts/cluster-administration/addons/

## (2)网络和网络策略.

## (3)服务发现:
- CoreDNS

## (4)可视化和控制:
- Dashboard
- Weave Scope

# 五 Master-Node通信:
## (1)Cluster到Master:
- 所有Cluster到Master的通信路径都终止在**apiserver**(其它master组件设计为不暴露远程服务).
- 在典型部署中, apiserver配置在HTTPs端口443监听远程服务, 有一种或多种client authentication方式.

## (2)Master到Cluster:
- Master(apiserver)到cluster有两种主要通信路径: **apiserver到运行在集群中每个节点的kubelet进程**和**通过apiserver的proxy功能, apiserver到任意node,pod或service**.
- apiserver到kubelet的**主要用途**: 获得pods的日志; attaching(通过kubectl)到运行中的pods; 提供kubelet的端口转发功能.
