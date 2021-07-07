# 一 概述:
## (1)概述:
- pkg/controller/volume

## (2)其它:
- kubelet/volumemanager

# 二 persistentvolume控制器:
## (1)概述:
- 将pvc绑定到pv.

# 三 attachdetach控制器:
## (1)概述:
- 检查pod对应pv和node的挂载情况,判断volume需要attach/detach.

## (2)备注:
- 相关pkg/volume

# 四 pvcprotection控制器:
## (1)概述:
- 功能: removes PVCProtectionFinalizer from PVCs that are used by no pods.

# 五 pvprotection控制器:
## (1)概述:
- 功能: removes PVProtectionFinalizer from PVs that are not bound to PVCs.

# 六 expand控制器
