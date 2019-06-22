# 一 概述:
## (1)概述:
- Pod是不稳定的, 例如ReplicaSets可以动态的创建和销毁pod(例如:扩容或缩容时), 虽然每个pod可以有自己的IP地址, 但是这些IP地址是不稳定的.
- **服务**: 定义了一个逻辑的Pod集合和一个访问它们的策略, Service通常通过**标签选择器**来选择Pod集.
- Service解耦了前端用户和后端服务, 当后端服务pod集合变化时, 前端用户不需要感知.
- 对于非Kubernetes-native应用, kubernetes提供一个virtual-IP-based网桥来把服务redirect到后端Pods.

## (2)Endpoints资源:
- Service不直接和Pods直接相连, 有一种资源介于两者之间, 即**Endpoint**.
- 针对kubernetes-native应用, kubernetes提供一个简单的**Endpoints** API, 当Service中Pod变化时被更新. 

## (3)服务创建:
- 服务是一个REST对象, 和其它REST对象(例如Pod)一样, 通过将服务的定义POST给apiserver来创建服务实例.
- 服务被分配一个IP地址(有时称为**cluster IP**), 用来进行服务代理.
- 服务的selector会被持续计算且结果会被POST给一个**Endpoints**对象;当服务没有selector时不会创建Endpoints对象(可以自己创建一个endpoint,并将服务映射过去).
- 服务可以将incoming端口映射到一个**targetPort**, 默认情况下**targetPort**和**port**一样.
- 支持协议: TCP(默认), UDP, HTTP等, 只能给一个服务设置一个port和protocol.

## (4)属性:
- apiVersion
- kind
- metadata
- spec
- status: 只读,最近的服务的状态,由系统来填充.

## (5)备注:
- 连接集群外部的服务方法: 自定义Endpoint或创建ExternalName类型服务.

# 二 spec属性:
## (1)clusterIP:
- 服务的IP地址,通常由master随机分配, 若手动指定IP则需保证没有被其它使用, 该属性不能通过update来改变.
- 合法值: **None**, 空字符串("")或一个合法ip地址.
- None: 针对不需要使用代理的headless服务.

## (2)externalIPs:
- a list of IP addresses for which nodes in the cluster will also accept traffic for this service.
- These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP.
- A common example is external load-balancers that are not part of the Kubernetes system.

## (3)externalName:
- externalName is the external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.

## (4)externalTrafficPolicy

## (5)healthCheckNodePort

## (6)ports:
- 服务暴露的端口列表.

## (7)selector:
- 将服务的流量路由到和选择器标签key和值match的pod.
- 若为空或者不指定, 则k8s会认为服务有有一个外部进程管理它的endpoints, 此时不会做修改.

## (8)负载平衡相关:
- loadBalancerIP
- loadBalancerSourceRanges

## (9)session相关:
- sessionAffinity
- sessionAffinityConfig

## (10)type:
- ClusterIP: 默认,以集群内部IP的形式暴露服务, 因此服务只在集群内部可访问.
- ExternalName: 将服务映射到**externalName**属性的内容, 可用于为外部服务创建别名等.
- NodePort: 集群中每个节点上都打开一个端口, 并将端口上接收到的流量重定向到后端服务.
- LoadBalancer: 是NodePort的一种扩展, 使得服务可以通过专门的负载平衡器来访问, 由k8s运行的云基础设施提供, 客户端通过负载均衡器的IP地址来连接到服务, 负载均衡器有一个公开的访问IP地址.

# 三 服务代理(proxy):
## (1)概述:
- k8s集群中每个node上都运行一个**kube-proxy**, 它负责为服务实现一种虚拟IP(virtual IP).
- 在k8s 1.0服务是四层构造(TCP/UDP over IP), proxy是完全在用户空间(userspace).
- 在k8s 1.1中, **Ingress**API被加入来表示七层(HTTP)服务, iptables proxy被加入, 并且作为默认的操作模式(从1.2起).
- 在k8s 1.8中, ipvs代理被加入.
- Proxy-mode: **userspace**, **iptables**, **ipvs**.

## (2)userspace模式:
- kube-proxy观察k8s master对**Service**和**Endpoints**对象的创建和删除.
- 在本地节点上, kube-proxy会为每个Service打开一个端口(随机选择), 任何到该代理端口的连接都会被proxy搭配服务的后端Pods(由EndPoints上报).
- 选择哪个Pod是基于Service的SessionAffinity.

## (3)iptables模式:
- kube-proxy观察k8s master对**Service**和**Endpoints**对象的创建和删除.
- 针对每个服务, 安装iptable规则用来capture到服务ClusterIP和Port的流量, 将这些流量redirect到Service后端集合.

## (4)ipvs模式:
- kube-proxy观察**Service**和**Endpoints**, 调用**netlink**接口来创建ipvs规则并周期性向Services和Endpoints同步.
-当Service被访问时, 流量被重定向到后端Pods.
- 与iptable类似, ipvs基于netfilter hook函数, 但是要hash表作为底层数据结构并且工作在内核空间, 即ipvs重定向速度更快, 在同步proxy rules时性能更好, ipvs也提供更多的load balance方法.

# 四 服务发现(discovery):
## (1)概述:
- K8s支持两种主要类型的服务发现: **环境变量**和**DNS**(推荐).

## (2)环境变量:
- 在pod开始运行时, k8s会初始化一系列环境变量指向**已存在的服务**.

## (3)DNS:
- DNS服务器watch创建Service的k8s API, 并且创建为每个服务创建一个DNS记录集.
- 若集群中开启DNS, 则所有Pod可以自动做Service的命名解析.
- DNS server是唯一的一种访问**ExternalName**类型服务的方式.
- 例如: 在k8s中的my-ns命名空间中有一个my-service服务, 则一个DNS记录my-service.my-ns会被创建, 在my-ns命名空间中的Pod可以通过对my-service的简单名字查询来找到服务; 其它命名空间中的Pod则需要使用my-service.my-ns来查询; 名字查询的结果是**cluster IP**.

## (4)备注:
- 参考: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
- CoreDNS和kube-dns.

# 五 Headless服务:
## (1)概述:
- 有时不需要load-balancing和唯一服务IP的场景, 可以通过指定服务的.spec.clusterIP为None来创建"headless"服务.
- Headless服务减少和K8s系统的耦合, 允许开发者自由选择服务发现机制.
- 针对Headless服务, clusterIP没有分配, 因此kube-proxy不会处理这些服务, 且没有负载平衡和proxy, 怎么动态配置DNS依赖服务是否定义selector.

## (2)使用selectors:
- 针对定义了selectors的headless服务, endpoint控制器会创建Endpoints记录, 并且会修改DNS配置来返回直接指向服务后端Pods的A记录.

## (3)不使用selectors:
- CNAME records for ExternalName-type services.
- A records for any Endpoints that share a name with the service, for all other types.

# 六 发布服务:
## (1)概述:
- 需要把服务暴露到一个外部(集群外部)ip地址.
- k8s使用**ServiceTypes**来指定需要的服务类型.

## (2)服务类型：
- ClusterIP: 默认,以集群内部IP的形式暴露服务, 因此服务只在集群内部可访问.
- ExternalName: 将服务映射到**externalName**属性的内容, 可用于为外部服务创建别名等.
- NodePort: 集群中每个节点上都打开一个端口, 并将端口上接收到的流量重定向到后端服务.
- LoadBalancer: 是NodePort的一种扩展, 使得服务可以通过专门的负载平衡器来访问, 由k8s运行的云基础设施提供, 客户端通过负载均衡器的IP地址来连接到服务, 负载均衡器有一个公开的访问IP地址.

## (3)ClusterIP和NodePort区别:
- ClusterIP只能集群内部访问.
- NodePort不仅可以集群内部IP访问, 也可以通过节点上任何的IP和预留端口来访问.

## (4)Ingress优点:
- 每个LoadBalance服务都需要有自己的负载均衡器和公网Ip地址.
- Ingress只需提供一个公网IP地址就可以为多个服务提供访问, 当客户请求到Ingress时会根据路径名将请求转发到对应服务.
- Ingress是基于HTTP层操作, 提供了其它服务类型不提供的功能, 例如基于cookie的会话黏贴(affinity).

## (5)备注:
- 相关:ingress.

# 七 服务和Pod的DNS:
## (1)概述:
- Kubernetes DNS schedules a DNS Pod and Service on the cluster, and configures the kubelets to tell individual containers to use the DNS Service’s IP to resolve DNS names.

## (2)服务

## (3)Pod
- DNS策略支持Pod级别的配置, Pod可通过Pod Spec的**dnsPolicy**属性来配置DNS策略.
- Pod的DNS Config为用户提供更多的更多的DNS配置, 通过**dnsConfig**配置.
- 当dnsPolicy设置为None时, 则dnsConfig必须指定.

## (4)Pod DNS策略:
- Default: Pod继承Pod运行的node上的名字解析配置.
- ClusterFirst
- ClusterFirstWithHostNet
- None

## (5)Pod DNS配置

## (6)备注:
- https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
- https://github.com/kubernetes/dns/blob/master/docs/specification.md

# 八 向pod(/etc/hosts)添加entries:
- https://kubernetes.io/docs/concepts/services-networking/add-entries-to-pod-etc-hosts-with-host-aliases/
