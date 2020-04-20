# Azure Test



- 部署cert-manager: `kubectl apply -f cert-manager.yaml`
- 部署cluster-api: `kubectl apply -f cluster-api-components.yaml`
- 部署Azure provider: `kubectl apply -f cluster-api-provider-azure-components.yaml`
- 部署master: `kubectl apply -f azure-cluster-master-xxx.yaml`
- 部署worker: `kubectl apply -f azure-cluster-worker-xxx.yaml`



### PS: 如果想自己测试，可以参考https://github.com/kubernetes-sigs/cluster-api-provider-azure下的Makefile create-cluster，env参数设置如下

```bash
# admin@az360.partner.onmschina.cn
export AZURE_SUBSCRIPTION_ID="593ea0a5-2089-4f6f-be30-ebe12fc78339"
export AZURE_TENANT_ID="c9572b54-e243-4caf-8684-cff70654c290"
export AZURE_CLIENT_ID="8736e6cc-dfd4-415f-98d4-c92f1de063f5"
export AZURE_CLIENT_SECRET="Jka@-iT8-H-yjyWvlzCOcbpo0huX36ns"
```

`templates/cluster-template.yaml 修改存储 diskSizeGB: 64`

```bash
# Cluster settings.
export CLUSTER_NAME="test"
export AZURE_VNET_NAME=${CLUSTER_NAME}-vnet

# Azure settings.
export AZURE_LOCATION="chinanorth"    # eastus chinanorth
export AZURE_RESOURCE_GROUP=${CLUSTER_NAME}
export AZURE_SUBSCRIPTION_ID_B64="$(echo -n "$AZURE_SUBSCRIPTION_ID" | base64 | tr -d '\n')"
export AZURE_TENANT_ID_B64="$(echo -n "$AZURE_TENANT_ID" | base64 | tr -d '\n')"
export AZURE_CLIENT_ID_B64="$(echo -n "$AZURE_CLIENT_ID" | base64 | tr -d '\n')"
export AZURE_CLIENT_SECRET_B64="$(echo -n "$AZURE_CLIENT_SECRET" | base64 | tr -d '\n')"

# Machine settings.
export CONTROL_PLANE_MACHINE_COUNT=1
export AZURE_CONTROL_PLANE_MACHINE_TYPE="Standard_A2"
export AZURE_NODE_MACHINE_TYPE="Standard_A2"
export WORKER_MACHINE_COUNT=1
export KUBERNETES_VERSION="v1.16.7"

# Generate SSH key.
# If you want to provide your own key, skip this step and set AZURE_SSH_PUBLIC_KEY to your existing file.
SSH_KEY_FILE=.sshkey
rm -f "${SSH_KEY_FILE}" 2>/dev/null
ssh-keygen -t rsa -b 2048 -f "${SSH_KEY_FILE}" -N '' 1>/dev/null
echo "Machine SSH key generated in ${SSH_KEY_FILE}"
export AZURE_SSH_PUBLIC_KEY=$(cat "${SSH_KEY_FILE}.pub" | base64 | tr -d '\r\n')
```



### 流程

- 集群中确认必要组件安装
  - cert-manager.yaml
  - cluster-api-components.yaml
  - 基础Azure components组件
- 租户创建集群
  - 部署租户相关Azure components组件，等待pod状态condition=Ready
  - 部署master相关crd (Cluster、AzureCluster、KubeadmControlPlane、AzureMachineTemplate)
    - cluster
      - cluster.spec.controlplaneEndpoint is not zero(v.Host == "" && v.Port == 0)
      - cluster.status.ready=true
      - 异常: cluster.status.failureDomains、FailureReason、FailureMessage
    - kubeadmcontrolplane
      - kubeadmcontrolplane.Status.Initialized = true 当新创建的集群可以从kube-system中获取kubeadm-config
      - kubeadmcontrolplane.status.ReadyReplicas 当新创建的集群中节点ready
      - kubeadmcontrolplane.Status.Ready = true 当ReadyReplicas>0
      - fail
        - FailureReason
        - FailureMessage
    - cluster.Status.ControlPlaneInitialized = true 当control-plane节点的machine.status.NodeRef != nil
    - cluster.Status.ControlPlaneReady = true
  - 部署worker相关crd (MachineDeployment、AzureMachineTemplate、KubeadmConfigTemplate)
  - 安装插件 (CNI)
    - kubeadmcontrolplane.status.Ready

machine
- machine.status.bootstrapReady
- machine.status.infrastructureReady
- machine.Status.NodeRef != nil 表示machine对应的node可以从kubectl get node获取到
- fail
  - machine.status.FailureReason
  - machine.status.FailureMessage



### 状态梳理

准备阶段：

1. kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.14.2/cert-manager.yaml

2. kubectl apply -f https://github.com/kubernetes-sigs/cluster-api/releases/download/v0.3.3/cluster-api-components.yaml

3. kubectl apply -f cluster-api-provider-azure-components-common.yaml



### 创建集群步骤：

1. kubectl apply -f cluster-api-provider-azure-components-tenant.yaml
    OK: pod ready
    namespace 每个租户绑定
    kubectl wait --for=condition=Ready --timeout=5m -n capz-system pod -l cluster.x-k8s.io/provider=infrastructure-azure

2. 创建cluster，kubectl apply -f azure-cluster-master-gzw-1.yaml
    OK: Cluster.Status.InfrastructureReady=true

3. 创建control-plane，kubectl apply -f azure-cluster-master-gzw-2.yaml
    OK: Cluster.Status.ControlPlaneInitialized=true
    多少个节点可用 根据control-plane角色的machine Status.NodeRef!=nil决定

4. 创建worker，kubectl apply -f azure-cluster-master-worker-gzw.yaml
    OK: MachineDeployment.Status
        ReadyReplicas
        AvailableReplicas

5. 网络插件
    OK: Cluster.Status.ControlPlaneReady KubeadmControlPlane.Status.Ready



Cluster Status:
- InfrastructureReady
    - 当AzureCluster.Status.Ready = true
- ControlPlaneInitialized
    - 当control-plane角色的machine Status.NodeRef != nil
- ControlPlaneReady
    - cluster.Spec.ControlPlaneRef != nil
    - 与KubeadmControlPlane.Status.Ready同步
- Phase
    - Pending 初始状态
    - Provisioning 当cluster对象provider infrastrure对象关联
    - Provisioned 当infrastructure创建并配置完毕
    - Deleting 当删除并且infrastructure还没有完全删除
    - Failed 可能需要人工介入
    - Unknown

Master

KubeadmControlPlane Status:
- Replicas
    - 从machine中获取匹配selector的non-terminated数量
- UpdatedReplicas
    - 根据machine对象中的label kubeadm.controlplane.cluster.x-k8s.io/hash，查看和KubeadmControlPlane.Spec hash值一样的machine数
- ReadyReplicas
    - 从workload cluster中获取control-plane的node对象
    - node.Status.Conditions包含type为ready，status为true
- UnavailableReplicas
    - replicas - readyReplicas
- Initialized
    - HasKubeadmConfig
    - 从workload cluster中获取kube-system/kubeadm-config的configmap对象，获取到即为true
    - 只会设置一次，不会变更
- Ready
    - KubeadmControlPlane.Status.ReadyReplicas > 0


Worker

MachineDeployment Status:
- Replicas
    - 从machine中获取匹配selector的non-terminated数量
- UpdatedReplicas
- ReadyReplicas
    - 获取对应的machine，根据machine获取node对象，判断node status.Conditions是否包含 ready=true
- AvailableReplicas
    - node对象要求ready status.Conditions包含ready=true
    - node Condition对象中的readyCondition LastTransitionTime=0 或者 minReadySeconds 已达
- UnavailableReplicas
- Phase
    - ScalingUp
    - ScalingDown
    - Running
    - Failed
    - Unknown

Machine Status:
- NodeRef
    - machine.Spec.ProviderID有值
    - 通过集群的kubeconfig获取client，查询新建集群中具有指定providerID的node对象，找到赋值给machine.Status.NodeRef
- BootstrapReady
    - machine.Spec.Bootstrap.DataSecretName != nil
- InfrastructureReady
    - 当AzureMachine.Status.Ready = true
- Phase
    - Pending
    - Provisioning
    - Provisioned
    - Running
    - Deleting
    - Deleted
    - Failed
    - Unknown
