# 一 概述:
## (1)资源种类:
- **Workloads**: are objects you use to manage and run your containers on the cluster.
- **Discovery&LB**: resources are objects you use to "stitch" your workloads together into an externally accessible, load-balanced Service.
- **Config&Storage**: resources are objects you use to inject initialization data into your applications, and to persist data that is external to your container.
- **Cluster**: resources objects define how the cluster itself is configured; these are typically used only by cluster operators.
- **Metadata**: resources are objects you use to configure the behavior of other resources within the cluster, such as HorizontalPodAutoscaler for scaling workloads.

## (2)资源组成:
- **资源ObjectMeta**: 资源的metadata, 例如:name,type,api verison, annotation和lables等, 这些属性可能会被终端用户或者系统修改.
- **ResourceSpec**: 被用户定义用来描述系统的期望状态, 在创建对象或更新对象时填写.
- **ResourceStatus**: 报告系统当前状态, 被server填写, 通常情况下用户不需要修改这个.

## (3)相关操作:
- **创建**: 创建操作会在后端存储创建资源, 资源创建后系统会apply到期望状态.
- **更新**: 两种形式: Replace和Patch.
- **读取**: 三种形式: Get,List和Watch. Get: 根据name来获取指定资源对象; List: 返回在一个namespace内指定类型的所有资源对象; Watch: Watch will stream results for an object(s) as it is updated.
- **删除**
- **额外操作**: Rollback, Read/Write Scale, Read/Write Status.

## (4)备注:
- https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/
