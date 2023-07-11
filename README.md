# RPC框架
> 手动实现RPC框架

目录 |二级目录
---|---
[理解](#理解) |[生活化类比](#生活化类比)
[实现](#实现) |[功能分解](#功能分解)，[类图](#类图)
[运行展示](#运行展示) |[启动服务注册中心](#启动服务注册中心)，[启动服务端](#启动服务端)，[客户端测试](#客户端测试)

***

## 理解
### 生活化类比
> 其他的 **服务模型**：操作系统处理请求（select/epoll将连接转发给进程）、Go语言中调度模型（将协程任务G转发给空闲的M处理）。

类别|生活例子|RPC框架|Web框架
---|---|---|---
客户端 |用户：可以保留商家电话/聊天窗口，下次不用在平台中重新查找（服务缓存），遇到人满为患的店家则换一家（负载均衡。此处考虑的是流量拥塞，而不是人多反映服务质量）|微服务中的调用方 |用户/浏览器/APP
通信协议 |电话/文字/语言 |考虑性能 - TCP，权衡性能与便捷 - http2 |http：客户端调用简单
转接方 |前台/客服：尽职在线、将请求转发给具体服务方，需要将交接完手头接收的请求才能下班 |Server中的监听模块 |Web服务器（Nginx）的路由模块、框架中的路由模块
服务平台 |菜市场/电商平台 |服务注册中心 |服务容器：存取「请求/响应/数据库实例」等实例
服务方 |商家 |服务端中定义的方法/功能 |MVC模型

## 实现
### 功能分解
> 要实现的是什么
- 客户端调用

类别|说明|TODO
---|---|---
调用哪个服务 |获取服务地址、缓存、请求负载均衡 |watch注册中心的服务、及时更新本地服务缓存
如何请求 |设置重试次数、超时时间、失败策略 |熔断

- 传输

类别|说明
---|---
思路 |网络传输TCP/HTTP/WebSocket，粘包（使用定长字段表示变长数据的长度）
实现 |消息头5字节长，保存校验数、协议版本、消息类型（请求/响应）、压缩类型（Gzip）、序列化类型（Protobuf/Json）

- 服务端

类别|说明|TODO
---|---|---
注册服务 |服务名与对应处理类保存在内存中	
处理请求 |监听端口-持久accept请求，读数据-解码，调用对应方法，编码-返回。每个步骤都可以插入AOP逻辑	|限流，认证授权
“负责任”-优雅启停 |先启动服务再注册到服务中心；先注销再关闭服务，处理中的请求数为0时才真正关闭服务	

### 类图
![RPC框架设计类图](https://i.imgtg.com/2023/05/27/OoNSag.jpg)

***

## 运行展示
### 启动服务注册中心
[服务注册中心](https://github.com/dingqing/registry)
### 启动服务端
```
cd demo
go run server.go -c config.yaml
```
### 客户端测试
```
go run client_proxy.go
```
