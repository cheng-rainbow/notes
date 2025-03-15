Elasticsearch 是一个分布式搜索和分析引擎。它常用于 **日志分析、搜索引擎、数据分析、监控系统等场景**。Elasticsearch 具有高扩展性，可以存储海量数据，并提供实时搜索和分析能力。



### 1、基本概念

| 概念                 | 说明                                                         |
| -------------------- | ------------------------------------------------------------ |
| **Index（索引）**    | 类似于数据库中的 "库"（Database），存储一类结构化或非结构化数据。 |
| **Mapping**（映射）  | 用于定义索引中字段的**数据结构**和**数据类型**的方式         |
| **Document（文档）** | 存储的基本单元，类似数据库中的一条记录（Row）。              |
| **Shard（分片）**    | 将索引拆分为多个小块，每个 Shard 负责存储部分数据，提高性能和扩展性。 |



```bash
// 拉取elasticsearch镜像	http://localhost:9200/
docker pull docker.elastic.co/elasticsearch/elasticsearch:8.6.2

sudo docker run -d --name elasticsearch \
  -p 9200:9200 -p 9300:9300 \
  -e "discovery.type=single-node" \
  -e "xpack.security.enabled=false" \
  docker.elastic.co/elasticsearch/elasticsearch:8.6.2




// 拉取kibana镜像    http://localhost:5601/
docker pull docker.elastic.co/kibana/kibana:8.6.2

sudo docker run -d --name kibana \
  -p 5601:5601 \
  --link elasticsearch:elasticsearch \
  docker.elastic.co/kibana/kibana:8.6.2


// 下载ki 
// 先手动在浏览器下载对应的zip,   https://release.infinilabs.com/analysis-ik/stable/
sudo docker cp ~/Downloads/elasticsearch-analysis-ik-8.6.2.zip elasticsearch:/tmp

sudo docker exec -it --user root elasticsearch /bin/bash
cd /tmp
unzip elasticsearch-analysis-ik-8.6.2.zip -d analysis-ik
mv analysis-ik /usr/share/elasticsearch/plugins/

chown -R elasticsearch:elasticsearch /usr/share/elasticsearch/plugins/analysis-ik

exit
sudo docker restart elasticsearch

// 添加和删除 词
vim IKAnalyzer.cfg.xml
ext.dic  stopword.dic
```



### 2. 对索引的操作

```bash
// 1. 创建带 Mapping 的索引
PUT my_index
{
  "mappings": {
    "properties": {
      "name": {
        "type": "text"
      },
      "age": {
        "type": "integer"
      },
      "created_at": {
        "type": "date",
        "format": "yyyy-MM-dd HH:mm:ss"
      }
    }
  }
}

// 2. 查看某个索引的详细信息
GET my_index

// 3. 查看 Mapping 结构
GET my_index/_mapping

// 4. 删除索引
DELETE my_index
```



### 3. 对文档的操作

```bash
// 1. 插入（新增）文档
POST my_index/_doc/1
{
  "name": "张三",
  "age": 25,
  "created_at": "2025-03-15 10:30:00"
}

// 2. 获取文档
GET my_index/_doc/1

// 3. 删除文档
DELETE my_index/_doc/1

// 4. 更新文档
POST my_index/_update/1
{
  "doc": {
    "age": 26
  }
}
// 4. 更新文档 （先删除id为1的，然后重新插入）
put my_index/_doc/1
```



### 4. SpringBoot中使用ES

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-data-elasticsearch</artifactId>
</dependency>
```

```java
@Data
@AllArgsConstructor
@NoArgsConstructor
@Document(indexName = "products") // 指定索引名称
public class Product {
    @Id
    private String id;

    @Field(type = FieldType.Text)
    private String name;

    @Field(type = FieldType.Text)
    private String description;

    @Field(type = FieldType.Double)
    private Double price;
}
```

```java
public interface ProductMapper extends ElasticsearchRepository<Product, String> {

}
```

```java
public interface ProductService {
    Product save(Product product);

    void delete(String id);

    Product findById(String id);

}
```

```java
@Service
public class ProductServiceImpl implements ProductService{
    @Autowired
    private ProductRepository productRepository;

    @Override
    public Product save(Product product) {
        return productRepository.save(product);
    }

    @Override
    public void delete(String id) {
        productRepository.deleteById(id);
    }

    @Override
    public Product findById(String id) {
        return productRepository.findById(id).orElse(null);
    }
}
```

```java
package com.es.learnes.controller;

import com.es.learnes.pojo.Product;
import com.es.learnes.service.ProductService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/products")
public class ProductController {

    @Autowired
    private ProductService productService;

    @PostMapping
    public Product create(@RequestBody Product product) {
        return productService.save(product);
    }

    @GetMapping("/{id}")
    public Product getById(@PathVariable String id) {
        return productService.findById(id);
    }

    @DeleteMapping("/{id}")
    public void delete(@PathVariable String id) {
        productService.delete(id);
    }
}
```



### 5. 查询语句

```json
// 1. 查询所有
// 可以返回索引中的所有文档
GET products/_search
{
    "query": {
        "match_all": {}
    }
}

// 2. 查询包含特定词的
// 查找名称中包含 “手表” 的产品
GET products/_search
{
    "query": {
        "match": {
            "name": "手表"
        }
    }
}

// 2. 多个字段中进行全文搜索
// name 和 description 字段中同时搜索包含 “无线” 的产品
GET products/_search
{
    "query": {
        "multi_match": {
            "query": "无线",
            "fields": ["name", "description"]
        }
    }
}

// 3. 精确匹配，一般是数值，keyword，id类型
// 查找 id 为 3 的产品
GET products/_search
{
    "query": {
        "term": {
            "id": "3"
        }
    }
}

// 3. 按照多个精确id匹配
// 查找 id 为 3 或 4 的产品
GET products/_search
{
    "query": {
        "terms": {
            "id": ["3", "4"]
        }
    }
}

// 4. 查询某个范围
// 查找价格在 200 到 500 之间的产品
GET products/_search
{
    "query": {
        "range": {
            "price": {
                "gte": 200,
                "lte": 500
            }
        }
    }
}

// 5. 组合多条件查询
// 查找名称中包含 “耳机” 且价格小于 300 的产品
GET products/_search
{
    "query": {
        "bool": {
            "must": [	// must 相当于 与；should相当于 或； must_not 非，不参与算分； filter 相当于 与，但不参与算分
                {
                    "match": {
                        "name": "耳机"
                    }
                },
                {
                    "range": {
                        "price": {
                            "lt": 300
                        }
                    }
                }
            ]
        }
    }
}

// 6. 排序查询，默认是按照权重分pai'de
// 将查询结果按价格升序排列
GET products/_search
{
    "query": {
        "match_all": {}
    },
    "sort": [
        {
            "price": {
                "order": "asc"
            }
        }
    ]
}

// 7. 分页查询
// 使用 from 和 size 参数实现分页，from 表示起始偏移量，size 表示每页显示的文档数量
GET products/_search
{
    "query": {
        "match_all": {}
    },
    "from": 2,
    "size": 2
}

// 8. 字段过滤查询
// 使用 _source 字段指定要返回的字段
GET products/_search
{
    "query": {
        "match_all": {}
    },
    "_source": ["name", "price"]
}

// 9. 聚合查询
// 按 category 字段分组并计算每组的平均价格
GET products/_search
{
    "query": {
        "match_all": {}
    },
    "aggs": {
        "category_groups": {
            "terms": {
                "field": "category"
            },
            "aggs": {
                "avg_price": {
                    "avg": {
                        "field": "price"
                    }
                }
            }
        }
    }
}
```

