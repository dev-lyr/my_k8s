# 一 概述:
## (1)概述:
- virtual-kubelet是一个开源kubernetes kubelet实现, 伪装(masquerade)成kubelet用于连接kubernetes到其它APIs.
- 允许nodes的backend为其它服务, 例如: aws fargate, aliyun eci等.
- virtual kubelet是个插件式架构, 直接使用kubernetes原语, 插件式provider接口允许开发者实现kubelet定义的actions.
- virtual kubelet专注于提供一个库, 用户构建自定义kubernetes node agent中使用.

## (2)可用的provider:
- aliyun eci provider
- aws fargate provider

