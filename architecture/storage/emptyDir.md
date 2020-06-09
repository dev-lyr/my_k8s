# 一 概述:
## (1)概述:
- emptyDir当Pod被分配到一个Node时创建, 初始为空.
- 和Pod有同样的生命周期, 当Pod被从node上删除时, emptyDir中的数据也被删除.
- emptyDir卷可以存储在后端node上的任何媒介上, 例如: disk,SSD或网络存储, 也可以设置emptyDir.midium为Memory告诉k8s挂载一个tmpfs使用.

## (2)EmptyDirVolumeSource:
- medium: 目录后端存储介质的类型, 可选: 空字符串(默认)或Memory(tmpfs), 空字符串表示使用node默认的介质.
- sizeLimit: 默认为nil, 无限制.

## (3)使用场景:
- 暂存空间
- 为耗时较长计算任务提供检查点, 从而方便的从奔溃前状态恢复执行.
- holding files that a content-manager Container fetches while a webserver Container serves the data.

## (4)备注:
- 若使用内存, 则node重启后数据会丢失; 同时写入文件会计算到容器的内存使用中, 受容器内存限制约束.
