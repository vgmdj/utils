# web tools

## chars
字符串处理包
- 身份证号解析
- 常用正则判断
- 数组判断重复与去重
- 格式转码
- 数字处理
- 随机数生成
- 时间与字符转换

## httplib
http 请求与接收，并根据返回Content-Type 将返回结果绑定到相应的结构体上<br>
支持的Content-type
- application/json
- application/xml
- text/xml
- text/plain (返回字符串)

## area
区域号码转地址
- 行政区划号码
- 邮编
- 车牌号

## logger
log输出相关

## config
配置文件读取与使用

## encrypt
加密相关
- md5及加盐
- 简单兑换码方案

## files
文件操作相关
- 将结构体信息解析到文本
- 将文本内信息解析到结构体

## db
对sql语句处理

## mq
消息队列相关

## sms
short message system，一些常用消息（短信）平台的消息推送接口<br>
支持平台
- 容联云通讯（支持）
- 阿里云（待开发）
- 腾讯云（待开发）
- 微信公众号（待开发）

## lifeApi
生活相关api
- 查油价，调用的是eolinker的接口
- 地图相关api，调用的是腾讯地图