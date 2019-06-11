# 一 概述:
## (1)功能:
- kubectl controls the Kubernetes cluster manager.
- 参考: https://kubernetes.io/docs/reference/kubectl/overview/

## (2)基本命令:
- create: Create a resource from a file or from stdin.
- expose: 使用replication controller, service, deployment 或者 pod 并暴露它作为一个 新的 Kubernetes Service.
- run: 在集群中运行一个指定的镜像.
- set: 为objects设置一个指定的特征.
- explain: 查看资源的文档.
- get: 显示一个或更多resources.
- edit: 在服务器上编辑一个资源.
- delete: Delete resources by filenames, stdin, resources and names, or by resources and label selector.

## (3)部署命令:
- rollout: Manage the rollout of a resource.
- scale: 为Deployment, ReplicaSet, Replication Controller 或者 Job 设置一个新的副本数量.
- autoscale: 自动调整一个Deployment, ReplicaSet, 或者 ReplicationController 的副本数量.

## (4)集群管理:
- certificate: 修改certificate 资源.
- cluster-info: 显示集群信息.
- top: Display Resource (CPU/Memory/Storage) usage.
- cordon: 标记node为unschedulable.
- uncordon: 标记node为schedulable.
- drain: Drain node in preparation for maintenance.
- taint: 更新一个或者多个node上的taints.

## (5)定位和debug命令:
- describe: 显示一个指定resource或者group的resources详情.
- logs: 输出容器在pod中的日志.
- attach: Attach到一个运行中的container.
- exec: 在一个container中执行一个命令..
- port-forward: Forward one or more local ports to a pod.
- proxy: 运行一个proxy到Kubernetes API server.
- cp: 复制files和directories到containers和从容器中复制files和directories..
- auth: Inspect authorization.

## (6)高级命令:
- apply: 通过文件名或标准输入流(stdin)对资源进行配置.
- patch: 使用strategic merge patch更新一个资源的field(s).
- replace: 通过filename或者stdin替换一个资源.
- wait: Experimental: Wait for a specific condition on one or many resources..
- convert: 在不同的API versions转换配置文件.

## (7)setting命令:
- label: 更新在这个资源上的labels.
- annotate: 更新一个资源的注解.
- completion: Output shell completion code for the specified shell (bash or zsh).

## (8)其它:
- alpha: Commands for features in alpha.
- api-resources: Print the supported API resources on the server.
- api-versions: Print the supported API versions on the server, in the form of "group/version".
- config: 修改kubeconfig文件.
- plugin: Provides utilities for interacting with plugins.
- version: 输出client和server的版本信息.

## (9)备注:
- Use "kubectl <command> --help" for more information about a given command.
- Use "kubectl options" for a list of global command-line options (applies to all commands).
