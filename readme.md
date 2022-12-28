<h1>简介</h1>

-     golang iris mvc 框架的一套基础项目架构，只实现了基础的后台账户管理登录，权限管理
-     后台基于layuiadmin嵌套，如不需要可自行修改
-     如需修改go mod 名称 全局替换 'iris-init'即可

- [后台 layuiadmin 模板下载地址](https://github.com/zqjzqj/layuiAdmin.git)

<h1>安装流程</h1>

### 正式数据请编译后替代[go run main.go]运行

---
- 1.将config.env复制更名为config.yml 然后填写对应的配置
---
- 2.运行迁移命令 go run main.go -migrate=run 创建数据库 初始账号 admin 123456
---
- 3.启动程序go run main.go [如未配置域名访问 请使用localhost:port访问本地]
---

### 支持参数
 * -config 指定配置文件位置 默认./config 【注：参数值不要带.yml】
 * -migrate (run, rollback) 运行或回滚迁移 
 * -mRollbackId 当 -migrate=rollback时 需要指定回滚版本号
    

<h3>开发命令</h3>

-     用于创建 services, repositories, repoInterface, migrate

      -alias为model缩写，空则默认为model首字母小写, -appRoot为项目代码root目录，空则默认为当前目录
      -migrate为迁移models, 使用'-'分割多个models 会生成对应的迁移文件
    
-     go run .\cmd\generateServTpl.go -model=Xxx -alias=xxx -appRoot=xxxx
-     go run .\cmd\generateServTpl.go -migrate=Xxx1-Xxx2-Xxx3
      