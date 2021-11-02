## GO语言实践

### 一. 用法
  1. 编译构建
        - 克隆代码 ```git clone git@github.com:2lovecode/graffito.git```
        - 进入项目目录 ```cd graffito```
        - 构建 ```go build```
  2. 使用各个功能
        - 查看命令帮助 ```./graffito -h```
        - 执行命令，例如 ```./graffito tools count abc``` 计算字符串abc的字符数
    
### 二. 目录说明

| 目录  | 说明  | 备注                       |
| :---:  | :------:  | :------------------------  |
| algorithm |   数据结构&算法  ||
| experiment |  一些实验性想法的实现     ||
| pattern |  设计模式实现     ||
| practice | 编程练习代码     ||
| tools | 特别好用的命令行工具集     ||

### 三. 模块列表

#### 工具列表
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [计算中文字符个数](tools/string_op)|  ```./graffito tools count abc``` |  |
| [简易redis命令行客户端](tools/redis)|  ```./graffito tools redis``` |  |

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
| [缓存 穿透/击穿/雪崩 解决方案](experiment/cache)| ```./graffito exp cache``` |  |
| [事件绑定与触发](experiment/event)|  |  |
| [有依赖关系的并行模型](experiment/depends)|  |  详情见：https://github.com/2lovecode/depends|
| [一种批量多样数据的处理模式](experiment/mode0)|  |  |

#### LeetCode
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [LeetCode](algorithm/leetcode) |  |  |

#### 设计模式
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [建造者](pattern/builder.go) |  |  |
| [工厂模式](pattern/factory.go) |  |  |
| [对象池模式](pattern/obj_pool.go) |  |  |
| [观察者模式](pattern/observer.go) |  |  |
| [单例](pattern/singleton.go) |  |  |
| [策略模式](pattern/strategy.go) |  |  |

#### 练习代码
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [数组](practice/array/main.go)|  |  |
| [通道](practice/channel/main.go)|  |  |
| [文件读写](practice/file/main.go)|  |  |
| [函数](practice/func/main.go)|  |  |
| [gin框架使用demo](practice/gin/main.go)|  |  |
| [http实现demo](practice/http/main.go)|  |  |
| [接口](practice/interface/main.go)|  |  |
| [json](practice/json/main.go)|  |  |
| [切片](practice/list/main.go)|  |  |
| [map](practice/map/main.go)|  |  |
| [协程](practice/routine/main.go)|  |  |
| [字符串](practice/string/main.go)|  |  |
| [结构体](practice/struct/main.go)|  |  |
| [异步简单实现](practice/sync/main.go)|  |  |
| [time包](practice/time/main.go)|  |  |
| [zookeeper连接demo](practice/zookeeper/main.go)|  |  |