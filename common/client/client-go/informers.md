# 一 概述:

# 二 factory.go:
## (1)概述:

## (2)sharedInformerFactory(struct):
- client kubernertes.Interface
- namespace
- tweakListOptions
- lock
- defaultResync
- customResync
- informers 
- startedInformers
- Start方法: 开始初始化所有请求的informer, 调用所有informer的Run方法, 详情参考cache.md.

## (3)SharedInformerFactory(interface):
- internalinterfaces.SharedInformerFactory
- ForResource
- WaitForCacheSync
- Apps
- 等等.

## (4)函数:
- NewSharedInformerFactory: kube-controller-manager调用创建所有namespace使用的sharedInformerFactory.
- NewSharedInformerFactoryWithOptions: 使用可选的选项来构造一个sharedInformerFactory实例.

