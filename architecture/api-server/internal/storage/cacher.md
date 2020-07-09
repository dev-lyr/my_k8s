# 一 概述:
## (1)功能:
- 基于cache为给定资源提供list和watch服务,并且在后台基于底层存储的内容更新cache.
- 实现storage.Interface接口.

# 二 Cache结构体:
## (1)概述:
- 实现storage.Interface接口, 是底层存储(etcd3)的装饰器, 从cache读, 调用底层storage方法更新.
- 创建路径: StorageWithCacher->NewCacherFromConfig

## (2)属性:
- incoming chan watchCacheEvent: 进来的事件, 需dispatch给各个watcher.
- ready: 表示cache是否就绪, 访问cache之前需ready.
- storage: 底层存储接口, 通过storagebackend.Create方法创建.

## (3)方法

# 三 watchCache
