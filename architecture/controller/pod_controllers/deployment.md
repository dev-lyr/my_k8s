# 一 概述:
## (1)功能:
- 提供Pods和ReplcaSets的声明式更新.
- 在Deployments对象中描述期望的状态, Deployments控制器负责更新当前状态到期望状态.

## (2)滚动更新(rolling updates)

## (3)deployment controller:
- watch的对象: deployment,rs和pod.
- rs的名字格式为:[deployment-name]-[random-string],随机字符串使用pod-template-hash作为seed随机产生.

## (4)使用场景:
- 金丝雀部署(Canary Deployment).

## (5)kubectl rollout:
- history: 显示rollout历史.
- pause
- resume
- status
- undo
- 备注: rollout的开始是通过修改pod template来触发的.

## (6)kubectl rolling-update(过时):
- 缺点: 会直接修改对象的配置; kubectl客户端负责执行滚动升级的复杂逻辑, 当出现网络断开等情况会处于中间状态.
- 备注: 已过时, 不建议使用, Deployment首选.

# 二 spec:
## (1)Pod相关:
- spec.template: 指定pod模板, 除了Pod必须的属性外, Deployment中的pod模板还需指定合适的label和restart策略.
- spec.selector: 必选属性, deployment用来选择pod的标签选择器,必须与spec.template.metadata.labels匹配,否则会被API拒绝.
- spec.replicas: 可选的属性, 指定期望的pod数量, 默认为1.
- 备注: 要选择合适的selector和pod模板的label, 不要与其它controller覆盖, kubernetes不阻止overlap, 但是多个controller有覆盖的selectors则可能会冲突并出现非预期行为.

## (2)spec.strategy:
- **type**: 指定替换旧的pod的策略,**RollingUpdate**(默认,滚动升级)和**Recreate**(所有已存在pod被kill,然后再新建出来).
- **rollingUpdate**: 滚动更新的参数, 包含maxSurge和maxUnAvailable.
- maxSurge: 决定了最多允许超出deployment副本期望的pod数量的比例, 默认为25%, 将百分比转换为绝对值时会进行四舍五入, 默认25%.
- maxUnAvailable: 决定了在滚动升级期间, 相对于期望副本数, 允许有多少比例的pod处于不可用状态, 默认为25%, 会进行四舍五入, 默认25%.
- 备注: maxUnAvailable和maxSurge用于决定一次替换多少pod, 控制滚动升级的速率.

## (3)spec.progressDeadlineSeconds:
- 默认值为600s, 在10min内不能完成滚动升级, 则被视为失败.
- Deployment控制器会继续处理失败的deployment, 并且在deployment的status会展示一个**ProgressDeadlineExceeded** reason.

## (4)spec.minReadySeconds:
- 指定新创建pod在成功运行多久后才视为可用, 当minReadySeconds时间内, 就绪探针出现了失败, 则新版本的滚动升级会被暂停.
- 使用正确的就绪探针和minReadySeconds可以预先阻止发布有问题的版本.
- 默认为0, pod是ready就认为是可用的.

## (5)spec.paused:
- 表示deployment暂停或恢复的一个属性.

## (6)spec.revisionHistoryLimit:
- 功能: 保存的用于回滚的rs的数量,默认10.

# 三 status:
## (1)概述:
- Deployment的状态有: Progressing, Complete和Failed.
- 可使用kubectl rollout status来查看.

## (2)Progressing

## (3)Complete

## (4)Failed

# 四 创建:
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

## (4)pod-template-hash标签:
- deployment控制器会给每个deployment创建的rs添加一个**pod-template-hash** label, 该label的值是rs的PodTemplate的hash值.
- 该label来确保一个deploy的子rs不会overlap.
- 该label会被添加到rs selector, pod template labels和rs管理的每个pod中.

# 五 滚动更新:
## (1)概述:
- 只有在Deployment的**pod template(spec.template)**更新时才会触发Deployment的rollout.
- 通过更新Deployment的PodTemplateSpec来声明Pods的新状态, 此时一个**新的ReplicaSet会被创建**, Deployment以可控的速度将Pods从旧的ReplicaSet移动到新的.

## (2)修改方式:
- kubectl set
- kubectl edit
- 备注: --record

## (3)备注:
- 旧的replicateset会保留, 可通过设置spec.revisionHistoryLimit来指定多少旧的ReplicaSets希望保存, 其它的会被后台垃圾回收, 默认是10.

# 六 回滚(rolling back):
## (1)概述:
- 默认情况下, Deployment的rollout历史记录被保存在系统中, 所以可以任意时刻rollback.
- Deployment的revision在Deployment rollout被触发时创建.
- kubectl rollout history可查询历史rollout.

## (2)回滚到指定version:
- kubectl rollout undo xxx: 回滚到上一个版本.
- kubectl rollout undo xxx --to-revision=n: 回滚到指定版本.

# 七 扩容:
## (1)概述:
- kubectl scale

# 八 暂停和恢复:
## (1)功能
- 通过暂停, 可以同时修改多次deployment, 在修改完毕后再通过恢复了触发一次滚动升级, 而不是每次修改都触发.
- 可以在滚动升级过程中, 暂停和恢复, 来进行金丝雀发布.

## (2)命令:
- kubectl rollout pause
- kubectl rollout resume

