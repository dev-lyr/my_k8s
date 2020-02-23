# 一 概述:
## (1)概述:
- 背景: 开发人员不需知道底层使用那些存储技术, 基础设施相关的应该有集群管理员来管理.
- PersistentVolume子系统提供一个API来抽象storage是怎么提供以及怎么消费的.
- persistentvolume controller: pkg/controller/volume/persistentvolume.
- 名字有点误导, 常规的volume也提供持久化能力.

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
- https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/

# 二 PV:
## (1)概述:
- PV是集群级别资源, 不属于任何namespace.

## (2)PersistentVolumeSpec:
- accessModes: ReadWriteOnce(RWO),ReadOnlyMany(ROM),ReadWriteMany(RWM).
- capacity: 目前只有设置大小,未来可能会包括IOPS和吞吐量.
- storageClassName: pv属于的StorageClass的名字,空值说明不属于任何StorageClass.
- claimRef: PV和PVC双向绑定的一部分.
- mountOptions: mount选项列表.
- volumeMode: 使用格式化文件系统(filesystem,默认)还是raw block状态(block), 默认是格式化的文件系统.
- nodeAffinity: NodeAffinity defines constraints that limit what nodes this volume can be accessed from.
- persistentVolumeReclaimPolicy: 当pv被pvc释放后的操作. 可选: Retain(手动创建pv的默认), Delete(动态提供pv的默认, 删除底层存储), Recycle(删除卷的内容可再次claim).
- 各种pv类型的属性, 例如: azureDisk,hostPath等.

## (3)PersistentVolumeStatus:
- message
- phase: Available(空闲的资源, 还未绑定到claim); Bound(已绑定到claim), Released(claim已删除,但资源还未被集群回收), Failed(the volume has failed its automatic reclamation).
- reason

# 三 PVC
## (1)概述:
- PVC是namespace资源. 

## (2)PersistentVolumeClaimSpec:
- accessModes
- resources: 对存储资源的request.
- selector: 标签选择器, 只有match的label才可以绑定到该claim.
- storageClassName: 指定只有特定StorageClass的pv才能绑定到claim; 设置为""则只能绑定没有class的pv; 若不设置, 依赖是否开启DefaultStorageClass admission插件.
- volumeMode
- volumeName
- dataSource

## (3)PersistentVolumeClaimStatus
- accessModes: PVC对应的volume的实际access mode.
- capacity: 底层volume的实际资源.
- conditions: volume claim的当前condition.
- phase 

## (4)备注:
- 若想使用已有的PV, 需将storageClassName设置为"", 否则将会动态提供新的PV.

# 四 StorageClass
## (1)概述:
- StorageClass为管理员提供一种描述他们提供存储资源类型的方式, 非namespaced对象.
- 不同类型StorageClass可能有不同的QoS或不同的备份策略等等.
- 若PVC不指定storage class则, 可通过DefaultStorageClass admission设置一个默认的.

## (2)属性:
- provisioner: 绝对那个volume插件来提供PV.
- parameters
- reclaimPolicy
- mountOptions
- volumeBindingMode
- allowVolumeExpansion
- allowedTopologies

## (3)DefaultStorageClass admission

# 五 volume和claim的lifecycle:
## (1)提供(Provisioning):
- **static**: 集群管理员创建一些PV, 他们负责真实存储的细节, 这些PV可被使用.
- **dynamic**: 当没有static PV匹配用于的PVC时, 集群会尝试为PVC动态提供一个卷, 基于**StorageClasses**, pvc请求中带上storageClassName; storageClassName为""则disable动态提供.
- 备注: To enable dynamic storage provisioning based on storage class, the cluster administrator needs to enable the **DefaultStorageClass admission controller** on the API server.

## (2)绑定(binding):
- A control loop in the master watches for new PVCs, finds a matching PV (if possible), and binds them together.
- Once bound, PersistentVolumeClaim binds are exclusive, regardless of how they were bound. A PVC to PV binding is a one-to-one mapping, using a ClaimRef which is a bi-directional binding between the PersistentVolume and the PersistentVolumeClaim.
- Claims will remain unbound indefinitely if a matching volume does not exist. Claims will be bound as matching volumes become available. 

## (3)使用
- Pods use claims as volumes. The cluster inspects the claim to find the bound volume and mounts that volume for a Pod.

## (4)存储对象保护:
- 若用户删除一个正在被Pod使用的PVC, 则PVC不会立即被删除, 延缓至在Pod不使用后删除.
- 若管理员删除一个绑定到PVC的PV, PV不会立即被删除, 延缓至PV不在绑定到PVC再删除.
- PVC被保护: status是Terminating, Finalizer有kubernets.io/pvc-protection.
- PV被保护: status是Terminating, Finalizer有kubernets.io/pv-protection.
