# 1. 行政区划
__GB2260 中华人民共和国行政区划代码<br>
因行政区划会不断调整，所以默认使用最新代码（2018.02）<br>
同时也支持指定时间，如身份证信息便返回当年的地址__

### 1.1 功能
- 6位行政区划号转地址
- 6位行政区划号转省份
- 6位行政区划号转市区


### 1.2 使用方式
- 先导入数据包，可以使用我已经准备好的

> import _ "github.com/vgmdj/gb2260/gbdata"

- 之后

```
    gb2260 := area.NewArea(area.GB2260)
    bj := gb2260.Get("110101")
    fmt.Println(bj.Province)    //北京市
    fmt.Println(bj.County)      //东城区
    fmt.Println(bj.FullName())  //北京市东城区
    
```

- 详细例子可以参考测试文件
https://github.com/vgmdj/utils/blob/master/area/area_test.go

# 2. 邮编
- 待更新

# 3. 区号
- 待更新

# 4. 车牌
- 待更新
