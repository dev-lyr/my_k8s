# 一 workqueue:
## (1)概述:
- 目录: client-go/util/workqueue

## (2)相关类型:
- queue.go: Interface接口和Type类型(实现了Interface), Type是一个工作队列.
- delaying_queue.go: DelayingInteface是一个可以延迟添加item的Interface, 内部实现类型为delayingType.
- rate_limiting_queue.go: RateLimitingInterface是一个提供限速功能的队列.

## (3)Interface接口:
- Add: 若dirty有则直接返回, 没有则插入dirty, 接着判断若processing有则直接返回, 若processing没有则放入queue.
- Get: 从queue取, 放入processing并从dirty删除.
- Done: 做完, 从processing删除, 然后判断是否在dirty中, 若是重新放入queue.
- Len
- ShutDown
- ShuttingDown

## (4)Type类型:
- queue: 一个slice, 包含待处理的item, slice中的所有item都在dirty而不在processing中.
- dirty: 一个用map表示的set, 定义所有需要处理的item.
- processing: 一个用map表述的set, 表示正在并发处理的item, 会存在一些item正在被处理且也在dirty队列中, 此时在处理完后, 检查它是不是在dirty队列, 若是则把它放入queue.
- shuttingDown
- metrics
- unfinishedWorkUpdatePeriod
- clock

