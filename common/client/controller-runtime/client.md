# 一 概述:
## (1)概述:
- pkg/client: 提供与API server交互的功能.
- k8s中常用模式是从cache读和写入api server, 该模式通过DelegatingClient类型来实现.

## (2)New方法:
- 功能: 返回一个直接与API server交互的client.
- 对于普通类型, 使用Scheme来查询给定类型group,version和kind.
- 对于unstructured类型, group,version和kind信息从对象的属性来获取.

## (3)Client接口:
- Reader: get和list.
- Writer: create,update和delete.
- StatusClient
- 备注: 实现类typedClient和unstructuredClient. 

## (4)client结构体:
- typedClient
- unstructuredClient

## (5)备注:
- https://godoc.org/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured#Unstructured
