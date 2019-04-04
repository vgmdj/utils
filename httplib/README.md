# httplib
http调用包，用来发送http请求，并解析返回结果

# 所需参数：
- URL
  目标地址

- body
  发送内容，如果使用PostJSON，则会将body结构体进行json.Marshal转化；同理，如果使用PostXML，则将body结构体进行xml.Marshal转化

- resp
  对返回内容进行解析，如果不进行指定，则默认根据返回的 Content-Type 进行格式化，
  如调用微信支付，返回xml格式内容，返回的Content-Type 为 application/xml，则默认进行xml解析，对resp进行赋值。
  需要注意的是，resp需传入指针类型。
  如果调用方不注重这些细节，返回Content-Type 与实际内容格式不符，则可以手动指定解析方式，指定方式是在headers里进行

- headers
  对发送的Header进行设置，可以在这里设置Cookie，方式为：

  ```
    httplib.NewClient().PostJSON(url, body, &resp,
     	map[string]string{
                "Cookie":"key=value",
            })

  ```
  也可以在这里对返回内容解析方式进行指定，方式为：

  ```
    httplib.NewClient().PostJSON(url, body, &resp, map[string]string{
                   httplib.ResponseResultContentType:httplib.ContentTypeAppJson,
                })


  ```

# 举例
接口信息

```

目标url： http://testapi.vgmdj.cn
输入参数： json格式
        {
            "timestamp":1234567890,
            "input":"test"
        }

输出参数： json 格式
        {
            "timestamp":1234567890,
            "output":"test"
        }


```

调用代码

```

type Input struct{
    Timestamp int64 `json:"timestamp"`
    Input string    `json:"input"`
}

type Output struct{
    Timestamp int64 `json:"timestamp"`
    Output string   `json:"output"`
}


func Test(){
    input := Input{
        Timestamp: time.Now().Unix(),
        Input:"test",
    }

    var output Output
    httplib.NewClient().PostJSON("http://apitest.vgmdj.cn",input,&output,nil)

    //或者是
    output2 := make(map[string]interface{})
    httplib.NewClient().PostJSON("http://apitest.vgmdj.cn",input,&output2,map[string]string{
            httplib.ResponseResultContentType:httplib.ContentTypeAppJson,
    })

    fmt.Println(output)   //1234567890 test
    fmt.Println(output2)  //1234567890 test

}


```
