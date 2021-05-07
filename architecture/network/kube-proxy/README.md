# 一 概述:
## (1)概述:
- k8s集群中**每个node**上都运行一个**kube-proxy**, 它负责为服务实现一种**虚拟IP(即cluster IP)**.
- Proxy-mode: **userspace**, **iptables**, **ipvs**.

## (2)userspace模式:
- kube-proxy**Service**和**Endpoints**对象的创建和删除.
- 在本地节点上, kube-proxy会为每个Service打开一个端口(随机选择), 任何到该代理端口的连接都会被proxy搭配服务的后端Pods(由EndPoints上报).
- 选择哪个Pod是基于Service的SessionAffinity.
- 路径: 客户端->iptables->kube-proxy->pod, kube-proxy配置iptables并通过kube-proxy来重定向.

## (3)iptables模式:
- kube-proxy观察**Service**和**Endpoints**对象的创建和删除.
- 针对每个服务, 安装iptable规则用来capture到服务ClusterIP和Port的流量, 将这些流量redirect到Service后端集合.
- 路径: 客户端->iptables->pod, kube-proxy配置iptables, 直接通过pod重定向.

## (4)ipvs模式:
- kube-proxy观察**Service**和**Endpoints**, 调用**netlink**接口来创建ipvs规则并周期性向Services和Endpoints同步.
- 当Service被访问时, 流量被重定向到后端Pods.
- 与iptable类似, ipvs基于netfilter hook函数, 但是要hash表作为底层数据结构并且工作在内核空间, 即ipvs重定向速度更快, 在同步proxy rules时性能更好, ipvs也提供更多的load balance方法.

## (5)选择:
- usernamespace模式性能太差已废弃.
- iptables适用于中小规模集群.
- ipvs模式适用于大规模集群.

## (6)备注:
- iptables -t nat -L查询规则.

# 二 kube-proxy:
## (1)概述:
- https://kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/

## (2)常用选项:
- --proxy-mode: userspace,iptables或ipvs, 当前默认iptables.
- --bind-addresss: proxy server的服务IP,默认0.0.0.0.
- --cleanup: 若为true则清理iptables和ipvs规则并退出.
- --cluster-cidr
- --config: 配置文件的路径.
- --config-sync-period
- --healthz-bind-address: 默认0.0.0.0:10256.
- --metrics-bind-address: 默认127.0.0.1:10249.
- --hostname-override
- --kube-api-burst

## (3)iptables相关

## (4)ipvs相关

## (5)conntrack相关

## (6)备注:
- cmd/kube-proxy
- pkg/proxy

# 三 ProxyServer:
## (1)概述:
- 表示启动proxyServer所需要的所有参数.
- 相关: Options表示创建和运行proxyServer所需的所有信息.

## (2)方法:
- Run: 运行ProxyServer.
- CleanupAndExit: 删除iptables规则并退出.

## (3)属性:
- Proxier(proxy.Provider): 代理的实现, 例如: iptables,ipvs,userspace. 

## (4)Run

# 三 Config:
## (1)概述:
- xxConfig对资源xx的informer的eventHandler进行封装.

## (2)EndpointsConfig:
- 属性: eventHandler []EndpointsHandler
- 方法: handleAddEndpoints, handleDeleteEndpoints, handleUpdateEndpoints, Run, RegisterEventHandler.
- EndpointsHandler接口: OnEndpointsAdd,OnEndpointsDelete,OnEndpointsUpdate,OnEndpointsSynced(只调用一次).

## (3)ServiceConfig:
- eventHandler []ServiceHandler
- 方法: handleAddService, handleDeleteService, handleUpdateService, Run, RegisterEventHandler.

## (4)NodeConfig:
- eventHandler []NodeHandler
- 方法: handleAddNode, handleDeleteNode, handleUpdateNode, Run, RegisterEventHandler.

## (5)EndpointSliceConfig

# 四 types.go:
## (1)Provider:
- config.EndpointsHandler
- config.EndpointSliceHandler
- config.ServiceHandler
- config.NodeHandler
- Sync: synchronizes the Provider's current state to proxy rules.
- SyncLoop: 执行periodic work.
