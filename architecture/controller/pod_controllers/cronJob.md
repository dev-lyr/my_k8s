# 一 概述:
## (1)功能:
- CronJob用于创建定期和循环的任务,例如:备份或发送邮件等.
- 一个CronJob对象类似crontab文件中的一行, 根据指定schedule(Cron形式)周期性运行job.

## (2)注意事项:
- cronjob资源的name需要是一个合法的dns subdomain name, 且长度不超过52字符,因为cronJob controller会自动给它创建的pod追加11个字符, 而job name的不能超过63字符.

# 二 CronJob资源
