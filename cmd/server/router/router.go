package router

import (
	"aihealth/cmd/server/router/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// Get user value
	r.POST("/accounts", handler.AddAccount)
	r.DELETE("/accounts/:user_id", handler.DeleteAccountByUserID)
	r.PUT("/accounts/:user_id", handler.UpdateAccountByUserID)
	r.GET("/accounts", handler.GetAccount)
	// r.GET("/accounts", handler.GetAccount)
	// r.GET("/accounts/id/:user_id", handler.GetAccount)

	// authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))
	// authorized.POST("", handler.AuthorizeAccount)

	r.POST("/medicals", handler.AddMedical)
	r.GET("/medicals", handler.GetMedicals)
	r.DELETE("/medicals/id/:medical_id", handler.DelMedicals)
	r.PUT("medicals/id/:medical_id", handler.UpdateMedicalByID)

	r.POST("/mtrs", handler.AddMTR)
	r.GET("/mtrs", handler.GetMTR)
	r.DELETE("/mtrs/id/:mtr_id", handler.DelMtr)
	r.PUT("mtrs/id/:mtr_id", handler.UpdateMtrByID)

	return r
}
