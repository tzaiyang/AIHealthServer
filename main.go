package main

import (
	"aihealth/common"
	"aihealth/controllers"
	"aihealth/databases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @title AIHealth API
// @version 1.0
// @description This is a AIHealth server.
// @host localhost:10086
// @BasePath /
func main() {
	// gin.SetMode(gin.ReleaseMode)
	log.Println("AIHealth start...")
	log.Println("Connecting MongoDB")
	common.LoadConfig()

	err := databases.Database.Init()
	if err != nil {
		log.Fatal("Connect MongoDB Failed")
		log.Println(err)
	} else {
		log.Println("Connected MongoDB Success")
	}

	defer databases.Database.Close()
	r := SetupRouter()
	r.Run(":10086")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// Get user value
	r.POST("/accounts", controllers.AddAccount)
	r.DELETE("/accounts/:user_id", controllers.DeleteAccountByUserID)
	r.PUT("/accounts/:user_id", controllers.UpdateAccountByUserID)
	r.GET("/accounts", controllers.GetAccount)
	r.GET("/accounts/name/:name", controllers.GetAccountByName)
	r.GET("/accounts/id/:user_id", controllers.GetAccountByID)

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))
	authorized.POST("", controllers.AuthorizeAccount)

	r.POST("/medicals", controllers.AddMedical)
	r.GET("/medicals", controllers.GetMedicals)
	r.DELETE("/medicals/id/:medical_id", controllers.DelMedicals)
	r.PUT("medicals/id/:medical_id", controllers.UpdateMedicalByID)

	r.POST("/mtrs", controllers.AddMTR)
	r.GET("/mtrs", controllers.GetMTR)
	r.DELETE("/mtrs/id/:mtr_id", controllers.DelMtr)
	r.PUT("mtrs/id/:mtr_id", controllers.UpdateMtrByID)

	return r
}
