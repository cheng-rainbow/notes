`Kubernetes`（简称 K8s）是一个开源的**容器编排**平台，用于自动化部署、扩展和管理容器化应用。由 Google 设计并开源（基于 Borg 系统）。

- **核心功能**：
  - 自动化容器部署与生命周期管理。
  - 负载均衡与服务发现。
  - 自动扩缩容（水平扩展）。
  - 自愈能力（如重启故障容器、替换节点）。
  - 存储编排（动态挂载存储卷）。
  - 配置与密钥管理。



## 一、Kubernetes 架构

Kubernetes 集群由 **Master 节点（控制平面）** 和 **Worker 节点（工作节点）** 组成。

### 1. Master 节点组件

(1) **API Server**

- 作用：这是 Kubernetes 集群的核心组件，负责提供 RESTful API 接口，供用户、客户端以及其他组件与集群交互。
- 特点：唯一与 etcd 通信的组件，负责状态存储和同步。

(2) **etcd**

- 作用：一个高可用、分布式键值存储数据库，用于存储 Kubernetes 集群的所有状态数据，包括所有资源的定义
- 特点：
  - 主节点依赖 etcd 来持久化集群状态
  - 不直接与用户交互，仅供 kube-apiserver 使用。

(3) **Controller Manager**

- 作用：运行多个控制器进程，确保集群的实际状态与期望状态一致。
- 特点: 所有控制器作为一个单一进程运行，但逻辑上独立。

  - **Deployment Controller**：管理 **Deployment**，确保 Pod 数量符合期望值，当pod挂掉的时候再启动一个。

  - **ReplicaSet Controller**：管理 **ReplicaSet**，保证 Pod 副本数正确。

  - **Node Controller**：检测 Node 是否存活，宕机时迁移 Pod。

  - **Job Controller**：管理 **Job** 任务，确保一次性任务能成功完成。

(4) **Scheduler**

- 作用：负责将 Pod 调度到合适的子节点上运行
- 特点：
  - 不直接管理 Pod 的生命周期，仅负责调度决策。
  - 可通过自定义调度策略扩展。

### 2. Worker 节点组件

(1) **kubelet**

- 作用：管理节点上的 Pod 生命周期（如创建、销毁容器），并与 API Server 通信。
- 特点: 
  - 不负责调度，只执行和管理。
  - 如果 kubelet 宕机，该节点上的 Pod 无法正常运行。

(2) **kube-proxy**

- 作用：维护节点网络规则，实现 Service 的负载均衡和网络路由。

(3) **容器运行时（Container Runtime）**

- 作用：负责实际运行容器。



### 3. 举例

入口（API Server）

1. 你执行 `kubectl apply -f nginx-deployment.yaml`，这个用户请求会首先发送到**API Server（kube-apiserver）**。
2. API Server 负责**验证请求**，然后把 Nginx 的部署信息存入 **etcd**（分布式存储）。

**💡 API Server 相当于 Kubernetes 的“前台”，所有请求都要经过它。**

数据存储（etcd）

1. API Server 将 `nginx-deployment.yaml` 解析后，把**Nginx 需要的副本数（2 个 Pod）等信息存入 etcd**。
2. etcd 是 Kubernetes 的数据库，负责存储**所有集群的状态信息**。

**💡 etcd 相当于 Kubernetes 的“大脑”，记录所有资源信息。**

控制器（Controller Manager）

1. **Deployment Controller** 监听到 etcd 里有一个新的 Nginx Deployment 需求。
2. 它发现“当前 Nginx Pod 数量=0，期望=2”，所以决定**创建 2 个 Pod**。
3. 它向 API Server 发送请求，要求创建 2 个 Nginx Pod。

**💡 Controller Manager 相当于“管理员”，负责让集群符合预期状态。**

调度（Scheduler）

1. **Scheduler 负责决定 Nginx Pod 应该运行在哪个 Node（工作节点）上。**
2. 它会检查**每个 Node 的资源情况（CPU、内存）**，选择最合适的节点运行 Pod。
3. 选好后，它告诉 API Server：
   - **Pod 1 → Node A**
   - **Pod 2 → Node B**

**💡 Scheduler 相当于“调度员”，决定任务在哪里执行。**

运行（工作节点 - Kubelet & 容器运行）

1. **Node A 和 Node B** 各自的 **kubelet** 发现：“哎，API Server 让我运行一个 Nginx Pod”。
2. **kubelet** 让本节点的 **Container Runtime（比如 Docker 或 Containerd）** 去拉取 `nginx:latest` 镜像。
3. 镜像下载完成后，**Container Runtime 运行 Nginx 容器**。
4. Nginx Pod 启动成功后，**kubelet 会向 API Server 汇报状态**（"Nginx Pod 运行成功"）。

**💡 kubelet 相当于“工人”，它会执行调度员的命令，实际启动容器。**



### 4. 资源

**资源（Resource）** 指的是集群内可管理的对象，比如 **Pod、Service、Deployment** 等。这些资源描述了 K8s 中的各种组件，并且可以通过 **YAML/JSON** 配置、`kubectl` 命令管理。

#### 4.1 Pod（最小运行单元）

- **Pod 是 Kubernetes 中最小的部署单元**，可以包含一个或多个容器（通常是一个）。
- **所有容器共享同一个网络和存储**，可以通过 `localhost` 互相通信。
- **Pod 不是直接创建的**，通常由 **Deployment** 或 **ReplicaSet** 管理。

#### 4.2 Deployment（无状态应用管理）

- **管理 Pod 的副本数，保证高可用**。
- **支持滚动更新、回滚**。
- 通过 **ReplicaSet** 机制确保 Pod 持续运行。

#### 4.3 Service（网络暴露）

- **Pod 的 IP 会变化，Service 提供稳定的访问方式**。
- **三种类型**：
  - `ClusterIP`（默认）：仅集群内可访问。
  - `NodePort`：通过 `NodeIP:端口` 访问。
  - `LoadBalancer`：云环境下提供外部负载均衡。

#### 4.4 Ingress（域名管理）

- **负责将外部流量引导到 Service**。
- **支持基于域名、路径的路由**，可以实现多个服务共用一个 LoadBalancer。

#### 4.5 namespace

Namespace **用于隔离资源**，比如不同的项目、环境（开发/测试/生产）。

#### 4.6 ConfigMap（配置管理）

ConfigMap 用于存储应用程序的**配置信息**，如环境变量、配置文件等。



### 5. 管理资源的方式

在 Kubernetes (K8s) 中，**管理资源** 有多种方式，包括 **手动管理、YAML 配置文件、kubectl 命令、Kubernetes API 等

#### **1. 直接使用 `kubectl` 命令管理资源**

1. 创建资源

可以使用 `kubectl create` 创建资源：

```sh
kubectl create deployment nginx --image=nginx
```

💡 **这个方式适用于快速测试**，但无法保存资源的定义，无法追踪变化。

2. 查看资源

```sh
kubectl get pods         # 查看所有 Pod
kubectl get services     # 查看所有 Service
kubectl get deployments  # 查看 Deployment
kubectl get nodes        # 查看所有 Node
```

3. 修改资源

```sh
kubectl edit deployment nginx-deployment
```

💡 **优点**：可以在线修改资源，例如增加副本数、更新环境变量等。
 💡 **缺点**：不易追踪修改历史，手动操作容易出错。



#### **2. 使用 YAML 配置文件管理资源**

相比 `kubectl create` 命令，我们更推荐**使用 YAML 文件管理 Kubernetes 资源**：

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:latest
```

**部署资源**

```sh
kubectl apply -f nginx-deployment.yaml
```

**更新资源**

```sh
kubectl apply -f nginx-deployment.yaml
```

💡 **优点**：

- **可重复部署**，同样的 YAML 可以在不同环境复用。
- **易于版本控制**，可以在 Git 里追踪变更。
- **声明式管理**，保证资源状态符合定义，而不是靠手动操作。



#### 3. 使用 Kubernetes API 管理资源

K8s 提供了 **REST API**，可以通过 HTTP 请求管理资源：

```sh
curl -X GET https://<k8s-api-server>/api/v1/pods
```

或者用 Python 代码调用：

```python
from kubernetes import client, config

config.load_kube_config()
v1 = client.CoreV1Api()
pods = v1.list_pod_for_all_namespaces()
for pod in pods.items:
    print(pod.metadata.name)
```

💡 **优点**：

- **适合自动化管理**，可以写程序批量管理 K8s 资源。
- **适用于 CI/CD 流程**，可以和 GitLab、Jenkins 结合。



## 二、Kubernetes 核心资源对象

### 1. Pod

- **最小部署单元**：一个或多个共享网络和存储的容器（如主容器+Sidecar）。
- **特点**：
  - 每个 Pod 有唯一 IP 地址。
  - 生命周期短暂（可被调度到不同节点）。

### 2. Deployment（Pod资源控制器）

- **作用**：声明式管理 Pod 副本集（ReplicaSet），支持滚动更新和回滚。
- **典型场景**：无状态应用（如 Web 服务）。

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
```

### 3. Service（服务发现资源）

- **作用**：为一组 Pod 提供稳定的访问入口（IP/DNS），支持负载均衡。
- **类型**：
  - ClusterIP（默认，集群内部访问）。
  - NodePort（通过节点端口暴露服务）。
  - LoadBalancer（云提供商负载均衡器）。

### 4. ConfigMap & Secret（配置资源）

- **ConfigMap**：存储非敏感配置数据（如环境变量、配置文件）。
- **Secret**：存储敏感信息（如密码、Token），以 Base64 编码存储。

### 5. Volume（存储资源）

- **作用**：为 Pod 提供持久化存储。
- **类型**：
  - 临时卷（emptyDir）。
  - 持久卷（PersistentVolume, PV）与持久卷声明（PersistentVolumeClaim, PVC）。

### 6. Namespace（集群级别资源）

- **作用**：逻辑隔离集群资源（如将开发、测试、生产环境分开）。如果所有的 Pod、Service、ConfigMap 等都放在一起，资源管理就会变得非常混乱。Namespace 允许你把这些应用**划分到不同的空间**，让它们各自管理自己的资源，不会互相影响。



## 三、Kubernetes 网络模型

### 1. 基本要求

- 所有 Pod 可直接通信（无需 NAT）。
- 所有节点可与所有 Pod 通信。
- Pod 看到的自身 IP 与其他 Pod 看到的 IP 一致。

### 2. 网络实现（CNI 插件）

- **常见插件**：Calico、Flannel、Cilium、Weave Net。
- **功能**：管理 Pod 网络（IP 分配、路由规则）。

### 3. Service 网络

- **ClusterIP**：虚拟 IP，通过 kube-proxy 的 iptables/IPVS 规则转发到 Pod。



## 四、存储管理

### 1. 持久卷（PersistentVolume, PV）

- **作用**：集群级别的存储资源（如云盘、NFS）。

### 2. 持久卷声明（PersistentVolumeClaim, PVC）

- **作用**：用户对存储资源的请求，绑定到 PV。

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```



## 五、本地环境测试

`Kubernetes` 一般都运行在大规模的计算集群上，管理很严格，这就对我们个人来说造成了一定的障碍，在官网上推荐使用 `minikube` 来学习Kubernetes ，它们都可以在本机上运行完整的 Kubernetes 环境。

### 1. 安装 `minikube`

1. 下载最新Minikube二进制文件：

   ```bash
   curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
   ```

2. 安装到系统路径：

   ```bash
   sudo install minikube-linux-amd64 /usr/local/bin/minikube
   ```

3. 验证安装：

   ```bash
   minikube version	
   ```

4. 创建集群

   ```bash
   minikube start --driver=docker
   ```

5. 查看集群状态

   ```bash
   minikube status
   ```

`minikube` 只能够搭建 `Kubernetes` 环境，要操作 `Kubernetes`，还需要另一个专门的客户端工具`kubectl`。



### 2. 安装 `kubectl`

1. 下载最新kubectl：

   ```bash
   curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
   ```

2. 安装：

   ```bash
   sudo install kubectl /usr/local/bin/
   ```

3. 验证：

   ```bash
   kubectl version --client
   ```

4. 查看k8s集群

   ```bash
   kubectl get nodes
   ```





## 六、部署与管理应用

### 1. 使用 `kubectl` 命令

- 常用命令：

  ```bash
  kubectl apply -f deployment.yaml  # 部署资源
  kubectl get pods                 # 查看 Pod 状态
  kubectl logs <pod-name>         # 查看日志
  kubectl exec -it <pod-name> -- sh  # 进入容器
  ```

### 2. 滚动更新与回滚

- **滚动更新**：逐步替换旧版本 Pod，确保零停机。

  ```bash
  kubectl set image deployment/nginx-deployment nginx=nginx:1.16.1
  ```

- **回滚**：

  ```bash
  kubectl rollout undo deployment/nginx-deployment
  ```

### 3. 自动扩缩容（HPA）

- 基于 CPU/内存使用率自动调整 Pod 数量：

  ```bash
  kubectl autoscale deployment nginx-deployment --cpu-percent=50 --min=1 --max=10
  ```



## 七、安全与权限管理

### 1. RBAC（基于角色的访问控制）

- **角色（Role）**：定义对某命名空间资源的权限。
- **角色绑定（RoleBinding）**：将角色赋予用户或服务账号。

### 2. Service Account

- **作用**：为 Pod 提供身份认证，用于访问 Kubernetes API。

### 3. 网络策略（NetworkPolicy）

- **作用**：控制 Pod 之间的网络流量（如允许特定端口的通信）。



## 八、生态系统与工具

### 1. Helm

- **作用**：Kubernetes 包管理器，通过 Charts 管理应用部署。

### 2. Operator 模式

- **作用**：通过自定义资源（CRD）扩展 Kubernetes，自动化复杂应用管理（如数据库）。

### 3. 监控与日志

- **工具**：
  - Prometheus（监控）。
  - Grafana（可视化）。
  - EFK（Elasticsearch + Fluentd + Kibana，日志收集）。



## 九、实际应用场景

### 1. 微服务架构

- 通过 Service 和 Ingress 管理服务间通信和外部访问。

### 2. CI/CD 流水线

- 与 Jenkins、GitLab CI 等工具集成，实现自动化构建和部署。

### 3. 混合云与多集群管理

- 使用 Kubefed 或 Karmada 管理跨云、跨地域的集群。



## 十、最佳实践

1. **资源限制**：为容器设置 CPU/内存的 `requests` 和 `limits`。
2. **健康检查**：配置 `livenessProbe` 和 `readinessProbe`。
3. **标签管理**：使用标签（Labels）高效组织资源。
4. **灾备策略**：定期备份 etcd 数据。
