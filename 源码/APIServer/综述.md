# 一 概述:
## (1)功能:
- 暴露Kubernetes API, 是Kubernetes的**控制面板**的前端.
- k8s系统组件间只能通过API服务器通信, 它们之间不会直接通信.
- API服务器是和etcd通信的唯一组件, 其它组件都通过api服务器来修改集群状态.
- 支持水平扩展(scale horizontally).

## (2)相关代码:
- kubernetes/cmd/kube-apiserver: 启动入口.
