npm install apidoc -g

```bash
go mod init
go mod tidy
go get -u

go get gopkg.in/mgo.v2
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

swag init
go run main.go

# localhost:8080/api/index.html
```

```mongo
db.users.find().pretty()
```

## Refernces
https://docs.mongodb.com/drivers/go/
https://github.com/swaggo/swag#declarative-comments-format
https://gin-gonic.com/docs/examples/bind-query-or-post/
