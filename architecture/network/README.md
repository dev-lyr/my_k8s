# 一 概述:
## (1)分类:
- 高度耦合的容器间(pod内容器): 通过pod和localhost来通信.
- pod-pod间: CNI插件等.
- Pod和Service间: 通过kube-porxy和dns等.
- 外界和Service间: Ingress和LoadBalance,NodePort服务.

## (2)备注:
- https://kubernetes.io/docs/concepts/cluster-administration/networking/
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/network/networking.md

# 二 kubernetes网络模型:
## (1)概述:
- 每个Pod有自己的IP地址. 

## (2)kubernetes对网络实现的要求:
- 一个node上的pods可以与所有node上的pods进行通信(不使用nat).
- node上的agent(例如:kubelet等)可以与该node上的所有pods通信.

## (3)实现:
- flannel
- cilium
- calico
- linux bridge
- OpenVSwitch
- 等等.

# 三 proxy类型:
## (1)概述:
- https://kubernetes.io/docs/concepts/cluster-administration/proxies/

# 四 插件:
## (1)类型:
- CNI插件: 遵守CNI规范.
- kubenet插件: 使用**bridge**实现**cbr0**和**host-local** CNI插件, 一个简单基本的网络插件, 只适用于linux.

## (2)安装:
- kubelet有单独的默认网络插件, 且默认网络插件全集群通用.
- kubelet在启动时候探测网络插件(参数**cni-bin-dir**指定的目录下), 并在pod生命周期内(只适用于docker)合适的时间执行选择的插件.
- kubelet的**network-plugin**参数: 指定使用的网络插件, 名字和cni-bin-dir指定的目录下插件的名字, 针对CNI插件名字为**cni**.

## (3)网络插件要求:
- 除了提供网络插件接口来配置和清除pod网络外, 查询要需要对kube-proxy进行特定支持, 例如: iptables proxy依赖iptables,插件需要确保容器流量对iptables可用.
- 默认情况下, 若没指定kubelet网络插件, 则**noop**插件被使用, 它将**net/bridge/bridge-nf-call-iptables=1**来确保iptables proxy可以正常使用.

## (4)cni插件:
- kubelet通过**--network-plugin=cni**来表明使用cni插件.
- kubelet从**--cni-conf-dir**指定的目录(默认为/etc/cni/net.d)来读取一个文件, 并通过该文件使用CNI规范来创建每个pod的网络.
- 若配置文件目录下有多个CNI配置文件, 则按字典顺序排序选择第一个.
- 配置文件中引用的插件必须在**--cni-bin-dir**目录(默认为/opt/cni/bin)存在.

## (5)kubenet插件:
- kubernetes/pkg/kubelet/dockershim/network/kubenet

## (6)网络插件种类:
- --cni-bin-dir目录.
- https://kubernetes.io/docs/concepts/cluster-administration/addons/
- https://github.com/dev-lyr/my_docker/blob/master/network/README.md

## (7)备注:
- https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/
- https://github.com/kubernetes/kubernetes/blob/v1.14.0/pkg/kubelet/dockershim/network/plugins.go
- https://chrislovecnm.com/kubernetes/cni/choosing-a-cni-provider/
- kubernetes/pkg/kubelet/dockershim/network
