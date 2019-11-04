# 一 概述:
## (1)配置:
- 默认情况下, kubectl以$HOME/.kube目录下的config文件作为配置文件, 可通过通过KUBECONFIG环境变量或设置--kubeconfig flag来设置.
- https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/

## (2)基本命令:
- create: Create a resource from a file or from stdin.
- expose: 使用replication controller, service, deployment 或者 pod 并暴露它作为一个 新的 Kubernetes Service.
- run: 在集群中运行一个指定的镜像.
- set: 为objects设置一个指定的特征.
- explain: 查看资源的文档.
- get: 显示一个或更多resources.
- edit: 在服务器上编辑一个资源, 使用默认编辑器打开资源配置, 修改保存后资源对象会被更新.
- delete: Delete resources by filenames, stdin, resources and names, or by resources and label selector.

## (3)部署命令:
- rollout: Manage the rollout of a resource.
- scale: 为Deployment, ReplicaSet, Replication Controller或者Job设置一个新的副本数量.
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
- apply: 通过文件名或标准输入流(stdin)对资源进行配置, 通过一个完整的YAML或JSON文件, 应用其中的值来更新对象, 若对象不存在则会被创建.
- patch: 使用strategic merge patch更新一个资源的field(s), 修改单个资源属性, 不用编辑配置文件.
- replace: 通过filename或者stdin替换一个资源, 将对象替换为YAML/JSON文件中定义的新对象, 若对象不存在, 会报错.
- wait: Experimental: Wait for a specific condition on one or many resources..
- convert: 在不同的API versions转换配置文件.

## (7)setting命令:
- label: 更新在这个资源上的labels.
- annotate: 更新一个资源的注解.
- completion: Output shell completion code for the specified shell (bash or zsh).

## (8)其它:
- alpha: Commands for features in alpha.
- api-resources: 显示支持的API资源类型.
- api-versions: Print the supported API versions on the server, in the form of "group/version".
- config: 修改kubeconfig文件.
- plugin: Provides utilities for interacting with plugins.
- version: 输出client和server的版本信息.

## (9)选项:
- --kubeconfig: kubectl使用的kubeconfig文件.
- -n,--namespace: kubectl请求的ns范围.
- -v,--v=0: 日志的级别.
- 其它.

## (10)备注:
- Use "kubectl <command> --help" for more information about a given command.
- Use "kubectl options" for a list of global command-line options (applies to all commands).
- 参考: https://kubernetes.io/docs/reference/kubectl/overview/

# 二 基本命令:
## (1)get:
- 功能: 显示一个或多个资源相关信息.
- 可通过**--selector**来过滤列表; 若查询的资源类型是命名空间化的, 则只能看到当前namespace的, 除非指定--all-namespaces.
- 默认情况, 未初始化的对象不会显示, 除非指定--include-uninitialized.
- 可通过指定-o [yaml|json|...]等格式来获得详细信息.
- 可通过**kubectl api-resources**来查询支持资源的完整列表.

## (2)run:
- 功能: 创建一个deployment或job来管理容器(s).

## (3)explain:
- 功能: 显示API资源详情, 包括属性字段等.
- 用法: kubectl explain RESOURCE [options]
- 默认只显示一层属性, 可使用--recursive=true来显示所有层次.

# 三 部署命令:
## (1)rollout:
- 管理资源的rollout, 合法的资源类型: deployments, daemonsets和statefulsets.
- 用法:  kubectl rollout SUBCOMMAND [options]
- SUBCOMMAND: history(显示rollout历史), pause(标记提供的资源为中止状态), resume(继续一个中止的资源), status(显示rollout的状态), undo(撤销上一次rollout).

## (2)scale:
- 为Deployment, ReplicaSet, Replication Controller或StatefulSet设置一个新的大小(副本数量).
- 用法: kubectl scale [--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME) [options]

## (3)autoscale:
- 根据一个name来查找Deployment, ReplicaSet, StatefulSet或ReplicationContoller并创建一个autoscaler, 该autoscaler根据系统需要自动增加和减少pod的数量.
- 用法: kubectl autoscale (-f FILENAME | TYPE NAME | TYPE/NAME) [--min=MINPODS] --max=MAXPODS [--cpu-percent=CPU] [options]

# 四 集群管理
## (1)cluster-info:
- 显示master和标签为kuberbetes.io/cluster-service=true的服务的地址.

## (2)top

# 五 定位和debug命令:
## (1)describe:
- 功能: 描述指定resource或group的resource的详情.
- 用法: kubectl describe (-f FILENAME | TYPE [NAME_PREFIX | -l label] | TYPE/NAME) [options]

## (2)logs:
- 功能: 打印pod或指定资源中container的日志, 若pod只有一个container, 则container名字是可选的.
- 用法: kubectl logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER] [options]

## (3)attach:
- 功能: attach到已存在容器内运行的一个进程.
- Attach to a process that is already running inside an existing container.

## (4)exec:
- 功能: 在pod内一个容器内执行命令.
- 用法: kubectl exec POD [-c CONTAINER] -- COMMAND [args...] [options]

## (5)cp:
- 功能: 向容器或从容器向外复制文件.

# 六 高级命令:
## (1)apply:
- 功能: 通过文件名或标准输入流(stdin)对资源进行配置, 若不存在会创建.
- 用法: kubectl apply (-f FILENAME | -k DIRECTORY) [options].

## (2)patch:
- 功能: 使用strategic merge patch更新一个资源的field(s).
- 用法: kubectl patch (-f FILENAME | TYPE NAME) -p PATCH [options]

## (3)replace:
- 功能: 通过filename或者stdin替换一个资源, 若资源不存在会报错.
- 用法: kubectl replace -f FILENAME [options].
- 可使用kubectl get 导出完整列表, 修改后调用kubectl replace.

# 七 设置命令:
## (1)lable:
- 功能: 更新一个资源的label.
- 用法: kubectl label [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options]
- 默认--overwrite=false不允许更新已存在label, 可设置为true来更新.

## (2)annotate:
- 功能: 更新一个资源的注解.
- 用法: kubectl annotate [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version] [options]
