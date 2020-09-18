## GO语言实践

### 一. 用法
  1. 克隆代码
        - git clone git@github.com:2lovecode/graffito.git
  2. 进入项目目录
        - cd graffito
  3. 进入某个模块目录 如: algorithm
        - cd algorithm
  4. 运行某个文件 如: heap.go
        - go run heap.go
    
### 二. 目录说明

| 目录  | 说明  | 备注                       |
| :---:  | :------:  | :------------------------  |
| algorithm |   数据结构&算法  ||
| experiment |  一些实验性想法的实现     ||
| pattern |  设计模式实现     ||
| practice | 编程练习代码     ||

### 三. 模块列表

#### 数据结构&算法
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [数组实现](algorithm/array.go) |  |  |
| [斐波那契数列](algorithm/fibonacci.go) |  | |
| [哈希表](algorithm/hash.go) |  |  |
| [堆实现](algorithm/heap.go) |  |  |
| [快排实现](algorithm/quicksort.go) |  |  |

#### 实验代码
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [缓存 穿透/击穿/雪崩 解决方案](experiment/cache.go)|  |  |
| [事件绑定与触发](experiment/event.go)|  |  |

#### 设计模式
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [建造者](pattern/builder.go) |  |  |
| [工厂模式](pattern/factory.go) |  |  |
| [单例](pattern/singleton.go) |  |  |

#### 练习代码
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [数组](practice/array.go)|  |  |
| [通道](practice/channel.go)|  |  |
| [文件读写](practice/file.go)|  |  |
| [函数](practice/func.go)|  |  |
| [gin框架使用demo](practice/gin.go)|  |  |
| [http实现demo](practice/http.go)|  |  |
| [接口](practice/interface.go)|  |  |
| [json](practice/json.go)|  |  |
| [切片](practice/list.go)|  |  |
| [map](practice/map.go)|  |  |
| [协程](practice/routine.go)|  |  |
| [字符串](practice/string.go)|  |  |
| [结构体](practice/struct.go)|  |  |
| [异步简单实现](practice/sync.go)|  |  |
| [time包](practice/time.go)|  |  |
| [zookeeper连接demo](practice/zookeeper.go)|  |  |

#### 做的小工具
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [计算中文字符个数](tools/count.go)|  |  |
