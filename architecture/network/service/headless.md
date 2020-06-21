# 一 概述:
## (1)概述:
- 有时不需要load-balancing和虚拟IP(服务clusterIP)的场景, 可以通过指定服务的.spec.clusterIP为None来创建"headless"服务.
- Headless服务减少和K8s系统的耦合, 允许开发者自由选择服务发现机制.
- 针对Headless服务, clusterIP没有分配, 因此kube-proxy不会处理这些服务, 且没有负载平衡和proxy, 怎么动态配置DNS依赖服务是否定义selector.

## (2)使用selectors:
- 针对定义了selectors的headless服务, endpoint控制器会创建endpoints记录, 并且会修改DNS配置来返回直接指向服务后端pods的**A记录**.

## (3)不使用selectors:
- **CNAME records**: **ExternalName-type** services.
- **A records**: any Endpoints that **share a name** with the service, for all other types.

