# 一 概述:
## (1)概述:
- 语法: kubebuilder [command]

## (2)command:
- init: 初始化一个新项目.
- create: 创建一个api资源或webhook.
- edit: 编辑项目配置.
- help
- version

## (3)通常开发流程:
- 调用**kubebuilder init**初始化一个project.
- 调用**kubebuilder create**创建一个或多个新的API资源.
- 修改spec和status, 实现reconcile.

# 二 kubebuilder init:
## (1)概述:
- 语法: kubebuilder init [flag]
- 功能: 初始化一个新项目.

## (2)选项:
- --domain sting: groups的domain, 默认为"my.domain".
- --fetch-deps: 下载依赖, 默认为true.
- --repo: go module使用的名称(例如:github.com/xxx/repo), 默认为当前工作目录的go package.
- --license
- --owner
- --project-version: 默认为2.
-  --skip-go-version-check: 是否跳过go版本检查.

## (3)输出文件:
- a boilerplate license file
- a PROJECT file with the project configuration
- a Makefile to build the project
- a go.mod with project dependencies
- a Kustomization.yaml for customizating manifests
- a Patch file for customizing image for manager manifests
- a Patch file for enabling prometheus metrics
- a cmd/manager/main.go to run

# 三 kubebuilder create:
## (1)概述:
- 语法: kubebuilder create [command]; command: api或webhook.
- 可通过--resource和--controller选择或者生成过程的交互输入来选择是否只生成资源或控制器.

## (2)create api:
- --group: 资源组.
- --version: 资源version.
- --kind: 资源种类.
- --controller: 生成controller的代码,默认为true.
- --namespaced: 资源是不是namespaced, 默认为true.
- --make: 在生成代码后运行make.
- --resource: 生成资源, 默认为true.
- --force: 强制创建资源,即使存在.
- --example: 默认为true,生成reconcile样例代码.

## (3)create webhook:
- --group: 资源组.
- --version: 资源version.
- --kind: 资源种类.
- --programmatic-validation
- --conversion
- --defaulting
