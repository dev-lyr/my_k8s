# 一 概述:
## (1)概述:
- k8s集群中每个node上都运行一个**kube-proxy**, 它负责为服务实现一种**虚拟IP(即cluster IP)**.
- Proxy-mode: **userspace**, **iptables**, **ipvs**.
- 备注: iptables -t nat -L查询规则.

## (2)userspace模式:
- kube-proxy观察k8s master对**Service**和**Endpoints**对象的创建和删除.
- 在本地节点上, kube-proxy会为每个Service打开一个端口(随机选择), 任何到该代理端口的连接都会被proxy搭配服务的后端Pods(由EndPoints上报).
- 选择哪个Pod是基于Service的SessionAffinity.

## (3)iptables模式:
- kube-proxy观察k8s master对**Service**和**Endpoints**对象的创建和删除.
- 针对每个服务, 安装iptable规则用来capture到服务ClusterIP和Port的流量, 将这些流量redirect到Service后端集合.

## (4)ipvs模式:
- kube-proxy观察**Service**和**Endpoints**, 调用**netlink**接口来创建ipvs规则并周期性向Services和Endpoints同步.
- 当Service被访问时, 流量被重定向到后端Pods.
- 与iptable类似, ipvs基于netfilter hook函数, 但是要hash表作为底层数据结构并且工作在内核空间, 即ipvs重定向速度更快, 在同步proxy rules时性能更好, ipvs也提供更多的load balance方法.

## (5)选择:
- usernamespace模式性能太差已废弃.
- iptables适用于中小规模集群.
- ipvs模式适用于大规模集群.