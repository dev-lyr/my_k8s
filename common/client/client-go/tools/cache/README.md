# 一 概述:
## (1)概述:
- 代码目录: client-go/tools/cache. 

# 二 controller.go:
## (1)概述:
- controller
- Config
- ResourceEventHandler

## (2)函数:
- NewInformer: 返回一个用于Get/List的Store和一个Controller.
- NewIndexerInformer: 返回一个用于Get/List的Index和一个Controller.

## (3)Controller:
- Controller(Interface): Run, HasSynced, LastSyncResourceVersions等方法.
- controller(struct): config(Config), **reflector**, reflectorMutex, clock等属性; 实现Controller接口; 提供processLoop方法.

## (4)Config:
- Queue
- ListerWatcher
- Process ProcessFunc
- ObjectType runtime.Object
- FullResyncPeriod time.Duration
- ShouldResync ShouldResyncFunc
- RetryOnError bool

## (5)ResourceEventHandler:
- ResourceEventHandler(interface): OnAdd, OnUpdate和OnDelete.
- ResourceEventHandlerFunc: 实现ResourceEventHandler接口.
- FilteringResourceEventHandler: 为ResourceEventHandler增加filter功能.

# 三 Reflector.go:
## (1)概述:
- Reflector watches a specified resource and causes all changes to be reflected in the given store.
- 属于controller struct的属性.

## (2)Relector属性:
- name
- metrics
- expectedType reflect.Type
- store Store
- listerWatcher ListerWatcher
- period time.Duration
- resyncPeriod time.Duration
- clock
- lastSyncResourceVersion
- lastSyncResourceVersionMutex
- WatchListPageSize

## (3)Reflector方法:
- Run: 启动一个watch并处理watch事件, 主要是调用ListAndWatch来实现.
- ListAndWatch: 首先list所有items并获得当前的resource version, 随后使用该version来watch.

# 四 shared_informer.go:
## (1)概述:
- SharedInformer(interface): 有一个共享数据cache并且可能将cache的changes通知多个通过AddEventHandler注册的listeners, 当接收到通知的时候, cache的内容可能已经改变, 例如: 首先一个create,再接着一个delete, 则cache中该对象可能已经不存在了.
- SharedIndexInformer(interface):扩展SharedInformer, 增加AddIndexers和GetIndexer方法.

## (2)SharedInformer(interface):
- AddEventHandler
- AddEventHandlerWithResyncPeriod
- GetStore
- GetController
- Run
- HasSynced
- LastSyncResourceVersion

## (3)sharedIndexInformer(struct):
- indexer Indexer(扩展Store)
- controller Controller
- processor `*sharedProcessor`
- cacheMutationDetector
- listerWatcher
- objectType
- resyncCheckPeriod
- defaultEventHandlerResyncPeriod
- 等等.

## (4)sharedProcessor(struct):
- listenersStarted bool
- listenersLock
- listeners: processorListener指针数组
- syncingListeners
- clock
- wg

## (5)processorListener:
- nextCh
- addCh
- handler ResourceEventHandler

## (6)函数:
- WaitForCacheSync: 等待cache构造(populate)完成.
- NewSharedIndexInformer
- NewIndexInformer

