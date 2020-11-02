# 一 概述:
## (1)功能:
- 一站式构造一个client, 可根据固定的配置或.kubeconfig文件或命令行参数或它们任意组合等.

# 二 client_config.go:
## (1)BuildConfigFromFlags:
- 根据masterUrl或者kubeconfig文件路径来够一个restclient.config.
- 若两者都没指定, 则使用restclient.InClusterConfig, 若InClusterConfig也失败, 则使用默认配置.
