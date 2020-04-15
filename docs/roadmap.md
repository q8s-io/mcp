# Docs For RoadMap

## v0.1.0
- 集群
  - 集群创建
    - attach cluster
  - 集群列表
  - 集群基础信息查看
    - name
    - create time
    - kubeconfig
    - k8s version, runtime
  - 集群删除
    - detach cluster

## v0.1.1
- 集群
  - 集群创建
    - create cluster
      - 不考虑plugin自动发布应用
      - 对接Azure cn
    - 集群基础信息查看
        - region
    - 集群删除

## v0.1.2
- 配置
  - 多租户
  - 用户管理
  - 基础计费
- 集群
  - 节点管理
    - label
    - annotation
    - taint
- 应用
  - helm chart仓库
  - helm client-go（helm的client-go，方法调用），部署plugin（公有云有可选插件列表，先适配网络，对接Azure、AWS），plugin部署yaml

## v0.1.3
- 资产
  - 管理
    - 节点
    - 网络
    - 存储
    - DNS、LB
- 集群
  - 状态维护
  - 集群信息补全
- 配置
  - 配置中心
  - 权限管理
