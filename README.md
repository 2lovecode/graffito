### GO语言实践

#### 一. 用法
##### 构建执行
  1. 克隆代码
        - git clone git@github.com:2lovecode/graffito.git
  2. 进入项目根目录
        - cd graffito
  3. 构建
        - go build
  4. 构建完成后，会在当前目录生成可执行文件
        - windows: 
            - graffito.exe
        - unix:
            - graffito
  5. 运行
        - windows: 
            - ./graffito.exe -m tools -t hello param1 param2
        - unix: 
            - ./graffito -m tools -t hello param1 param2

##### 开发执行           
  1. 克隆代码
    - git clone git@github.com:2lovecode/graffito.git
  2. 进入项目根目录
    - cd graffito
  3. 运行
    - go run main.go -m tools -t hello param1 param2
    
#### 二. 功能列表

##### 1.工具类
| 功能  | 所属模块  | 用法                       |简介       |
| :---  | :------  | :------------------------  | :-------------- |
| hello |  tools   | -m tools -t hello          |                  |
| ccnum |  tools     | -m tools -t ccnum abc 你好 | 统计字符串字符数 |

##### 2.算法与数据结构
| 功能  | 所属模块  | 用法                       |简介       |
| :---  | :------  | :------------------------  | :-------------- |
| fibo | alg | -m alg -t fibo  | 两种Fibonacci生成方式对比 |
| qsort | alg | -m alg -t qsort  | 快速排序实现 |
| heap | alg | -m alg -t heap  | 最大堆实现 |
| hash | alg | -m alg -t hash  | 哈希表实现 |

##### 3.设计模式
|  功能 | 命令                       |简介       |
| :---  | :------------------------  | :-------------- |
| singleton | -m pattern -t singleton          |   单例模式实现 |
| factory | -m pattern -t factory          |   工厂模式实现 |
| builder | -m pattern -t builder          |   构建者模式实现 |

##### 4.一些实验性代码
| 功能  | 所属模块  | 用法                       |简介       |
| :---  | :------  | :------------------------  | :-------------- |
| parallel | exp | -m exp -t parallel  | 具有依赖关系的并行调用解决方案 |

##### 5.练习demo
| 功能  | 所属模块  | 用法                       |简介       |
| :---  | :------  | :------------------------  | :-------------- |
| pattern | practice | -m practice -t pattern | 设计模式实践 |
| zktest | practice | -m practice -t zktest test1 | zookeepr连接demo |
| cache | practice | -m practice -t cache | 缓存穿透及解决方案对比 |
| list | practice | -m practice -t list | list包使用 |
| string | practice | -m practice -t string | string使用 |
| json | practice | -m practice -t json | json使用 |
| map | practice | -m practice -t map | map使用 |
| struct | practice | -m practice -t struct | struct使用 |
| interface | practice | -m practice -t interface | interface使用 |
| func | practice | -m practice -t func | func使用 |
| arr | practice | -m practice -t arr | array使用 |
| fcopy | practice | -m practice -t fcopy | file使用 |
| channel | practice | -m practice -t channel | channel使用 |
| routine | practice | -m practice -t routine | routine使用 |
| gin | practice | -m practice -t gin | gin使用 |
| time | practice | -m practice -t time | time使用 |
