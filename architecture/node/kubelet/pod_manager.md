# 一 概述:
## (1)概述:
- 存储和管理对pod的访问,以及static pods和mirror pods间的映射.
- Kubelet relies on the pod manager as the source of truth for the desired state. If a pod does not exist in the pod manager, it means that it has been deleted in the apiserver and no action (other than cleanup) is required.

## (2)备注:
- pkg/kubelet/pod
