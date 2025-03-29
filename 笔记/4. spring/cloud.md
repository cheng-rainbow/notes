Spring Cloud 是一套基于 Spring Boot 的框架集合，用于构建分布式微服务架构。它提供了一系列工具和库，帮助开发者更轻松地管理分布式系统中的关键问题，比如服务注册与发现、负载均衡、分布式配置管理、熔断与降级、链路追踪等。

下图展示了微服务架构中每个主要功能模块的常用解决方案。
![在这里插入图片描述](./../../../笔记/笔记图片/ae970a9b6b8441d1b51f707fd30a5237.png)

## 一、相关功能的介绍

1. `服务注册与发现`

**服务注册**：服务注册与发现用于让各个服务在启动时自动注册到一个中央注册中心（如 Nacos、Eureka），并且能让其他服务通过注册中心找到并调用它们的地址。  
**发现**：每个服务启动后会将自身的地址和端口信息注册到注册中心；其他服务要调用它时，通过注册中心获取服务实例的地址，而**不需要固定的地址**。

2. `分布式配置管理`

分布式配置管理用于集中管理各服务的配置文件，支持动态更新，不需要重启服务。  可以在配置更新后自动推送至各服务节点，使它们能实时更新配置信息，提升了系统的灵活性和一致性。

3. `服务调用和负载均衡`

**服务调用**：服务之间的通信方式，可以通过 HTTP（如 RESTful API）或 RPC（远程过程调用）进行服务之间的请求。  
**负载均衡**：在微服务架构中，通常会有多个相同的服务实例分布在不同的服务器上。负载均衡用于在多个实例间分配请求，常见的策略有轮询、随机、最小连接数等，从而提升系统的处理能力和容错性。

3. `分布式事务`

分布式事务用于保证多个服务在处理同一个业务操作时的一致性。例如，用户下单时，需要支付服务和库存服务同时完成，如果某一方失败，整个操作需要回滚。  

4. `服务熔断和降级`

**服务熔断**：用于防止一个服务的故障传导到其他服务。如果某个服务在短时间内出现大量的错误或响应缓慢，熔断机制会自动切断对该服务的调用，避免对系统造成更大影响。  
**服务降级**：在服务出现问题时，提供降级策略，比如返回默认值或简化响应内容，使系统能够在部分服务不可用的情况下继续运行。

5. `服务链路追踪`

服务链路追踪用于跟踪分布式系统中一次请求的完整路径，分析其跨多个服务的执行情况，方便发现延迟或错误。  

6. `服务网关`

服务网关作为服务的统一入口，处理所有外部请求，提供认证授权、负载均衡、路由分发、监控等功能。它还能对请求进行限流、熔断、降级等保护。  



## 二、引入springcloud依赖

### 1. dependencyManagement

`dependencyManagement` 是 Maven 构建工具中的一个元素，用于定义项目中依赖的管理方式。


1. **统一依赖版本管理：**  
   `dependencyManagement` 能够帮助你统一管理多个模块中某个依赖的版本。例如，你可以在一个中央位置（通常是父 POM 文件）声明 Spring Boot 的版本号，这样所有子项目都会使用这个版本，而不需要每个项目中都定义。

2. **简化子项目中的依赖声明：**  
   子项目无需在每个 `pom.xml` 文件中声明依赖的版本号，只需要定义依赖的 `groupId` 和 `artifactId`，版本号将从 `dependencyManagement` 中继承。

根项目 pom 文件
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>dubbo-parent</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <packaging>pom</packaging>

    <properties>
        <java.version>17</java.version>
        <spring-boot.version>3.2.5</spring-boot.version>
        <spring-cloud.version>2023.0.3</spring-cloud.version>
        <spring-cloud-alibaba.version>2023.0.1.3</spring-cloud-alibaba.version>
    </properties>

    <!-- 添加 dependencyManagement 并导入依赖 -->
    <dependencyManagement>
        <dependencies>
            <!-- Spring Boot 版本管理 -->
            <dependency>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-dependencies</artifactId>
                <version>${spring-boot.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            <!-- Spring Cloud 版本管理 -->
            <dependency>
                <groupId>org.springframework.cloud</groupId>
                <artifactId>spring-cloud-dependencies</artifactId>
                <version>${spring-cloud.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            <!-- Spring Cloud Alibaba 版本管理 -->
            <dependency>
                <groupId>com.alibaba.cloud</groupId>
                <artifactId>spring-cloud-alibaba-dependencies</artifactId>
                <version>${spring-cloud-alibaba.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
    </dependencyManagement>

    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>

    <modules>
        <module>Consumer</module>
    </modules>
</project>
```

子项目 pom 文件

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <parent>
        <groupId>com.example</groupId>
        <artifactId>dubbo-parent</artifactId>
        <version>1.0.0-SNAPSHOT</version>
        <relativePath>../pom.xml</relativePath> <!-- 指向父模块的POM文件 -->
    </parent>

    <artifactId>Proverder</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>Proverder</name>
    <description>Proverder</description>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter</artifactId>
        </dependency>
		<dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>
    </dependencies>
</project>
```

[maven官网](https://mvnrepository.com/)

- `Spring Cloud Dependencies`
![在这里插入图片描述](./../../../笔记/笔记图片/ec1575eac380470db52dc41f2a0c6698.png)

- `spring-cloud-alibaba-dependencies`
![在这里插入图片描述](./../../../笔记/笔记图片/b193dc266b6b4f19a7b015d75d952af4.png)

### 2. 版本兼容性问题

- 参考文章 [地址](https://github.com/alibaba/spring-cloud-alibaba/blob/2023.x/README-zh.md)
![在这里插入图片描述](./../../../笔记/笔记图片/3270c7b52d494b38858c3a55e3594378.png)
这里我引入的是 springcloud 2023 , springcloudalibaba 2023, springboot 3.2, jdk 17



## 三、服务注册与发现 nacos

### 1. 配置

1. 添加依赖到子模块

在需要被nacos注册的模块中加入下面配置，
```xml
<!-- Nacos 服务注册和发现 -->
 <dependency>
     <groupId>com.alibaba.cloud</groupId>
     <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
 </dependency>
```

2. 简单配置

```properties
spring.cloud.nacos.discovery.server-addr=localhost:8848
spring.cloud.nacos.discovery.namespace=命名空间id # 指定命令空间, 不同空间下的实例不可互相访问
spring.cloud.nacos.discovery.group=DEFAULT_GROUP # 默认是DEFAULT_GROUP，指定group，不同group下的实例不可互相访问

# 下面是其他常见配置
spring.cloud.nacos.discovery.cluster-name=BeiJing # 指定当前实例是哪个集群，一般按照地区划分，讲请求发送距离近的实例
spring.cloud.loadbalancer.nacos.enabled=true # 开启 优先向跟当前发送请求的实例 在同一集群的实例发送请求
spring.cloud.nacos.discovery.weight=1 # 当前实例的权值，权值在1-100，默认是1，权值越大越容易接收请求，一般给配置高的服务器权值高一些
```

3. 在启动类上加上`@EnableDiscoveryClient`注解
   `@EnableDiscoveryClient` 是一个注解，启动项目即可在 `localhost:8848/nacos` 中查看到已经被注册到中央注册中心



### 2. 以单例启动

解压进入nacos的 bin目录，以单例模式启动（通过docker）
```bash
docker run -d --name nacos \
    -p 8848:8848 \
    -p 9848:9848 \
    -p 9849:9849 \
    -p 7848:7848 \
    -e MODE=standalone \
    nacos/nacos-server:v2.3.2
```



### 3. Nacos 集群架构

对于 Nacos 集群，主要的作用是 **实现高可用和数据一致性**，保证服务注册和配置管理的可靠性。

-  **集群架构**
   Nacos 集群通常包含多个节点，部署在不同的机器或虚拟机上，以提供服务注册、配置管理的冗余和高可用性。当一个节点发生故障时，其他节点可以继续提供服务，从而保证系统的稳定运行。

-  **数据一致性（RAFT 协议）**
   Nacos 集群内部使用 **RAFT 协议** 来实现服务数据的强一致性。这种一致性保证了同一服务的多个实例在集群中都能被正确地注册和发现。（在集群中的任意一个nacos中注册，即可被整个集群的nacos访问）
   - **Leader 选举**：在 Nacos 集群中，一个节点会被选举为 Leader，其它节点作为 Follower。Leader 负责处理写请求并同步数据到 Follower 节点。
   - **数据同步**：当服务实例的注册、注销或配置变更等写请求发生时，Leader 会将数据同步到所有 Follower，确保数据在集群中的一致性。
   
- **高可用性和故障恢复**
   - 当 Leader 节点故障时，集群中的其他节点会重新选举一个新的 Leader，继续处理写请求和同步数据。
   - Follower 节点接收 Leader 的数据更新请求并定期与 Leader 进行心跳检测，以保证集群稳定运行。

可以通过nginx反向代理，实现只暴漏一个nacos服务地址，nginx内容实现负载均衡
也可以通过loadbalancer或是在 application 中添加集群的所有地址实现简单的负载均衡

```java
# 会选其中一个地址注册服务
spring.cloud.nacos.discovery.server-addr: 172.20.10.2:8870,172.20.10.2:8860,172.20.10.2:8848
```





## 四、分布式配置管理 nacos

分布式配置管理功能的主要作用是在不同的服务之间**集中管理和统一分发配置**。这使得系统在配置变更时无需重启服务，可以实时更新配置，从而达到快速响应的效果。

### 1. 基本概念

- **Data ID（数据 ID）**：表示每个配置的唯一标识。在 Nacos 中，一个配置项通常用 Data ID 表示，通常为字符串形式，代表唯一的配置文件名。

- **Group（组）**：用于将不同的配置项进行分组管理，方便区分开发、测试、生产环境等场景。

- **Namespace（命名空间）**：用于逻辑隔离配置数据。不同命名空间内的配置是互相隔离的，这在多租户场景中非常有用。

- **配置项**：每个具体的配置信息称为配置项，可以是一个或多个键值对。



### 2. 引入 Nacos 配置管理

1. **引入依赖**

   ```xml
   <dependency>
       <groupId>org.springframework.cloud</groupId>
       <artifactId>spring-cloud-starter-bootstrap</artifactId>
   </dependency>
   <!--这里需要指定一个版本，我用默认的不行, 无法从配置中心获取配置-->
   <dependency>
       <groupId>com.alibaba.cloud</groupId>
       <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
       <version>2023.0.1.2</version>
   </dependency>
   ```

   创建一个 bootstrap.yaml 的文件

   ```yaml
   spring:
     application:
       name: module1
     profiles:
       active: dev
   
     cloud:
       nacos:
         config:
           server-addr: 192.168.227.128:8848 # 服务地址
           file-extension: yaml
           namespace: public
           group: DEFAULT_GROUP
           name: module1-dev	# 手动指定配置文件名，module1-dev.yaml
   ```

   data_id 一般命名采用 `application.name`-`profiles.active`.`filex-extension`，根据上面的配置，我的dataid就是 module1-dev.yaml
   ![在这里插入图片描述](./../../../笔记/笔记图片/7270be8d65044c86aaa0ca506ecf282d.png)

2. **获取配置**
   可以使用 Spring Boot 的 `@Value` 注解来获取 Nacos 中的配置项，使用 `@RefreshScope` 注解，自动刷新配置：(当我们配置中心修改时，不需要重启项目，hello 就会自动更新)

   ```java
   @RestController
   @RefreshScope
   public class TestController {
   
       @Value("${hello}")
       private String hello;
   
       @GetMapping("/nacos/config")
       public String nacosConfig() {
           return hello;
       }
   }
   ```

 **Nacos 配置中心优先级高于本地配置，最终生效的是 Nacos 的配置**。



## 五、服务调用和负载均衡 LoadBalancer

`Spring Cloud LoadBalancer` 是 Spring Cloud 中的一个客户端负载均衡模块，用于在服务调用者和多个实例之间分配流量。它通过服务发现（比如使用 Nacos）获取可用服务实例的列表，并根据不同的负载均衡策略（如轮询、随机等）选择一个实例进行请求分发。

### 1. 配置环境

在`调用者`项目中添加  `Spring Cloud LoadBalancer`

```xml
<dependency>
    <groupId>org.springframework.cloud</groupId>
    <artifactId>spring-cloud-starter-loadbalancer</artifactId>
</dependency>
```

配置默认策略（轮询）
```xml
spring.cloud.loadbalancer.configurations=default
```

### 2. 负载均衡的使用方式

Spring Cloud LoadBalancer 支持使用 `RestTemplate` 、`WebClient` 、`OpenFeign ` 进行负载均衡。这里我们使用的是RestTemplate

1. **定义 RestTemplate Bean** 并标注 `@LoadBalanced` 注解：
   
    ```java
    @Configuration
    public class AppConfig {
        
        // 我们同一个服务名称一般会有多个实例, 通过带有 `@LoadBalanced`注解的 `RestTemplate` 可以实现负载均衡, 让请求根据我们的配置分别发送到不同的实例
        @Bean
        @LoadBalanced  // 启用 RestTemplate 的负载均衡
        public RestTemplate restTemplate() {
            return new RestTemplate();
        }
    }
    ```
    
2. **发起请求**
   使用 `@LoadBalanced` 的 `RestTemplate` 时，可以直接通过服务名称调用服务，而不需要手动指定服务的 IP 地址和端口，避免了ip和端口写死, 只向一个实例发送请求的情况。
   
    ```java
    @Autowired
    private RestTemplate restTemplate;
   
    public String callService() {
        // 使用服务名代替实际地址
        return restTemplate.getForObject("http://module2/api/v1/data", String.class);
    }
    ```



### 3. 使用不同的负载均衡器

上面的配置是所有服务都是用默认的负载均衡器，即轮询的负载均衡器。我们也可以让调用不同的服务，使用不同的负载均衡器

- 创建两个配置文件，把轮询负载均衡器和随机负载均衡器注册为bean
	```java
	@Configuration
	public class DefaultLoadBalancerConfig {
	
	    @Bean
	    ReactorLoadBalancer<ServiceInstance> roundRobinLoadBalancer(Environment environment,
	                                                                LoadBalancerClientFactory loadBalancerClientFactory) {
	        String name = environment.getProperty(LoadBalancerClientFactory.PROPERTY_NAME); // 获取负载均衡器名称
	        return new RoundRobinLoadBalancer(loadBalancerClientFactory.getLazyProvider(name, ServiceInstanceListSupplier.class), name);
	    }
	}
	```
	```java
	@Configuration
	public class RandomLoadBalancerConfig {
	    @Bean
	        // 定义一个Bean
	    ReactorLoadBalancer<ServiceInstance> randomLoadBalancer(Environment environment, // 注入环境变量
	                                                            LoadBalancerClientFactory loadBalancerClientFactory) { // 注入负载均衡器客户端工厂
	        String name = environment.getProperty(LoadBalancerClientFactory.PROPERTY_NAME); // 获取负载均衡器的名称
	
	        // 创建并返回一个随机负载均衡器实例
	        return new RandomLoadBalancer(loadBalancerClientFactory.getLazyProvider(name, ServiceInstanceListSupplier.class), name);
	    }
	}
	```
- 创建restTempalte
	```java
	// module2 使用默认，module3 使用随机
	@Configuration
	
	@LoadBalancerClients({
	        @LoadBalancerClient(value = "module2", configuration = RandomLoadBalancerConfig.class),
	        @LoadBalancerClient(value = "module3", configuration = DefaultLoadBalancerConfig.class)
	})
	public class AppConfig {
	
	    @Bean
	    @LoadBalanced  // 启用 RestTemplate 的负载均衡
	    public RestTemplate restTemplate() {
	        return new RestTemplate();
	    }
	}
	```



## 六、服务网关 gateway

Gateway（网关）是微服务架构中的一个重要组件，它通常用作客户端和多个微服务之间的中介，负责请求的路由、负载均衡、认证、限流、安全控制等功能。它通常部署在前端，起到了“入口”作用，是微服务的前端统一访问点。

**Spring Cloud Gateway** 基于 **WebFlux + Netty + Reactor**，可以更高效地处理大量请求，适用于微服务架构。

### 1. 网关的核心功能

网关的核心职责是将外部请求路由到相应的微服务，同时提供一些重要的功能：

- **请求路由：** 网关根据请求的路径、请求头、参数等信息，将请求转发到对应的微服务。
- **负载均衡：** 网关能够实现请求的负载均衡，将请求分发到多个后端服务实例，提升服务的可用性和性能。
- **安全性：** 网关通常是整个系统的第一道防线，可以进行请求的身份验证、授权控制、加密等。
- **服务发现：** 通过与服务注册中心集成，网关可以动态地获取微服务的实例信息，实现动态路由。
- **过滤器：** 允许在请求处理过程中添加自定义逻辑。过滤器分为“全局过滤器”和“局部过滤器”。
- **动态路由：** 可以动态添加路由规则，无需重新启动网关。

### 2. 准备工作

使用gateway的模块不能引入 `spring-boot-starter-web`，Spring MVC（基于 Servlet） 和 Spring Cloud Gateway（基于 WebFlux）会冲突。

- 引入依赖

  ```xml
  <!-- gateway 依赖 -->
  <dependency>
      <groupId>org.springframework.cloud</groupId>
      <artifactId>spring-cloud-starter-gateway</artifactId>
  </dependency>
  <!-- 需要基于注册中心转发请求的话，加上 nacos 依赖 -->
  <dependency>
     <groupId>com.alibaba.cloud</groupId>
      <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
  </dependency>
  <!-- 如果用到 lb: 的话需要在引入getaway的pom中引入loadbalancer -->
  <dependency>
      <groupId>org.springframework.cloud</groupId>
      <artifactId>spring-cloud-starter-loadbalancer</artifactId>
  </dependency>
  ```

- 配置路由规则：

  ```yaml
  server:
    port: 8002
  
  spring:
    application:
      name: gateway
    cloud:
      nacos:
        discovery:
          server-addr: localhost:8848
  
      gateway:
        routes: # 网关的路由规则
          - id: module1 # 路由的唯一标识，可以随意命名,仅用于区分不同的路由规则。
            uri: lb://module1 # 表示使用负载均衡，将请求转发给注册中心的 module1 服务实例。
            predicates: # 断言规则，表示请求的路径必须以 /module1 开头。
              - Path=/module1/**
            filters: # 用于修改请求或响应，这里的 AddRequestHeader 过滤器会给请求添加一个请求头, test: hello world
              - AddRequestHeader=test, hello world
  ```

  

### 3. 断言和内置过滤器

常见的 `predicates`（路由匹配条件）：

1. **`Path`**：根据请求路径进行匹配
   - `Path=/users/**`：匹配以 `/users/` 开头的路径。
   - 示例：`/users/123`、`/users/details`。

2. **`Method`**：根据 HTTP 方法进行匹配
   - `Method=GET`：只匹配 `GET` 请求。
   - `Method=POST`：只匹配 `POST` 请求。
   - 示例：`GET /users/123`、`POST /users`。

3. **`Host`**：根据请求的 Host 进行匹配
   - `Host=*.example.com`：匹配所有请求的 Host 名称为 `example.com` 的请求。
   - 示例：`GET /users` 请求的 Host 为 `api.example.com`。

4. **`Query`**：根据查询参数进行匹配
   - `Query=username={value}`：匹配查询参数 `username` 的值。
   - 示例：`GET /users?username=john`。

5. **`Header`**：根据请求头进行匹配
   - `Header=Authorization=Bearer {token}`：匹配带有特定 Authorization 头的请求。
   - 示例：`GET /users/123`，并且 `Authorization=Bearer <token>`。



常见的 `filters`（过滤器）：

`filters` 用于在请求和响应之间进行处理，通常用于修改请求头、响应体、重定向等。这里的过滤器是局部的过滤器

1. **`AddRequestHeader`**：添加请求头
   - `AddRequestHeader=X-Request-Foo, Bar`：向请求中添加 `X-Request-Foo` 头，值为 `Bar`。
   - 示例：请求中会包含 `X-Request-Foo: Bar`。

2. **`AddResponseHeader`**：添加响应头
   - `AddResponseHeader=X-Response-Foo, Baz`：向响应中添加 `X-Response-Foo` 头，值为 `Baz`。

3. **`SetPath`**：修改请求路径
   - `SetPath=/newpath/{segment}`：将请求的路径设置为新的路径。
   - 示例：请求 `/users/123` 会被设置为 `/newpath/123`。

4. **`RedirectTo`**：重定向请求到其他地址
   - `RedirectTo=301, /new-location`：将请求重定向到 `/new-location`。
   - 示例：会发出 `301` 重定向到 `/new-location`。

上面的`predicates`和`filters`只写了一部分，具体可以参考spring官网 [地址](https://docs.spring.io/spring-cloud-gateway/docs/current/reference/html/#gateway-request-predicates-factories)



### 4. 自定义全局和局部过滤器

下面是自定义过滤器的实现方式，其中后两个是gateway提供的

| 方式                            | 作用               | 适用范围          | **执行**                                               |
| ------------------------------- | ------------------ | ----------------- | ------------------------------------------------------ |
| **WebFilter**                   | 低级别的请求拦截   | **基于 WebFlux**  | 最早执行，**拦截所有请求**                             |
| **Spring Security FilterChain** | 权限认证           | **基于 Security** | 如果 **认证不通过**，请求不会进入 Gateway 的 `filters` |
| **GlobalFilter**                | 拦截所有请求       | **全局过滤**      | Security 通过后，**作用于所有 Gateway 处理的请求**     |
| **GatewayFilterFactory**        | 针对单个路由的过滤 | **局部过滤**      | 仅针对匹配的 **某个路由** 生效                         |



#### 3.1 自定义全局过滤器

在 Spring Cloud Gateway 中，全局过滤器（Global Filters）用于在请求和响应过程中对所有路由进行处理。

- 过滤器的作用：
  - **请求过滤：** 在请求到达后端微服务之前对请求做一些处理，比如增加请求头、日志记录、权限校验等。
  - **响应过滤：** 在响应从后端微服务返回到客户端之前对响应做一些修改，比如修改响应内容、加密、日志记录等。

1. 创建一个全局过滤器

首先，需要创建一个实现 `GlobalFilter` 接口的类。在这个类中，你可以定义过滤器的逻辑。

```java
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

@Component
public class AddHeaderGlobalFilter implements GlobalFilter, Ordered {

   @Override
   public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
       return chain.filter(exchange);
   }

   @Override
   public int getOrder() {
       return 1;
   }
}
```

2. 全局过滤器的工作原理

- **`filter()`**：在这里你可以获取到 `ServerWebExchange` 对象，它包含了请求和响应的所有信息。你可以在这里操作请求和响应的内容，进行一些预处理或后处理。最后，调用 `chain.filter(exchange)` 以传递请求继续向下执行其他过滤器或路由。

- **`getOrder()`**：返回一个整数值，用来决定过滤器的执行顺序。**值越小的过滤器会优先执行**。如果你有多个全局过滤器，它们会按照 `getOrder()` 返回值的顺序执行。



3. 示例：添加请求头

假设我们需要为每个请求添加一个特定的请求头。

```java
package com.cloud.gateway.config;

import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.http.HttpHeaders;
import org.springframework.stereotype.Component;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

@Component
public class AddHeaderGlobalFilter implements GlobalFilter, Ordered {

    @Override
    public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
        // 获取请求的 headers
        HttpHeaders headers = exchange.getRequest().getHeaders();

        // 打印原始请求头
        System.out.println("Request Headers: " + headers);

        // 为请求添加一个新的头部
        exchange.getRequest().mutate()
                .header("test", "hello world") // 添加请求头
                .build();

        // 继续传递到下一个过滤器
        return chain.filter(exchange);
    }

    @Override
    public int getOrder() {
        return 1;
    }
}

```

#### 3.1 自定义局部过滤器

1. 自定义局部过滤器（`GatewayFilter`）：

```java
import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.factory.AbstractGatewayFilterFactory;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Component;

@Component
public class AuthFilter  extends AbstractGatewayFilterFactory<AuthFilter.Config> {

    public AuthFilter () {
        super(Config.class);
    }

    @Override
    public GatewayFilter apply(Config config) {
        return (exchange, chain) -> {
            System.out.println("AuthFilter"); 
            // 下面实现自己的逻辑
            String token = exchange.getRequest().getHeaders().getFirst("token");
            if (token == null || !token.equals(config.getToken())) {
                exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
                return exchange.getResponse().setComplete();
            }
            return chain.filter(exchange);  // 继续处理链中的其他过滤器
        };
    }


    public static class Config {
        private String token;

        public String getToken() {
            return token;
        }

        public void setToken(String token) {
            this.token = token;
        }
    }
}
```

2. 使用局部过滤器

局部过滤器通常在路由配置中使用，你可以将它应用于特定的路由，例如：

```yaml
spring:
  cloud:
    gateway:
      routes:
        - id: user-service
          uri: lb://USER-SERVICE
          predicates:
            - Path=/users/**
          filters:
            - AuthFilter  # 这里引用自定义的局部过滤器
```




## 五、分布式事务 seata

Seata 是一款开源的分布式事务解决方案，旨在解决微服务架构中跨服务的事务一致性问题。它提供了易于使用、性能高效的分布式事务管理功能，帮助开发者在分布式系统中保持数据一致性。

### 1. Seata 的概要

Seata 由阿里巴巴发起，最初的目的是为了解决微服务场景下的数据一致性问题，尤其是在分布式数据库事务中。Seata 提供了全局事务、分支事务以及资源管理等功能。

- **全局事务**：在分布式事务中，通常一个事务可能会涉及多个服务或数据库，Seata 通过全局事务 ID 来保证跨服务事务的原子性。
- **分支事务**：每个服务或数据库在执行全局事务时会变成一个分支事务，Seata 负责协调各个分支事务的提交与回滚。
- **AT模式**：Seata 提供了基于数据库的 AT (Automatic Transaction) 模式，它通过对数据库的预存操作和恢复操作来实现事务的一致性。
- **TCC模式**：TCC (Try, Confirm, Cancel) 是 Seata 支持的另一种事务模式，适用于需要显式操作的场景。
- **SAGA模式**：SAGA 是另一种事务模型，适合于长事务和复杂业务场景。



![在这里插入图片描述](./../../../笔记/笔记图片/00f91ac47ccb4af59545cd0539e69d09.png)
- RM（ResourceManager）：用于直接执行本地事务的提交和回滚。
- TM（TransactionManager）：TM是分布式事务的核心管理者。比如现在我们需要在借阅服务中开启全局事务，来让其自身、图书服务、用户服务都参与进来，也就是说一般全局事务发起者就是TM。
- TC（TransactionManager）这个就是我们的Seata服务器，用于全局控制，一个分布式事务的启动就是由TM向TC发起请求，TC再来与其他的RM进行协调操作

>TM请求TC开启一个全局事务，TC会生成一个XID作为该全局事务的编号，XID会在微服务的调用链路中传播，保证将多个微服务的子事务关联在一起；RM请求TC将本地事务注册为全局事务的分支事务，通过全局事务的XID进行关联；TM请求TC告诉XID对应的全局事务是进行提交还是回滚；TC驱动RM将XID对应的自己的本地事务进行提交还是回滚；

下面以官网的案例来演示整个使用过程

### 2. 准备工作

- **下载seata服务（tc）**
  下载seata-service压缩包，解压进入`/bin`  ([下载地址](https://github.com/apache/incubator-seata/releases))， 执行 `./seata-server.bat`


- **配置 Seata**：
   在 `conf` 目录下，你会找到一个配置文件 `application.yml`。编辑该文件以配置 Seata Server 的各种参数，比如数据库连接、事务日志等。

   ```yaml
	server:
	  port: 7091
	
	spring:
	  application:
	    name: seata-server
	
	logging:
	  config: classpath:logback-spring.xml
	  file:
	    path: ${log.home:${user.home}/logs/seata}
	  extend:
	    logstash-appender:
	      destination: 127.0.0.1:4560
	    kafka-appender:
	      bootstrap-servers: 127.0.0.1:9092
	      topic: logback_to_logstash
	
	console:
	  user:
	    username: seata
	    password: seata
	seata:
	  config:
	    # support: nacos, consul, apollo, zk, etcd3
	    type: nacos
	    nacos:
	      server-addr: 127.0.0.1:8848
	      namespace:
	      group: SEATA_GROUP
	      username: nacos
	      password: nacos
	  registry:
	    # support: nacos, eureka, redis, zk, consul, etcd3, sofa
	    type: nacos
	    nacos:
	      application: seata-server
	      server-addr: 127.0.0.1:8848
	      group: SEATA_GROUP
	      namespace:
	      cluster: default
	      username: nacos
	      password: nacos
	  store:
	    # support: file 、 db 、 redis 、 raft
	    mode: db
	  db:
	      datasource: druid
	      db-type: mysql
	      driver-class-name: com.mysql.jdbc.Driver
	      url: jdbc:mysql://127.0.0.1:3306/seata?rewriteBatchedStatements=true
	      user: mysql
	      password: mysql
	      min-conn: 10
	      max-conn: 100
	      global-table: global_table
	      branch-table: branch_table
	      lock-table: lock_table
	      distributed-lock-table: distributed_lock
	      query-limit: 1000
	      max-wait: 5000
	  #  server:
	  #    service-port: 8091 #If not configured, the default is '${server.port} + 1000'
	  security:
	    secretKey: SeataSecretKey0c382ef121d778043159209298fd40bf3850a017
	    tokenValidityInMilliseconds: 1800000
	    ignore:
	      urls: /,/**/*.css,/**/*.js,/**/*.html,/**/*.map,/**/*.svg,/**/*.png,/**/*.jpeg,/**/*.ico,/api/v1/auth/login,/metadata/v1/**
   ```

- **数据库配置**
 创建 `seata` 数据库，并访问 [该地址](https://github.com/apache/incubator-seata/blob/develop/script/server/db/mysql.sql)，在seata数据库中执行里面所有的sql脚本
  创建 `seata_account`, `seata_order`, `seata_storage`数据库，访问 [该地址](https://seata.apache.org/zh-cn/docs/v2.0/user/quickstart), 在上面三个数据库中都创建 `UNDO_LOG` 表，在对应的数据库中创建对应的 业务表。（地址中有sql脚本）

-  **启动 Seata Server**：
   使用以下命令启动 Seata Server：
   ```bash
   sh bin/seata-server.sh
   ```



最终效果如下
![在这里插入图片描述](./../../../笔记/笔记图片/03c1b7ff52024fab80c47167a34c7668.png)

创建三个模块(account, order, storage)，加入相关依赖(数据库驱动，mybatis...)，然后按照下面加入 `nacos` 和 `seata` 依赖和配置



### 3. Seata 与 Spring Boot 集成

#### 3.1 添加 Seata 依赖

```xml
<dependency>
   <groupId>com.alibaba.cloud</groupId>
    <artifactId>spring-cloud-starter-alibaba-seata</artifactId>
    <exclusions>
        <exclusion>
            <groupId>io.seata</groupId>
            <artifactId>seata-spring-boot-starter</artifactId>
        </exclusion>
    </exclusions>
</dependency>
```

#### 3.2 配置 Seata

在 Spring Boot 项目的 `application.yml` 中配置 Seata。

```yaml
seata:
  registry:
    type: nacos
    nacos:
      server-addr: localhost:8848
      namespace: ""
      group: SEATA_GROUP # 组名  跟我们之前配置的seata的配置文件的是对应的
      application: seata-server # 服务名 跟我们之前配置的seata的配置文件的是对应的
  tx-service-group: default_tx_group
  service:
    vgroup-mapping:
      default_tx_group: default
  data-source-proxy-mode: AT
```

#### 3.3 开启事务管理

`@GlobalTransactional` 是 Seata 提供的注解，用于实现分布式事务的管理。它是 Seata 的全局事务控制器，通过这个注解，你可以在一个跨多个微服务的操作中，确保数据的一致性和事务的回滚。

1. 作用：

   - **开启全局事务**：使用 `@GlobalTransactional` 注解可以标记方法为全局事务，Seata 会在这个方法执行时开启一个全局事务。

   - **事务的提交与回滚**：在方法执行过程中，如果发生异常，Seata 会自动回滚所有与该全局事务相关的子事务。相反，如果方法执行成功，Seata 会提交所有子事务。


2. 基本语法：

```java
@GlobalTransactional(name = "your-global-tx-name", rollbackFor = Exception.class)
public void yourMethod() {
    // Your business logic
}
```

3.  参数：

   - `name`：指定全局事务的名称，通常为了区分不同的事务，可以给它一个有意义的名字。

   - `rollbackFor`：指定哪些异常类型会导致事务回滚，默认是 `RuntimeException` 和 `Error`，如果需要捕获其他异常，可以通过此参数指定。





## 六、服务熔断和降级 sentinel

Sentinel 是阿里巴巴开源的分布式系统流量控制组件，主要用于保护微服务在高并发、突发流量等场景下的稳定性和可靠性。Sentinel 提供了 流量控制、熔断降级、系统自适应保护等机制

### 1. Sentinel 的核心功能

Sentinel 提供了以下核心功能：

- **流量控制**：根据预定义的规则限制访问频率或数量，避免系统过载。
- **熔断降级**：在请求失败率或响应时间过高时自动降级，防止雪崩效应。
- **系统自适应保护**：基于系统的负载情况自动进行保护，如 CPU 使用率、内存占用等。
- **热点参数限流**：对请求中的参数进行限流，比如按不同用户 ID 或商品 ID 限制请求。

### 2. Sentinel 的基本概念

在 Sentinel 中，核心是“资源”，可以是服务、接口或方法。每一个资源都会绑定一套限流规则或熔断规则。以下是 Sentinel 的几个基本概念：

- **资源 (Resource)**：受保护的对象，如 API 接口或数据库操作。
- **规则 (Rule)**：用于定义流量控制、熔断降级的条件。
- **Slot Chain**：Sentinel 使用 Slot Chain 机制来应用限流、降级、系统保护等规则。每一个 Slot Chain 包含多个 Slot，不同 Slot 处理不同的规则。
### 3. Sentinel 的监控和控制台

Sentinel 提供了 Dashboard 管理控制台，可以用来监控各个资源的访问情况，并动态配置流量控制和熔断规则。([下载地址](https://github.com/alibaba/Sentinel/releases), 下载`jar`包)

- 启动 jar 包。
![在这里插入图片描述](./../../../笔记/笔记图片/a971cdcfa04c4a7a905eb2337eb104c3.png)
访问 `http://localhost:8888/` (这里我设置的端口是 8888) 进入控制台。
账号密码都是 sentinel

- 引入 sentinel
	```xml
	 <dependency>
	    <groupId>com.alibaba.cloud</groupId>
	    <artifactId>spring-cloud-starter-alibaba-sentinel</artifactId>
	</dependency>
	```
	
	```java
	spring.cloud.sentinel.transport.dashboard=localhost:8888
	```
- 访问module任意controller后就可以看到 module1 加入到控制台
![在这里插入图片描述](./../../../笔记/笔记图片/1419e7c5b247457d8e1be87e893ff734.png)

### 4. @SentinelResource注解

`@SentinelResource` 是 **Sentinel** 中用于标记 `方法` 的注解，它允许你在方法调用时应用 Sentinel 的规则和控制策略。通过使用这个注解，你可以将 **Sentinel** 的流量控制、熔断降级、热点参数限流等功能集成到业务逻辑中。

#### 4.1 `value` 参数：资源名称

- **作用**：指定该方法的 **资源名称**，这个名称会作为 Sentinel 资源的标识，用来应用流控、熔断、降级等策略。通常可以是方法名或其他具代表性的字符串。
- **默认值**：方法名（如果未指定）。



#### 4.2 `blockHandler` 参数：限流或熔断处理方法

- **作用**：当资源受到 **流量控制**（如限流）或 **熔断降级**（如调用失败）时，会调用指定的 `blockHandler` 方法。(`handelBlock` 方法一定要写上 `BlockException ex `)
- **类型**：`blockHandler` 的方法可以是当前类中的一个静态方法，或是其他类的静态方法。

示例：
```java
@SentinelResource(value = "myResource", blockHandler = "handleBlock")
public String someMethod() {
    // 你的业务逻辑
    return "Hello, Sentinel!";
}

// 流控或熔断时的处理方法
public static String handleBlock(BlockException ex) {
    return "Service is currently unavailable due to high traffic. Please try again later.";
}
```
在这个例子中，当 `myResource` 资源被流控或熔断时，`handleBlock` 方法将被调用，返回一条友好的错误信息。

#### 4.3 `fallback` 参数：降级处理方法

- **作用**：当资源发生 **异常** 或 **超时** 时，触发降级逻辑，会调用指定的 `fallback` 方法。这个方法需要与原始方法的签名一致。（**流量超限**或**熔断**触发时也会调用 `fallback`，如果同时有 `fallback` 和 `blockHandler`, 会优先调用 `blockHandler`）
- **类型**：`fallback` 的方法可以是当前类中的一个静态方法，或是其他类的静态方法。

示例：
```java
@SentinelResource(value = "myResource", fallback = "fallbackMethod")
public String someMethod() {
    // 可能会抛出异常的业务逻辑
    throw new RuntimeException("Something went wrong");
}

// 降级方法
public static String fallbackMethod(Throwable ex) {
    return "Service is temporarily unavailable due to internal error. Please try again later.";
}
```
在这个例子中，当 `someMethod` 方法抛出异常时，`fallbackMethod` 会被调用，返回一个默认的降级响应。

#### 4.4 `exceptionsToIgnore` 参数：忽略的异常类型

- **作用**：指定不触发 **降级** 或 **熔断** 的异常类型。即使这些异常发生，也不会进入 `fallback` 或 `blockHandler`。
- **默认值**：没有忽略的异常，所有的异常都会触发降级逻辑。

#### 4.5 `blockHandlerClass` 参数：自定义流控处理类

- **作用**：指定一个类，其中包含 `blockHandler` 处理方法。这样可以将流控处理方法与业务逻辑分离，便于管理。
- **默认值**：如果未指定，`blockHandler` 方法会在当前类中查找。

示例：
```java
@SentinelResource(value = "myResource", blockHandler = "handleBlock", blockHandlerClass = BlockHandler.class)
public String someMethod() {
    // 业务逻辑
    return "Hello, Sentinel!";
}

// 自定义流控处理类
public class BlockHandler {
    public static String handleBlock(BlockException ex) {
        return "Service is temporarily unavailable due to traffic control.";
    }
}
```

#### 4.6 `fallbackClass` 参数：自定义降级处理类

- **作用**：指定一个类，其中包含 `fallback` 处理方法。这样可以将降级处理方法与业务逻辑分离，便于管理。
- **默认值**：如果未指定，`fallback` 方法会在当前类中查找。

示例：
```java
@SentinelResource(value = "myResource", fallback = "fallbackMethod", fallbackClass = FallbackHandler.class)
public String someMethod() {
    // 业务逻辑
    throw new RuntimeException("Something went wrong");
}

// 自定义降级处理类
public class FallbackHandler {
    public static String fallbackMethod(Throwable ex) {
        return "Service is temporarily unavailable due to an error.";
    }
}
```


### 5. 流量控制
#### 5.1 流控模式
- 直接：只针对于当前接口。
- 关联：如果指定的关联接口超过qps（每秒的请求数），会导致当前接口被限流。
- 链路：更细粒度的限流，能精确到具体的**方法**, 直接和关联只能对接口进行限流
	

链路流控模式指的是，当从指定接口过来的资源请求达到限流条件时，开启限流。

```
# 关闭Context收敛，这样被监控方法可以进行不同链路的单独控制
spring.cloud.sentinel.web-context-unify: false
# 比如我 /test1 和 /test2 接口下都调用了 test 方法，那么我可以在sentinel 控制台中分别对 /test1 和 /test2 的 test 进行设置
```

![在这里插入图片描述](./../../../笔记/笔记图片/ee34af3e6ded420f914243b058bad341.png)



 #### 5.2 流控效果

**快速拒绝：** 既然不再接受新的请求，那么我们可以直接返回一个拒绝信息，告诉用户访问频率过高。
**预热:** 依然基于方案一，但是由于某些情况下高并发请求是在某一时刻突然到来，我们可以缓慢地将阈值提高到指定阈值，形成一个缓冲保护
**排队等待：** 不接受新的请求，但是也不直接拒绝，而是进队列先等一下，如果规定时间内能够执行，那么就执行，要是超时就算了。

#### 5.3 实现对方法的限流控制
我们可以使用到`@SentinelResource`对某一个**方法**进行限流控制，无论是谁在何处调用了它，，一旦**方法**被标注，那么就会进行监控。
```java
@RestController
public class TestController {

    @Autowired
    RestTemplate restTemplate;

    @Autowired
    MyService myService;

    @GetMapping("/test1")
    public String test1() {
        return myService.test();
    }

    @GetMapping("/test2")
    public String test2() {
        return myService.test();
    }
}
// ------------------------
@Service
public class MyService {

    @SentinelResource("mytest")	// 标记方法
    public String test() {
        return "test";
    }
}
```
添加 `spring.cloud.sentinel.web-context-unify=false`, 可以对 /test1 和 /test2 的 mytest单独控制
![在这里插入图片描述](./../../../笔记/笔记图片/d8c7a8edab6048ff94c6f139ae15ed54.png)
不添加的, 无法单独控制
![在这里插入图片描述](./../../../笔记/笔记图片/767ba7e5356d4f1eb80a7845601feeab.png)
~~一些坑~~
如果在方法上添加了 `@SentinelResource` 注解，但是不在控制台中显示的话(不显示mytest), 可能是因为添加的方法没有加入到spring容器中进行管理。比如我当时下面这样写就出现了不在控制台显示的情况。
```java
@RestController
public class TestController {
    
    @GetMapping("/test1")
    public String test1() {
        return test();
    }
    
    @SentinelResource("mytest")
    public String test() {
        return "test";
    }
}
```

#### 5.4 处理限流情况
当访问某个 `接口` 出现限流时，会抛出限流异常，重定向到我们添加的路径
```java
@RequestMapping("/blockPage")
    public String blockPage() {
        return "blockPage";
    }
```
```xml
# 用于设置当 流控 或 熔断 触发时，返回给用户的 阻塞页面
spring.cloud.sentinel.block-page=/blockPage
```

当访问的某个 `方法` 出现了限流，为 `blockHander` 指定方法，当限流时会访问 `test2`, 并且可以接受 `test` 的参数
```java
@Service
public class MyService {

    @SentinelResource(value = "mytest", blockHandler = "test2")
    public String test(int id) {
        return "test";
    }

    public String test2(int id, BlockException ex) {
        return "test 限流了 " + id;
    }
}
```

#### 5.5 处理异常的处理
当调用 `test` 时抛出了异常的话，那么就会执行 fallback 指定的方法, `exceptionsToIgnore` 指定的是忽略掉的异常。（限流的话会抛出限流的异常，也会被捕获，执行 fallback 指定的方法）
如果 fallback 和 blockHandler 都指定了方法，出现限流异常会优先执行 blockHandler 的方法
```java
@Service
public class MyService {

    @SentinelResource(value = "mytest", fallback = "except", exceptionsToIgnore = IOException.class)
    public String test(int id) {
        throw new RuntimeException("hello world");
    }
// 注意参数, 必须是跟 test 的参数列表相同, 然后可以多加一个获取异常的参数
    public String except(int id, Throwable ex) { 
        System.out.println("hello = " + ex.getMessage());
        return ex.getMessage();
    }
}
```

#### 5.6 热点参数限流
对接口或方法的某个参数进行限流

```java
@GetMapping("/test1")
@SentinelResource(value = "test1")
public String test1(@RequestParam(required = false) String name, @RequestParam(required = false) String age) {
    return name + " " + age;
}
```

![在这里插入图片描述](./../../../笔记/笔记图片/418ed3656ddf48a0930583ae3c973c5b.png)

### 6. 熔断降级

#### 6.1 **服务熔断**

当某个服务的错误率或失败次数（如异常或超时）超过预设的阈值时，熔断器就会“断开”，该服务的调用将不会继续执行。熔断器“断开”后，所有请求都会被 拒绝，直到熔断器进入 **恢复阶段**，然后根据预设规则恢复正常工作。

#### 6.2 熔断规则

- 2.1 熔断规则
![在这里插入图片描述](./../../../笔记/笔记图片/e2369f121582458aa34ebbcd6f8f15e3.png)

熔断规则定义了什么情况下触发熔断。常见的触发条件有：

- **异常比例（异常比例熔断）**：当某个资源的调用异常（如超时、返回错误等）比例超过设定的阈值时，触发熔断。
- **错误数（错误数熔断）**：当某个资源的调用错误次数超过设定的阈值时，触发熔断。
- **RT（响应时间熔断）**：当某个资源的调用响应时间超过设定的阈值时，触发熔断。

这些规则是可以配置的，通常会在应用初始化时进行设置。

- 2.2 熔断恢复机制
	熔断后，Sentinel 会进入 **自恢复机制**，通过设定的时间窗口，逐渐恢复正常的服务调用。这一恢复过程通常包括以下两个阶段：

	- **半开状态**：熔断器进入半开状态，允许一定数量的请求尝试访问该资源。此时系统会验证该资源是否已恢复，若恢复则正常调用，若失败则继续熔断。
	- **正常状态**：如果半开状态中的请求均成功，则熔断器恢复为正常状态，恢复正常的请求处理。

#### 6.3 **服务降级 (Service Degradation)**

服务降级是指在某些情况下，通过返回一个默认值、简单处理或快速失败来减少服务的压力，避免因资源超载或异常请求导致服务崩溃。Sentinel 通过配置不同的降级策略，使得系统能够在流量激增或服务不稳定时自动切换到降级模式。

最后简单介绍一下限流，熔断，降级之间的联系。 根据降级的概念，当出现`限流`或`熔断`的时候都会触发降级的方法，只不过`熔断`会根据自己的配置，来跟熔断的服务断开联系，不再接受请求。而`限流`的话不会断开服务，而是继续接受请求，如果请求不满足 限流规则的话，还是会进入到降级的方法

## 七、服务链路追踪 

在zipkin官网下载 `zipkin.jar` 包。[下载地址](https://zipkin.io/pages/quickstart.html)

**Micrometer Tracing** 是 Spring 官方在现代可观测性（Observability）体系中的新工具，用于实现分布式链路追踪。

---

### 1. 核心概念

1. **Span**  
   - 每个 Span 表示一个操作的执行过程（如调用某个方法、处理一个 HTTP 请求）。
   - 包含操作的开始时间、持续时间、名称等信息。
   - Trace 是由多个 Span 组成的分布式调用链。

2. **Trace**  
   - 表示分布式系统中多个服务协作完成的整体调用链。  
   - 包括一个根 Span 和多个子 Span。

3. **Context Propagation**  
   - 在分布式系统中，需要在服务之间传递上下文（如 traceId 和 spanId），以实现跨服务的追踪。
   - Micrometer Tracing 通过支持 W3C Trace Context 标准（OpenTelemetry 默认标准）完成上下文传递。

---

### 2. 工作原理

1. **生成和传播追踪信息**
   - 服务 A 生成 traceId 和 spanId，并通过 HTTP header 或消息队列传递给服务 B。
   - 服务 B 接收这些信息，继续生成新的 span 并关联到 trace。

2. **采样策略**
   - 决定是否记录 trace 数据（例如只采样部分请求以减少性能开销）。
   - Micrometer Tracing 提供多种采样策略，可以通过配置控制。

3. **导出追踪数据**
   - Micrometer Tracing 支持将追踪数据导出到多个后端（如 Jaeger、Zipkin 或 Prometheus）。这里我们用 `zipkin`
   - 通过 OpenTelemetry Bridge 提供对多种观察后端的支持。

---

### 3. 简单配置

在 Spring Boot 3+ 项目中，Micrometer Tracing 的依赖配置如下：

-  **引入依赖**
	在 `pom.xml` 中添加以下依赖：
	
	```xml
	<dependency>
	   <groupId>org.springframework.boot</groupId>
	    <artifactId>spring-boot-starter-actuator</artifactId>
	</dependency>
	<dependency>
	    <groupId>io.micrometer</groupId>
	    <artifactId>micrometer-tracing</artifactId>
	</dependency>
	<dependency>
	    <groupId>io.micrometer</groupId>
	    <artifactId>micrometer-tracing-bridge-brave</artifactId>
	</dependency>
	<dependency>
	    <groupId>io.micrometer</groupId>
	    <artifactId>micrometer-observation</artifactId>
	</dependency>
	<dependency>
	    <groupId>io.zipkin.reporter2</groupId>
	    <artifactId>zipkin-reporter-brave</artifactId>
	</dependency>
	<!-- 使用feign的话加入下面这个依赖 -->
	<!--        <dependency>-->
	<!--            <groupId>io.github.openfeign</groupId>-->
	<!--            <artifactId>feign-micrometer</artifactId>-->
	<!--        </dependency>-->
	```


-  **配置**
	```yaml
	management:
	  zipkin:
	    tracing:
	      endpoint: http://localhost:9411/api/v2/spans
	  tracing:
	    sampling:
	      probability: 1.0 #采样率默认0.1(10次只能有一次被记录)，值越大手机越及时
	```

之后就可以访问 `http://127.0.0.1:9411` 

---





