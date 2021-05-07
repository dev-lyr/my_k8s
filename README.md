# 一 概述:
## (1)功能:
- Kubernetes是一个开源的容器编排系统, 包括:自动化部署和扩容, 容器化应用的管理等.

## (2)交互方式:
- 命令行工具: kubectl.
- 直接通过Kubernetes API.

## (3)特性:
- **服务发现和负载平衡**: No need to modify your application to use an unfamiliar service discovery mechanism. Kubernetes gives containers their own IP addresses and a single DNS name for a set of containers, and can load-balance across them.
- **自动装箱(binpacking)**: Automatically places containers based on their resource requirements and other constraints, while not sacrificing availability. Mix critical and best-effort workloads in order to drive up utilization and save even more resources.
- **存储编排(orchestration)**: Automatically mount the storage system of your choice, whether from local storage, a public cloud provider such as GCP or AWS, or a network storage system such as NFS, iSCSI, Gluster, Ceph, Cinder, or Flocker.
- **自愈(Self-healing)**: Restarts containers that fail, replaces and reschedules containers when nodes die, kills containers that don’t respond to your user-defined health check, and doesn’t advertise them to clients until they are ready to serve.
- **Automated rollouts and rollbacks**: Kubernetes progressively rolls out changes to your application or its configuration, while monitoring application health to ensure it doesn’t kill all your instances at the same time. If something goes wrong, Kubernetes will rollback the change for you. Take advantage of a growing ecosystem of deployment solutions.
- **Secret and configuration management**: Deploy and update secrets and application configuration without rebuilding your image and without exposing secrets in your stack configuration.
- **Batch execution**: In addition to services, Kubernetes can manage your batch and CI workloads, replacing containers that fail, if desired.
- **Horizontal scaling**: Scale your application up and down with a simple command, with a UI, or automatically based on CPU usage.

## (4)接口标准:
- CRI(Container Runtime Interface): 容器运行时接口.
- CSI(Container Storage Interface): 容器存储接口.
- CNI(Container Network Interface): 容器网络接口.

## (5)备注:
- https://github.com/kubernetes/community: 各个模块.
- CCNF: https://www.cncf.io/.
- Knative: https://github.com/knative/
- k8s API: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/

# 二 相关生态:
## (1)knative

## (2)serverless

