# 一 概述:
## (1)概述:
- 容器网络接口(CNI)

## (2)常用实现:
- flannel
- calico
- terway: aliyun

## (3)备注:
- https://github.com/containernetworking/cni
- https://github.com/containernetworking/plugins: 实现CNI规范的CNI插件.
- https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/
- https://github.com/containernetworking/cni/tree/master/cnitool
- awk cni: https://github.com/aws/amazon-vpc-cni-k8s

# 二 CNI规范:
## (1)概述:
- 每个CNI插件必须以可执行方式实现, 被容器管理系统(例如:k8s)调用.
- CNI插件负责向容器网络空间插入一个网络接口(例如: veth pair的一端), 并且可以在host上执行必要的修改(例如: attaching另外一端veth到bridge).
- CNI可以通过调用合适的**IPAM**插件来给该网络接口分配IP地址, 搭建与IP Address Management一致的路由信息.

## (2)CNI支持的操作:
- ADD: 添加容器到网络.
- DEL: 从网络中删除容器.
- CHECK: 检查容器的网络是否符合预期.
- VERSION: 查看verison.

## (3)IPAM插件:
- 被CNI插件调用, 用来生成接口IP/subnet, Gateway和Routes信息.
- IPAM插件可以通过dbcp,本地文件系统,网络配置文件中的ipam部分等方式来获得信息.

## (4)错误码：
- 1: 不兼容CNI版本.
- 2: 网络配置不支持的field.
- 3: 未知容器或不存在.
- 11: 重试.
- 备注: 错误码1-99不要使用, 除非在CNI规范中定义了用途.

## (5)备注:
- https://github.com/containernetworking/cni/blob/master/SPEC.md

# 三 libcni
## (1)概述:
- https://github.com/containernetworking/cni/tree/master/libcni

## (2)CNI接口:
- AddNetworkList
- CheckNetworkList
- DelNetworkList
- GetNetworkListCachedResult
- GetNetworkListCachedConfig
- AddNetwork
- CheckNetwork
- DelNetwork
- GetNetworkCachedResult
- GetNetworkCachedConfig
- ValidateNetworkList
- ValidateNetwork

## (3)CNIConfig:
- Path []string
- exec invoke.Exec
- cacheDir string
- 备注: 实现CNI接口.

# 四 网络配置:
## (1)概述:
- https://github.com/containernetworking/cni/blob/master/pkg/types/types.go

## (2)NetConf:
- Name
- Type
- Capabilities
- IPAM
- DNS
- RawPrevResult
- PreResult

## (3)DNS:
- NameServers
- Domain
- Search
- Options

## (4)Route
- Dst net.IPNet
- GW net.IP
