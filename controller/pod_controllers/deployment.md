# 一 概述:
## (1)功能:
- 提供Pods和ReplcaSets的声明式更新.
- 在Deployments对象中描述期望的状态, Deployments控制器负责更新当前状态到期望状态.

## (2)滚动更新(rolling updates)

## (3)deployment controller:
- watch的对象: deployment, rs和pod.

## (4)使用场景:
- 金丝雀部署(Canary Deployment).

## (5)kubectl rolling-update:
- 缺点: 会直接修改对象的配置; kubectl客户端负责执行滚动升级的复杂逻辑, 当出现网络断开等情况会处于中间状态.
- 备注: 已过时, 不建议使用, Deployment首选.

# 二 Deployment的spec:
## (1)Pod相关:
- spec.template: 指定pod模板, 除了Pod必须的属性外, Deployment中的pod模板还需指定合适的label和restart策略.
- spec.selector: 必选属性, deployment用来选择pod的标签选择器.
- spec.replicas: 可选的属性, 指定期望的pod数量, 默认为1.

## (2)spec.strategy:
- 指定替换旧的pod的策略, 可选**Recreate**或**RollingUpdate(默认)**.
- Recreate: 所有已存在的pod会被kill, 然后再创建新的出来.
- RollingUpdate: maxUnAvailable(可选的属性, 指定在更新过程中最大的不可用的pod的数量); maxSurge.

## (3)spec.progressDeadlineSeconds.

## (4)spec.minReadySeconds

## (5)spec.rollbackTo:
- 已废弃, 使用kubectl rollout undo.

## (6)spec.revisionHistoryLimit

## (7)spec.paused

# 三 创建:
## (1)概述:
- 创建Deployment时指定的Pod由**ReplicaSet**负责在后端创建.

## (2)kubectl get deployments:
- Name: 集群中Deployment的names.
- DESIRED: 期望的副本的数量, 对应spec.replicas.
- CURRENT: 当前运行中的副本数量, 对应status.replicas.
- UP-TO-DATE: 已被更新到期望状态的副本数量, 对应status.updatedReplicas.
- AVAILABLE: 当前可用的副本数量, 对应status.availableReplicas.
- AGE: 显示应用运行的总时间.

## (3)kubectl rollout status:
- 查看Deployment的rollout状态.

## (4)pod-template-hash label

# 四 更新:
## (1)概述:
- 只有在Deployment的pod template更新时才会触发Deployment的rollout.
- 通过更新Deployment的PodTemplateSpec来声明Pods的新状态, 此时一个**新的ReplicaSet会被创建**, Deployment以可控的速度将Pods从旧的ReplicaSet移动到新的.

## (2)修改方式:
- kubectl set
- kubectl edit

## (3)备注:
- 可通过设置spec.revisionHistoryLimit来指定多少旧的ReplicaSets希望保存, 其它的会被后台垃圾回收, 默认是10.

# 五 回滚(rolling back):
## (1)概述:
- 默认情况下, Deployment的rollout历史记录被保存在系统中, 所以可以任意时刻rollback.
- kubectl rollout history可查询历史rollout.

## (2)rollback到指定version:
- kubectl rollout undo xxx: 回滚到上一个版本.
- kubectl rollout undo xxx --to-revision=n: 回滚到指定版本.

# 六 扩容:
## (1)概述:
- kubectl scale

# 七 暂停和恢复:
## (1)概述:
- kubectl rollout pause
- kubectl rollout resume

# 八 Deployment的status:
## (1)概述:
- Deployment的状态有: Progressing, Complete和Failed.
- 可使用kubectl rollout status来查看.

## (2)Progressing

## (3)Complete

## (4)Failed
