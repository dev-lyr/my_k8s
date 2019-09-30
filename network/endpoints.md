# 一 概述:
## (1)概述:
- Service不直接和Pods直接相连, 有一种资源介于两者之间, 即**Endpoints**, Endpoint和Service的name一样.
- 针对kubernetes-native应用, kubernetes提供一个简单的**Endpoints** API, 当Service中Pod变化时被更新. 
- 当定义service没有指定selector时不会自动创建endpoint, 随后可以根据不同情况手动创建.

## (2)subsets属性:
- addresses: IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
- notReadyAddresses: IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
- ports: Port numbers available on the related IP addresses.
