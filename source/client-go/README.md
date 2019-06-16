# 一 概述:
## (1)概述:
- k8s集群的Go client.

## (2)组成:
- kubernetes目录: 包含访问k8s API集合.
- discovery: 用来发现k8s API服务器支持的API.
- dynamic: 包含一个动态client, 用来在k8s API东西上执行通用操作.
- plugin/pkg/client/auth: 包含可选的authentication插件.
- transport: 建立auth和开始一个连接.
- tools/cache: 用来实现contoller.

## (3)备注:
- https://github.com/kubernetes/client-go
