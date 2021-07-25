package controller

import (
	"log"
	"net/http"

	"aihealth/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func DelMedicals(c *gin.Context) {
	medical_id := c.Params.ByName("medical_id")
	log.Println(medical_id)
	var data model.Medical
	if results, err := data.Delete(bson.M{"_id": medical_id}); err == nil {
		c.JSON(http.StatusOK, gin.H{"deletedCount": results})
	} else {
		log.Println("DeleteMedical error: " + err.Error())
		c.JSON(500, err.Error())
	}
}

func UpdateMedicalByID(c *gin.Context) {

	medical_id := c.Params.ByName("medical_id")
	log.Println(medical_id)
	objID, _ := primitive.ObjectIDFromHex(medical_id)
	var updateMedical model.Medical
	if err := c.ShouldBindJSON(&updateMedical); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	if updateResults, err := updateMedical.UpdateByID(objID); err == nil {
		c.JSON(http.StatusOK, gin.H{"ModifiedCount": updateResults.ModifiedCount})
	} else {
		log.Println("ModifiedMedical error: " + err.Error())
		c.JSON(500, err.Error())
	}
}
