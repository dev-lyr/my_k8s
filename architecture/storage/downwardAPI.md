# 一 概述:
## (1)使用场景:
- 使用configMap和secret卷应用传递Pod调度, 运行前的数据是可行的, 但对于不能预先知道的数据, 比如: Pod的Ip,主机名或Pod自身的名字名称等, 则需要使用downward API来解决.
- downward API允许用户通过**环境变量**或**文件(downward API卷)**来将Pod和容器的元数据传递给它们内部运行的进程.
- 详细: https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/

## (2)downward API可传递的数据:
- pod的名称,ip,所在namespace,运行节点的名称,所归属服务账号的名称.
- 每个容器请求的CPU和内存使用量, CPU和内存的限制.
- Pod标签和注解.
- 备注: Pod标签和注解只能通过downward API**卷**来暴露, 其它的可以使用环境变量或downward API卷来暴露.

## (3)DownwardAPIVolumeSource:
- defaultMode: 创建文件使用的mode bits, 默认为0644, 可选范围:0-0777.
- items(DownwardAPIVolumeFile数组): downward API volume文件的数组.

## (4)DownwardAPIVolumeFile:
- fieldRef: 选择pod的一个field, 当前支持: 注解,labels, name和namespace.
- mode: 该文件上使用的mode bits, 若不指定则使用volume上defaultMode.
- path: 需要创建文件的相对路径, 不能是绝对路径也不能包含..路径.
- resourceFiledRef(ResourceFiledSelector): 选择容器的一个资源, 当前只支持resource limit和requests.

## (5)备注:
- https://kubernetes.io/docs/tasks/inject-data-application/downward-api-volume-expose-pod-information/

