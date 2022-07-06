## 使用Gin框架构建一个web项目的基础

[![OSCS Status](https://www.oscs1024.com/platform/badge/MarkGaoGit/skeleton.svg?size=small)](https://www.oscs1024.com/project/MarkGaoGit/skeleton?ref=badge_small)


### 完成项
- 参数绑定
- 中间件的使用
- 数据库的连接【基础】
- 基本的增删改查
- 验证器的封装
- redis的使用
- 一个简单的用户信息基础

### 包引用
- go get -u github.com/gin-gonic/gin
- go get -u github.com/spf13/viper
- go get -u gorm.io/gorm
- go get -u gorm.io/driver/mysql 
- go get github.com/go-playground/validator/v10 [验证器](https://godoc.org/github.com/go-playground/validator)
- go get github.com/go-redis/redis/v8 使用了v8版本


### 待改进项
1. 添加日志
2. 错误统一处理捕捉
3. redis的初始化优化


### 编译
#### Linux 
- CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o LinuxWebServer cmd/web/main.go

#### mac
- go build -o skeleton cmd/web/main.go