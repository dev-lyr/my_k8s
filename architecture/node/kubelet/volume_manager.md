# 一 概述:
## (1)功能:
- VolumeManager runs a set of asynchronous loops that figure out which volumes need to be attached/mounted/unmounted/detached based on the pods scheduled on this node and makes it so.

## (2)备注:
- pkg/kubelet/volumemanager
- pkg/kubelet/volume_host.go
- pkg/volume

# 二 volumeManager:
## (1)概述:
- 在kubelet创建时创建和启动时候启动.
- 实现VolumeManager接口.

## (2)volumeManager属性:
- kubeClient: DesiredStateOfWorldPopulator用它和API服务器通信来fetch PV和PVC对象.
- volumePluginMgr
- desiredStateOfWorld
- desiredStateOfWorldPopulator
- actualStateOfWorld
- reconciler
- operationExecutor: 用来执行异步的attach, detach, mount和umount操作.
- csiMigratedPluginManager
- intreeToCSITranslator

## (3)Run方法:
- 若kubeClient不为nil, 则volumePluginMgr.Run.
- desiredStateOfWorldPopulator.Run
- reconciler.Run

## (4)WaitForAttachAndMount方法:
- kubelet的syncPod方法中调用.
- 代码写死等待2分3秒, 超时则返回失败, 等待重试.

# 三 desiredStateOfWorldPopulator:
## (1)概述:
- monitor and keep the states of the caches in sync with the "ground truth.
- 间隔100ms执行一次populatorLoop.
- 被reconciler使用.

## (2)populatorLoop:
- findAndAddNewPods: Iterate through all pods and add to desired state of world.
- findAndRemoveDeletedPods: calls out to the container runtime to determine if the containers for a given pod are terminated.

# 四 reconciler
## (1)功能:
- reconciler runs an asynchronous periodic loop to reconcile the desiredStateOfWorld with the actualStateOfWorld by triggering attach,detach, mount, and unmount operations using the operationExecutor.

## (2)reconcile:
- unmountVolumes
- mountAttachVolumes
- unmountDetachDevices
