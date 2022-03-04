### websocket-机器人

### 实现功能
```.env
机器人根据指令进行发送信息
```

### 技术方案+使用说明文档
```.env
技术方案：https://docs.qq.com/doc/DV0J4Ukp6b0xGZmxp
说明文档：https://docs.qq.com/doc/DV1NSUERGSkNwU3Za
```

### websocket底层技术
```.env
    1.websocket:github.com/gorilla/websocket
    2.goland
    3.sdk:1.16.8
```

### 目录树
```.env
D:.
│  go.mod
│  go.sum
│  main.go  主函数-唯一
│  README.md
│
├─bootstrap 初始化包
│  │  init.go 初始化函数
│  │
│  ├─config 配置文件
│  │      config.yaml
│  │      config_init.go
│  │
│  ├─exception  异常捕捉
│  │      exception_init.go
│  │
│  ├─response   返回结构+常规返回信息
│  │      response.go
│  │      response_init.go
│  │
│  ├─router 路由配置
│  │      router_init.go
│  │
│  ├─websocket  websocket连接
│  │  │  websockett_connect.go  底层github.com/gorilla/websocket实现
│  │  │
│  │  └─botGo   引用botGo的官方项目++
│  │          bot.yaml  
│  │          ip.go
│  │          processor.go
│  │          websocket_init.go
│  │
│  └─xlogger    日志配置
│          logCore.go
│          log_init.go
│
├─http  请求接口
│      get.go
│
├─log   日志打印所在地
│      error.log
│      info.log
│
├─pojo  结构体
│      websocket_url.go
│
├─service   业务层
│      qq_service.go    QQ小程序业务
│      robot_service.go 机器人业务
│
└─web   接口层
        qq_controller.go    QQ小程序接口
        robot_controller.go 机器人接口
```
