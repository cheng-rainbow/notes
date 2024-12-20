前段时间看了下《Java核心技术 I》，算是对之前学的进行查漏补缺，下面是对书中Java的一些不常见或笔者初学时遗漏的一些知识点汇总。



## 一、Java的基本程序设计结构（三）

### 1. 整数表示

可以为数字字面量加上下划线，这些下划线只是为了让人更易读。Java编译器会去除这些下划线。

```java
int n = 1_000_000_000;
System.out.println(n);
// 输出
1000000000
```



### 2. 无符号类型

Java 中没有原生的无符号整数类型。所有的整数类型（`byte`、`short`、`int`、`long`）都是带符号的。不过，Java 提供了某些方法来模拟无符号行为，特别是对于 `int` 和 `long` 类型。



Java 8 引入的 `Integer` 类中的 `toUnsignedLong(int x)` 和 `divideUnsigned(int x, int y)` 方法主要用于支持无符号整数的操作，尽管 `int` 类型本身是带符号的。

#### 2.1 `toUnsignedLong(int x)`

该方法将一个带符号的 `int` 转换为无符号的 `long` 类型。由于 `int` 类型是 32 位，且在 Java 中是有符号的，它的范围是 `-2^31` 到 `2^31-1`，而无符号 `int` 的范围是 `0` 到 `2^32-1`。通过该方法，可以将带符号的 `int` 值转换为无符号的 `long` 值来处理。

```java
int signedInt = -1;  // 带符号的 int 值
long unsignedLong = Integer.toUnsignedLong(signedInt);  // 转换为无符号 long
System.out.println(unsignedLong);  // 输出：4294967295
// -1 二进制是 11111111 11111111 11111111 11111111，转成无符号就是 2^32-1 即4294967295
```



#### 2.2 `divideUnsigned(int x, int y)`

该方法执行两个无符号 `int` 值的除法操作，并返回无符号的结果。它的作用类似于普通的 `x / y`，但是它是以无符号的方式来处理运算，即忽略符号位。

```java
int x = 10;
int y = 3;
int result = Integer.divideUnsigned(x, y);  // 无符号除法
System.out.println(result);  // 输出：3
```



#### 2.3 `remainderUnsigned(int x, int y)`

这个方法与 `divideUnsigned` 类似，但是它返回的是无符号的余数。

```java
int x = 10;
int y = 3;
int remainder = Integer.remainderUnsigned(x, y);  // 无符号余数
System.out.println(remainder);  // 输出：1
```

无符号数的**加减乘法**在二进制层面与有符号数一致，因为它们使用相同的二进制表示，只是解释方式不同。所以加减乘法可以直接使用。



### 3. double 类型

float 类型的数值有一个后缀 F 或 f（例如，3.14F）。没有后缀 F 的浮点数值（如3.14）总是默认为 double 类型。可选地，也可以在double数值后面添加后缀 D 或 d。



所有浮点数计算都遵循 IEEE 754 规范，具体来说，有3个特殊的浮点数值表示溢出和出错的情况。

1. 正无穷大（POSITIVE_INFINITY）
2. 负无穷（NEGATIVE_INFINITY）
3. NaN（不是一个数）

例如，一个正整数除以0的结果为正无穷大，计算0/0或负数的平方根结果为NaN。

```java
Double x = 1.0;
if (x != Double.POSITIVE_INFINITY) {
    System.out.println("x is not POSITIVE_INFINITY");
}
```



### 4. 右移操作

`>>` 和 `>>>` 都是按照位模式右移的运算符，但有略微区别。

1. `>>` 是带符号的右移操作符，叫做 **算数右移**。会保持符号位不变。

2. `>>>` 是无符号右移操作符，叫做 **逻辑右移**。不管正负数，都用 0 填充最高位。

```java
int a = -8; // 二进制：11111111 11111111 11111111 11111000 （补码形式）
int b = a >> 2; // 右移 2 位
System.out.println(b); // 输出：-2
```

`-8` 右移两位后是 `11111111 11111111 11111111 11111110`，负数在计算机中是以补码的形式存储，其原码是各个位取反，再+1。`00000000 00000000 00000000 00000001` + 1 = `00000000 00000000 00000000 00000010`（原码是2）



```java
int a = -8; // 二进制：11111111 11111111 11111111 11111000 （补码形式）
int b = a >>> 2; // 右移 2 位，高位补 0
System.out.println(b); // 输出：1073741822
```

`-8` 右移两位后是 `00111111 11111111 11111111 11111110`，(1073741822)



### 5. char类型

UTF-8 和 UTF-16 是两种常用的字符编码方式。

1. **UTF-8**是一种可变长度的字符编码，它将 Unicode 字符集中的字符编码成 1 到 4 个字节（8 位），使用 1 个字节到 4 个字节表示一个字符。其中前128个字符（ASCII）采用1字节

2. **UTF-16**也是一种可变长度的字符编码方式，但它将 Unicode 字符集中的字符编码成 2 个字节或 4 个字节，并且对于大多数常用字符（前128个字节），使用 2 个字节来表示。

Java中的 **char类型** 使用的就是 UTF-16，其中 UTF-16 有一个代码单元的概念，即表示一个字符的最小存储单位，UTF-16的代码单元是2字节。

观察下面代码可以发现，`charAt` 返回的是一个代码单元，那么对于占用两个代码单元的字符串就会出现奇怪的问题。所以尽量避免使用 Char 类型。

```java
String s = "\uD83D\uDE00hello"; // 这个字符串实际是 😀hello
System.out.println(s);	
System.out.println(s.charAt(0));
System.out.println(s.charAt(1));	// 这里输出的并不是 h
System.out.println(s.charAt(2));
// 输出如下
😀hello
?
?
h
```



### 6. trim 和 strip

1. `trim()` 只会去除 **ASCII 空白字符**（即 **`\u0020`** 空格、**`\u0009`** 制表符、**`\u000A`** 换行符、**`\u000D`** 回车符）以及一些常见的控制字符。

2. `strip()` 在 Java 11 中引入，除了去除 **ASCII 空白字符** 外，它还会去除一些其他的 Unicode 空白字符（如 **`\u2000`** 到 **`\u200B`** 等）。

strip 去除的范围更广，一般采用这个。

```java
String s = "Hello World!\u3000   ";  // 包含零宽空格（\u200B）和全角空格（\u3000）
String trimmed = s.trim();
String stripped = s.strip();

System.out.println("Original: '" + s + "'");
System.out.println("Trimmed: '" + trimmed + "'");
System.out.println("Stripped: '" + stripped + "'");
```



### 7. switch

传统的switch存在 **直通行为**。（即当 `switch` 语句中的某个 `case` 匹配到时，它会继续执行下一个 `case` 的代码，直到遇到 `break` 语句或 `switch` 语句结束为止。这种行为叫做 **fall-through**，即“穿透”到下一个 `case`。）



Java 12 引入了增强的 `switch` 表达式，支持 **非直通行为**。可以返回值并消除掉传统 `switch` 语句中的直通行为，避免不小心出现的错误。引入了 `->` 和 `yield`，其中用 `->` 代替传统的 `:`，并且还可以返回一个值，`yield` 用来指定返回什么，相当于return。

```java
int day = 3;

String result = switch (day) {
    case 1 -> "Monday";	// 对于没有花括号包围的，相当于是 yield "monday";
    case 2 -> "Tuesday";
    case 3 -> {
        // 使用 yield 返回值
        System.out.println("Processing Wednesday...");
        yield "Wednesday"; // 必须使用 yield 来返回值
    }
    case 4 -> "Thursday";
    case 5 -> "Friday";
    default -> "Invalid day";
};

System.out.println(result);  // 输出返回的字符串
// 输出
Processing Wednesday...
Wednesday
```

对于 `switch cast ->` 这种语句只能使用 **yield** 返回结果，不能用 **break**、**continue** 和 **return**。



## 二、对象与类

### 1. final 的用途

1. `final` 修饰一个字段时，意味着该字段的值 **一旦被赋值后就不能再改变**。

2. `final`  修饰方法时，意味着该方法不能被子类重写（覆写）。

3. `final` 修饰类时，意味着该类不能被继承。

final 修饰符对于**基本类型**或者**不可变类型**的字段尤其有用。如果修饰的是**可变类型**的字段，那么表示存储该字段指向的对象引用不用再指向其他的，但是引用本身是**可以改变**的。



### 2. 静态方法构造对象

**使用静态工厂方式来构造对象**是指通过类的**静态方法**来创建并返回对象，而不是直接通过类的构造器 (`new`) 来实例化对象。这种设计模式可以提供更多的灵活性和可扩展性。

#### 2.1 与构造器的对比

| **构造器**                                       | **静态工厂方法**                               |
| ------------------------------------------------ | ---------------------------------------------- |
| 名称固定，必须与类名相同。                       | 方法名可以自由命名，能更清晰地表达意图。       |
| 每次调用都会返回一个新对象。                     | 可以返回缓存对象或共享对象，提升性能。         |
| 直接调用类的构造器，不能对返回的对象做额外处理。 | 可以自定义逻辑，如参数校验、对象池管理等。     |
| 无法隐藏实现类。                                 | 可以返回接口类型或不同的实现类，隐藏具体实现。 |



#### 2.2 静态工厂的优势

1. **更具可读性和语义性**
    方法名可以描述创建对象的意图，而构造器只能使用类名。

   ```java
   // 使用构造器
   LocalDate today = new LocalDate();
   
   // 使用静态工厂
   LocalDate today = LocalDate.now(); // 可读性更强
   ```

2. **可以缓存和复用实例**
    静态工厂方法可以返回已有实例，而不总是创建新对象。

   ```java
   public class Boolean {
       public static final Boolean TRUE = new Boolean(true);
       public static final Boolean FALSE = new Boolean(false);
   
       public static Boolean valueOf(boolean b) {
           return b ? TRUE : FALSE;
       }
   }
   ```

3. **隐藏实现类，返回接口或父类类型**
    静态工厂方法可以返回接口类型或其父类，隐藏具体实现，从而提高灵活性和扩展性。

   ```java
   public interface Animal {}
   public class Dog implements Animal {}
   public class Cat implements Animal {}
   
   public class AnimalFactory {
       public static Animal createAnimal(String type) {
           if ("dog".equals(type)) {
               return new Dog();
           } else if ("cat".equals(type)) {
               return new Cat();
           }
           throw new IllegalArgumentException("Unknown type");
       }
   }
   ```



### 3. 方法参数传递方式

1. **基本数据类型的传递**（传递值）

2. **引用类型的传递**（传递引用的副本）

对于对象的传递，方法得到的是**对象引用的副本**，而不是像C++那样按引用调用，原来的对象引用和这个副本都是引用同一个对象。下面两个示例都可以证明这一点。



1. String 类型

对于String类型，参数传递是按照引用类型传递，但是String是不可变对象，所以在 `testString` 方法中重新赋值a，b，实际上是对方法中的a和b重新创建String对象，并不会改变main中的a，b。

```java
public static void main(String[] args) throws InterruptedException {
    String a = "hello";
    String b = "world";
    System.out.println(a + " " + b);
    testString(a, b);
    System.out.println(a + " " + b);
}
public static void testString(String a, String b) {
    a = "testa";
    b = "testb";
};
// 输出
hello world
hello world
```

2. 对象类型

我们打算让两个对象 person1 和 person2 互换一下对象的引用，可以发现并没有成功。实际是方法中的 p1 指向了 person2 引用的对象，p2指向了person1引用的对象。

```java
public class Test {
    private static final Logger log = LoggerFactory.getLogger(Main.class);
    public static void main(String[] args) throws InterruptedException {
        Person person1 = new Person("name1", 18);
        Person person2 = new Person("name2", 18);
        System.out.println(person1);
        System.out.println(person2);
        testObject(person1, person2);
        System.out.println(person1);
        System.out.println(person2);
    }

    public static void testObject(Person p1, Person p2) {
        Person person = p1;
        p1 = p2;
        p2 = person;
    }
}

class Person {
    String name;
    Integer age;
    public Person(String name, Integer age) {
        this.name = name;
        this.age = age;
    }
}
// 输出
other.Person@27c6e487
other.Person@49070868
other.Person@27c6e487
other.Person@49070868
```



