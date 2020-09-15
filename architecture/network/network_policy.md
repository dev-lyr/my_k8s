# 一 概述:
## (1)概述:
- 网络策略(NetworkPolicies): 指定pod如何与其它网络实体通信.
- 网络策略是网络插件(CNI)来实现的, 使用时需要网络插件也支持NetworkPolicy(flannel不支持,calico支持).
- 默认情况下, pod是non-isolated的, 可以接收任意source的流量.
- 当有NetworkPolicy选择某个pod后, pod会变为isolated, pod会拒绝NetworkPolicy不允许的连接.

## (2)备注:
- https://kubernetes.io/docs/concepts/services-networking/network-policies/
