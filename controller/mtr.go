package controller

import (
	"log"
	"net/http"

	"aihealth/model"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
)

// @Description Add MTR detail information
// @Accept json
// @Produce json
// @Param mtr body model.MedicalTreatmentRecord true "Add MTR"
// @Success 200
// @Router /mtrs [post]
func AddMTR(c *gin.Context) {
	// path c.Param
	// @Accept formData c.PostForm
	log.Println(c)
	var data model.MedicalTreatmentRecord
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	data.InsertOne()
	c.JSON(200, data)
}

// @Description View all MTR information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} mtrs "mtrs name"
// @Router /mtrs [get]
func GetMTR(c *gin.Context) {
	var data model.MedicalTreatmentRecord

	if results, err := data.Find(bson.D{}); err == nil {
		c.JSON(http.StatusOK, results)
	} else {
		c.JSON(500, err.Error())
		log.Print("getMTR error: " + err.Error())
	}
}
