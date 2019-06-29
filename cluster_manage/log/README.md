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
- cluster-level-logging: 要求一个独立的后端来存储,分析和查询日志.
- k8s没有提供集群级别日志的解决方案, 可以将已存在的日志方案集成到k8s中.

## (2)可选方案:
- 在每个节点上使用一个node级别的logging agent,推荐使用.
- 使用一个显式的sidecar容器来记录pod中应用的日志.
- 应用直接推送日志到后端存储.

## (3)node logging agent:
- agent通常作为DaemonSet, Pod或节点上原始进程来运行, 后两种不推荐.
- logging agent只用于应用的stdout和stderr.
- 可选的logging agent: Elasticsearch等.

## (4)sidecar容器:
- 两种使用方式: sidecar容器streams应用的日志到它自己的stdout; sidecar容器运行一个logging agent, 从应用容器pick up日志.

