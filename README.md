### GO语言实现的一些功能

#### 一. 用法
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
    
#### 二.模块列表
| 功能  | 所属模块  | 用法                       |简介       |
| :---  | :------  | :------------------------  | :-------------- |
| hello |  tools   | -m tools -t hello          |                  |
| ccnum |    -     | -m tools -t ccnum abc 你好 | 统计字符串字符数 |