package main

import (
	"aihealth/cmd/server/router"
	"aihealth/internal/pkg/cfg"
	"aihealth/internal/pkg/db"

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
	cfg.LoadConfig()

	err := db.Database.Init()
	if err != nil {
		log.Fatal("Connect MongoDB Failed")
		log.Println(err)
	} else {
		log.Println("Connected MongoDB Success")
	}

	defer db.Database.Close()
	r := router.SetupRouter()
	r.Run(":" + cfg.Config.Service.Port)
}
