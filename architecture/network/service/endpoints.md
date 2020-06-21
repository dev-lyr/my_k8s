# 一 概述:
## (1)概述:
- Service不直接和Pods直接相连, 有一种资源介于两者之间, 即**Endpoints**, Endpoint和Service的name一样.
- 针对kubernetes-native应用, kubernetes提供一个简单的**Endpoints** API, 当Service中Pod变化时被更新. 
- 当定义service没有指定selector时不会自动创建endpoint, 随后可以根据不同情况手动创建.
- Endpoints只有一个subsets属性, 是一个EndpointSubset数组.

## (2)EndpointSubset:
- **addresses(EndpointAddress数组)**: IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.
- **notReadyAddresses(EndpointAddress数组)**: IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.
- **ports(EndpointPort数组)**: Port numbers available on the related IP addresses.

## (3)EndpointAddress:
- hostname
- ip: endpoint的IP, 例如: pod的IP.
- nodeName: 可选, endpoint的节点.
- targetRef(ObjectReference): 到提供endpoint的对象的引用, 例如:指向一个pod对象.

## (4)EndpointPort:
- name
- port
- protocol: 默认TCP, 可选TCP, UDP, SCTP.

## (5)手动创建endpoint:
- 手动创建ep时, ep的名字需与service名字相同, 并包含服务目标IP地址和端口.
- 使用场景: 访问外部服务,为service的使用方提供一致的访问方法, 底层目标机器可以是pod或者外部机器, 可用于将外部服务迁入k8s集群或者反向.

# 二 endpoint控制器:
## (1)EndpointController struct:
- queue: 存放需要update的service.
- 等等.

## (2)NewEndpointController:
- watch service
- watch pod

## (3)调和(syncService):
- 从队列中取出待调和的service的ns和name, 然后从service的cache中查询service对象.
- 若出错, 若错误是没有找到, 则删除同名endpoint; 其它错误则返回由handleErr处理.
- 若service对象没有selector则不处理, 直接返回(没有selector的service不会自动创建ep).
- 从pod cache中取出匹配selector的所有pod.
- 判断service是否tolerateUnreadyEndpoints(spec指定或使用TolerateUnreadyEndpointsAnnotation注解(已废弃)).
- 根据list出的pod和tolerateUnreadyEndpoints构造ep对象, 然后查询当前的ep, 比较是需要create还是update.
- 等等.

# 三 endpointSlice:
## (1)概述:

## (2)备注:
- https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/
