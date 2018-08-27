## logger
模仿beego/logs模块

## 配置
#### 设置异步
> SetAsync()

#### 设置log等级
> SetLevel(l int)

- LevelEmergency 
- LevelAlert
- LevelCritical
- LevelError
- LevelWarning
- LevelNotice
- LevelInformational
- LevelDebug
    

#### 设置logger
> SetLogger(adaptername string, config string)

目前支持的 adapter

- Console

- File

- ElasticSearch

- Ali Log Service

## 计划
新添配置，使出现不同等级的log时，可自定义处理方式