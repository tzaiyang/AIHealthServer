# Introduction
https://localhost:8080/api/index.html

# Build and Deployment
1. local test
```bash
swag init
go run main.go
```

2. docker deployment
```bash
swag init
GOOS=linux GOARCH=amd64 go build -o aihealth main.go
docker build -t aihealth:0.1 ./
docker run -d -p 8080:8080 aihealth:0.1
```

# Development
1. Requirement
```bash
go mod init
go mod tidy
go get -u

go get gopkg.in/mgo.v2
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

# Swagger API 
```go
// @Accept json 表示请求数据格式为json
// @Produce json 表示响应的数据格式为json
// @Param  user_id path string true "User Id" 表示user_id将被放入URI中传送给服务器
// @Param account body model.User true "Add account" 表示account的值将用json字符串格式传送给服务器。
// @Param  user_id formData string true "User Id" 表示user_id将以-d "user_id=1234&name=Ryan"的形式传给服务器
```

cURL test 
```bash
curl -X POST "http://localhost:8080/accounts" -H "accept: application/json" -H "Content-Type: application/json" -d '{"user_id":"13","name":"1","gender":"2"}'
curl -X POST "http://localhost:8080/accounts" -d 'name=1&gender=2&user_id=14'
```

# MongoDB Command
```json
db.users.find().pretty()
```

`eread` is a C-S architecture online ebook web application. 

# Requirements
```
Gin  
MySQL
MongoDB

sqlite3==3.27.2
# wget https://sqlite.org/2019/sqlite-tools-linux-x86-3270200.zip
elastics
```

## Refernces
https://docs.mongodb.com/drivers/go/
https://github.com/swaggo/swag#declarative-comments-format
https://gin-gonic.com/docs/examples/bind-query-or-post/


- docker login registry.cn-hangzhou.aliyuncs.com
  - docker pull registry.cn-hangzhou.aliyuncs.com/mango9102/aihealth:0.1
  - docker run -d -p 8080:8080 registry.cn-hangzhou.aliyuncs.com/mango9102/aihealth:0.1
