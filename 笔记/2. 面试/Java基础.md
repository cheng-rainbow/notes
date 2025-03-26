## 一、基本概念

### 1. Java为什么是跨平台的

我们编写的 Java 源码，编译后会生成的是**字节码文件**，而不是机器码。Jvm 是一种软件，不同平台有不同的版本，Jvm 可以把字节码文件编译成对应平台的**机器码**然后运行。所以，只要在不同平台上安装对应的 Jvm 就可以实现跨平台。



### 2. JVM、JDK、JRE三者关系？

1. **JVM：**运行 Java 字节码（`.class` 文件），提供跨平台能力。

2. **JRE：**提供运行 Java 程序的环境，包含 **JVM + 必要的核心类库**。
3. **JDK：**开发 Java 应用的完整工具包**，包含 **JRE + 开发工具（编译器、调试工具等）**。

| 组件    | 是否包含 JVM | 是否包含标准库 | 是否包含开发工具 |
| ------- | ------------ | -------------- | ---------------- |
| **JVM** | ✅            | ❌              | ❌                |
| **JRE** | ✅            | ✅              | ❌                |
| **JDK** | ✅            | ✅              | ✅                |

如果 **只运行 Java 程序**，安装 **JRE 即可**；如果 **需要开发 Java 程序**，必须安装 **JDK**。



### 3. 为什么Java解释和编译都有？

**解释性：**Java代码编译生成字节码文件，jvm逐行解释并执行字节码。

**编译执行：**Jvm通过 JIT ，把热点代码直接编译成对应平台的机器码，提高代码运行效率。

解释执行保证跨平台，编译执行提高性能。



## 二、数据类型

### 1.  Integer相比int有什么优点？

| 特性                  | `int`                  | `Integer`                                              |
| --------------------- | ---------------------- | ------------------------------------------------------ |
| **存储方式**          | **栈**（Stack）        | **堆**（Heap），存对象                                 |
| **是否可以为 `null`** | ❌ 不能                 | ✅ 可以                                                 |
| **支持的方法**        | ❌ 无方法，只能直接运算 | ✅ 提供 `parseInt()`、`valueOf()`、`compareTo()` 等方法 |
| **支持泛型**          | ❌ 不能用于泛型         | ✅ 适用于泛型容器，如 `List<Integer>`                   |
| **自动装箱 & 拆箱**   | ❌ 不支持               | ✅ 支持 `int` 和 `Integer` 之间的自动转换               |
| **缓存机制**          | ❌ 无缓存               | ✅ `-128` 到 `127` 之间有缓存                           |
| **性能**              | ✅ 速度快，占用内存少   | ❌ 速度较慢，开销大                                     |



## 三、面向对象

### 1. 怎么理解面向对象？简单说说封装继承多态

面向对象（Object-Oriented Programming, OOP）是一种 **程序设计思想**，核心是 **将现实世界的事物抽象成对象**，通过 **封装、继承、多态** 组织代码，提高代码的**复用性、可维护性、可扩展性**。

1. **封装**：把数据和方法封装到对象中，隐藏实现细节，只暴露必要的接口。
2. **继承**：子类继承父类的属性和方法，避免重复代码，提高复用性。
3. **多态**：同一个方法，不同的对象有不同的表现。



### 2. 多态体现在哪几个方面？

1. 方法重写：子类**重写**父类的方法，父类引用指向子类对象时，调用的是子类的方法。
2. 方法重载：在 **同一个类** 中，多个方法**方法名相同，参数列表不同**，编译器会根据调用方法时提供的参数类型和数量来决定具体调用哪个版本的方法。



### 3.  多态解决了什么问题？

| **问题**               | **多态的解决方案**         | **带来的好处**               |
| ---------------------- | -------------------------- | ---------------------------- |
| 代码耦合严重           | **父类或接口引用指向子类** | 低耦合，新增功能不影响旧代码 |
| 代码复用性低           | **统一接口，子类实现**     | 代码更通用，避免重复代码     |
| 代码不通用             | **方法重载（Overload）**   | 适配不同参数，提高灵活性     |
| 代码臃肿，不适用于框架 | **泛型 + 多态**            | 适配各种类型，提高框架通用性 |

🚀 **一句话总结**： **多态解决了代码耦合、复用性低、通用性差的问题，让代码更加灵活、可扩展、易维护！**



### 4. 面向对象的设计原则你知道有哪些吗

面向对象编程（OOP）的六大设计原则主要有：单一职责原则（SRP）、开放封闭原则（OCP）、里氏替换原则（LSP）、依赖倒置原则（DIP）、接口隔离原则（ISP） 和 迪米特法则（LoD）。



1. **单一职责原则**（SRP）

定义：一个类应该仅有一个职责。

简单来说，一个类应该只负责一项任务。

```java
// 不遵守SRP：类做了太多的事情
class UserService {
    public void addUser(String username, String password) {
        // 逻辑：添加用户
    }

    public void sendEmail(String email) {
        // 逻辑：发送邮件
    }
}

// 遵守SRP：分开职责
class UserService {
    private EmailService emailService;

    public UserService(EmailService emailService) {
        this.emailService = emailService;
    }

    public void addUser(String username, String password) {
        // 添加用户逻辑
    }
}

class EmailService {
    public void sendEmail(String email) {
        // 发送邮件逻辑
    }
}
```

在这个例子中，`UserService` 类负责添加用户，而 `EmailService` 类负责发送邮件。这样，代码更容易维护和扩展。

2. **开放封闭原则**（OCP）

定义：软件实体（类、模块、函数等）应该对扩展开放，对修改封闭。

简单来说，当需要功能扩展时，可以通过新增代码来实现，而不是修改原有代码。

```java
// 不遵守OCP：每新增一种支付方式都要修改原有代码
class PaymentService {
    public void processPayment(String paymentType) {
        if ("credit".equals(paymentType)) {
            // 处理信用卡支付
        } else if ("paypal".equals(paymentType)) {
            // 处理PayPal支付
        }
    }
}

// 遵守OCP：通过继承和多态扩展
interface PaymentMethod {
    void processPayment();
}

class CreditCardPayment implements PaymentMethod {
    public void processPayment() {
        // 处理信用卡支付
    }
}

class PayPalPayment implements PaymentMethod {
    public void processPayment() {
        // 处理PayPal支付
    }
}

class PaymentService {
    public void processPayment(PaymentMethod paymentMethod) {
        paymentMethod.processPayment();
    }
}
```

通过使用接口和多态，我们可以轻松扩展支付方式，而无需修改 `PaymentService` 类。

3. **里氏替换原则**（LSP）

定义：子类对象应该能够替换父类对象并正常工作。

简单来说，如果一个程序依赖于某个基类（或接口），那么用这个基类的任何一个子类替换后，程序仍然应该正常运行。

1. **子类不能抛出父类未声明的异常**：异常行为必须符合父类的契约。

2. **置条件不能加强，后置条件不能削弱**：方法执行前对输入参数的要求。子类的前置条件不能比父类更严格。方法执行后对输出结果的保证。子类的后置条件不能比父类更弱。

```java
// 不遵守LSP：子类不能替换父类的行为
class Bird {
    public void fly() {
        System.out.println("Bird can fly");
    }
}

class Ostrich extends Bird {
    @Override
    public void fly() {
        throw new UnsupportedOperationException("Ostrich cannot fly");
    }
}

// 遵守LSP：确保子类符合父类的期望行为
class Bird {
    public void move() {
        System.out.println("Bird moves");
    }
}

class Sparrow extends Bird {
    @Override
    public void move() {
        System.out.println("Sparrow flies");
    }
}

class Ostrich extends Bird {
    @Override
    public void move() {
        System.out.println("Ostrich runs");
    }
}
```

在遵守LSP的例子中，`Sparrow` 和 `Ostrich` 都可以替换 `Bird` 类，并且行为符合预期。

4. **依赖倒置原则**（DIP）

定义：高层模块不应该依赖低层模块，两者都应该依赖抽象；抽象不应该依赖细节，细节应该依赖抽象。

简单来说，DIP 要求模块之间的依赖关系通过抽象（如接口或抽象类）来建立，而不是直接依赖具体的实现类。这样可以降低耦合，提高系统的灵活性和可扩展性。

```java
// 不遵守DIP：高层模块直接依赖低层模块
class LightBulb {
    public void turnOn() {
        System.out.println("LightBulb turned on");
    }
}

class Switch {
    private LightBulb bulb;

    public Switch(LightBulb bulb) {
        this.bulb = bulb;
    }

    public void operate() {
        bulb.turnOn();
    }
}

// 遵守DIP：引入抽象接口
interface Switchable {
    void turnOn();
}

class LightBulb implements Switchable {
    public void turnOn() {
        System.out.println("LightBulb turned on");
    }
}

class Fan implements Switchable {
    public void turnOn() {
        System.out.println("Fan turned on");
    }
}

class Switch {
    private Switchable device;	// 接口

    public Switch(Switchable device) {
        this.device = device;
    }

    public void operate() {
        device.turnOn();
    }
}
```

在遵守DIP的例子中，`Switch` 类依赖于 `Switchable` 接口，而不是具体的 `LightBulb` 或 `Fan`，这样我们可以轻松添加其他可操作的设备而不需要修改 `Switch` 类。

5. **接口隔离原则**（ISP）

定义：客户端不应该被迫依赖它不使用的接口。（客户端（Client）指的是**使用某个接口或类的代码实体**。）

简单来说，一个类应该只依赖它需要的接口方法，而不是被迫去实现或使用一个功能过多、包含无关方法的‘大而全’接口

```java
// 不遵守ISP：接口过大，客户端实现了不需要的方法
interface Worker {
    void work();
    void eat();
}

class Robot implements Worker {
    public void work() {
        System.out.println("Robot is working");
    }

    public void eat() {
        throw new UnsupportedOperationException("Robot does not eat");
    }
}

class Human implements Worker {
    public void work() {
        System.out.println("Human is working");
    }

    public void eat() {
        System.out.println("Human is eating");
    }
}

// 遵守ISP：拆分接口
interface Workable {
    void work();
}

interface Eatable {
    void eat();
}

class Robot implements Workable {
    public void work() {
        System.out.println("Robot is working");
    }
}

class Human implements Workable, Eatable {
    public void work() {
        System.out.println("Human is working");
    }

    public void eat() {
        System.out.println("Human is eating");
    }
}
```

遵守ISP后，`Robot` 类只实现了 `Workable` 接口，不需要去实现不必要的 `eat()` 方法。

6. **迪米特法则**（LoD）

定义：一个对象应该对其他对象有尽可能少的了解。

简单来说，减少对象之间的依赖关系，避免直接访问其他类的内部数据。

```java
// 不遵守LoD：类与类之间的依赖过多
class Car {
    private Engine engine;

    public Car(Engine engine) {
        this.engine = engine;
    }

    public void start() {	// 直接调用 Engine 的 ignite 方法
        engine.ignite();
    }
}

class Engine {
    public void ignite() {
        System.out.println("Engine started");
    }
}

// 遵守LoD：通过方法间接访问，减少直接依赖
class Car {
    private Engine engine;

    public Car(Engine engine) {
        this.engine = engine;
    }

    public void start() {
        engine.startEngine();	// 调用 Engine 的高层接口
    }
}

class Engine {
    public void startEngine() {
        ignite();	 // 内部实现细节被封装
    }

    private void ignite() {
        System.out.println("Engine started");
    }
}
```

在遵守LoD的例子中，`Car` 类通过 `Engine` 的公有方法 `startEngine()` 来间接启动引擎，而不直接调用 `ignite()` 方法，从而减少了对 `Engine` 类内部实现的依赖。Car 知道引擎启动需要调用 `ignite`，这超出了它应该知道的范围。理想情况下，Car 只需发出“启动”指令，具体怎么启动应由 Engine 自己负责。



### 5. 重载与重写有什么区别？

**重载（Overloading）**指的是在同一个类中，可以有**多个同名方法**，它们具有**不同的参数列表**（参数类
型、参数个数或参数顺序不同），**编译器**根据调用时的参数类型来决定调用哪个方法。
**重写（Overriding）**指的是子类可以重新定义父类中的方法，**方法名、参数列表和返回类型**必须与父类
中的方法一致，通过@override注解来明确表示这是对父类方法的重写。



### 6. 抽象类和普通类有什么区别

1. 抽象类不能实例化，普通类可以。

2. 抽象类可以包含抽象方法（无方法体），普通类不能。

3. 子类必须实现抽象方法，否则也要定义为抽象类。

4. abstract 不能和 final 同时使用



### 7. 非静态内部类和静态内部类的区别

| 特性               | 非静态内部类                    | 静态内部类                               |
| ------------------ | ------------------------------- | ---------------------------------------- |
| 定义关键字         | 无 static                       | 使用 static                              |
| 是否依赖外部类实例 | 是                              | 否                                       |
| 实例化方式         | outer.new Inner()               | Outer.StaticInner()                      |
| 访问外部类成员     | 可以访问 `实例成员`和`静态成员` | 不可以访问`实例成员`，可以访问`静态成员` |
| 访问外部类方法     | 可以访问 `实例方法`和`静态方法` | 不可以访问`实例方法`，可以访问`静态方法` |
| 使用场景           | 与外部实例紧密相关              | 独立于外部实例的逻辑                     |
| 内存泄漏风险       | 有（因持有外部类引用）          | 无                                       |



### 8. 非静态内部类可以直接访问外部方法，编译器是怎么做到的？

编译器在生成字节码时会为非静态内部类维护一个指向外部类实例的引用。编译器会在生成非静态内部类的构造方法时，将外部类实例作为参数传入



### 9. 在我new一个子类对象的时候加载顺序是怎么样的？

1. 首先会加载父类的静态成员变量和静态代码块

2. 接下来，加载子类的静态成员变量和静态代码块。
3. 调用父类构造方法，然后是子类的构造方法。



## 四、深拷贝和浅拷贝

### 1. 浅拷贝和深拷贝有什么区别

浅拷贝是指只复制`对象本身`和其内部的`值类型字段`，但不会复制对象内部的`引用类型`字段。

深拷贝是复制`对象本身`和其内部的`值类型字段`和内部的`引用类型`字段，引用类型字段不再是共享的



### 2. 实现深拷贝的三种方法是什么？

1. 实现 Cloneable 接口并重写 clone() 方法

这种方法要求对象及其所有引l用类型字段都实现 `Cloneable 接口`，并且重写cloneO方法。在 cloneO方法中，通过递归克隆引用类型字段来实现深拷贝。

2. 使用序列化和反序列化

通过将对象序列化为`字节流`，再从字节流反序列化为对象来实现深拷贝。要求对象及其所有引I用类型字段都实现 `Serializable 接口`。

3. 使用 Json

通过将对象转换为 JSON 或其他中间格式，再解析为新对象，或者使用工具类直接复制。 



## 五、泛型

### 1. 什么是泛型

泛型 是一种在编译时支持参数化类型的`特性`，允许类、接口和方法使用占位符（类型参数）来指定数据类型，从而提高代码的**类型安全性**和**复用性**。



## 六、对象

### 1. 除了new创建对象还有什么方法

1. 通过反射（newInstance）
2. 通过反序列化
3. 通过 clone 方法
4. 通过类提供的静态方法（value.of）



## 七、反射

### 1. 反射是什么

反射是一种运行时动态获取和操作类信息的机制。它允许程序在运行时访问和修改类的成员，调用方法或创建对象，而无需在编译时知道具体的类名或结构。



### 2. 反射在你平时写代码或者框架中的应用场景有哪些?

Spring框架：

1. 依赖注入：Spring 使用反射根据 `@Autowired 或 XML 配置`动态创建和注入 Bean。
2. 属性填充：将配置文件的值通过反射注入到 Bean 的字段中。
3. AOP：通过反射动态调用目标方法或代理方法。

ORM框架：

1. 字段赋值：反射用于`读取实体类的注解`，动态`映射`数据库字段到对象属性。

Spring MVC：

1. 控制器方法调用：根据 URL 映射，通过反射调用控制器中的方法。



## 八、代理

### 1. 什么是代理

代理是一种**设计模式**，也是一种**机制**，通过在方法调用前后加入额外的行为，增加了灵活性。

1. **静态代理**：手动编写代理类。
2. **动态代理**：通过 Java 提供的 API 动态生成代理类。

静态代理是指代理类和被代理类必须在**编译时**就确定，代理类通常会显式实现被代理类的接口并转发方法调用。

动态代理是通过 Java 的 `Proxy` 类和 `InvocationHandler` 接口在**运行时**创建代理类。它不需要显式创建代理类，代码更加灵活。



## 九、异常

### 2. 抛出异常不用 throws 的情况

1. 抛出的是运行时异常
2. 方法内部捕获了异常



## 十、Object

### 1. == 与 equals 有什么区别？

== ： 是一个**关系运算符**，用于比较两个操作数的基本值。

equals：来自 **Object** 类，Object 类中的是比较对象的引用地址是否相同，但是不可以对这个方法进行重写，比如 String, Inter 类对 equals 进行了重写，它们的类型比较的是内容是否相同。



### 2. StringBuffer和StringBuild区别是什么？

**StringBuffer**: 是线程安全的。它的方法都使用了 **synchronized** 关键字进行同步，确保在多线程环境下多个线程不会同时修改同一个 StringBuffer 对象。但由于同步带来的开销，性能较低。

**StringBuffer**:因为有同步机制，每次操作都会有额外的锁开销，因此执行速度较慢。



## 十一、序列化

### 1. 序列化和反序列化让你自己实现你会怎么做?

Java 默认的序列化虽然实现方便，但却存在安全漏洞、不跨语言以及性能差等缺陷。我会考虑用主流序列化框架，比如FastJson、Protobuf来替代Java 序列化。

