# 一 概述:
## (1)功能:
- k8s集群联邦(kubefed)允许用户通过在hosting集群中的单个API集来协调多个kubernetes集群的配置.

## (2)kubefed配置的信息类型:
- **Type配置**: 声明kubefed需处理的API类型.
- **Cluster配置**: 声明kubefed对应的目标(target)集群.

## (3)Type配置三个基础概念:
- **Templates**: define the representation of a **resource** common across clusters.
- **Placement**: defines which **clusters** the resource is intended to appear in.
- **Overrides**: define **per-cluster field-level** variation to apply to the template.

## (4)High-level API相关概念:
- **Status**: collects the status of resources distributed by KubeFed across all federated clusters.
- **Policy**: determines which subset of clusters a resource is allowed to be distributed to.
- **Scheduling**: refers to a decision-making capability that can decide how workloads should be spread across different clusters similar to how a human operator would.

## (5)相关概念:
- Host Cluster: 用来expose Kubefed API和kubefed控制面板的集群.
- Cluster Registration: 一个集群通过kubefedctl join到host cluster.
- ServiceDNSRecord: A resource that associates one or more Kubernetes Service resources and how to access the Service, with a scheme for constructing Domain Name System (DNS) resource records for the Service.
- IngressDNSRecord: A resource that associates one or more Kubernetes Ingress and how to access the Kubernetes Ingress resources, with a scheme for constructing Domain Name System (DNS) resource records for the Ingress.
- DNSEndpoint: 一个Endpoint资源的CR封装.
- Endpoint: 一个表示DNS resource record的资源.
- 备注: https://github.com/kubernetes-sigs/kubefed/blob/master/docs/concepts.md

## (6)kubefedctl

## (7)备注:
- https://github.com/kubernetes-sigs/kubefed
