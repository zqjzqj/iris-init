<h1>安装流程</h1>

### 正式数据请编译后替代[go run main.go]运行

---
- 1.将config.env复制更名为config.yml 然后填写对应的配置
---
- 2.运行迁移命令 go run main.go -migrate=run 创建数据库 初始账号 admin 123456
---
- 3.启动程序go run main.go
---

### 支持参数
 * -config 指定配置文件位置 默认./config 【注：参数值不要带.yml】
 * -migrate (run, rollback) 运行或回滚迁移 
 * -mRollbackId 当 -migrate=rollback时 需要指定回滚版本号
    