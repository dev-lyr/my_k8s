# 一 概述:
## (1)功能:
- Job创建一个或多个Pod来执行, 当pod成功运行结束时Job也完成.
- Job中Pod模板的RestartPolicy只能为Never或OnFailure, 防止容器正常结束时重启容器.

## (2)JobSpec:
- activeDeadlineSeconds: Job的超时时间, 超时系统会尝试终止pod并将Job标记为失败.
- backoffLimit: 指定Job被标记为失败之前重试的次数, 默认为6.
- completions: 希望Job中作业运行多少次.
- parallelism: 允许多少个Pod并行执行.
- template: pod template.
- ttlSecondsAfterFinished: 限制Job结束执行后的生存时间, 不设置则Job不会被自动删除, 若为0则Job执行完后立即被删除.

## (3)CronJobSpec:
- concurrentPolicy: 设置如何对待Job的并发执行,合法值: Allow(默认), Forbid(禁止, 若上次未完成则跳过), Replace(取消当前执行并开始新的执行).
- schedule: crom格式的表达式.
- suspend: 告诉控制器暂停后续的执行, 对已开始的没有影响.
- failedJobHistoryLimit: 保留的失败的Job数量.
- successJobHistoryLimit: 保留的成功的Job数量.
- startingDeadlineSeconds
