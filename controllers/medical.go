package controllers

import (
	"aihealth/daos"
	"aihealth/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MedicalTreatmentRecord manages
type Medical struct {
	data models.Medical
	DAO  daos.Medical
}

// @Description Add Medical detail information
// @Accept json
// @Produce json
// @Param medical body models.Medical true "Add Medical"
// @Success 200
// @Router /medicals [post]
func AddMedical(c *gin.Context) {
	// path c.Param
	// @Accept formData c.PostForm
	log.Println(c)
	var medical Medical
	if err := c.ShouldBindJSON(&medical.data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	_, err := medical.DAO.InsertOne(medical.data)
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, medical.data)
}

// @Description View all medicals information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} medical "medical name"
// @Router /medicals [get]
func GetMedicals(c *gin.Context) {
	var medical Medical
	if results, err := medical.DAO.Find(bson.M{}); err == nil {
		c.JSON(http.StatusOK, results)
	} else {
		log.Println("getMedical error: " + err.Error())
		c.JSON(500, err.Error())
	}
}

func DelMedicals(c *gin.Context) {
	medical_id := c.Params.ByName("medical_id")
	log.Println(medical_id)
	var medical Medical
	if results, err := medical.DAO.Delete(bson.M{"_id": medical_id}); err == nil {
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

	var updateMedical Medical
	if err := c.ShouldBindJSON(&updateMedical.data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	if updateResults, err := updateMedical.DAO.UpdateByID(objID, updateMedical.data); err == nil {
		c.JSON(http.StatusOK, gin.H{"ModifiedCount": updateResults.ModifiedCount})
	} else {
		log.Println("ModifiedMedical error: " + err.Error())
		c.JSON(500, err.Error())
	}
}
