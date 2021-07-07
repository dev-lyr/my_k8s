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

## (3)资源类型:
- 集群范围资源
- 命名空间内的资源
- 子资源: 一些资源类型可能有1个或多个子资源, 格式: /apis/GROUP/VERSION/namespaces/NAMESPACE/RESOURCETYPE/NAME/SUBRESOURCE, 例如: eviction,scale等.

## (4)备注:
- https://kubernetes.io/docs/concepts/overview/kubernetes-api/
- https://kubernetes.io/docs/reference/using-api/api-concepts/
- https://kubernetes.io/docs/reference/using-api/api-concepts/#alternate-representations-of-resources
- https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md

# 二 API version:
## (1)概述:
- 为了删除属性或重构资源表示, k8s支持多个API版本, 每个在不同的API路径, 例如:/api/v1或/apis/extensions/v1beta1.

## (2)版本类型：
- Alpha: version名字包含alpha,例如v1alpha1.
- Beta: version名字包含beta, 例如v2beta3.
- Stable: 例如vX, X是整数.

## (3)备注:
- 可使用kuectl api-verisons查询.

# 三 API group:
## (1)概述:
- API组使扩展k8s API更加容易, API组在REST path中指定.
- Group names are typically in **domain name form**. 
- Kubernetes reserves use of the empty group, all single word names ("extensions", "apps"), and any group name ending in ".k8s.io" for its sole use.
- 当自定义资源选择group name时, 建议使用你所在组织的子域名, 例如:"xx.mycompany.com"

## (2)当前使用的API组:
- 核心组(core): **/api/版本**, 不用组名.
- 命名组(named group): **/apis/组名/版本**
- 备注: 可在API手册看具体属于哪个组.

## (3)开启API组和资源:
- 特定资源和API组默认是开启的, 可通过在apiserver上设置--runtime-config来开启或关闭.
- DaemonSets, Deployments, Ingress, Jobs, ReplicaSets, HorizontalPodAutoscalers是默认开启的, 可通过--runtime-config来开启或关闭.

## (4)支持自定义资源的路径:
- CRD: apiextensions.k8s.io/v1(1.16以上版本)或apiextensions.k8s.io/v1beta1
- AA: apiregistration.k8s.io/v1

## (5)备注:
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/api-machinery/api-group.md

# 四 资源变更的检测(watch):
## (1)概述:
- 为了让client可以构建集群当前的状态模型, 所有k8s对象资源类型支持**一致性list(consistent list)**和**增量更新通知(watch)**.
- 每个k8s对象都有一个**resourceVersion**属性表示当前资源在底层存储的版本, 该字段通过服务器返回, 可用于初始化watch中的版本.
- 服务器会返回在指定resourceVersion之后的所有修改(创建,删除和更新).
- 该特性允许客户获取当前状态, 并且不会丢失任何修改.
- 若客户端watch断开, 则可以从最近的resourceVersion重新一个新的watch, 或者执行一个新的连接.

## (2)Watch bookmarks:
- 功能: 减少由于较小历史窗口(resourceVersion过旧)而造成的relist的次数.
- **bookmark**事件: a special kind of event to mark that all changes up to a given resourceVersion the client is requesting have already been sent.
- 使用方法: watch请求带上allowWatchBookmarks=true.

## (3)相关问题:
- k8s只会保存有限时间的历史变更, etcd3默认保存5分钟, 因此watch时候带的resourceVersion太旧就会返回失败, 客户端则必须通过状态码410 Gone识别出来, 清理本地cache, 执行list, 然后根据list返回的resourceVersion重新watch, 多数client库提供该逻辑的, 例如:client-go中的Reflector.

## (4)备注:
- https://kubernetes.io/docs/reference/using-api/api-concepts/

# 五 分批查询:
## (1)概述:
- 背景: 为了防止一些资源类型的集合查询返回数据太大影响client和server, 例如:pod数据太大.
- k8s 1.9后开始支持将一个大的集合请求划分为多个小的chunks,同时保证一致性, 通过两个参数**limit**和**continue**来支持, 请求时使用limit和continue, 返回时metadata属性中返回continue属性, 当continue为空是表示没有更多资源返回.
- 在每次分批查询的请求中, resourceVersion不变(一致性读), 在resourceVersion之后的修改并不会返回.

# 六 其它表示形式
## (1)概述:
- 默认kube将对象序列化为JSON返回(content type为applicaiton/json).

# 七 资源删除:
## (1)概述:
- 资源删除分为两个阶段: finalization和removal.
- 一些对象是其它对象的owner, 被owner的对象是owner对象的附属, 每个附属对象有一个**metadata.ownerReferences**属性指向owner对象, 在一些情况下, 该属性会被kube自动设置, 也可以手动设定.

## (2)流程:
- 当client删除一个资源时, .metadata.deletionTimestamp被设置为当前时间.
- Once the .metadata.deletionTimestamp is set, external controllers that act on finalizers may start performing their cleanup work at any time, in any order.
- 一旦最后的finalizer被删除,该资源才会实际的从etcd中删除.

## (3)控制删除依赖:
- 级联删除(cascading deletion): 删除对象时自动删除它的附属, 级联删除分为2种: backgroup和foregroup.
- 若删除对象时不删除它的附属, 则依赖会变成孤儿(orphaned).

## (4)备注:
- https://kubernetes.io/docs/concepts/workloads/controllers/garbage-collection/
