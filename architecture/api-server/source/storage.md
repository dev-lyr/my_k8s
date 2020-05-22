# 一 概述:
## (1)概述:
- https://godoc.org/k8s.io/apiserver/pkg/storage: 提供数据库相关操作接口.

## (2)相关结构:
- storage.Interface: 提供对象操作的接口并隐藏底层存储相关细节.

## (3)相关问题:
- 不传resourceVersion会直接调用底层存储进行查询, 量大时会击穿apiserver的cache.

# 二 cacher:
## (1)功能:
- 基于cache为给定资源提供list和watch服务,并且在后台基于底层存储的内容更新cache.
- 实现storage.Interface接口.

## (2)Cache结构体:
- incoming chan watchCacheEvent: 进来的事件, 需dispatch给各个watcher.
- ready: 表示cache是否就绪, 访问cache之前需ready.
- storage: 底层存储接口.

## (3)cacheWatcher:
- 实现watch.Interface接口.

# 三 etcd3

# 四 storagebackend
