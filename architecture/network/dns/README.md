# 一 概述:
## (1)概述:
- CoreDNS是推荐的DNS服务器, 用来替代kube-dns.
- CoreDNS和kube-dns的名字都为kube-dns.

## (2)kubelet相关:
- --cluster-dns=<dns-service-ip>: 将DNS传递到每个容器.
- --cluster-domain=<default-local-domain>
- --resolv-conf

## (3)备注:
- https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
- https://kubernetes.io/docs/tasks/administer-cluster/dns-custom-nameservers/
- https://kubernetes.io/docs/tasks/administer-cluster/dns-horizontal-autoscaling/
- https://github.com/kubernetes/dns/blob/master/docs/specification.md

# 二Service和Pod的DNS:
## (1)概述:
- Kubernetes DNS schedules a DNS Pod and Service on the cluster, and configures the kubelets to tell individual containers to use the DNS Service’s IP to resolve DNS names.

## (2)服务:
- A记录: 普通服务(非headless),解析到服务的ClusterIp; headless服务,解析到服务选择的Pods的IP的集合.
- SRV记录: 用于有named port的普通或headless服务.

## (3)Pod主机名和subdomain:
- 当前pod的hostname就是metadata.name, 除非指定hostname字段.
- podSpec还有个可选的subdomain字段用来指定它的subdomain, 若指定则the fully qualified Pod hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>",  若不指定, pod不会有domainname.
- 若存在一个headless service的name和pod的subdomain的name一致, kube的dns服务器会为Pod的全限定hostname提供一个A记录, **statefuleset**在使用该特性.

## (4)Pod DNS策略:
- Default: Pod继承Pod运行的node上的名字解析配置.
- ClusterFirst: 所有不match配置的cluster domain suffix的,会被转发给从node上继承的upstream nameserver.
- ClusterFirstWithHostNet: 使用hostNetwork.
- None: 允许Pod忽略kubernetes环境中的DNS配置, 使用podSpec.dnsConfig提供的DNS配置.

## (5)Pod DNS配置(dnsConfig):
- nameservers
- searches
- options

## (6)备注:
- DNS策略支持Pod级别的配置, Pod可通过Pod Spec的**dnsPolicy**属性来配置DNS策略.
- Pod的DNS Config为用户提供更多的更多的DNS配置, 通过**dnsConfig**配置.
- 当dnsPolicy设置为None时, 则dnsConfig必须指定.
- https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
- https://github.com/kubernetes/dns/blob/master/docs/specification.md

# 三 kubernete基于DNS服务发现的规范:
## (1)概述:
- https://github.com/kubernetes/dns/blob/master/docs/specification.md

## (2)有ClusterIP的service

## (3)Headless Service

## (4)External Name Service
