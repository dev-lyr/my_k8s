# 一 概述:
## (1)功能:
- Job创建一个或多个Pod来执行并确保它们中的指定数量成功终止, 当pod成功运行结束时Job也完成.
- 删除od会清除它创建的pod.
- Job中Pod模板的RestartPolicy只能为Never或OnFailure, 防止容器正常结束时重启容器.

## (2)JobSpec:
- activeDeadlineSeconds: Job的超时时间, 超时系统会尝试终止pod并将Job标记为失败.
- backoffLimit: 指定Job被标记为失败之前重试的次数, 默认为6.
- completions: 希望Job中作业运行多少次.
- parallelism: 允许多少个Pod并行执行.
- template: pod template.
- ttlSecondsAfterFinished: 限制Job结束执行后的生存时间, 不设置则Job不会被自动删除, 若为0则Job执行完后立即被删除.
