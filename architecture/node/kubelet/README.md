# 一 概述:
## (1)功能:
- kubelet是每个node上运行的"node agent", kubelet将node注册到apiserver(使用hostname或一个覆盖hostname的flag或一个cloud provider的特定逻辑).
- 持续watchAPI服务器是否把Node分配给Pod, 然后通过运行时启动Pod容器, 随后持续健康运行中的容器向API服务器汇报它们状态,事件和资源消耗.
- kubelet运行容器的探针, 当探针失败时会重启容器; 当Pod删除时, kubelet会终止容器.
- Kubelet和PodSpec一起工作, 保证PodSpec描述的容器健康运行, kubelet不管理非kebernetes创建的容器.

## (2)PodSpec的来源:
- apiserver(主要).
- File: 通过命令行flag传递文件路径, 该路径下的文件会被周期性监控更新, 默认监控周期为20s.
- HTTP endpoint
- HTTP server

## (3)源码:
- kubernetes/cmd/kubelet
- kubernetes/pkg/kubelet/

## (4) kubelet结构体:
- 主要的kubelet实现.
- pkg/kubelet/kubelet.go
- NewMainKubelet方法: 创建一个Kubelet对象以及它需要的所有内部module.

# 二 Kubelet结构属性:
## (1)pod相关:
- podManager
- podWorkers
- evictionManager
- nodeLister
- serviceLister

## (2)Node相关:

## (3)容器相关

# 三 Kubelet结构方法
## (1)kubelet.go
## (2)kubelet_pod.go
## (3)kubelet_network.go
## (4)kubelet_volume.go
## (5)kubelet_resource.go
