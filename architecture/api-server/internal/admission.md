# 一 概述:
## (1)概述:
- https://github.com/kubernetes/apiserver/tree/master/pkg/admission

## (2)相关接口:
- Interface: Handles(若admission controller可以处理给定操作(create,update,delete或connect)则返回true).
- MutationInterface: 扩展Interface并新增Admit方法.
- ValidationInterface: 扩展Interface并新增Validate方法.
- 备注:interfaces.go

## (3)chain:
- 备注: chain.go

## (4)decorator:
- 备注:decorator.go

# 二 admission controller类型:
## (1)概述:
- --admission-control-config-file
- --enable-admission-plugins stringSlice: 开启admission controller.
- --disable-admission-plugins stringSlice: 关闭admission controller.
- 查询默认开启的admission controller: kube-apiserver -h | grep enable-admission-plugins

## (2)备注:
- AlwaysPullImages
- EventRateLimit
- ExtendedResourceToleration
- MutatingAdmissionWebhook
- ValidatingAdmissionWebhook
- 等等.

# 三 webhook相关:
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
