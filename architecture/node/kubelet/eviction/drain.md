# 一 概述:
## (1)使用方式:
- kubectl drain node(驱逐)和kubectl uncordon(恢复).
- POST /api/v1/namespaces/{namespace}/pods/{name}/eviction: 驱逐单个pod, kubectl drain也是调用该接口.

## (2)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/safely-drain-node/
