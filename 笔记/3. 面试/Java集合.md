## 一、 概念

### 1. 集合基本内容

Java 的集合框架主要分为三大类：List、Set 和 Map（严格来说 Map 不属于 Collection 接口，但通常归类为集合框架的一部分）。

**1. List（列表）**

List 是有序的 Collection，这里的“有序”指的是元素按照插入顺序排列，或者通过索引精确指定插入位置。用户可以通过索引（从 0 开始）访问 List 中的元素，也可以通过索引插入或删除元素（例如 add(int index, E element) 和 remove(int index)）。List 允许存储重复元素。

- **ArrayList**：底层是一个数组，每次增加元素时，如果数组空间不足，ArrayList 会创建一个更大的数组并将原数组中的元素复制过来。**支持随机访问**（通过索引访问元素效率高），但插入、删除、扩容性能差
- **LinkedList**：底层是由节点（Node）构成的双向链表。每个节点包含一个数据元素和两个引用（指向前一个节点和后一个节点）。**不支持高效随机访问**，但插入和删除效率高、没有扩容问题。
- **Vector**：底层与 ArrayList 类似，是一个动态数组，但通过 `synchronized` 关键字实现了线程安全。**支持随机访问**，但由于同步机制，性能比 ArrayList 低，插入、删除、扩容性能也较差。现已较少使用。
- **Stack**：是 Vector 的子类，实现了后进先出（LIFO）的栈结构。**支持随机访问**（继承自 Vector），但插入和删除仅限于栈顶操作。由于性能和设计问题，不推荐直接使用，建议用 Deque 的实现（如 ArrayDeque）作为栈替代。

------

**2. Set（集合）**

Set 是一个不允许重复元素的 Collection，元素唯一性由 equals() 和 hashCode() 方法决定。与 List 不同，Set 的元素顺序取决于具体实现，而不是一概无序。

- **HashSet**：使用哈希表来存储元素，它通过哈希值来**快速查找和存取元素**（时间复杂度平均为 O(1)）。不保证元素的顺序，元素的存储顺序与它们的插入顺序无关，因为哈希表使用哈希值来决定元素的存储位置。
- **TreeSet**：使用红黑树（内部基于 TreeMap 实现）。会根据元素的自然顺序（通过元素的 Comparable 接口）或提供的自定义比较器（Comparator）对元素进行排序。**适合需要排序的场景**，但查找和插入性能（O(log n)）不如 HashSet。
- **LinkedHashSet**：是 HashSet 的子类，底层使用哈希表和双向链表结合的方式，既保持了哈希表的快速查找（O(1)），又通过链表维护插入顺序。**适合需要保持插入顺序的场景**，但性能略低于 HashSet（因为需要额外维护链表）。

------

**3. Queue（队列）**

Queue 是 Java 集合框架中的一个接口，继承自 Collection，通常用于表示一种“队列”结构，元素按照特定规则（通常是先进先出 FIFO）添加和移除。

- **PriorityQueue**：是一个基于优先级堆（默认是最小堆）的实现类，它按照元素的自然顺序（通过元素的 Comparable 接口）或者提供的自定义比较器（Comparator）来决定元素的优先级。**适合需要按优先级排序的场景**，出队时返回优先级最高（默认最小）的元素，插入和删除性能为 O(log n)。
- **ArrayDeque**：基于动态数组实现的双端队列（Deque 的子接口实现），支持从两端高效添加和移除元素。比 LinkedList 在大多数场景下性能更好（因为数组连续内存），是 Java 推荐的双端队列和栈实现。**适合 FIFO 或 LIFO 场景**，不支持 null 元素，插入和删除性能为 O(1)。
- **LinkedList**：实现了 Queue 和 Deque 接口，底层是双向链表。**适合频繁插入和删除的场景**，支持从两端操作，作为队列时效率高（O(1)），但不支持高效随机访问（需要遍历）。
- **ConcurrentLinkedQueue**：是线程安全的无界队列，基于链表实现，采用 `CAS`机制实现高效并发。**适合高并发场景**，不支持 null 元素，插入和删除性能为 O(1)，但不保证严格的 FIFO（可能因并发重排序）。
- **ArrayBlockingQueue**：是 BlockingQueue 接口的实现类，基于固定大小的数组实现的有界阻塞队列。支持线程安全的阻塞操作（put() 和 take() 会阻塞线程）。**适合生产者-消费者模型**，支持 null 元素检查，但容量固定，插入和删除性能为 O(1)。
- **LinkedBlockingQueue**：是 BlockingQueue 接口的实现类，基于链表实现的可选有界阻塞队列（可以指定容量上限，也可以无界）。支持线程安全的阻塞操作。**适合生产者-消费者模型**，插入和删除性能为 O(1)，支持 null 元素检查。

---

**4. Map（映射）**

Map 是一个键值对映射集合，存储键（key）、值（value）及它们之间的映射关系。Key 必须唯一（重复的 key 会覆盖旧值），Value 允许重复。Map 不继承 Collection 接口，但属于 Java 集合框架。通过 key 可以快速检索对应的 value（例如 get(Object key) 方法）。Key 和 Value 的顺序取决于具体实现。

- **HashMap**：键值对存储，键唯一，基于哈希表实现，**适合快速查找**（时间复杂度平均为 O(1)）。不保证键的顺序，允许 null 键和 null 值。插入和查找效率高，但无序性可能不适合需要顺序的场景。
- **TreeMap**：使用红黑树（TreeMap 自身就是基于红黑树实现的）。会根据 key 的自然顺序（通过 key 的 Comparable 接口）或提供的自定义比较器（Comparator）对键进行排序。**适合需要按键排序的场景**，但查找和插入性能（O(log n)）不如 HashMap，不允许 null 键。
- **LinkedHashMap**：是 HashMap 的子类，内部通过哈希表和双向链表结合的方式实现，既保持了哈希表的快速查找（O(1)），又维护插入顺序（或访问顺序）。**适合需要保持插入顺序的场景**，允许 null 键和值，性能略低于 HashMap。
- **Hashtable**：类似于 HashMap，基于哈希表实现，但通过 `synchronized` 关键字实现线程安全。**适合简单的线程安全场景**，但性能较低（同步开销大），且不允许 null 键或 null 值。**现已较少使用，推荐用 ConcurrentHashMap 替代**。
- **ConcurrentHashMap**：是线程安全的 HashMap 替代品，基于分段锁（Java 8 后改为 CAS 和同步块）实现高并发。**适合高并发场景**，查找和插入效率高（O(1)），但不允许 null 键或 null 值，性能优于 Hashtable。

> java.util 包中直接提供的线程安全类包括：Vector、Hashtable、Stack、StringBuffer(基于`synchronized` 实现的线程安全 )，其他的线程安全类都来自 **JUC** 包中。



### 2. 数组与集合区别

1. 数组是一种固定长度的、存储同类型元素的数据结构。定义时需要指定长度，且长度不可变。集合是 Java 提供的一组动态数据结构，可以动态增长或缩小
2. 数组可以存储基本数据类型和引用类型变量，集合只能存储引用类型变量
3. 数组可以直接访问元素，而集合需要通过迭代器或其他方法访问元素



### 3. Collections和Collection的区别

**Collection**：接口，定义集合的基本行为，定义了一组通用的操作和方法，如添加、删除、遍历等，用于操作和管理一组对象。Collection接口有许多实现类，如List、Set和Queue等。

**Collections**：工具类，提供一系列静态方法，用于对集合进行高级操作（如排序、查找、替换等）。这些方法可以对实现了Collection接口的集合进行操作



### 4. 集合遍历的方法有哪些？

1. 使用普通 for 循环（仅限 List）

```java
List<String> list = Arrays.asList("A", "B", "C");
for (int i = 0; i < list.size(); i++) {
    System.out.println("Index " + i + ": " + list.get(i));
}
```

2. 增强 `for` 循环

适用于所有实现了 Iterable 接口的集合（Collection 继承了 Iterable，所以 List、Set、Queue 都适用等）。

```java
List<String> list = Arrays.asList("A", "B", "C");
Set<String> set = new HashSet<>(list);

// 遍历 List
for (String item : list) {
    System.out.println(item);
}

// 遍历 Set
for (String item : set) {
    System.out.println(item);
}
```

3. 使用 Iterator（Map不可以）

```java
List<String> list = new ArrayList<>(Arrays.asList("A", "B", "C"));
Set<String> set = new HashSet<>(list);

// 遍历 List 并删除特定元素
Iterator<String> listIterator = list.iterator();
while (listIterator.hasNext()) {
    String item = listIterator.next();
    if ("B".equals(item)) {
        listIterator.remove(); // 安全删除
    } else {
        System.out.println(item);
    }
}


// 遍历 Set
System.out.println("Traversing Set with Iterator:");
Iterator<String> setIterator = set.iterator();
while (setIterator.hasNext()) {
    System.out.println(setIterator.next());
}
```

4. 使用 ListIterator（仅限 List，支持双向遍历）

```java
List<String> list = new ArrayList<>(Arrays.asList("A", "B", "C"));
ListIterator<String> listIterator = list.listIterator();
while (listIterator.hasNext()) {
    System.out.println(listIterator.next());
}

while (listIterator.hasPrevious()) {
    System.out.println(listIterator.previous());
}
```

5. 使用 forEach 方法

Java 8 引入的 forEach 方法，基于 Lambda 表达式或方法引用，简洁且支持函数式编程。

```java
List<String> list = Arrays.asList("A", "B", "C");
Set<String> set = new HashSet<>(list);

// 遍历 List
list.forEach(item -> System.out.println(item));

// 遍历 Set
set.forEach(System.out::println);
```



## 二、List

### 1. list可以一边遍历一边修改元素吗？

`List` 可以一边遍历一边修改元素，但需要根据使用的遍历方式和修改操作的具体情况来判断是否可行

1. 可以使用普通 `for` 循环来遍历 `List` 并修改元素，因为普通 `for` 循环是通过索引来访问元素的，在遍历过程中可以直接通过索引修改元素的值。
2. 使用迭代器遍历 `List`，可以使用迭代器的 `set` 方法来修改当前迭代的元素，但不能使用 `List` 的 `add`、`remove` 等方法，否则会抛出异常。
3. 不建议使用增强 `for` 循环遍历 `List` 并修改元素，因为增强 `for` 循环本质上是使用迭代器实现的，它不允许在遍历过程中修改集合的结构（添加、删除元素），可能会抛出异常。



### 2. 为什么增强 for 不可以修改

```java
for (Integer num : list) {
    // 循环体
}

// 上面代码本质上是下面这个
Iterator<Integer> iterator = list.iterator();
while (iterator.hasNext()) {
    Integer num = iterator.next();
    // 循环体
}
```

Java 的迭代器有一个并发修改检查机制，它会维护一个 `expectedModCount` 变量，用于记录集合结构的预期修改次数。当集合的结构发生改变（例如添加、删除元素）时，`modCount` 变量会增加。在每次调用 `iterator.next()` 方法时，迭代器会检查 `expectedModCount` 和 `modCount` 是否相等，如果不相等，就会抛出 `ConcurrentModificationException` 异常。

对于下面这个代码，并不会抛出异常，但这只是一个巧合。因为 `set` 方法并没有改变集合的结构，`modCount` 没有变化。然而，这种做法是不推荐的，因为它违反了 `foreach` 循环的设计初衷，可能会导致代码的可读性和可维护性变差。

```java
import java.util.ArrayList;
import java.util.List;

public class ModifyListInForeach {
    public static void main(String[] args) {
        List<Integer> list = new ArrayList<>();
        list.add(1);
        list.add(2);
        list.add(3);

        int index = 0;
        for (Integer num : list) {
            list.set(index, num * 2); // 尝试修改元素
            index++;
        }
    }
}
```



### 3. list如何快速删除某个指定下标的元素？

`ArrayList` 是基于数组实现的动态数组，remove 删除指定下标元素时，需要将该下标之后的所有元素向前移动一位

`LinkedList` 是基于双向链表实现的，remove 删除指定下标元素时，需要先遍历到该下标对应的节点，删除节点本身的操作时间复杂度为 *O*(1)。



### 4. 为什么ArrayList不是线程安全的，具体来说是哪里不安全？

扩容时有多个线程同时进行，可能会造成元素丢失或覆盖。

`ArrayList` 使用 `size` 变量来记录当前存储的元素个数，而 `size++` 不是一个原子操作，多个线程同时修改 `size` 可能会导致数组越界

`ArrayList` 的 `Iterator` 不是线程安全的，如果在迭代过程中有另一个线程修改了 `ArrayList`（例如 `add()` 或 `remove()`），`modCount` 变量的变化会触发 `ConcurrentModificationException`。



### 5. 线程安全的 List， CopyonWriteArraylist是如何实现线程安全的

CopyOnWriteArrayList底层也是通过一个数组保存数据，使用volatile关键字修饰数组，保证当前线程对数组对象重新赋值后，其他线程可以及时感知到。在写入操作时，加了一把互斥锁ReentrantLock以保证线程安全。

```java
public boolean add(E e) {
    final ReentrantLock lock = this.lock;
    lock.lock(); // 加锁，保证写操作的线程安全
    try {
        Object[] elements = getArray();
        int len = elements.length;
        Object[] newElements = Arrays.copyOf(elements, len + 1); // 复制新数组
        newElements[len] = e; // 添加新元素
        setArray(newElements); // 替换原数组
        return true;
    } finally {
        lock.unlock(); // 解锁
    }
}
```

首先创建一个新数组，长度是原数组的len + 1，然后将原数组中的数据从头拷贝过来并将要添加的元素放在新数组最后的位置后，用新数组的地址替换掉老数组的地址就能得到最新的数据了。



## 三、Set

### 1. Set集合有什么特点？如何实现key无重复的？

`Set` 是一种 **不允许存储重复元素** 的集合

HashSet和LinkedHashSet都采用 hasCode 和 equals 来去重。插入元素时，计算元素的 `hashCode()`，确定存放位置。如果该位置已有元素（哈希冲突），调用 equals() 方法检查是否相等，相等则认为是重复元素，不存储；不相等则使用链地址法或红黑树存储。

TreeSet依赖  `compareTo()`，相同 `compareTo()` 结果的元素被视为重复，不会添加。



### 2.有序的Set是什么？记录插入顺序的集合是什么？

`TreeSet` 按元素大小排序（**自然排序**或**自定义排序**），`LinkedHashSet` 按插入顺序存储

记录插入顺序的集合通常指的是 `LinkedHashSet`，它不仅保证元素的唯一性，还可以保持元素的插入顺序。



## 四、Map





