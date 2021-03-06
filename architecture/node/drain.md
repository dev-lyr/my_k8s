# 一 概述:
## (1)使用方式:
- kubectl drain node(驱逐)和kubectl uncordon(恢复).
- POST /api/v1/namespaces/{namespace}/pods/{name}/eviction: 驱逐单个pod, kubectl drain也是调用该接口.

## (2)kubectl drain:
- --delete-local-data=false: 为true表示即使pod使用emptyDir也删除(当node被drain时本地数据会被删除).
- --disable-eviction=false: 为true表示drain使用delete, 即使支持使用Eviction, 使用delete会绕过PodDisruptionBudgets, 需小心.
- --force=false: 为true表示即使遇到不被controller管理的pod时也继续执行.
- --grace-period=-1: 表示等待pod平滑结束的时间, 若未负则使用pod指定的默认.
- --ignore-daemonsets=false: 忽略ds的pods.
- -l, --selector='': 过滤pods.
- --skip-wait-for-delete-timeout=0
- --timeout=0s
- --dry-run='none'

## (3)Eviction API:
- pod的子资源Eviction可认为是pod的**策略控制的删除操作**.
- 情况一: eviction被批准,则pod会被删除并返回200,就像针对pod发送了一个DELETE请求.
- 情况二: 当前状态不允许驱逐(不满足pdb)则会返回429 Too Many Requests.
- 情况三: 其它类型错误配置(例如:多个pdb关联同一pod), 500 Internal Server Error.
- 备注: pkg/registry/core/pod/storage/eviction.go

## (4)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/
- https://github.com/kubernetes/kubectl
- registry/core/pod/storage/eviction.go


