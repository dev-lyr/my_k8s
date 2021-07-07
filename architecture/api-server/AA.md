# 一 概述:
## (1)概述:
- Aggregated APIs are subordinate APIServers that sit behind the primary API server, which acts as a proxy. This arrangement is called **API Aggregation(AA)**.
- The aggregation layer allows you to provide specialized implementations for your custom resources by writing and deploying **your own standalone API server**. The main API server **delegates** requests to you for the custom resources that you handle, making them available to all of its clients.

## (2)Aggregaton Layer:
- 聚合层与apiserver运行在一个进程内, 没有扩展资源时聚合层不做任何事情.
- 可通过添加APIService对象来注册API.
- 常用方式是在集群内的pod上运行extension API server来实现APIService.

## (3)备注:
- https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/
- https://github.com/kubernetes-sigs/apiserver-builder-alpha
- https://github.com/kubernetes/sample-apiserver
- 服务目录: https://github.com/kubernetes-incubator/service-catalog/blob/master/README.md

# 二 APIService:

# 三 配置:
## (1)概述

## (2)备注:
- https://kubernetes.io/docs/tasks/access-kubernetes-api/configure-aggregation-layer/

# 四 部署
## (1)概述

## (2)备注:
- https://kubernetes.io/docs/tasks/access-kubernetes-api/setup-extension-api-server/ 
