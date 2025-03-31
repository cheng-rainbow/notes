### 1. Cookie 和 Session 的区别

Cookie存储在浏览器的本地，记录用户信息，服务器用它来识别用户的 Session。（临时的，或持久的）

Session存放在服务器，默认存储在内存，也可以存 Redis，通过 Session ID 关联用户的 Cookie，保持用户的登录状态

1. **用户登录时，服务器创建 Session 并返回 Session ID**。（springboot会自动设置响应头）

2. **Session ID 存储在 Cookie 中（`JSESSIONID`）**。

3. **客户端请求时，浏览器会自动带上 `JSESSIONID`**。

4. **服务器通过 `JSESSIONID` 找到对应的 Session，判断用户身份**。



### 2.  `@Transactional` 事务

**事务（Transaction）** 是对同一个数据库的操作集合，这些操作要么**全部成功提交**（`commit`），要么**全部失败回滚**（`rollback`）。

Spring 提供了 **`@Transactional` 注解** 来管理事务

常见的事务管理器有：

- `DataSourceTransactionManager`（适用于单个数据库，默认是这个）
- `JpaTransactionManager`（适用于 JPA）
- `HibernateTransactionManager`（适用于 Hibernate）
- `JtaTransactionManager`（适用于分布式事务）



Spring 默认只对 **`RuntimeException`（运行时异常） 和 `Error` 进行回滚**，如果是 `CheckedException`（受检异常，如 `Exception`），不会触发回滚。如果希望 **受检异常（如 `Exception`） 也回滚**，可以手动哪些可以回滚，哪些不会滚：

```java
@Transactional(rollbackFor = Exception.class, noRollbackFor = RuntimeException.class)
```



Spring 事务支持 7 种传播行为（`Propagation`），用于控制事务的执行方式，也就是 带有事务的方法A 中调用了 方法B 时：

| 传播级别           | 说明                                     |
| ------------------ | ---------------------------------------- |
| `REQUIRED`（默认） | 如果当前有事务，则加入；否则新建事务     |
| `REQUIRES_NEW`     | 无论是否有事务，都创建新事务，老事务挂起 |
| `SUPPORTS`         | 有事务就加入，没有就以非事务方式执行     |
| `NOT_SUPPORTED`    | 有事务就挂起，非事务方式执行             |
| `MANDATORY`        | 必须有事务，否则抛异常                   |
| `NEVER`            | 不能有事务，否则抛异常                   |
| `NESTED`           | 如果当前事务存在，则嵌套事务执行         |



Spring 提供 **5 种事务隔离级别**，用于控制 **多个事务之间的可见性**，防止并发问题（脏读、不可重复读、幻读）。

| 隔离级别                 | 说明                                           |
| ------------------------ | ---------------------------------------------- |
| `DEFAULT`                | 由数据库驱动决定（一般是 `READ_COMMITTED`）    |
| `READ_UNCOMMITTED`       | 允许脏读（不安全）                             |
| `READ_COMMITTED`（默认） | 只能读取已提交数据（防止脏读）                 |
| `REPEATABLE_READ`        | 相同事务内，多次读取数据不变（防止不可重复读） |
| `SERIALIZABLE`           | 完全串行化（最高级别，性能最低）               |



事务失效的几种情况

1. 方法不是 `public`（Spring AOP 只会代理 `public` 方法）。

2. 事务管理器未正确配置

3. `@Transactional` 标注的方法在同一个类中直接调用：

原因：`@Transactional` 通过代理来实现事务管理。但是，如果在 **同一个类** 内部直接调用的方法，Spring **不会走代理机制**。

```java
@Transactional
public void methodA() {
    methodB(); // 事务不会生效，因为 methodB 不是通过代理调用
}

@Transactional
public void methodB() {
    // 事务不会生效
}
```



