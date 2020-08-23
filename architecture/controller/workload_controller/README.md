# 一 概述:
## (1)类型:
- **ReplicaSet**: 是ReplicationController的继承者, 支持基于set的选择器(Selector).
- **ReplicationController**: 保证指定数量的pod副本运行, 多了删除, 少了启动.
- **Deployments**: 优先使用, 提供Pods和ReplicaSets的声明式更新.
- **StetefulSets**: 管理有状态应用.
- **DaemonSet**
- **Garbage Collection**
- **TTL Controller for Finished Resources**
- **Jobs**
- **CronJob**

## (2)选择:
- Deployment是一个高层次API对象, 通过类似kubectl rolling update形式来更新底层ReplicateSet和Pods, 若需要rolling update则推荐使用Deployment, 与kubectl rolling-update不同, Deployment是声明式,服务侧的且有其它特性.
- ReplicaSet是ReplicationController的继承, 目的类似, ReplicaSet支持基于集合的选择器, 因此优先使用ReplicaSet.
- Bare Pods: 不建议直接创建Pod, 使用controller创建的Pod在节点失败时候可以在其它节点重新启动, 而直接创建的Pod不支持. 即使应用只需要一个Pod, 此种情况下controller类似进程supervisor.
- Job: 当期望程序自己执行完终止(例如:批量jobs), 应该使用Job替换ReplicaSet.
- DaemonSet: 当需要Pod提供一个机器级别功能时使用DaemonSet替换ReplicaSet, 例如: 监控和日志.

## (3)经验:
- 原生workload缺少合适的分组策略且不支持分组暂停.
- 原地更新.

# 二 ReplicationController:
## (1)功能:
- 确保指定数量的pod副本的执行, 多了删, 少了创建.

# 三 ReplicaSet:
## (1)功能:
- 用来维护一个稳定的pod副本运行集合, 通常用来确保指定pod的副本的可用性.
- ReplicaSet根据期望的pod数量, 创建或销毁pod, 创建时候使用pod模板创建.

## (4)属性(fields):
- 选择器(selector): 指定如何识别可以获取的pod.
- 副本(replica)数量: 指定需要保持多少pod.
- pod模板: 新创建pod时使用的数据.

# 四 Deployments:
## (1)功能:
- 提供Pods和ReplcaSets的声明式更新.
- 在Deployments对象中描述期望的状态, Deployments控制器负责更新当前状态到期望状态.

## (2)使用场景:
- **Create a Deployment to rollout a ReplicaSet**. The ReplicaSet creates Pods in the background. Check the status of the rollout to see if it succeeds or not.
- **Declare the new state of the Pods by updating the PodTemplateSpec of the Deployment**. A new ReplicaSet is created and the Deployment manages moving the Pods from the old ReplicaSet to the new one at a controlled rate. Each new ReplicaSet updates the revision of the Deployment.
- **Rollback to an earlier Deployment revision if the current state of the Deployment is not stable**. Each rollback updates the revision of the Deployment.
- **Scale up the Deployment to facilitate more load**.
- **Pause the Deployment to apply multiple fixes to its PodTemplateSpec and then resume it to start a new rollout**.
- **Use the status of the Deployment as an indicator that a rollout has stuck**.
- **Clean up older ReplicaSets that you don’t need anymore**.

# 五 StatefuleSets:
## (1)概述:
- 管理pods集合的部署和扩容, 保证这些pod**有序**和**唯一**, 用来管理**有状态**应用.

# 六 DeamonSet:
## (1)概述:
- 保证所有node上都运行一个pod的副本, 当node新加入cluster时, pod会被添加到它, 当node离开集群时, pod被垃圾回收.
- 删除一个DeploySet会清除它创建的Pods.

## (2)用途:
- 在每个节点需要运行一个日志收集deamon, 例如:fluentd或logstash.
- 在每个节点需要运行一个监控daemon, 例如: Prometheus Node Exporter, collectd, Dynatrace OneAgent, AppDynamics Agent, Datadog agent, New Relic agent, Ganglia gmond or Instana agent.
- 在每个节点运行一个集群存储daemon, 例如: gluster,ceph.

# 七 Jobs:
## (1)概述:
- Job创建一个或多个Pod并保证指定数量的Pods成功终止.
- 当到达指定数量的Pod成功完成时, 则Job完成, 删除Job会清除它创建的Pod.
- 可以使用Job来并行运行多个Pods.

# 八 CronJob
## (1)概述:
- CronJob创建基于时间调度的Jobs, CronJob对象类似crontab文件, 它周期性根据一个给定的调度逻辑来执行Job.
