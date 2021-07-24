# AIHealth
[![Build Status](https://www.travis-ci.com/tzaiyang/AIHealthServer.svg?branch=master)](https://www.travis-ci.com/tzaiyang/AIHealthServer)

AIHealth is an app recording your health data, including base information(height,weight,ABO,etc.), sports data, medical record, etc. The app will analyse these data to help you improve your body health.

# Build and Deployment
1. Docker deployment

    ```bash
    docker pull tzaiyang/aihealth:0.1
    docker run --rm --name aihealth -d -p 8080:8080 tzaiyang/aihealth:0.1
    ```

2. Local build and deployment

    Requirements
    ```
    Gin  
    MySQL
    MongoDB

    sqlite3==3.27.2
    # wget https://sqlite.org/2019/sqlite-tools-linux-x86-3270200.zip
    elastics
    ```

    ```bash
    git clone https://github.com/tzaiyang/AIHealthServer.git
    go mod init AIHealthServer
    go get github.com/swaggo/swag/cmd/swag
    swag init
    go get -u 
    go run main.go
    ```



# API Document
https://localhost:8080/api/index.html

Swagger API annotaion format

```go
// @Accept json (request data format)
// @Produce json (response data format)
// @Param  user_id path string true "User Id" (put user_id into URI)
// @Param account body model.User true "Add account" (`account` value will be sent to server with json string format)
// @Param  user_id formData string true "User Id" (`user_id` will be sent with `-d "user_id=123&name=Ryan" format to server
```

cURL test 
```bash
curl -X POST "http://localhost:8080/accounts" -H "accept: application/json" -H "Content-Type: application/json" -d '{"user_id":"13","name":"1","gender":"2"}'
curl -X POST "http://localhost:8080/accounts" -d 'name=1&gender=2&user_id=14'
```

## Refernces
https://docs.mongodb.com/drivers/go/  
https://github.com/swaggo/swag#declarative-comments-format  
https://gin-gonic.com/docs/examples/bind-query-or-post/
