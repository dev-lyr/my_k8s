# 一 NetworkPolicy:
## (1)概述:
- 网络策略是一个specification, 用来指定pods组间以及与其它网络endpoints间如何通信.
- **NetworkPolicy**使用lables来选择pod, 并且为被选择的pods定义rule(指定哪些流量可以到pods).
- 网络策略是网络插件实现的, 所以必须使用一个支持**NetworkPolicy**的网络方案.

## (2)isolated和non-isolated Pods:
- 默认情况下, Pod是非孤立的, 可以接收任意源的流量.
- 当Pod被同一namespace下的NetworkPolicy选择时, 就变为独立的, pod会拒绝点NetworkPolicy不允许的连接.

## (3)NetworkPolicy资源:
- 参考: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#networkpolicy-v1-networking-k8s-io

## (4)备注:
- https://kubernetes.io/docs/concepts/services-networking/network-policies/
