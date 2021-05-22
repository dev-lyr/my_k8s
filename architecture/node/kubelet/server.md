# 一 概述:
## (1)概述:
- ListenAndServe(Http, 依赖enable-serve(默认为true)).
- ListenAndServeReadOnly(Http, 需readOnlyPort>0).
- ListenAndServePodResources(grpc, 依赖KubeletPodResources特性开启(默认为true)).

## (2)对外提供API:
- /pods
- /metrics
- /metrics/cadvisor
- 等等

## (3)备注:
- pkg/kubelet/server

# 二 Server(struct):
## (1)概述:
- 一个Http handler, 用来通过http暴露kubelet的功能.

## (2)属性:
- auth AuthInterface
- host HostInterface
- restfulCont containerInterface
- metricsBuckets sets.String
- metricsMethodBuckets sets.String
- resourceAnalyzer stats.ResourceAnalyzer

## (3)方法:
- InstallAuthFilter
- InstallDefaultHandlers
- InstallDebuggingHandlers
- InstallSystemLogHandler
- InstallProfilingHandler
- InstallDebugFlagsHandler

## (4)HostInterface:
- GetRunningPods
- RunInContainer
- ServeLogs
- GetKubeletContainerLogs
- GetExec
- GetAttach

## (5)备注:
- --enable-debugging-handlers(默认为true).

