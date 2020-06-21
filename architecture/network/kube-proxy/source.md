# 一 概述:
## (1)相关目录:
- cmd/kube-proxy: kube-proxy命令相关.
- pkg/proxy

## (2)cmd/kube-proxy:
- ProxyServer结构
- Options结构

# 二 ProxyServer:
## (1)概述:
- 表示启动proxyServer所需要的所有参数.
- 相关: Options表示创建和运行proxyServer所需的所有信息.

## (2)方法:
- Run: 运行ProxyServer.
- CleanupAndExit: 删除iptables规则并退出.

## (3)属性:

# 三 Config:
## (1)EndpointsConfig:
- 属性: eventHandler []EndpointsHandler
- 方法: handleAddEndpoints, handleDeleteEndpoints, handleUpdateEndpoints, Run, RegisterEventHandler.
- 接口: EndpointsHandler.

## (2)ServiceConfig:
- eventHandler []ServiceHandler
- 方法: handleAddService, handleDeleteService, handleUpdateService, Run, RegisterEventHandler.
- 接口: ServiceHandler

## (3)NodeConfig:
- eventHandler []NodeHandler
- 方法: handleAddNode, handleDeleteNode, handleUpdateNode, Run, RegisterEventHandler.
- 接口: Nodehandler

# 四 iptables:
## (1)概述:
- Proxier: 基于iptables的后端代理, 实现EndpointsHandler,ServiceHandler等接口, 作为各种Config的eventHandler.
 
## (2)syncProxyRules
