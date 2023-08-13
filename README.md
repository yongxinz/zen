# zen

使用 go-zero 实现的前后端分离工作流引擎系统

本项目为前后端分离项目，后端使用微服务框架 go-zero 开发，前端使用 Vue + Element UI 开发。

**前端项目地址：** https://github.com/yongxinz/zen-fe

## 开发环境

后端开发环境基于 [gonivinck](https://github.com/nivin-studio/gonivinck) 搭建，一个基于 docker 的 go-zero 运行环境。

1、拉取 gonivinck 项目，按需修改 .env 配置

2、启动基础环境服务

```shell
docker-compose up -d
```

3、启动项目 rpc 服务

```shell
# 进入 golang 容器
docker exec -it gonivinck_golang_1 bash
# 进入项目目录
cd service/sys/rpc
# 启动服务
go run sys.go -f etc/sys.yaml
```

4、启动项目 http 服务

```shell
# 进入 golang 容器
docker exec -it gonivinck_golang_1 bash
# 进入项目目录
cd service/sys/api
# 启动服务
go run sys.go -f etc/sys.yaml
```