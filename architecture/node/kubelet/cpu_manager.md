# 一 概述:
## (1)背景:
- By default, the kubelet uses **CFS quota** to enforce pod CPU limits.  When the node runs many CPU-bound pods, the workload can move to different CPU cores depending on whether the pod is throttled and which CPU cores are available at scheduling time. Many workloads are not sensitive to this migration and thus work fine without any intervention.
- However, in workloads where **CPU cache affinity** and **scheduling latency** significantly affect workload performance, the kubelet allows alternative CPU management policies to determine some placement preferences on the node.

## (2)排它(exclusive)使用条件:
- kubelet的--cpu-manager-policy设置为static
- Pod是Guaranteed
- Request为整数
- 备注: 底层基于cgroup的cpusets控制器实现.

## (3)kubelet相关配置:
- --cpu-manager-policy: none(默认)或static. 
- --cpu-manager-reconcile-period

## (4)备注:
- https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/
- pkg/kubelet/cm/cpumanager
