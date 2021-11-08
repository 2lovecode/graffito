# Golang代码涂鸦
## 一. 用法
  1. 编译构建
        - 克隆代码 ```git clone git@github.com:2lovecode/graffito.git```
        - 进入项目目录 ```cd graffito```
        - 构建 ```go build```
  2. 使用各个功能
        - 查看命令帮助 ```./graffito -h```
        - 执行命令，例如 ```./graffito tools count abc``` 计算字符串abc的字符数
  ![image](images/build.png)
    
## 二. 目录说明

| 目录  | 说明  | 备注                       |
| :---:  | :------:  | :------------------------  |
| tools | 百宝工具箱     ||
| algorithm |   数据结构&算法  ||
| experiment |  一些实验性想法的实现     ||
| pattern |  设计模式实现     ||
| practice | 编程练习代码     ||

## 三. 模块说明

### 百宝工具箱
#### [计算字符串字符个数](tools/string_op)
 - 命令：```./graffito tools count abc```
#### [简易redis命令行客户端](tools/redis)
 - 命令：```./graffito tools redis 127.0.0.1 6379```
 - ![image](images/redis.png)
 - 目前支持的命令: ``` set get hset hget hgetall```
#### [自定义命令帮助工具](tools/helper)
 - 命令：``` ./graffito tools helper  ~/Development/Tools/tips.json  ```
 - tips.json文件 
 ```
     {
    "name": "系统使用手册",
    "description": "",
    "comment": "",
    "tips": [
        {
        "group": "别名定义",
        "description": "在~/.bash_profile中定义了别名",
        "items": [
            {
            "name": "进入指定目录",
            "usage": [
                {
                "name": "2c",
                "description": "cd到~/Development/Code目录",
                "comment": ""
                },
                {
                "name": "2d",
                "description": "cd到~/Downloads目录",
                "comment": ""
                },
                {
                "name": "2l",
                "description": "cd到/usr/local目录",
                "comment": ""
                },
                {
                "name": "2t",
                "description": "cd到~/Development/ThirdPartyCode目录",
                "comment": ""
                },
                {
                "name": "2w",
                "description": "cd到~/Development/Work目录",
                "comment": ""
                },
                {
                "name": "2log",
                "description": "cd到~/Development/logs目录",
                "comment": ""
                }
            ],
            "comment": ""
            },
            {
            "name": "php",
            "usage": [
                {
                "name": "php-start",
                "description": "启动php开发docker环境",
                "comment": ""
                },
                {
                "name": "php-stop",
                "description": "停止php开发docker环境",
                "comment": ""
                }
            ],
            "comment": ""
            },
            {
            "name": "sftp",
            "usage": [
                {
                "name": "fdev",
                "description": "sftp登陆开发机",
                "comment": ""
                }
            ],
            "comment": ""
            },
            {
            "name": "cheat",
            "usage": [
                {
                "name": "cheat",
                "description": "运行cht程序查询一些开发使用文档",
                "comment": ""
                }
            ],
            "comment": ""
            },
            {
            "name": "proxy",
            "usage": [
                {
                "name": "proxy",
                "description": "当前shell临时打开代理",
                "comment": ""
                },
                {
                "name": "unproxy",
                "description": "当前shell临时关闭代理",
                "comment": ""
                }
            ],
            "comment": ""
            }
        ]
        },
        {
        "group": "工具",
        "description": "安装的工具的使用",
        "items": [
            {
            "name": "gitui",
            "usage": [
                {
                "name": "gitui",
                "description": "在已经git init的目录中运行 gitui",
                "comment": "https://github.com/extrawurst/gitui"
                }
            ],
            "comment": ""
            },
            {
            "name": "flameshot",
            "usage": [
                {
                "name": "flameshot",
                "description": "",
                "comment": "https://github.com/flameshot-org/flameshot"
                }
            ],
            "comment": ""
            },
            {
            "name": "duf",
            "usage": [
                {
                "name": "duf",
                "description": "磁盘空间查看",
                "comment": "https://github.com/muesli/duf"
                }
            ],
            "comment": ""
            }
        ]
        }
    ]
    }

 ```
 - ![image](images/helper.png)


### 数据结构&算法
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [数组实现](algorithm/array.go) |  |  |
| [斐波那契数列](algorithm/fibonacci.go) |  | |
| [哈希表](algorithm/hash.go) |  |  |
| [堆实现](algorithm/heap.go) |  |  |
| [快排实现](algorithm/quicksort.go) |  |  |

### 实验代码
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [缓存 穿透/击穿/雪崩 解决方案](experiment/cache)| ```./graffito exp cache``` |  |
| [事件绑定与触发](experiment/event)|  |  |
| [有依赖关系的并行模型](experiment/depends)|  |  详情见：https://github.com/2lovecode/depends|
| [一种批量多样数据的处理模式](experiment/mode0)|  |  |

### LeetCode
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [LeetCode](algorithm/leetcode) |  |  |

### 设计模式
| 功能  | 说明  | 备注                       |
| :---:  | :------: | :------------------------:  |
| [建造者](pattern/builder.go) |  |  |
| [工厂模式](pattern/factory.go) |  |  |
| [对象池模式](pattern/obj_pool.go) |  |  |
| [观察者模式](pattern/observer.go) |  |  |
| [单例](pattern/singleton.go) |  |  |
| [策略模式](pattern/strategy.go) |  |  |

### 练习代码
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