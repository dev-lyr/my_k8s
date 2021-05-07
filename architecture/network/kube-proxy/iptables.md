# 一 概述:
## (1)概述:
- pkg/proxy/iptables

## (2)chain:
- KUBE-SERVICES
- KUBE-EXTERNAL-SERVICES
- KUBE-NODEPORTS
- KUBE-POSTROUTING
- KUBE-MARK-MASQ
- KUBE-MARK-DROP
- KUBE-FORWARD
- KUBE-PROXY-CANARY

# 二 Proxier:
## (1)概述

## (2)属性:
- iptablesData
- existingFilterChainsData
- filterChains
- filterRules
- natChains
- natRules

## (3)方法:
- syncProxyRules
