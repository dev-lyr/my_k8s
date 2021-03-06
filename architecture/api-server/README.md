# 一 概述:
## (1)功能:
- 暴露Kubernetes API, 是Kubernetes的**控制面板**的前端.
- k8s系统组件间只能通过API服务器通信, 它们之间不会直接通信.
- API服务器是和etcd通信的唯一组件, 其它组件都通过api服务器来修改集群状态.
- API server提供一致性的对象访问, 也对对象进行校验(避免存入非法对象), 同时提供支持并发更新的乐观锁.

## (2)访问方式:
- 各种client库.
- kubectl proxy命令.
- Pod内访问: 使用每个容器内挂载的secret目录下的ca.crt,token和namespace来访问(例如:curl); 使用ambassador容器, 在该容器内执行kubectl proxy命令.

## (3)访问流程:
- Human User/Pod(k8s service account)(使用证书)->Authentication(认证)->Authorization(授权)->Admission Control->资源.

## (4)优化点:
- 自定义list-watch的store来减少存储数据的大小来避免内存占用过多.
- 分片list-watch, 实现参考: https://github.com/kubernetes/kube-state-metrics

# 二 访问控制:
## (1)概述:
- 访问方式: 可通过kubectl, client库或构造REST请求来访问API.
- 访问对象: human users和kubernetes service accounts可以被授权访问API.
- 访问流程: Human User/Pod(k8s service account)(使用证书)->Authentication(认证)->Authorization(授权)->Admission Control->资源.

## (2)传输层安全(TLS):
- 通常集群中, API使用端口443服务, API server提供一个certificate.
- 该证书是自签名(self-signed)的.
- 默认情况下, API Server在两个端口支持HTTP请求: localhost port(默认端口8080, 主要用于测试)和secure port.

## (3)认证:
- 一旦TLS连接建立, HTTP请求移动到Authentication(认证)阶段, 集群admin可以配置API服务器来运行一个或多个认证模块.
- 可指定使用多个认证模块, 顺序的尝试认证, 直至任一成功, 若认证不通过, 则以HTTP状态码**401**拒绝; 若认证通过,用户被认证为一个指定**username**.

## (4)授权:
- 当请求通过了认证, 则该请求必须被授权.
- 请求中必须包含: 请求者的username, 请求action和该action影响的对象.
- 若存在一个**policy**对象声明user有权限完成请求的action, 则该请求被授权.
- 若配置多个授权模块, 则Kubernetes检查每个模块, 只要有一个授权该请求则该请求可被执行; 若所有模块都拒绝该请求, 则该请求被拒绝(HTTP状态码**403**).

## (5)admission控制:
- admission模块是一个软件模块, 用于修改或拒绝请求, 还可以set complex defaults for fields.
- 可配置运行多个admission模块, 每个依序调度; 与认证和授权不同, 若任一admission模块拒绝请求, 则请求立马被拒绝.
- Once a request passes all admission controllers, it is validated using the validation routines for the corresponding API object, and then written to the object store.

## (6)webhook类型:
- 认证webhook
- 授权webhook
- 准入webhook: 支持动态配置.
- 审计webhook

## (7)备注:
- https://kubernetes.io/docs/reference/access-authn-authz/controlling-access/
