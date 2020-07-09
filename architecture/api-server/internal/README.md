# 一 概述:
## (1)概述:
- kubernetes/cmd/kube-apiserver: 启动入口.
- apiserver: https://godoc.org/k8s.io/apiserver
- apiextensions: https://github.com/kubernetes/apiextensions-apiserver
- kube-aggregator: https://github.com/kubernetes/kube-aggregator

## (2)apiextensions:
- 提供注册CRD的API, 在kube-apiserver中作为delegate服务器运行.

## (3)kube-aggregator:
- Provide an API for registering API servers.
- Summarize discovery information from all the servers.
- Proxy client requests to individual servers.

# 二 apiserver:
## (1)概述:
- https://github.com/kubernetes/apiserver
- https://godoc.org/k8s.io/apiserver

## (2)目录:
- admission
- apis
- audit
- authentication
- authorization
- endpoints
- features
- registry
- server
- storage
- util
