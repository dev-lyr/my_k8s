# 一 概述:
## (1)功能:
- Kustomize introduces a **template-free** way to customize application configuration that simplifies the use of off-the-shelf applications.
- kustomize提供一个插件架构允许用户写他们自己的资源generators和transformers.

## (2)术语
- Kustomization: 可以指kustomization.yaml文件, 通常是一个包含kustomization.yaml以及其所引用文件的相对路径.
- Kustomization root: 包含kustomization.yaml文件的目录.
- Base: 会被其他kustomization引用的kustomization.
- Overlay: 是一个依赖其它kustomization的kustomization, 没有bases则overlay也无用.
- Plugin
- 备注: https://github.com/kubernetes-sigs/kustomize/blob/master/docs/glossary.md

## (3)使用方式:
- kustomize命令
- kubectl kustomize
- kubectl apply/create/patch/get -k(--kustomize).

## (4)备注:
- https://github.com/kubernetes-sigs/kustomize/tree/master/docs
- https://kustomize.io/

# 二 kustomization.yaml
## (1)四个类型:
- resources: 待定制的现存资源.
- generator: 将要创建的新资源.
- transformer: 对新旧资源进行处理的方式.
- meta: 对上面有影响的属性.

## (2)resources:
- resources(list): 包含k8s API对象的文件或者包含其它kustomization的目录.
- CRDS(list): CRD文件.

## (3)generators:
- configMapGenerator(list): 列表中每项会导致创建一个ConfigMap资源.
- secretGenerator(list): 列表中每项会导致创建一个secret资源.
- generatorOptions(string): 改变所有ConfigMap和Secret generator的行为.
- generators(list): 插件配置文件.

## (4)transformers:
- commonLabels: 添加labels和一些响应label选择器到所有资源.
- commonAnnotations: 添加注解到所有资源.
- 等等.

## (5)meta:
- vars: Vars capture text from one resource's field and insert that text elsewhere.
- apiVersion
- kind

# 三 使用步骤:
## (1)构造kustomization文件:
- 根据(二)构造kustomization文件.

## (2)使用overlays:
- overlay是一个依赖其它kustomization的kustomization, overlay依赖的kustomization被称作bases(通过文件路径, URL或其它方法).
- 若没有bases, 则overlays也不起作用, 需在overlay的kustomization.yaml指定bases, bases在v2.1.0中被废弃, 将相关选项移到resources字段,
- 一个overlay也可以作为其它overlay的bases.

# 四 插件:
## (1)概述:
- 备注: https://github.com/kubernetes-sigs/kustomize/tree/master/docs/plugins

# 五 工作流:
## (1)概述:
- 备注: https://github.com/kubernetes-sigs/kustomize/blob/master/docs/workflows.md
