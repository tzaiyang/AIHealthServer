package controller

import (
	"log"
	"net/http"

	"aihealth/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Description Add Medical detail information
// @Accept json
// @Produce json
// @Param medical body model.Medical true "Add Medical"
// @Success 200
// @Router /medicals [post]
func AddMedical(c *gin.Context) {
	// path c.Param
	// @Accept formData c.PostForm
	log.Println(c)
	var data model.Medical
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	_, err := data.InsertOne()
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, data)
}

// @Description View all medicals information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} medical "medical name"
// @Router /medicals [get]
func GetMedicals(c *gin.Context) {
	var data model.Medical
	if results, err := data.Find(bson.M{}); err == nil {
		c.JSON(http.StatusOK, results)
	} else {
		log.Println("getMedical error: " + err.Error())
		c.JSON(500, err.Error())
	}
}
