# 一 网络插件:
## (1)类型:
- CNI插件: 遵守CNI规范.
- Kubenet插件: 使用**bridge**实现**cbr0**和**host-local** CNI插件, 一个简单基本的网络插件, 只适用于linux.

## (2)安装:
- kubelet有单独的默认网络插件, 且默认网络插件全集群通用.
- kubelet在启动时候探测网络插件(参数**cni-bin-dir**指定的目录下), 并在pod生命周期内(只适用于docker)合适的时间执行选择的插件.
- kubelet的**network-plugin**参数: 指定使用的网络插件, 名字和cni-bin-dir指定的目录下插件的名字, 针对CNI插件名字为**cni**.

## (3)cni插件:
- kubelet通过**--network-plugin=cni**来表明使用cni插件.
- kubecet从**--cni-conf-dir**指定的目录(默认为/etc/cni/net.d)来读取一个文件, 并通过该文件使用CNI规范来创建每个pod的网络.
- 若配置文件目录下有多个CNI配置文件, 则按字典顺序排序选择第一个.

## (4)kubenet插件

## (5)网络插件种类:
- --cni-bin-dir目录.
- https://kubernetes.io/docs/concepts/cluster-administration/addons/
- https://github.com/dev-lyr/my_docker/blob/master/network/README.md

## (6)备注:
- https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/
- https://github.com/kubernetes/kubernetes/blob/v1.14.0/pkg/kubelet/dockershim/network/plugins.go
- https://chrislovecnm.com/kubernetes/cni/choosing-a-cni-provider/

# 二 NetworkPolicy:
## (1)概述:
- 网络策略是一个specification, 用来指定pods组间以及与其它网络endpoints间如何通信.
- **NetworkPolicy**使用lables来选择pod, 并且为被选择的pods定义rule(指定哪些流量可以到pods).
- 网络策略是网络插件实现的, 所以必须使用一个支持**NetworkPolicy**的网络方案.

## (2)isolated和non-isolated Pods:
- 默认情况下, Pod是非孤立的, 可以接收任意源的流量.
- 当Pod被同一namespace下的NetworkPolicy选择时, 就变为独立的, pod会拒绝点NetworkPolicy不允许的连接.

## (3)NetworkPolicy资源:
- 参考: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#networkpolicy-v1-networking-k8s-io

## (4)备注:
- https://kubernetes.io/docs/concepts/services-networking/network-policies/
