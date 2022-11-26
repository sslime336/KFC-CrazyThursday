# KFC-CrazyThursday

找到~~同城~~ KFC

## 食用方法

*需要高德和百度的 AppKey*

1. 新建 `.env`，参考 [example.env](example.env) 填写 AppKey
2. `go get github.com/sslime336/kfct`

```golang
kfc.Postion() // 简单封装
kfc.PosRaw()  // 高德接口返回值
```

:P
