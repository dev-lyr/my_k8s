# 一 概述:
## (1)概述:
- 容器存储接口(CSI)为容器编排系统(例如:k8s)定义一个标准的接口, 用来将任意存储系统暴露给它们的容器workload.
- 容器编排系统和CSI插件通过**RPC**服务通信.

## (2)相关名词:
- Volume
- Block Volume
- Mounted Volume
- PV
- PVC
- StorageClass
- SP: storage provider.

## (3)使用方式:
- 通过pvc(CSIPersistentVolumeSource)
- CSI ephemeral volume(pod.spec.volume:csi(CSIVolumeSource))
- generic ephemeral volumes

## (4)备注:
- https://kubernetes.io/docs/concepts/storage/volumes/#csi
- https://kubernetes-csi.github.io/docs/introduction.html
- https://github.com/container-storage-interface/spec/blob/master/spec.md
- kubernetes/pkg/volume/csi

# 二 使用:
## (1)概述:
- 需要在kubernetes中部署csi compatible volume driver, 用户通过使用**csi**卷类型来attach,mount等, 该卷通过csi driver暴露.
- csi卷类型不支持直接通过Pod引用,Pod只能通过pvc对象来引用.

## (2)CSIPersistentVolumeSource:
- driver: 卷驱动的名字,与CSI驱动GetPluginInfoResponse返回的值一致.
- volumeHandle: 一个字符串,唯一标记一个卷.
- volumeAttributes
- fsType
- readOnly
- controllerExpandSecretRef
- controllerPublishSecretRef
- nodePublishSecretRef
- nodeStageSecretRef
