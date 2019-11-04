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

## (3)SharedInformerFactory(interface):
- internalinterfaces.SharedInformerFactory
- ForResource
- WaitForCacheSync
- Apps
- 等等.

## (4)函数:
- NewSharedInformerFactoryWithOptions: 使用可选的选项来构造一个sharedInformerFactory实例.

