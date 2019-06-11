# 一 概述:
## (1)概述:
- REST API是k8s的基础组成部分, 所有组件间的操作和通信以及外部用户的命令都以REST API调用的形式被**API Server**处理.
- 通过**watch**支持高效的资源change通知, 一致性list可以让其它组件高效的cache和同步资源的状态.
- K8s平台中的所有事务都被当做API对象, 并且在API有一个对应的实体.
- 大部分通过kubectl命令行接口或者其它命令行工具(例如:kubeadm)都使用API, 也可以直接通过REST调用来访问API.
- 若你的应用需要调用k8s API, 则可以考虑使用客户端库.

## (2)客户端库:
- Java
- Go
- 等等.
- 备注:https://kubernetes.io/docs/reference/using-api/client-libraries

## (3)备注:
- https://kubernetes.io/docs/concepts/overview/kubernetes-api/
- API手册: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/

# 二 API version:
## (1)概述:
- 为了删除属性或重构资源表示, k8s支持多个API版本, 每个在不同的API路径, 例如:/api/v1或/apis/extensions/v1beta1.

## (2)版本类型：
- Alpha: 例如v1alpha1.
- Beta: 例如v2beta3
- Stable: 例如vX, X是整数.

# 三 API group:
## (1)概述:
- API组使扩展k8s API更加容易, API组在REST path中指定.

## (2)当前使用的API组:
- 核心组: **/api/版本**, 不用组名.
- 命名组: **/apis/组名/版本**
- 备注: 可在API手册看具体属于哪个组.

## (3)开启API组和资源:
- 特定资源和API组默认是开启的, 可通过在apiserver上设置--runtime-config来开启或关闭.
- DaemonSets, Deployments, Ingress, Jobs, ReplicaSets, HorizontalPodAutoscalers是默认开启的, 可通过--runtime-config来开启或关闭.

## (4)备注:
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md

# 四 相关概念:
## (1)概述:
- https://kubernetes.io/docs/reference/using-api/api-concepts/#alternate-representations-of-resources

## (2)标准API术语

## (3)资源变更的检测:
- 为了让client可以构建集群当前的状态模型, 所有k8s对象资源类型支持**一致性list(consistent list)**和**增量更新通知(watch)**.
- 每个k8s对象都有一个**resourceVersion**属性表示当前资源在底层存储的版本, 该字段通过服务器返回, 可用于初始化watch中的版本.
- 服务器会返回在指定resourceVersion之后的所有修改(创建,删除和更新).
- 该特性允许客户获取当前状态, 并且不会丢失任何修改.
- 若客户端watch断开, 则可以从最近的resourceVersion重新一个新的watch, 或者执行一个新的连接.
- 问题1: k8s只会保存有限时间的历史变更, etcd3默认保存5分钟, 因此watch时候带的resourceVersion太旧就会返回失败, 客户端则必须通过状态码410 Gone识别出来, 清理本地cache, 执行list, 然后根据list返回的resourceVersion重新watch, 多数client库提供该逻辑的, 例如:client-go中的Reflector.

## (4)分块获取大量结果:
- 背景: 为了防止一些资源类型的集合查询返回数据太大影响client和server, 例如:pod数据太大.
- k8s 1.9后开始支持将一个大的集合请求划分为多个小的chunks,同时保证一致性, 通过两个参数**limit**和**continue**来支持, 请求时使用limit和continue, 返回时metadata属性中返回continue属性.

