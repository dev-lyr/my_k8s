# 一 概述:
## (1)概述:
- PersistentVolume子系统提供一个API来抽象storage是怎么提供以及怎么消费的.
- persistentvolume controller: pkg/controller/volume/persistentvolume.

## (2)相关资源:
- **PersistentVolume(PV)**: a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using **Storage Classes**.
- **PersistentVolumeClaim(PVC)**: a request for storage by a user.
- **StorageClass**: provides a way for administrators to describe the “classes” of storage they offer.

## (3)PV的类型:
- CSI
- NFS
- 等等.
- 备注: PV类型以插件方式实现.

## (4)备注:
- https://kubernetes.io/docs/concepts/storage/persistent-volumes/
- https://kubernetes.io/docs/concepts/storage/storage-classes/

# 二 PVC

# 三 PV

# 四 volume和claim的lifecycle:
## (1)提供(Provisioning):
- **static**: 集群管理员创建一些PV, 他们负责真实存储的细节, 这些PV可被使用.
- **dynamic**: 当没有static PV匹配用于的PVC时, 集群会尝试为PVC动态提供一个卷, 基于**StorageClasses**, pvc请求中带上storageClassName; storageClassName为""则disable动态提供.
- 备注: To enable dynamic storage provisioning based on storage class, the cluster administrator needs to enable the **DefaultStorageClass admission controller** on the API server.

## (2)绑定(binding)

## (3)使用

## (4)存储对象保护
