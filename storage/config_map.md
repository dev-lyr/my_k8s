# 一 概述:
## (1)概述:
- configMap资源提供一种向Pod内注入配置数据的方式.
- 数据存储在ConfigMap对象中, 可通过configMap类型卷来引用, 被运行在Pod内的容器化应用消费.

## (2)使用方式:
- 通过环境变量形式传递给容器.
- 通过configMap卷.

## (3)属性:
- binaryData
- data

## (4)configMap和secret:
- 采用configMap来存储非敏感的文件配置数据.
- secret存储敏感的数据, 若配置文件同时包含敏感和不敏感数据, 则该文件应该被存储到secret中.
