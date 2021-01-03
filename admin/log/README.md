# 一 概述:
## (1)概述:

## (2)备注:
- https://kubernetes.io/docs/concepts/cluster-administration/logging/
- kubectl logs

# 二 节点级别的日志
## (1)概述:
- 容器写入stdout和stderr的日志可以通过容器引擎被处理和重定向.
- 默认情况下, 若Pod被从Node上删除, 则所有容器和相关日志都会被删除.
- node上的日志需考虑日志rotate.

## (2)系统组件日志:
- 运行在容器内: scheduler和kube-proxy.
- 运行在容器外: kubelet和容器运行时, 在使用systemd的系统上, kubelet和容器运行时的日志写到journald, 若没有systemd, 则写入/var/log目录.

## (3)备注:
- https://docs.docker.com/config/containers/logging/configure/

# 三 集群级别日志:
## (1)概述:
- cluster-level-logging: 要求一个**独立的后端**来存储,分析和查询日志.
- k8s没有提供集群级别日志的解决方案, 可以将已存在的日志方案集成到k8s中.

## (2)可选方案:
- 在每个节点上使用一个node级别的logging agent,**推荐使用**.
- 使用sidecar容器.
- 应用直接expose日志或者推送日志到后端存储.

# 四 logging agent:
## (1)概述:
- agent expose logs或者push logs到一个backend, 通常agent是个可以访问某目录的容器, 该目录用来存放node上所有应用容器日志文件.
- agent需要在每个node上运行, 因此通常作为DaemonSet, Pod或节点上原始进程来运行, 后两种不推荐.

## (2)优缺点:
- 优点: 每个node上只需要一个agent, 且应用不需要修改.
- 缺点:logging agent只用于应用的stdout和stderr.

## (3)实现:
- fluentd

# 五 使用sidecar容器:
## (1)类型:
- 方案1:使用sidecar容器将应用的日志stream到自己的stdout或stderr, 从而可以被node上的logging agent来处理.
- 方案2:使用sidecar容器来运行一个logging agent, 该agent来从应用里pick up日志.

## (2)方案1优缺点:
- 优点: 可利用kubelet来处理stdout或stderr; 可以使用kubectl logs.
- 缺点: 将日志写入文件然后再stream到stdout会double disk usage.

## (3)方案2优缺点:
- 优点: 灵活.
- 缺点: 比较大资源开销; 且不能使用kubectl logs.

# 六 应用直接expose/push日志:
## (1)概述:
- 应用直接exposing日志或者pushing日志.


















