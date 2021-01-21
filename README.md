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
| [数组实现](algorithm/array/main.go) |  |  |
| [斐波那契数列](algorithm/fibonacci/main.go) |  | |
| [哈希表](algorithm/hash/main.go) |  |  |
| [堆实现](algorithm/heap/main.go) |  |  |
| [LeetCode Top100](algorithm/leetcode_top100) |  |  |
| [快排实现](algorithm/quicksort/main.go) |  |  |

#### 实验代码
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [缓存 穿透/击穿/雪崩 解决方案](experiment/cache/main.go)|  |  |
| [事件绑定与触发](experiment/event/main.go)|  |  |
| [有依赖关系的并行模型](experiment/httpdeps/main.go)|  |  详情见：https://github.com/2lovecode/depends|
| [一种批量多样数据的处理模式](experiment/mode0/main.go)|  |  |

#### 设计模式
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [建造者](pattern/builder/main.go) |  |  |
| [工厂模式](pattern/factory/main.go) |  |  |
| [对象池模式](pattern/obj_pool/main.go) |  |  |
| [观察者模式](pattern/observer/main.go) |  |  |
| [单例](pattern/singleton/main.go) |  |  |
| [策略模式](pattern/strategy/main.go) |  |  |

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

#### 做的小工具
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [计算中文字符个数](tools/count/main.go)|  |  |
