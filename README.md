# RPC框架
## 功能
- 客户端调用

类别|说明|TODO
---|---|---
调用哪个服务	|获取服务地址、缓存、请求负载均衡，	|长连接
如何请求	|设置重试次数、超时时间、失败策略	|熔断

- 传输

类别|说明
---|---
思路	|网络传输TCP/HTTP/WebSocket，粘包（使用定长字段表示变长数据的长度）
实现	|消息头5字节长，保存校验数、协议版本、消息类型（请求/响应）、压缩类型（Gzip）、序列化类型（Protobuf/Json）

- 服务端

类别|说明|TODO
---|---|---
注册服务	|服务名与对应处理类保存在内存中	
处理请求	|监听端口-持久accept请求，读数据-解码，调用对应方法，编码-返回。每个步骤都可以插入AOP逻辑	|限流，认证授权
“负责任”-优雅启停	|先启动服务再注册到服务中心；先注销再关闭服务，处理中的请求数为0时才真正关闭服务	

## 实现
### 类图
![RPC框架设计类图](https://github.com/dingqing/rpc/blob/main/RPC.jpg?raw=true)

## 运行
- 启动服务端
  
  先启动 [服务注册中心](https://github.com/dingqing/registry)
```
 cd demo
 go run server.go -c config.yaml
```
- 客户端测试
```
cd demo
go run client_proxy.go

```