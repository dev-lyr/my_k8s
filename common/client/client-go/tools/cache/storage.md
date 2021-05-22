# 一 概述:
## (1)概述:
- store.go
- index.go
- thread_safe_store.go

## (2)cache:
- cacheStorage ThreadSafeStore
- keyFunc KeyFunc
- 备注: 实现Indexer接口.

# 二 Store:
## (1)Store接口:
- Add
- Update
- Delete
- 等.

# 三 Indexer:
## (1)概述:
- Indexer extends Store with multiple indices and restricts each accumulator to simply hold the current object.

## (2)Indexer接口:
- 扩展Store接口.
- Index
- IndexKeys
- ListIndexFuncValues
- ByIndex
- GetIndexers

# 四 ThreadSafeStore:
## (1)概述:
-  allows concurrent indexed access to a storage backend.

## (2)ThreadSafeStore接口:
- Add
- Update
- Delete
- Index
- 等等.
