Protocol Buffers (简称 protobuf) 是 Google 开发的一种高效的数据序列化格式，用于结构化数据的序列化和反序列化。

工作原理：protobuf 使用 `.proto` 文件定义数据结构，然后通过 protobuf 编译器生成对应语言的代码

基本消息定义示例：

```protobuf
syntax = "proto3";

//option java_multiple_files = true;
//option java_package = "com.proto.message";
option java_outer_classname = "ContactsMessage";

message PeopleInfo {
  string name = 1;
  int32 age = 2;
}
```



## 一、消息定义规则

- 字段类型分为：标量数据类型和特殊类型（包括枚举、其他消息类型等）
- 字段唯一编号：用来标识字段，一旦开始使用就不能够再改变。



### 1. 标量数据类型

| protobuf 类型 | Java 对应类型 |                         说明                         |
| :-----------: | :-----------: | :--------------------------------------------------: |
|    double     |    double     |                      双精度浮点                      |
|     float     |     float     |                      单精度浮点                      |
|     int32     |      int      |               可变长度编码，负数效率低               |
|     int64     |     long      |               可变长度编码，负数效率低               |
|    uint32     |      int      | 无符号32位整数（Java中没有无符号类型，使用int表示）  |
|    uint64     |     long      | 无符号64位整数（Java中没有无符号类型，使用long表示） |
|    sint32     |      int      |            有符号32位整数，负数编码效率高            |
|    sint64     |     long      |            有符号64位整数，负数编码效率高            |
|    fixed32    |      int      |          固定4字节，大于2^28时比uint32高效           |
|    fixed64    |     long      |          固定8字节，大于2^56时比uint64高效           |
|   sfixed32    |      int      |                 固定4字节有符号整数                  |
|   sfixed64    |     long      |                 固定8字节有符号整数                  |
|     bool      |    boolean    |                        布尔值                        |
|    string     |    String     |                  UTF-8或ASCII字符串                  |
|     bytes     |  ByteString   |                     任意字节序列                     |



### 2. 特殊类型

1. `repeated` 表示一个 **数组字段**，它可以通过 `addPhone(value)，addAllPhone(collection)，setPhone(index, value)，clearPhone()` 等控制这个列表

2. `map`表示键值对集合，键可以是 `string` 或整型，值可以是任意类型。

```java
UserProfile.Builder builder = UserProfile.newBuilder();

// 添加键值对
builder.putAttributes("role", "admin");
builder.putAttributes("language", "java");
builder.putAllAttributes(Map.of("key1", "value1", "key2", "value2"));

// 获取 Map
Map<String, String> attrs = builder.getAttributesMap();  // 返回不可修改的 Map
```

3. `any`可以存储 **任意类型的 Protobuf 消息**，类似 Java 的 `Object`。

```java
import "google/protobuf/any.proto";

message Wrapper {
  google.protobuf.Any data = 1;
}
------
// 包装一个消息
User user = User.newBuilder().setName("Alice").build();
Wrapper wrapper = Wrapper.newBuilder()
    .setData(Any.pack(user))  // 将 User 消息打包到 Any
    .build();

// 解包 Any
if (wrapper.getData().is(User.class)) {
  User unpackedUser = wrapper.getData().unpack(User.class);
  System.out.println(unpackedUser.getName());  // 输出 "Alice"
}
```



### 3. 字段唯一编号

`1~536,870,911`(2^29-1)，其中 **19000～19999** 不可用。

19000～19999不可用是因为：在Protobuf协议的实现中，对这些数进行了预留。如果非要在 `.proto` 文件中使用这些预留标识号，例如将name字段的编号设置为19000，编译时就会报警



值得一提的是，范围为1～15的字段编号需要一个字节进行编码，16～2047 内的数字需要两个字节进行编码。编码后的字节不仅只包含了编号，还包含了字段类型。所以1～15要用来记出现非常频繁的字段，要为将来有可能添加的、频繁出现的字段预留一些出来。



|   特性   | protobuf |  JSON  |   XML    |
| :------: | :------: | :----: | :------: |
|   编码   |  二进制  |  文本  |   文本   |
|   大小   |    小    |   大   |   最大   |
| 解析速度 |    快    |   慢   |   最慢   |
|  可读性  |    差    |   好   |    好    |
| 模式演进 |   支持   | 不支持 | 有限支持 |
| 语言支持 |    多    |  所有  |   所有   |



## 二、编译

```bash
protoc --java_out=common/pojo contacts.proto -I common/proto
[java输出路径] [proto文件名] [.proto所在目录]

protoc --java_out=src/main/java/com/common/java contacts.proto -I src/main/java/com/common/proto
```



```java
public final class ContactsMessage {
    public static final class PeopleInfo {
        
        // 创建一个Builder
        public static Builder newBuilder() {}	
        
        //序列化
        public byte[] toByteArray(){} 
        ByteString toByteString();
        public void writeTo(final CodedOutputStream output)
        
        // 反序列化
        public static ContactsMessage.PeopleInfo parseFrom() {} 
        
        public final boolean isInitialized() // 检查所有字段是否已经初始化
        
        public static final class Builder {
             get、set、clear	// 用set方法填充PeopleInfo的字段
            	
             @java.lang.Override	
             public ContactsMessage.PeopleInfo build() {}	// 用于创建 PeopleInfo
            
             // 合并字段，保留目标消息的未冲突字段，对于 `repeated字段` 将源消息的列表元素追加到目标消息的列表中
             public BuilderType mergeFrom(final CodedInputStream input)
		}
    }
}
```

`newBuidler()` -> `set、get` -> `build` -> `序列化，反序列化`

