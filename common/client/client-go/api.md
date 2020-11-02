# 一 概述:
## (1)目录:
- rest: 提供访问kubernetes api的底层rest client.
- kubernetes: k8s API的封装, 代码都是通过client-gen生成, 不要修改.
- listers: 用来帮助从cache中list和get资源, 代码通过lister-gen生成, 不要修改.
- tools/clientcmd: provides one stop shopping for building a working client from a fixed config,from a .kubeconfig file, from command line flags, or from any merged combination.

## (2)相关:
- kubernetes/pkg/controller/client_builder.go

# 二 rest:
## (1)概述:
- 提供访问kubernetes的rest api.
- rest.Interface: captures the set of operations for generically interacting with Kubernetes REST apis.
- rest.Config: 初始化kubernetes client需要的通用属性.

## (2)client.go:
- RESTClient

## (3)config.go

# 三 kubernetes:
## (1)概述:
- 提供访问K8s api的clientset, 底层使用rest.Interface.
- Clientset: contains the clients for groups.
- 各controller使用Clientset调用apiserver.

## (2)clientset.go:
- New: 根据指定RESTClient创建clientset.
- NewForConfig: 根据给定config创建clientset.
- NewForConfigOrDie

## (3)typed:
- 包含对各类API资源的curd.

# 四 lister:
## (1)概述:
- 用来从cache.Indexer中list和get资源, 返回的对象**不能被修改(只读)**.
- 各controller通过lister来从cache中查询资源.
