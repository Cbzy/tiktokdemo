## 项目结构
```shell
|--app
    |-api
    |-model
    |-router
    |-service
|--boot
    |-global
        |-db.go
        |-global.go
    |-viper
       |setting.go
|--config
    |-config.yaml
|--utils
    |--regexp  
    |--jwt
|--go.mod
|--main.go 
```


| 文件夹       | 说明       | 描述                 |
| ------------ |----------|--------------------|
| `api`        | api层     | api层               |
| `config`     | 配置包      | config.yaml对应的配置结构体 |
| `global`     | 全局对象     | 全局对象               |
| `middleware` | 中间件层     | 用于存放 `gin` 中间件代码  |
| `model`      | 模型层      | 模型对应数据表            |  |
| `router`     | 路由层      | 路由层                |
| `service`    | service层 | 存放业务逻辑问题           |
| `utils`     | 放常用的包    | 目前有两个              |


