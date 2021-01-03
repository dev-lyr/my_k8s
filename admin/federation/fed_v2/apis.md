# 一概述:
## (1)概述:
- 目录: kubefed/pkg/apis.
- 子目录: core, multiclusterdns和scheduling.

## (2)core:
- FederatedTypeConfig
- KubeFedCluster: 让kubeFed感知kubernete集群并封装与该集群通信的细节信息.
- KubeFedConfig
- FederatedServiceStatus
- PropagatedVersion
- ClusterPropagatedVersion

## (3)scheduling:
- ReplicaSchedulingPreference

## (4)multiclusterdns:
- Domain
- DNSEndpoint
- IngressDNSRecord
- ServiceDNSRecord
