# 一 概述:
## (1)概述:
- client-go/util/flowcontrol

## (2)用处:
- kubelet的imageBackOff和backOff.

# 二 Backoff:
## (1)属性:
- sync.RWMutex
- Clock
- defaultDuration
- maxDuration
- perItemBackoff

## (2)方法:
- Next
- Get
- IsInBackOffSince
- IsInBackOffSinceUpdate
- Reset

# 三 RateLimiter
