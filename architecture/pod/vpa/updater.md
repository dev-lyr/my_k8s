# 一 概述:
## (1)概述:
- Updater运行在kubernetes内并决定哪个pod应该被重启(基于Recommender计算的推荐值), 若pod应该被updated, 则updater会**驱逐**pod.
- Updater遵守PDB, 通过Eviction API来驱逐pod.

## (2)备注:
- https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/pkg/updater/README.md
