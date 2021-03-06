# 一 概述:
## (1)概述:
- 背景: Pod是不稳定的, 例如ReplicaSets可以动态的创建和销毁pod(例如:扩容或缩容时), 虽然每个pod可以有自己的IP地址, 但是这些IP地址是不稳定的.
- **服务**: 定义了一个逻辑的Pod集合和一个访问它们的策略, Service通常通过**标签选择器**来选择Pod集.
- Service解耦了前端用户和后端服务, 当后端服务pod集合变化时, 前端用户不需要感知.

## (2)Endpoints资源:
- Service不直接和Pods直接相连, 有一种资源介于两者之间, 即**Endpoints**, Endpoint和Service的name一样.
- 针对kubernetes-native应用, kubernetes提供一个简单的**Endpoints** API, 当Service中Pod变化时被更新. 
- 当定义service没有指定selector时不会自动创建endpoint, 随后可以根据不同情况手动创建.

## (3)创建方式:
- kubectl expose
- kubectl create

## (4)集群内访问外部服务方法:
- 创建ExternalName类型服务
- 自定义Endpoint: 有固定IP地址的则直接创建对应EP, 若没有可自己实现controller来不断更新ep的地址.

## (5)外部访问服务的方法:
- NodePort
- LoadBalancer
- Ingress

# 二 spec
## (1)clusterIP:
- 服务的IP地址,通常由master随机分配, 若手动指定IP则需保证没有被其它使用, 该属性不能通过update来改变.
- 合法值: **None**,空字符串("")或一个合法ip地址.
- None: 针对不需要使用代理的**headless服务**.

## (2)externalIPs:
- externalIPs is a list of IP addresses for which **nodes in the cluster** will also accept traffic for this service.
- 这些IP不被kubernetes管理, 用户负责确保流量可以访问到该IP的Node.
- 备注: 适用于所有服务type, kube-proxy会创建对应listen端口.

## (3)externalName:
- externalName is the external reference that kubedns or equivalent will return as a **CNAME record** for this service. 
- No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.
- 为外部服务创建一个别名, Pod可以使用别名而不是外部服务的实际FQDN来访问外部服务, 隐藏了实际服务的名称以及使用该服务的pod的位置, 运行修改服务定义.

## (4)externalTrafficPolicy:
- 功能: denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints.
- Local: preserves the client source IP and avoids a second hop for **LoadBalancer and Nodeport** type services, but risks potentially imbalanced traffic spreading.
- Cluster: obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.

## (5)healthCheckNodePort

## (6)ports(ServicePort数组):
- 服务暴露的端口列表.
- name: pod在服务中的name, 在spec中需唯一, 若只有一个ServicePort则不填. 
- nodePort: 服务type是NodePort或LoadBalancer时在每个node上暴露的端口, 通常由系统分配.
- port: 服务暴露的端口.
- protocol: 默认是TCP, 支持TCP, UDP和SCTP.
- targetPort: pod的端口的数字或name.

## (7)selector:
- 将服务的流量路由到和选择器标签key和值match的pod.
- 若为空或者不指定, 则k8s会认为服务有有一个外部进程管理它的endpoints, 此时不会做修改.
- 适用于: type=ClusterIP,NodePort和LoadBanlancer, 若type为ExternalName则忽略.

## (8)负载平衡相关:
- loadBalancerIP: 只用于type=LoadBalancer, 部分云厂商支持指定该属性, 则会根据指定的loadbalanerIP来创建load-balancer, 若云提供商不支持, 则该忽略该属性.
- loadBalancerSourceRanges
- 相关: ServiceStatus.

## (9)session相关:
- sessionAffinity
- sessionAffinityConfig

## (10)type:
- ClusterIP: 默认,以集群内部IP的形式暴露服务, 因此服务只在集群内部可访问.
- ExternalName: 将服务映射到**externalName**属性的内容, 可用于为外部服务创建别名等.
- NodePort: 集群中每个节点上都打开一个端口, 并将端口上接收到的流量重定向到后端服务.
- LoadBalancer: 是NodePort的一种扩展, 使得服务可以通过专门的负载平衡器来访问, 由k8s运行的云基础设施提供, 客户端通过负载均衡器的IP地址来连接到服务, 负载均衡器有一个公开的访问IP地址.

# 三 发布服务:
## (1)概述:
- 需要把服务暴露到一个外部(集群外部)ip地址.
- k8s使用**ServiceTypes**来指定需要的服务类型.

## (2)服务类型：
- ClusterIP: 默认,以集群内部IP的形式暴露服务, 因此服务只在**集群内部**可访问.
- ExternalName: 将服务映射到**externalName**属性的内容, 可用于为外部服务创建别名等.
- NodePort: 集群中**每个节点**上都打开一个端口, 并将端口上接收到的流量重定向到后端服务.
- LoadBalancer: 是NodePort的一种扩展, 使得服务可以通过专门的负载平衡器来访问, 由k8s运行的云基础设施提供, 客户端通过负载均衡器的IP地址来连接到服务, 负载均衡器有一个公开的访问IP地址.
- 备注: 也可以使用Ingress来对外暴露服务.

## (3)ClusterIP和NodePort区别:
- ClusterIP只能集群内部访问.
- NodePort不仅可以集群内部IP访问, 也可以通过节点上任何的IP和预留端口来访问.

## (4)Ingress优点:
- 每个LoadBalance服务都需要有自己的负载均衡器和公网Ip地址.
- Ingress只需提供一个公网IP地址就可以为多个服务提供访问, 当客户请求到Ingress时会根据路径名将请求转发到对应服务.
- Ingress是基于HTTP层操作, 提供了其它服务类型不提供的功能, 例如基于cookie的会话黏贴(affinity).

## (5)LoadBalancer:
- 在支持外部loadbalancer的云提供商上, 可以通过设置type=LoadBalancer来为service提供一个load balancer.
- Load balancer的创建是异步的, 创建成功后, balancer的信息会在service的status.loadBalancer属性发布出来.
- LoadBalancer类型服务是一个具有额外基础设施提供的NodePort服务, 会给该类服务随机选择一个NodePort, 也可以通过节点访问服务.

## (6)ExternalName:
- 需指定spec.externalName.
- 为服务创建一个CNAME到指定的externalName.

## (7)备注:
- 相关:ingress.

# 四 服务发现(discovery):
## (1)概述:
- K8s支持两种主要类型的服务发现: **环境变量**和**DNS**(推荐).

## (2)环境变量:
- 在pod开始运行时, k8s会初始化一系列环境变量指向**已存在的服务**.
- 环境变量: {SVCNAME}_SERVICE_HOST和{SVCNAME}_SERVICE_HOST.

## (3)DNS:
- 可以(通常需要)为kubernete集群搭建DNS service, 此时需要安装CoreDNS等.
- DNS服务器(例如:CoreDNS)watch创建Service的k8s API, 并且创建为每个服务创建一个DNS记录集.
- 若集群中开启DNS, 则所有Pod可以自动做Service的命名解析.
- DNS server是唯一的一种访问**ExternalName**类型服务的方式.
- 例如: 在k8s中的my-ns命名空间中有一个my-service服务, 则一个DNS记录my-service.my-ns会被创建, 在my-ns命名空间中的Pod可以通过对my-service的简单名字查询来找到服务; 其它命名空间中的Pod则需要使用my-service.my-ns来查询; 名字查询的结果是**cluster IP**.

## (4)备注:
- 参考: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
- https://github.com/kubernetes/dns/blob/master/docs/specification.md
- CoreDNS和kube-dns.
