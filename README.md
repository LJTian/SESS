# SESS

*********
## 自述
SESS(Super Easy Service Scaffolding)超级简易的服务脚手架,本人练手的的东西。 
想法是写一个简单的符合微服务的框架，还得简单易用，不是很重。
最好就是在PC电脑上也可以使用。
突然发现，这个写的不像是脚手架，更像一个服务了，先写出来，然后再提炼吧。
*********
## 原则
- 1、使用开源中间件
- 2、容器化部署所有内容
- 3、每一个组件都有md文件介绍
- 4、代码注释完善
*********
## 目录介绍
遵循[Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)
*********
## 中间件列表
- 1、消息队列[rabbitmq](https://www.rabbitmq.com/)
- 2、缓存数据库[redis](https://redis.io/)
- 3、数据存储库[mariadb](https://mariadb.org/)
- 4、远程调用[grpc](https://www.grpc.io/)
- 5、服务注册中心[consul](https://www.consul.io/)
- 6、配置管理[nacos](https://nacos.io/zh-cn/)
- 7、日志库[zap](https://pkg.go.dev/go.uber.org/zap)
- 8、web框架[gin](https://gin-gonic.com/)

***各个中间件的[测试代码](https://github.com/LJTian/go-test/tree/master/demo-test)***

*********