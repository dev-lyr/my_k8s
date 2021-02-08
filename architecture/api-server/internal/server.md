# 一 概述:
## (1)概述:
- server包用来创建kubernetes-like API server command, 被kube-apiserver使用.

## (2)相关目录或文件:
- options: public flags and options used by a generic api server.
- config: 用来配置一个GenericApiServer.
- handler: API Server的http handlers.

# 二 config:
## (1)概述:
- Config结构
- CompletedConfig结构: New方法创建apiserver.

## (2)New:
- NewAPIServerHandler
- GenericAPIServer

## (3)DefaultBuildHandlerChain:
- 功能: http handler filters链.

