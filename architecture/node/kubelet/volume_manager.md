# 一 概述:
## (1)功能:
- VolumeManager runs a set of asynchronous loops that figure out which volumes need to be attached/mounted/unmounted/detached based on the pods scheduled on this node and makes it so.

## (2)备注:
- pkg/kubelet/volumemanager

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

# 三 reconciler


# 四 desiredStateOfWorldPopulator
