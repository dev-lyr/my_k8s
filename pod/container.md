# 一 镜像:
## (1)概述:
- 在kube pod中使用前, 需创建镜像并push到仓库.
- 参考: https://kubernetes.io/docs/concepts/containers/images

## (2)更新镜像:
- **imagePullPolicy**: IfNotPresent(默认), 当使用:latest tag时默认为Always; Always; Never.
- 开启**AlwaysPullImages** admission controller.
- 需避免使用:latest.

## (3)私有镜像仓库.

# 二 Container资源

# 三 容器环境变量:
## (1)概述:
- kube容器环境为容器提供一些重要的资源: 一个文件系统(镜像和一个或多个volume的组合); 容器自身相关信息; 集群中其它对象相关信息.
- https://kubernetes.io/docs/concepts/containers/container-environment-variables/

## (2)容器信息:
- hostname: 容器的hostname是包含该容器的pod的name.
- pod的name和namespace可通过downward API作为环境变量使用. 
- 在pod定义中自定义的环境变量对容器也是可用的.

## (3)集群信息:
- services: 在容器创建时所有运行中的服务作为容器的环境变量使用.
- Services have dedicated IP addresses and are available to the Container via DNS, if DNS addon is enabled. 

# 四 运行时类(runtime class):
## (1)概述:
- RuntimeClass用于选择容器运行时的配置.

## (2)备注:
- https://kubernetes.io/docs/concepts/containers/runtime-class/

# 五 容器生命周期:
## (1)概述:
- kube为容器提供了生命周期hooks, 可以让容器感知它们管理生命周期的事件, 并且hook被执行时运行对应的handler中实现的代码.

## (2)hook类型:
- PostStart: 当容器被创建后立即执行, 不保证会在容器的ENTRYPOINT前执行.
- PreStop

## (3)hook handler实现类型:
- Exec: 执行一个指定命令, 例如: pre-stop.sh.
- HTTP: 执行一个HTTP请求.

## (4)备注:
- https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/
