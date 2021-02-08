# 一 概述:
## (1)概述:
- https://github.com/kubernetes/apiserver/tree/master/pkg/admission
- 在endpoints的handler中直接调用.

## (2)相关接口:
- Interface: Handles(若admission controller可以处理给定操作(create,update,delete或connect)则返回true).
- MutationInterface: 扩展Interface并新增Admit方法.
- ValidationInterface: 扩展Interface并新增Validate方法.
- 备注:interfaces.go

## (3)plugins:
- Plugins结构: registry map[string]Factory.
- Register: 注册plugin, 被各插件调用.
- NewFromPlugins: 生成所有Plugin实例.

## (4)chain:
- 备注: chain.go

## (5)decorator:
- 备注:decorator.go

## (6)备注:
- apiserver/pkg/server/options/admission.go

# 二 webhook相关:
## (1)相关结构:
- WebhookAccessor: 为validating和mutating类型webhook提供公共接口, 实现类:validatingWebhookAccessor和mutatingWebhookAccessor.
- validatingWebhookConfigurationManager 
- mutatingWebhookConfigurationManager

## (2)WebhookAccessor:
- apiserver/pkg/admission/plugin/webhook/accessors.go

## (3)mutatingWebhookConfigurationManager:
- list/watch MutatingWebhookConfigurations资源的CUD并调用updateConfiguration进行更新.

## (4)validatingWebhookConfigurationManager:
- list/watch validatingWebhookConfigurations资源的CUD并调用updateConfiguration进行更新.

## (5)备注:
- apiserver/pkg/admission/plugin/webhook
