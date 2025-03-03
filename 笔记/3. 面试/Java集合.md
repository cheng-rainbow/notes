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

### 1. 如何对map进行快速遍历？

1. 增强for + entrySet/keySet/values

```java
// 使用 entrySet() 遍历
for (Map.Entry<String, Integer> entry : map.entrySet()) {
    String key = entry.getKey();
    Integer value = entry.getValue();
    System.out.println("Key: " + key + ", Value: " + value);
}

// 使用 keySet() 遍历
for (String key : map.keySet()) {
    Integer value = map.get(key);
    System.out.println("Key: " + key + ", Value: " + value);
}

// 使用 values() 遍历
for (Integer value : map.values()) {
    System.out.println("Value: " + value);
}
```

2. 使用 forEach() 方法

```java
// 使用 forEach() 遍历
map.forEach((key, value) -> System.out.println("Key: " + key + ", Value: " + value));
```

3. 使用迭代器（Iterator）

```java
// 使用 Iterator 遍历 entrySet
Iterator<Map.Entry<String, Integer>> iterator = map.entrySet().iterator();
while (iterator.hasNext()) {
    Map.Entry<String, Integer> entry = iterator.next();
    System.out.println("Key: " + entry.getKey() + ", Value: " + entry.getValue());
}
```



### 2. HashMap实现原理介绍一下？

HashMap 的底层实现基于哈希表，核心是一个数组，每个数组元素称为一个“桶”（bucket）,键值对通过哈希计算映射到某个桶。当多个键映射到同一个桶（哈希冲突）时，这些键值对以链表形式存储。Java 8 引入，当某个桶的链表过长（默认超过 8 个节点），会将链表转为红黑树以优化查找性能。



### 3. 了解的哈希冲突解决方法有哪些？

**分离链接法**：每个桶存储一个链表（或其他数据结构，如红黑树）。当多个键映射到同一个桶时，这些键值对被添加到该桶的链表中。查找时，先通过哈希函数定位桶，然后遍历链表查找目标键。

**开放寻址法**：当发生冲突时，按照某种探测策略在数组中寻找下一个空槽插入。查找时也按照相同的策略，直到找到目标键或空槽（表示未找到），常见的探测策略有线性探测，平方探测，双重哈希。

**再哈希**：使用多个哈希函数。如果第一个哈希函数发生冲突，使用另一个哈希函数重新计算位置，直到找到空位。

**哈希桶扩容**：当哈希冲突过多时，可以动态地扩大哈希桶的数量，重新分配键值对，以减少冲突的概率。



### 4. HashMap是线程安全的吗？

HashMap 是线程不安全的。

1. 数据覆盖和丢失：多个线程同时调用 put 方法可能导致数据覆盖。

2. 扩容时的并发问题：在扩容（resize）时，HashMap 需要重新分配所有元素到新的数组，可能出现链表死循环或数据丢失

3. 迭代时的 `ConcurrentModificationException`：HashMap 的迭代器是 fail-fast 的。如果在迭代过程中，另一个线程修改了 HashMap 的结构（增删元素），会导致程序抛出 `ConcurrentModificationException` 异常



### 5. hashmap的put过程介绍一下

1.  根据要添加的键的哈希码计算在数组中的位置（索引）。
2. 检查该位置是否为空（即没有键值对存在）
   - 槽位为空：直接创建一个新的节点（Node），放入该槽位。
   - 槽位不为空
     - 如果是链表遍历链表，查找是否有键相同的节点。如果找到相同的键，就覆盖旧值；如果没找到，就在链表尾部插入新节点，插入后链表长度超过阈值时，链表会转为红黑树。
     - 如果是红黑树，查找是否有键相同的节点。如果找到相同的键，就覆盖旧值；如果没找到，调用红黑树的插入逻辑，插入新节点。

3. 检查前元素数量是否超过扩容阈值，如果超过了，进行扩容操作



### 6. hashmap 调用get方法一定安全吗？

多线程下存在问题：在 get 方法遍历链表时，另一个线程可能正在调整链表结构（比如插入或删除节点），导致 get 方法遍历出错，甚至抛出异常。如果另一个线程触发了扩容（resize），HashMap 内部数组会被替换，get 方法可能访问到错误的槽位或旧数据。



### 7. HashMap一般用String 做为Key，为什么？

1. String 是不可变的（immutable）。一旦创建，其内容无法更改。果键对象在插入后被修改，导致 hashCode() 改变，那么后续的 get 操作可能无法找到原来的值（因为哈希值变了，槽位索引变了）。

2. String 的 hashCode() 方法实现经过精心设计，能够较好地分布哈希值，减少冲突。
3. String 的 equals() 方法基于内容比较，保证了基于内容的正确比较
4. String 对象在 JVM 中可以通过字符串常量池复用。减少了内存开销



### 8. 为什么HashMap要用红黑树而不是平衡二叉树？





### 9. hashmap key可以为null吗？

可以为 null。

在 HashMap 中，键的哈希值通过 hash() 方法计算。对于 null 键，HashMap 直接返回 0，取余后依然为0，放在槽位 0。



### 10. 自定义类作为 HashMap 的键时，如何正确重写 equals() 和 hashCode() 方法

（1）遵守 equals() 和 hashCode() 的契约

Java 的 Object 类定义了 equals() 和 hashCode() 的契约，必须严格遵守，否则 HashMap 的行为会不正确。契约如下：

- 自反性：对于任何非 null 的对象 x，x.equals(x) 必须返回 true。
- 对称性：对于任何非 null 的对象 x 和 y，如果 x.equals(y) 为 true，则 y.equals(x) 也必须为 true。
- 传递性：对于任何非 null 的对象 x、y 和 z，如果 x.equals(y) 为 true，且 y.equals(z) 为 true，则 x.equals(z) 也必须为 true。
- 一致性：对于任何非 null 的对象 x 和 y，只要 x 和 y 的状态（用于比较的字段）不变，多次调用 x.equals(y) 必须始终返回相同的结果。
- hashCode 一致性：如果 x.equals(y) 返回 true，则 x.hashCode() 必须等于 y.hashCode()。

（2）equals() 和 hashCode() 必须基于相同的字段

equals() 和 hashCode() 方法必须基于相同的字段来计算，否则会破坏契约。例如：

- 如果 equals() 只比较 id，但 hashCode() 使用 id 和 name，可能导致两个 equals() 为 true 的对象有不同的哈希值。



### 11. HashMap的扩容机制介绍一下

容量：HashMap 底层数组的长度，即槽位（bucket）的数量。默认初始容量是 16。

负载因子：决定何时触发扩容的阈值比例，默认是 0.75。表示当元素数量达到容量的 75% 时，触发扩容。

阈值：触发扩容的临界值，计算公式为 容量 x 负载因子



触发条件：

1. 当 `HashMap` 数组中的元素数量 `size` 超过阈值时

2. 插入元素后如果链表长度大于阈值，将要变成红黑树时，如果数组容量小于 64，也会触发扩容，先扩容到64，然后将链表变成红黑树

`HashMap` 会将数组的容量扩大为原来的两倍。将原来数组中的元素重新分配到新数组中，新数组索引是 `index = hash & (newCapacity - 1)`，由于 `newCapacity` 是 `oldCapacity` 的 2 倍，`newCapacity - 1` 的二进制形式比 `oldCapacity - 1` 多了一位高位的1。因此，元素要么留在原索引，要么移动到 原索引 + oldCapacity 的位置。最后数组引用和内部字段元素个数，容量等。



### 12. HashMap的大小为什么是2的n次方大小呢？

1. 方便通过位运算计算索引：当数组长度是 2 的 n 次方时，length - 1 的二进制表示是一串全 1 的位，哈希值的索引计算公式为：index = hash & (length - 1)（相当于hash % length），但位运算比取模运算快得多。

2. 扩容时的均匀性：如果容量是 2 的幂次方，元素的重新分配只需要根据哈希值的高位进行简单调整，而不需要重新计算整个哈希映射。



### 13. 说说hashmap的负载因子

负载因子默认值 0.75 是一个经验折中，适用于大多数场景。当 HashMap 中的元素个数超过了容量的 75% 时，就会进行扩容。

负载因子较小可以提高性能，但增加内存开销。负载因子较大可以节省内存，但可能导致冲突增多，性能下降。



### 14. Hashmap和Hashtable有什么不一样的？

- HashMap线程不安全，效率高一点，可以存储null的key和value，null的key只能有一个，null的value可以有多个。默认初始容量为16，每次扩充变为原来2倍。**底层数据结构为数组+链表**，插入元素后如果链表长度大于阈值（默认为8），先判断数组长度是否小于64，如果小于，则扩充数组，反之将链表转化为**红黑树**，以减少搜索时间。
- HashTable线程安全，效率低一点，其**内部方法基本都经过synchronized修饰**，不可以有null的key和value。默认初始容量为11，每次扩容变为原来的2n+1。创建时给定了初始容量，会直接用给定的大小。**底层数据结构为数组+链表**。它基本被淘汰了，要保证线程安全可以用ConcurrentHashMap。



### 15. ConcurrentHashMap怎么实现的？

采用 CAS 操作和 synchronized 锁来实现线程安全的，通过 volatile 修饰数组元素，保证读线程能看到最新的数据。



### 16. 已经用了synchronized，为什么还要用CAS呢？

CAS 不需要加锁，避免了 synchronized 带来的阻塞和上下文切换开销。在低竞争场景下，CAS 通常只需一次操作即可成功，性能远高于加锁。synchronized 提供了细粒度的锁，适合复杂操作（如链表/红黑树操作），保证线程安全。必要时使用 synchronized（加锁），确保复杂操作的正确性。避免了单一机制的局限性，最大化了并发性能，同时保持了线程安全。



### 17. ConcurrentHashMap用了悲观锁还是乐观锁?

同时使用了乐观锁和悲观锁

乐观锁（CAS）：优先使用，体现其高并发设计思想，适合简单操作、低竞争场景。

悲观锁（synchronized）：作为补充，用于复杂操作或高竞争场景，确保线程安全和数据一致性。

总体倾向：从设计目标（高并发、无锁读、细粒度写）来看，ConcurrentHashMap 更倾向于乐观锁的思想，但通过结合两种锁实现了性能和安全性的平衡。



### 18. HashTable线程安全是怎么实现的？

Hashtable 的所有公共方法都使用 synchronized 关键字修饰（全局锁）。synchronized方法或者访问synchronized代码块时，这个线程便获得了该对象的锁，其他线程暂时无法访问这个方法，只有等待这个方法执行完毕或者代码块执行完毕，这个线程才会释放该对象的锁，其他线程才能执行这个方法或者代码块。

原子性：全局锁保证了每个操作（如 put、get）是原子的，防止多线程并发修改导致数据不一致。

一致性：在操作期间，其他线程无法访问 Hashtable，避免了脏读、幻读等问题。

互斥性：同一时间只有一个线程可以执行任何同步方法，其他线程会被阻塞。





