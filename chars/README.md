# chars
常用字符处理包

## 格式转换
- ToString()
  将任意类型转换为string类型

- ToInt()
  将任意类型转换为int类型

## 放大缩小
  将任意格式数字，放大n倍，并返回指定格式类型

  ```
    缩小时，会默认按照倍数进行小数点位数的保留，如10倍，就是一位小数，1000倍就是三位小数
    也可以强制指定保留位数，就是在ToString()时，加入倍数参数

    chars.NewConversion("2.342",1000).ZoomOut().ToInt()   //2342
    chars.NewConversion(2342,100).ZoomIn().ToString()     //23.42
    chars.NewConversion(2342,100).ZoomIn().ToString(10)   //23.4

  ```

## ID生成
可以使用uuid生成方式，同时指定是否保留"-"，也可以使用bson objectid的生成方式

```
    chars.NewUUID(false)  //生成不带"-" 的uuid
    chars.NewBsonID()     //生成BsonID

```

## time 时间相关计算
- 计算下一个零点，有时，我们需要在零点来更新一些东西，这时，就需要计算出当前离下一个零点还有多长时间

  ```
    chars.Time24Sub(time.Now())
  ```

- 营业时间计算
  有时，我们需要在一些接口的营业时间进行调用，非营业时间则进行等待。

  ```
    //设置休息时间为 22：50 - 0：50， 由于中间隔了一个零点，所以crossMidNight 为true
    rt := RestTime(22, 50, 0, 50, true)
    t := time.Date(2018, time.July, 1, 23, 55, 0, 0, time.UTC)

    //设置额外等待时间为30秒
    rt.SetExtWaitTime(time.Second*30)
    rt.IsWorkingTime(t)                 // false
    rt.IsRestTime(t)                    // true
    rt.WaitTime(t)                      // 所需等待时间为 56分钟30秒

  ```

## 身份证解析
身份证号上有着大量的个人信息，可以通过 chars.ParseIDCard("xxxxxxxxxx") 来进行解析
- 可以判断是否是符合格式身份证，即通过最后一位校验位，来检验身份证号是否正确
- 可以判断出男，女
- 可以获取出生日期
- 可以判断出生地，出生地默认使用出生时那年发布的GB2260行政区划对应表，范围在1986-2018,因每年调整次数可能不止一次，这里查询使用的是最后一次发布的信息对应表，所以可能存在些许误差

## 编码
GBK转UTF8等