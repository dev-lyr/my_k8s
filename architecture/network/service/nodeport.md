# 一 概述:
## (1)概述:
- nodePort服务既可以在集群内通过clusterIP访问, 也可以通过node ip和port被外部client访问.
- 使用方式: Service的type=NodePort, 可使用nodePort属性来指定端口号(若忽略则随机选择).
- 集群中**每个节点**上都打开一个端口(端口一样), 并将端口上接收到的流量重定向到后端service.
- 客户端通过任意node的ip和port来访问服务, 请求被随机发送到一个pods, 不一定是客户端连接的node上的pod, 可通过externalTrafficPolicy调整.
- 流量路径: 客户端->node->service->pod.

## (2)相关配置:
- apiserver的service-node-port-range选项: 给NodePort服务预留的端口范围, 默认为30000-32767.
- kube-proxy的nodeport-addresses.
- externalTrafficPolicy

## (3)注意:
- client指定某个具体node访问时, 当node出现问题时就不能访问服务, 所以建议是在所有健康的node前添加一个load balancer, 若集群支持, 可使用LoadBanlancer类型服务.
