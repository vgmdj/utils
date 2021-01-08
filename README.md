# web tools

not good enough to use directly


## [chars](https://github.com/vgmdj/utils/tree/master/chars)

字符串处理包

- 身份证号解析
- 常用正则判断
- 数组判断重复与去重
- 格式转码
- 数字放大缩小和类型转换
- 随机数生成
- 时间与字符转换

## [httplib](https://github.com/vgmdj/utils/tree/master/httplib)

http 请求与接收，并根据返回Content-Type 将返回结果绑定到相应的结构体上
支持的Content-type

- application/json
- application/xml
- text/xml
- text/plain (返回字符串)

## [area](https://github.com/vgmdj/utils/tree/master/area)

区域号码转地址

- 行政区划号码,历年来GB2260号码转换
- 邮编
- 车牌号

## logger

推荐使用zap/logger

## [config](https://github.com/vgmdj/utils/tree/master/config)

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

## sync

fork from golang.org/x/sync/errgroup
