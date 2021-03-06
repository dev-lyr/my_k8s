# 一 概述:
## (1)概述:
- 背景: 开发人员不需知道底层使用那些存储技术, 基础设施相关的应该有集群管理员来管理.
- PersistentVolume子系统提供一个API来抽象storage是怎么提供以及怎么消费的.
- PV类型通过插件形式实现.

## (2)相关资源:
- **PersistentVolume(PV)**: a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using **Storage Classes**.
- **PersistentVolumeClaim(PVC)**: a request for storage by a user.
- **StorageClass**: provides a way for administrators to describe the “classes” of storage they offer.
- **VolumeAttachment**

## (3)PV的类型:
- CSI: CSIPersistentVolumeSource.
- NFS
- hostpath: 测试使用.
- 备注: PV类型以插件方式实现.

## (4)流程:
- bind: pv控制器将pvc绑定到一个pv.
- attach/detach: attachdetach控制器.
- mount/unmount: kubelet的VolumeManagerReconciler, mount到node上一个目录(注意不是mount到容器里).

## (5)备注:
- persistentvolume controller: pkg/controller/volume/persistentvolume.
- https://kubernetes.io/docs/concepts/storage/persistent-volumes/
- https://kubernetes.io/docs/concepts/storage/storage-classes/
- https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/persistent-storage.md

# 二 PV:
## (1)概述:
- PV是集群级别资源, 不属于任何namespace.
- PV是volume plugin(与Volumes类似), 但是独立于使用该pv的pod的生命周期.

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
- phase: Pending(pv当前不可用), Available(空闲的资源, 还未绑定到claim); Bound(已绑定到claim), Released(claim已删除,但资源还未被集群回收), Ffailed(the volume has failed its automatic reclamation).
- reason

# 三 PVC
## (1)概述:
- PVC是namespace资源, 是用户对存储资源的请求.

## (2)PersistentVolumeClaimSpec:
- accessModes
- resources: 对存储资源的request.
- selector: 标签选择器, 只有match的label才可以绑定到该claim.
- storageClassName: 指定只有特定StorageClass的pv才能绑定到claim; 设置为""则只能绑定没有class的pv; 若不设置, 依赖是否开启DefaultStorageClass admission插件.
- volumeMode
- volumeName: the binding reference to the PersistentVolume backing this claim.
- dataSource

## (3)PersistentVolumeClaimStatus
- accessModes: PVC对应的volume的实际access mode.
- capacity: 底层volume的实际资源.
- conditions: volume claim的当前condition.
- phase: Pending, Bound, Lost(pvc失去了底层pv, 该pvc bound过但是pv不存在所有数据已丢失).

## (4)备注:
- 若想使用已有的PV, 需将storageClassName设置为"", 否则将会动态提供新的PV.

# 四 StorageClass
## (1)概述:
- StorageClass为管理员提供一种描述他们提供存储资源类型的方式, 非namespaced对象.
- 不同类型StorageClass可能有不同的QoS或不同的备份策略等等.
- 若PVC不指定storage class,则可通过DefaultStorageClass admission设置一个默认的.
- 通过注解(storageclass.kubernetes.io/is-default-class)来标记某StorageClass为default.

## (2)属性:
- provisioner: 决定那个volume插件来提供PV.
- parameters
- reclaimPolicy
- mountOptions
- volumeBindingMode
- allowVolumeExpansion
- allowedTopologies

## (3)DefaultStorageClass admission:
- 观察PVC的创建, 若未指定storageClass则自动加上默认.
- 若没有配置默认storageClass则不干活, 配置多个默认storageClass会报错.

## (4)volume-provision规范:
- https://github.com/kubernetes/community/blob/master/contributors/design-proposals/storage/volume-provisioning.md

# 五 volume和claim的lifecycle:
## (1)提供(Provisioning):
- **static**: 集群管理员创建一些PV, 他们负责真实存储的细节, 这些PV可被使用.
- **dynamic**: 当没有static PV匹配用于的PVC时, 集群会尝试为PVC动态提供一个卷, 基于**StorageClasses**, pvc请求中带上storageClassName; storageClassName为""则disable动态提供.
- 备注: To enable dynamic storage provisioning based on storage class, the cluster administrator needs to enable the **DefaultStorageClass admission controller** on the API server.

## (2)绑定(binding):
- A control loop in the master watches for new PVCs, finds a matching PV (if possible), and binds them together.
- Once bound, PersistentVolumeClaim binds are exclusive, regardless of how they were bound. 
- A PVC to PV binding is a **one-to-one mapping**, using a ClaimRef which is a bi-directional binding between the PersistentVolume and the PersistentVolumeClaim.
- Claims will remain unbound indefinitely if a matching volume does not exist. Claims will be bound as matching volumes become available. 

## (3)使用
- Pods use claims as volumes. The cluster inspects the claim to find the bound volume and mounts that volume for a Pod.

## (4)存储对象保护:
- 若用户删除一个正在被Pod使用的PVC, 则PVC不会立即被删除, 延缓至在Pod不使用后删除.
- 若管理员删除一个绑定到PVC的PV, PV不会立即被删除, 延缓至PV不再绑定到PVC再删除.
- PVC被保护: status是Terminating, Finalizer有kubernets.io/pvc-protection.
- PV被保护: status是Terminating, Finalizer有kubernets.io/pv-protection.

## (5)回收(reclaiming):
- 当用户用完卷时, 可通过API删除pvc对象来允许回收资源, pv的回收策略决定pv在被pvc释放后如何处理该pv, 策略包括: Retain,Recycle,Delete.
- Retain: 当PVC删除时, PV依旧存在, 该卷被认为是Released,但是不能被其它pvc使用, 因为还有前个pvc的数据存在.
- Delete: 对于支持Delete回收策略的volume plugin, 删除会删除pv对象同时存储资源也会被删除.
- Decycle: 已被废弃, 推荐使用动态provisioning.

## (6)expanding pvc:
- 部分volume类型支持expanding pvc.

# 六 VolumeAttachment
