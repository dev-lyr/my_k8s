# 一 概述:
## (1)Horizontal Pod Autoscaler:
- 横向Pod自动伸缩通过Horizontal控制器来对Pod控制器(replicaSet等)管理的pod的副本数据进行自动伸缩.
- 通过创建**HorizontalpodAutoscaler(HPA)资源**来启动和配置Horizontal控制器.
- Horizontal控制器**周期**性检查Pod度量, 计算满足HPA资源所配置目标数量所需的副本数量, 从而调整对应Pod控制器的replicas字段.
