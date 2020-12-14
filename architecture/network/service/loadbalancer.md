# 一 概述:
## (1)概述:
- 是NodePort的一种扩展, 使得服务可以通过专门的负载平衡器来访问, 由k8s运行的云基础设施提供.
- 客户端通过**负载均衡器的IP地址**来连接到服务, 负载均衡器有一个公开的访问IP地址.
- 流量路径: 客户->负载均衡器->node->service->pod

## (2)使用方式:
- 将service的type=LoadBalancer,支持外部负载平衡器的云provider异步创建load balancer, 将balancer的信息写入service的status.loadbalancer属性.
- 可通过loadBalancerIP指定loadbalancer的地址, 依赖云provider是否支持, 不支持则忽略该属性.
