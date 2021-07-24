package router

import (
	"aihealth/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// Get user value
	r.POST("/accounts", controller.AddAccount)
	r.DELETE("/accounts/:user_id", controller.DeleteAccountByUserID)
	r.PATCH("/accounts/:user_id", controller.UpdateAccountByUserID)
	r.GET("/accounts", controller.GetAccount)
	r.GET("/accounts/name/:name", controller.GetAccountByName)
	r.GET("/accounts/id/:user_id", controller.GetAccountByID)
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	// 	"foo":  "bar",
	// 	"manu": "123",
	// }))
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	authorized.POST("", controller.AuthorizeAccount)
	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/

	r.POST("/medicals", controller.AddMedical)
	r.GET("/medicals", controller.GetMedicals)

	r.POST("/mtrs", controller.AddMTR)
	r.GET("/mtrs", controller.GetMTR)

	return r
}
