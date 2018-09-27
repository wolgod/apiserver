### 依赖的软件

| 软件 | 版本|  
|:---------|:-------:|
| node.js     |  8.4.0 (及以上) |
| golang  |  1.9 (及以上) |
| mysql  |  5.6.35 (及以上) |
| redis  |  4.0.1 (及以上) |

### 前端地址 
https://github.com/shaoheshan/vueadmin

### 后端依赖的库

本项目使用dep来管理依赖的包，请先安装dep, 执行以下命令即完成安装

```
go get -u github.com/golang/dep/cmd/dep
```

然后，在 **apiserver** 项目目录下运行以下命令来安装依赖(需要翻墙,有些包国内无法直接下载)

```
dep ensure
```
### ⚙️ 配置
  #### config.yaml
   修改config.yaml中数据库和redis地址
  #### db.sql
   创建mysql数据库,初始化脚本
 
## 🚕 运行
### 运行前端项目
注意一点运行前端项目的时候需要修改调用的后端的端口

### 运行后端项目
进入`apiserver`目录，然后运行

```
go run main.go
```
### 访问
 https://127.0.0.1:{'前端项目对应的端口'}  默认密码 admin/1

