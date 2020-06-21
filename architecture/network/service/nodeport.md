# 一 概述:
## (1)概述:
- 集群中每个节点上都打开一个端口(端口一样), 并将端口上接收到的流量重定向到后端服务.
- 使用方式: Service的type=NodePort, 可使用nodePort属性来指定端口号(若忽略则随机选择).

## (2)相关配置:
- apiserver的service-node-port-range选项: 给NodePort服务预留的端口范围, 默认为30000-32767.
- kube-proxy的nodeport-addresses.
