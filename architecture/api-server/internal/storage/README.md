# 一 概述:
## (1)概述:
- https://godoc.org/k8s.io/apiserver/pkg/storage: 提供数据库相关操作接口.

## (2)相关问题:
- 不传resourceVersion会直接调用底层存储进行查询, 量大时会击穿apiserver的cache.

## (3)相关配置:
- --storage-backend: 默认etcd3.

# 二 Interface接口:
## (1)概述:
- 提供对象marshaling/unmarshaling操作的公共接口, 隐藏底层存储相关细节.

## (2)CURD:
- Create
- Delete
- Watch
- WatchList
- Get
- GetToList
- Count
- GuaranttedUpdate

## (3)实现类:
- Cacher
- etcd3的store

# 三 storagebackend
