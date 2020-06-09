# 一 概述:
## (1)概述:
- hostPath卷挂载host node文件系统上的一个文件或者目录到pod中.

## (2)HostPathVolumeSource属性:
- path: 必选.
- type(可选): DirectoryOrCreate, Directory, FileOrCreate, File, Socket, CharDevice, BlockDevice, 默认为空(表示在挂载卷前不会执行任何check).

## (3)使用场景:
- 某些系统级别的Pod(通常是DaemonSet)需要读取节点文件或者使用节点文件系统来访问设备.
- 在容器中运行cAdvisor,以hostPath方式挂载/sys.

## (4)谨慎使用:
- 和节点绑定, 当pod被调度到其它节点时就访问不到数据.

# 二 实现:
## (1)hostPathPlugin:
- 实现VolumePlugin,PersistentVolumePlugin等接口.
