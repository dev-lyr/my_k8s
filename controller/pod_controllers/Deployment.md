# 一 概述:
## (1)功能:
- 提供Pods和ReplcaSets的声明式更新.
- 在Deployments对象中描述期望的状态, Deployments控制器负责更新当前状态到期望状态.

## (2)滚动更新(rolling updates)

## (3)Deployments对象属性

## (4)使用场景:
- 金丝雀部署(Canary Deployment).

# 二 创建:
## (1)概述:
- 创建Deployment时指定的Pod由ReplicaSet负责在后端创建.

## (2)kubectl get deployments:
- Name: 集群中Deployment的names.
- DESIRED: 期望的副本的数量, 对应spec.replicas.
- CURRENT: 当前运行中的副本数量, 对应status.replicas.
- UP-TO-DATE: 已被更新到期望状态的副本数量, 对应status.updatedReplicas.
- AVAILABLE: 当前可用的副本数量, 对应status.availableReplicas.
- AGE: 显示应用运行的总时间.

## (3)kubectl rollout status:
- 查看Deployment的rollout状态.

# 三 更新:
## (1)概述:
- 只有在Deployment的pod template更新时才会触发Deployment的rollout.
- 通过更新Deployment的PodTemplateSpec来声明Pods的新状态, 此时一个新的ReplicaSet会被创建, Deployment以可控的速度将Pods从旧的ReplicaSet移动到新的.

## (2)修改方式:
- kubectl set
- kubectl edit

## (3)备注:
- 可通过设置spec.revisionHistoryLimit来指定多少旧的ReplicaSets希望保存, 其它的会被后台垃圾回收, 默认是10.

# 四 回滚(rolling back):
## (1)概述:
- 默认情况下, Deployment的rollout历史记录被保存在系统中, 所以可以任意时刻rollback.
- kubectl rollout history可查询历史rollout.

## (2)rollback到指定version:
- kubectl rollout undo xxx: 回滚到上一个版本.
- kubectl rollout undo xxx --to-revision=n: 回滚到指定版本.

# 五 扩容:
## (1)概述:
- kubectl scale

# 六 暂停和恢复:
## (1)概述:
- kubectl rollout pause
- kubectl rollout resume

# 七 Deployment的status:
## (1)概述:
- Deployment的状态有: Progressing, Complete和Failed.
- 可使用kubectl rollout status来查看.

## (2)Progressing

## (3)Complete

## (4)Failed
