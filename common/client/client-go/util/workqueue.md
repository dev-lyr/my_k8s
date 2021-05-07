# 一 概述:
## (1)概述:
- 目录: client-go/util/workqueue
- k8s的controller使用的queue为: workqueue.DefaultControllerRateLimiter()返回的队列.

## (2)相关类型:
- Interface: 定义队列的基本操作, 内部实现为Type.
- DelayingInterface: 是一个可以延迟添加item的Interface的接口, 内部实现类型为delayingType.
- RateLimitingInterface: an interface that **rate limits** items being added to the queue.

# 二 Interface接口:
## (1)概述:
- 定义队列的基本操作, 被DelayingInteface和RateLimitingInterface接口扩展.
- 内部实现类型Type.

## (2)方法(以Type实现解读):
- Add: 若dirty有则直接返回, 没有则插入dirty, 接着判断若processing有则直接返回, 若processing没有则放入queue.
- Get: 从queue取, 放入processing并从dirty删除, 若没有元素则阻塞.
- Done: 做完, 从processing删除, 然后判断是否在dirty中, 若是重新放入queue.
- Len
- ShutDown
- ShuttingDown

## (3)Type类型:
- queue: 一个slice, 包含待处理的item, slice中的所有item都在dirty而不在processing中.
- dirty: 一个用map表示的set, 定义所有需要处理的item.
- processing: 一个用map表述的set, 表示正在并发处理的item, 会存在一些item正在被处理且也在dirty队列中, 此时在处理完后, 检查它是不是在dirty队列, 若是则把它放入queue.
- shuttingDown
- metrics
- unfinishedWorkUpdatePeriod
- clock

## (4)创建方法:
- New
- NewNamed

# 三 DelayingInterface接口:
## (1)概述:
- 提供**延时添加**功能的队列, 对Interface接口的扩展.
- 内部实现类型为delayType.

## (2)方法:
- Interface
- AddAfter: 在指定duration到达时将其添加到workqueue.

## (3)delayType
- Interface
- clock
- stopCh
- waitForAddCh
- metrics
- deprecatedMetrics

# 四 RateLimitingInterface:
## (1)概述:
- 提供**限流**功能的队列, 对DelayingInterface的扩展.
- 内部实现类型为rateLimitingType.

## (2)方法:
- DelayingInterface
- AddRateLimited: 在rate limiters觉得可以的时候将item添加到workqueue.
- Forget
- NumRequeues: 返回item重入队列的次数.

## (3)rateLimitingType:
- DelayingInterface
- rateLimiter(RateLimiter类型)

## (4)RateLimiter接口:
- When: When gets an item and gets to decide how long that item should wait.
- Forget: 从rate limiter的failures队列中删除.
- NumRequests: NumRequeues returns back how many failures the item has had.

## (5)RateLimiter类型:
- ItemExponentialFailureRateLimiter: 做了一个baseDelay的2重试次数次方倍数的limit, 当超过maxDelay时以maxDelay为准.
- ItemFastSlowRateLimiter: 先快速重试指定次数, 然后再慢速重试.
- BucketRateLimiter: 引用到一个golang.org/x/time/rate.
- MaxOfRateLimiter: 维护一个RateLimiter列表, 返回最坏情况.
- 备注: default_rate_limiters.go

# 五 metrics:

# 六 ParallelizeUntil:
## (1)概述:
- 调度器使用.
