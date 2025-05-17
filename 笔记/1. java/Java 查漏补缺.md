## 一、Java IO流

Java 的 I/O（输入/输出）流是 Java 提供的一种处理数据输入和输出的机制，用于从各种来源（如文件、网络、内存、键盘等）读取数据或将数据写入目标（如文件、网络、屏幕等）。

Java 的 I/O 流体系结构灵活且功能强大，位于 java.io 包中，分为 **字节流** 和 **字符流** 两大类，并支持多种流类型（如缓冲流、数据流、对象流等）以满足不同场景的需求。

按数据单位：

- 字节流：以字节（8 bit）为单位，处理二进制数据（如图片、视频、文本等）。
- 字符流：以字符（16 bit，基于 Unicode）为单位，处理文本数据。



每次调用读取或写入都会发起一次系统调用（如 Linux 的 read 或 Windows 的 ReadFile）。系统调用涉及用户态到内核态的切换，成本较高（包括上下文切换、权限检查等）。

即使 FileOutputStream 没有 Java 层的缓冲区，操作系统通常也会将数据缓存到**内核缓冲区**，而不是立即写入物理磁盘。`close()方法调用时` 会自动调用 flush()，确保内部缓冲区（8KB）和内核缓冲区的数据写入磁盘。



### 1. 字节流

字节流以字节（8 bit）为单位，适合处理所有类型的二进制数据。字节流的抽象基类是 `InputStream` 和 `OutputStream`。



常见的字节流有：

`FileInputStream`：从文件读取字节数据。

`FileOutputStream`：向文件写入字节数据。

```java
import java.io.*;

public class TestFileInputStream {
    public static void main(String[] args) {
        try (FileInputStream fis = new FileInputStream("1.jpg");
             FileOutputStream fos = new FileOutputStream("./tmp/output.jpg");) {

            int bt = 0;
            while ((bt = fis.read()) != -1) {
                fos.write(bt);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
```

也可以自己弄一个缓冲区，减少系统调用次数

```java
import java.io.*;

public class TestFileInputStream {
    public static void main(String[] args) {
        try (FileInputStream fis = new FileInputStream("1.jpg");
             FileOutputStream fos = new FileOutputStream("./tmp/output.jpg");) {
            byte[] buffer = new byte[1024];

            int len = 0;
            while ((len = fis.read(buffer)) != -1) {
                fos.write(buffer, 0, len);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
```



`BufferedInputStream`：为输入流添加缓冲区，减少对底层系统的直接访问，提升性能。默认会是一个8KB的缓冲区

`BufferedOutputStream`：为输出流添加缓冲区，提升性能。

```java
import java.io.*;

public class TestBufferInputStream {
    public static void main(String[] args) {
        try (BufferedInputStream bis = new BufferedInputStream(new FileInputStream("1.jpg"));
             BufferedOutputStream bos = new BufferedOutputStream(new FileOutputStream("output.jpg"))) {
            int byteRead;   // 按照字节一个一个读, 字节的值是0-255
            while ((byteRead = bis.read()) != -1) {
                bos.write(byteRead);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```



`ObjectInputStream`：反序列化对象（从字节流恢复对象）。

`ObjectOutputStream`：序列化对象（将对象转为字节流）。

```java
public class TestObjectInput {
    public static void main(String[] args) {
        Person person1 = new Person("ld", 22);
        try (
        ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream("./person.data"))) {
            System.out.println(person1.name + " " + person1.age);
            oos.writeObject(person1);
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }

        Person person2 = null;
        try (ObjectInputStream ois = new ObjectInputStream(new FileInputStream("./person.data"))) {
            person2 = (Person) ois.readObject();
            System.out.println(person1.name + " " + person1.age);
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        } catch (IOException e) {
            throw new RuntimeException(e);
        } catch (ClassNotFoundException e) {
            throw new RuntimeException(e);
        }
    }
}

class Person implements Serializable {
    public String name;
    public int age;

    public Person(String name, int age) {
        this.name = name;
        this.age = age;
    }
}
```



字节流的主要方法：

输入流（InputStream）：

- `int read()`：读取一个字节，返回 -1 表示流结束。
- `int read(byte[] b)`：读取多个字节到数组，返回读取的字节数。
- `int read(byte[] b, int off, int len)`：读取指定长度字节到数组。
- `void close()`：关闭流，释放资源。

输出流（OutputStream）：

- `void write(int b)`：写入一个字节。
- `void write(byte[] b)`：写入字节数组。
- `void write(byte[] b, int off, int len)`：写入数组的部分字节。
- `void flush()`：强制刷新缓冲区。
- `void close()`：关闭流。



### 2. 字符流

字符流以字符（Unicode，16 bit）为单位，专门用于处理文本数据，能够自动处理字符编码。字符流的抽象基类是 `Reader` 和 `Writer`。



常见的字符流有：

`FileReader`：从文本文件读取字符数据。

`FileWriter`：向文本文件写入字符数据。

```java
import java.io.*;

public class TestFileReader {
    public static void main(String[] args) {
        try (FileReader fr = new FileReader("./1.txt");
             FileWriter fw = new FileWriter("./tmp/output.txt")) {

            int c;
            while((c = fr.read()) != -1) {
                fw.write(c);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
```

也可以自己弄一个缓冲区

```java
import java.io.*;

public class TestFileReader {
    public static void main(String[] args) {
        try (FileReader fr = new FileReader("./1.txt");
             FileWriter fw = new FileWriter("./tmp/output.txt")) {

            int len = 0;
            char[] buffer = new char[1024];
            while((len = fr.read(buffer)) != -1) {
                fw.write(buffer, 0, len);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
```



`BufferedReader`：为输入流添加缓冲区，支持按行读取（readLine()）。

`BufferedWriter`：为输出流添加缓冲区，支持按行写入（newLine()）。

```java
import java.io.*;

public class TestBufferReader {
    public static void main(String[] args)  {
        try (BufferedReader br = new BufferedReader(new FileReader("./1.txt"));
             BufferedWriter bw = new BufferedWriter(new FileWriter("./tmp/output.txt"))) {
            String s;
            while ((s = br.readLine()) != null) {
                bw.write(s);
                bw.newLine();   // 换行，默认没有
            }

        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
```



`InputStreamReader`：将字节输入流转换为字符输入流，支持指定编码。

`OutputStreamWriter`：将字符输出流转换为字节输出流，支持指定编码。

```java
import java.io.*;
import java.nio.charset.StandardCharsets;

public class TestInputStreamReader {
    public static void main(String[] args) {
        try (InputStreamReader isr = new InputStreamReader(new FileInputStream("./1.txt"), StandardCharsets.UTF_8);
             OutputStreamWriter osw = new OutputStreamWriter(new FileOutputStream("./tmp/output2.txt"))) {
            int c;
            while ((c = isr.read()) != -1) {
                osw.write(c);
            }
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
```



字符流主要方法：

输入流（Reader）：

- `int read()`：读取一个字符，返回 -1 表示流结束。
- `int read(char[] cbuf)`：读取多个字符到数组，返回读取的字符数。
- `void close()`：关闭流。

输出流（Writer）：

- `void write(int c)`：写入一个字符。
- `void write(char[] cbuf)`：写入字符数组。
- `void write(String str)`：写入字符串。
- `void flush()`：刷新缓冲区。
- `void close()`：关闭流。



## 二、集合

### 1. 集合的遍历

```java
public class TestCollection {
    public static void main(String[] args) {
        List<String> fruitList = new ArrayList<>();
        fruitList.add("Apple");
        fruitList.add("Banana");
        fruitList.add("Orange");
        fruitList.add("Apple");

//        for (int i = 0; i < fruitList.size(); i ++) {
//            System.out.println(fruitList.get(i));
//        }

//        for (String item : fruitList) {
//            System.out.println(item);
//        }

//        Iterator<String> iterator = fruitList.iterator();
//        while(iterator.hasNext()) {
//            System.out.println(iterator.next());
//        }

//        fruitList.forEach(item -> System.out.println(item));

        Set<String> colorSet = new HashSet<>();
        colorSet.add("Red");
        colorSet.add("Blue");
        colorSet.add("Green");
        colorSet.add("Red");

//        for(String item : colorSet) {
//            System.out.println(item);
//        }

//        colorSet.forEach(item -> System.out.println(item));

//        Iterator<String> iterator = colorSet.iterator();
//        while (iterator.hasNext()) {
//            System.out.println(iterator.next());
//        }

        Map<Integer, String> studentMap = new HashMap<>();
        studentMap.put(101, "Alice");
        studentMap.put(102, "Bob");
        studentMap.put(103, "Charlie");

//        for (Integer key : studentMap.keySet()) {
//            System.out.println(key);
//        }
//        for (String val : studentMap.values()) {
//            System.out.println(val);
//        }
//        for (Map.Entry<Integer, String> entry : studentMap.entrySet()) {
//            System.out.println(entry.getKey() + " " + entry.getValue());
//        }

//        studentMap.forEach((key, val) -> System.out.println(key + " " + val));
    }
}
```



## 三、多线程

### 1. 基础

```java
// 1. 匿名内部类
new 接口名() {
    // 实现接口中的抽象方法
}
new 父类名() {
    // 重写或实现父类中的方法
}

// 2. Lambda 表达式，仅适用于函数式接口（只有一个抽象方法的接口，如 Runnable）。
new Thread(() -> System.out.println("Running!"))
```



### 2. 创建线程

创建线程，执行**普通的任务**

```java
public class TestThread {
    private static int sum = 0; // 共享变量

    public static void main(String[] args) {
        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 10000; i++) {
                sum += 1;
            }
        });

        Runnable task = new Runnable() {
            @Override
            public void run() {
                for (int i = 0; i < 10000; i++) {
                    sum += 1;
                }
            }
        };
        Thread t2 = new Thread(task);

        t1.start();
        t2.start();

        try {
            t1.join();
            t2.join();
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        System.out.println(sum);
    }
}
```

创建线程，执行**带有返回值的任务**

```java
import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

public class Main {

    static int sum = 0;

    public static void main(String[] args) throws ExecutionException, InterruptedException {
        FutureTask<Integer> ft1 = new FutureTask<>(() -> {
            for (int i = 0; i < 10000; i++) {
                sum += 1;
            }
            return sum;
        });

        FutureTask<Integer> ft2 = new FutureTask<>(() -> {
            for (int i = 0; i < 10000; i++) {
                sum += 1;
            }
            return sum;
        });

        Thread t1 = new Thread(ft1);
        Thread t2 = new Thread(ft2);
        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println(ft1.get());
        System.out.println(ft2.get());
    }
}
```

对于一个 `Future` 对象。`Future` 是一个用于表示异步计算结果的接口，提供了如下几个重要方法：

- **`get()`**：获取任务的执行结果。如果任务未完成，会阻塞直到任务完成并返回结果。如果任务执行过程中抛出了异常，`get()` 会抛出 `ExecutionException`。

  ```java
  Integer result = future.get(); // 阻塞直到任务完成
  ```

- **`cancel()`**：取消任务的执行。如果任务正在运行，取消任务会尝试停止任务的执行。需要注意，`cancel()` 不一定能立即终止任务，任务执行的具体中断方式取决于任务本身的实现。

  ```java
  boolean isCancelled = future.cancel(true); // 取消任务，尝试中断
  ```

- **`isDone()`**：检查任务是否已经完成，返回 `true` 表示任务已经完成，不论是正常完成还是抛出异常。

  ```java
  boolean done = future.isDone();
  ```

- **`isCancelled()`**：检查任务是否已被取消，返回 `true` 表示任务已经取消。

  ```java
  boolean cancelled = future.isCancelled();
  ```



### 3. 线程池

线程池（Thread Pool）是 Java 中用于管理线程的一种机制，旨在提高多线程任务的执行效率，减少线程创建和销毁的开销，并有效控制并发资源的使用。

JUC提供了一些预设好的线程池，如下：

1. `FixedThreadPool`：固定线程数的线程池。

   - 实现：`Executors.newFixedThreadPool(int nThreads)`
   - 特点：
     - 固定线程数（corePoolSize = maximumPoolSize）。
     - 任务队列为 LinkedBlockingQueue（无界）。
     - 适合任务量稳定、需要限制线程数的场景。

   - 风险：无界队列可能导致内存溢出。

2.  `CachedThreadPool`：动态调整线程数的线程池。
   - 实现：`Executors.newCachedThreadPool()`
   - 特点：
     - 线程数动态调整（corePoolSize = 0, maximumPoolSize = Integer.MAX_VALUE）。
     - 任务队列为 SynchronousQueue（无缓冲）。
     - 空闲线程存活 60 秒。
     - 适合短时间、轻量级任务。
   - 风险：高并发下可能创建过多线程，导致资源耗尽。

3. `SingleThreadExecutor`：只有一个线程的线程池。
   - 实现：`Executors.newSingleThreadExecutor()`
   - 特点：
     - 只有一个线程（corePoolSize = maximumPoolSize = 1）。
     - 任务队列为 LinkedBlockingQueue（无界）。
     - 任务按提交顺序执行。
     - 适合需要严格顺序执行的场景。
   - 风险：无界队列可能导致内存溢出。

4.  `ScheduledThreadPool`：持定时任务的线程池。
   - 实现：`Executors.newScheduledThreadPool(int corePoolSize)`
   - 特点：
     - 基于 ScheduledThreadPoolExecutor，支持延迟和周期性任务。
     - 使用 DelayedWorkQueue（优先级队列）管理任务。
     - 适合定时任务或周期性任务。

5. `ForkJoinPool`：于工作窃取的线程池。
   - 实现：`Executors.newWorkStealingPool() 或 new ForkJoinPool()`
   - 特点：
     - 基于工作窃取算法，线程从自己的任务队列或偷取其他线程的任务。
     - 适合递归分治任务（如并行排序、树遍历）。
     - 使用 ForkJoinTask（RecursiveTask 或 RecursiveAction）定义任务。



除了上面预设好的线程池，JUC也支持我们自定义线程池

```java
public ThreadPoolExecutor(
    int corePoolSize,                      // 核心线程数
    int maximumPoolSize,                   // 最大线程数
    long keepAliveTime,                    // 非核心线程空闲存活时间
    TimeUnit unit,                         // 时间单位
    BlockingQueue<Runnable> workQueue,     // 任务队列
    ThreadFactory threadFactory,           // 线程工厂
    RejectedExecutionHandler handler       // 拒绝策略
)
```



### 4. 任务队列

任务队列是线程池中用于存储待执行任务的容器，通常是一个实现了 `BlockingQueue<Runnable>` 接口的队列。当线程池中的线程繁忙（核心线程已满）时，新提交的任务会被放入任务队列，等待空闲线程处理。



任务队列的作用：

1. 缓冲任务：当所有核心线程（corePoolSize）都在执行任务时，新任务进入队列，而不是立即创建新线程。
2. 控制并发：通过队列容量限制任务积压，防止系统因过多任务而过载。
3. 支持线程池扩展：队列满时，线程池可能创建非核心线程（直到 maximumPoolSize），动态适应任务负载。
4. 异常情况处理：当队列和线程池都满时，触发拒绝策略（如抛异常、丢弃任务），确保系统稳定性。



任务队列的类型：

`LinkedBlockingQueue`：基于链表实现，FIFO（先进先出）顺序，默认是容量是无界的（`Integer.MAX_VALUE`），也可指定容量

`ArrayBlockingQueue`：基于数组实现，有界队列，固定容量

`SynchronousQueue`：无缓冲队列，不存储任务。每个 put 操作必须等待一个 take 操作（任务必须立即交给线程执行）。

`PriorityBlockingQueue`：基于堆实现，非 FIFO 顺序，无界优先级队列，任务按优先级排序（需实现 Comparable 或提供 Comparator）。

`DelayedWorkQueue`：专为 ScheduledThreadPoolExecutor 设计，基于优先级队列（最小堆）。任务按延迟时间排序，支持延迟和周期性任务。



### 5. 给任务分配线程

`execute()` ：接受一个实现了 `Runnable` 接口的任务，并将任务提交给线程池执行。（没有返回值且没有异常处理的任务。）

```java
Executor executor = Executors.newFixedThreadPool(3);
executor.execute(new RunnableTask());
```

`submit()` ：接受一个实现了 `Runnable` 或 `Callable` 接口的任务，并将任务提交给线程池。（可以返回结果或者需要捕获异常的任务。）

```java
ExecutorService executor = Executors.newFixedThreadPool(3);

// 提交一个返回结果的任务
Future<Integer> future = executor.submit(new CallableTask());

// 获取任务的执行结果
try {
    Integer result = future.get(); // 阻塞，直到任务执行完成
    System.out.println("Task result: " + result);
} catch (InterruptedException | ExecutionException e) {
    e.printStackTrace();
}
```

`invokeAll()` 方法用于提交一个任务集合，执行所有任务，并等待所有任务完成。（可以返回每个任务的结果，可设置超时时间）

```java
ExecutorService executor = Executors.newFixedThreadPool(3);
List<Callable<Integer>> tasks = Arrays.asList(
    new CallableTask(),
    new CallableTask(),
    new CallableTask()
);

// 提交任务并等待所有任务完成
try {
    List<Future<Integer>> results = executor.invokeAll(tasks);
    for (Future<Integer> result : results) {
        System.out.println("Task result: " + result.get());
    }
} catch (InterruptedException | ExecutionException e) {
    e.printStackTrace();
}
```

`invokeAny()` 方法用于提交一组任务，并返回第一个完成的任务的结果。如果有多个任务可以并行执行，那么 `invokeAny()` 会在第一个任务执行完成时返回该任务的结果，而其他任务会被取消执行。

```java
ExecutorService executor = Executors.newFixedThreadPool(3);
List<Callable<Integer>> tasks = Arrays.asList(
    new CallableTask(),
    new CallableTask(),
    new CallableTask()
);

// 提交任务并返回第一个完成的任务结果
try {
    Integer result = executor.invokeAny(tasks);
    System.out.println("First completed task result: " + result);
} catch (InterruptedException | ExecutionException e) {
    e.printStackTrace();
}
```



### 6. 关闭线程池

`shutdown()` 调用此方法后，线程池会停止接收新的任务，但是已经提交的任务会继续执行，直到所有任务执行完毕。

```java
ExecutorService executor = Executors.newFixedThreadPool(3);
// 提交任务
executor.submit(new Task());
executor.submit(new Task());

// 正常关闭线程池
executor.shutdown();
```

`shutdownNow()` 方法也用于关闭线程池，但它的行为与 `shutdown()` 方法不同。`shutdownNow()` 会尽量停止正在执行的任务，尝试中断正在运行的任务，并把尚未执行的任务列表返回。

```java
ExecutorService executor = Executors.newFixedThreadPool(3);
// 提交任务
executor.submit(new Task());
executor.submit(new Task());

// 尝试立即关闭线程池
List<Runnable> notExecutedTasks = executor.shutdownNow();
```



## 四、多线程带来的问题

### 1. 原子性

**概念**：一个操作或一组操作要么全部执行完成，要么完全不执行，不会出现部分执行的状态。

**解决办法**：

`volatile`关键字：synchronized 通过对象锁（monitor）确保同一时间只有一个线程进入临界区，保护整个代码块的操作不被中断。

`synchronized`关键字：ReentrantLock 提供显式锁机制，类似 synchronized，通过锁的获取和释放确保临界区的原子性。



### 2. 可见性

**概念**：一个线程修改了共享变量的值，但对另一个线程不可见

**产生原因**：

1. 每个线程有自己的工作内存（类似于CPU缓存或寄存器），用于缓存共享变量的副本，以提高性能。线程对共享变量的读/写操作通常先在工作内存中进行，然后再与主内存同步。

2. 现代CPU和JVM为了性能优化，可能会延迟变量的写回（store buffer）或提前读取（load buffer）。这种优化在单线程中无害，但在多线程中可能导致线程间数据不同步。

**解决方法**：

`volatile`关键字：写入 volatile 变量时，JVM插入写屏障，立即将值刷新到主内存。读取 volatile 变量时，JVM插入读屏障，直接从主内存获取最新值。

`synchronized`关键字：进入 synchronized 块时，清空线程工作内存中的变量副本，从主内存加载最新值。退出 synchronized 块时，将工作内存中的变量值刷新到主内存。

`ReentrantLock` 等锁实现基于AQS（AbstractQueuedSynchronizer），在获取和释放锁时会同步内存。



### 3. 有序性

**概念**：编译器和CPU可能会对代码的执行顺序进行调整，以提高指令流水线效率。

**示例**：

如果线程A执行的操作被重排序，线程B可能看到不一致的执行顺序。

```java
// 双检锁单例模式
public class Singleton {
    private static Singleton instance;

    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton(); // 可能被重排序
                }
            }
        }
        return instance;
    }
}
```

`instance = new Singleton()` 包含三个步骤：

1. 分配内存。
2. 初始化对象。
3. 将 instance 指向内存地址。

如果步骤2和3被重排序，线程B可能看到未初始化的对象（instance != null 但对象未完成初始化）。

**解决方法**：

`volatile`关键字：写入 volatile 变量时，JVM插入写屏障，确保之前的写操作完成后才写入 volatile 变量。读取 volatile 变量时，JVM插入读屏障，确保读取 volatile 变量后才执行后续操作。

`synchronized `关键字：synchronized 块内的代码不会与块外的代码重排序，进入和退出 synchronized 块时会插入内存屏障。

`ReentrantLock` 在加锁和解锁时插入内存屏障，防止锁内代码与外部代码重排序。



### 4. 深入理解

1. `volatile`

```java
public class Singleton {
    private volatile static Singleton instance;

    public static Singleton getInstance() {
        if (instance != null) { // 读取 volatile 变量
            return instance;   // 后续操作
        }
        synchronized (Singleton.class) {
            if (instance == null) {
                instance = new Singleton(); // 写 volatile 变量
            }
        }
        return instance;
    }
}
```

```java
1. LoadLoadBarrier;           // 屏障：确保从主内存读取最新值
2. value = instance;          // 读取 volatile 变量
3. LoadStoreBarrier;          // 屏障：确保后续写操作不提前
4. // 后续操作（如 return instance）
```

```java
// JVM实际执行（概念性表示）
1. memory = allocate();        // 分配内存
2. init(memory);              // 初始化对象
3. StoreStoreBarrier;         // 屏障：确保步骤1、2完成
4. instance = memory;         // 写 volatile 变量
5. StoreLoadBarrier;          // 屏障：确保写入主内存
```

2. `synchronized`

```java
public class Counter {
    private int count = 0;

    public void increment() {
        synchronized (this) {
            count++; // 读-修改-写操作
        }
    }

    public int getCount() {
        synchronized (this) {
            return count; // 读操作
        }
    }
}
```

```java
1. LoadLoadBarrier;           // 屏障：进入 synchronized 块，从主内存加载变量
2. LoadStoreBarrier;          // 屏障：确保后续写操作基于最新值
3. count = count + 1;         // 读-修改-写操作
4. StoreStoreBarrier;         // 屏障：确保修改完成
5. StoreLoadBarrier;          // 屏障：退出 synchronized 块，刷新到主内存
```

进入时（读屏障）：

- LoadLoadBarrier 和 LoadStoreBarrier 确保从主内存加载 count 的最新值，防止读取到工作内存中的旧值。
- 保证临界区内的读操作（如读取 count）基于最新数据。

退出时（写屏障）：

- StoreStoreBarrier 确保 count++ 的修改完成。
- StoreLoadBarrier 将 count 的新值刷新到主内存，并防止临界区外的读/写操作与临界区内的操作重排序。
