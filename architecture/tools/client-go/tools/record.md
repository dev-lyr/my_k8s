# 一 概述:
## (1)功能:

## (2)接口:
- EventRecorder: 记录事件.
- EventBroadcaster: 接收事件并发送事件给事件接收器,watcher或日志.
- EventSink: 事件接收器.

## (3)k8s Event资源(core/v1):
- action
- apiVersion
- count: 事件发生的次数.
- eventTime: 事件首次被observed到的时间.
- firstTimestamp: 事件首次recorded的时间.
- involvedObject
- lastTimestamp
- message: 可读的操作的状态的描述.
- **reason**: 简短可读描述状态改变的原因.
- related
- **sources**: 上报事件的模块.
- **type**: 事件类型, 目前有Normal和Warning.

# 二 EventBroadcaster:
## (1)概述:
- 接收事件并将他们发送到任意EventSink,watcher或log.

## (2)方法:
- StartEventWatcher: 将接收到事件发送给指定事件处理器.
- StartRecordingToSink: 将接收到事件发送给指定事件接收器.
- StartLogging: 将接收到事件发送给指定日志函数.
- NewRecorder: 生产一个可以将指定事件源事件发生给该EventBroadcaster的EventRecorder.

## (3)eventBroadcasterImpl:
- watch.Broadcaster指针
- sleepDuration
- options CorrelatorOptions

# 三 EventRecorder:
## (1)概述:
- 代表EventSource来记录事件.

## (2)方法:
- Event: 根据参数信息构造一个event并放到发送队列.
- Eventf
- PastEventf
- AnnotatedEventf

## (3)recorderImpl:
- scheme
- source v1.EventSource
- watch.Broadcaster指针
- clock

# 四 Broadcaster
## (1)概述:
- https://godoc.org/k8s.io/apimachinery/pkg/watch
