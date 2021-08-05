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
type MedicalTreatmentRecord struct {
	data models.MedicalTreatmentRecord
	DAO  daos.MedicalTreatmentRecord
}

// @Description Add MTR detail information
// @Accept json
// @Produce json
// @Param mtr body models.MedicalTreatmentRecord true "Add MTR"
// @Success 200
// @Router /mtrs [post]
func AddMTR(c *gin.Context) {
	// path c.Param
	// @Accept formData c.PostForm
	log.Println(c)
	var mtr MedicalTreatmentRecord
	if err := c.ShouldBindJSON(&mtr.data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	mtr.DAO.InsertOne(mtr.data)
	c.JSON(200, mtr)
}

// @Description View all MTR information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} mtrs "mtrs name"
// @Router /mtrs [get]
func GetMTR(c *gin.Context) {
	var mtr MedicalTreatmentRecord

	if results, err := mtr.DAO.Find(bson.D{}); err == nil {
		c.JSON(http.StatusOK, results)
	} else {
		c.JSON(500, err.Error())
		log.Print("getMTR error: " + err.Error())
	}
}

func DelMtr(c *gin.Context) {
	mtr_id := c.Params.ByName("mtr_id")
	log.Println(mtr_id)
	objID, _ := primitive.ObjectIDFromHex(mtr_id)

	var mtr MedicalTreatmentRecord
	if count, err := mtr.DAO.Delete(bson.M{"_id": objID}); err == nil {
		c.JSON(http.StatusOK, gin.H{"deletedCount": count})
	} else {
		log.Println("DeleteMtr error: " + err.Error())
		c.JSON(500, err.Error())
	}
}

func UpdateMtrByID(c *gin.Context) {

	mtr_id := c.Params.ByName("mtr_id")
	log.Println(mtr_id)
	objID, _ := primitive.ObjectIDFromHex(mtr_id)

	var mtr MedicalTreatmentRecord
	if err := c.ShouldBindJSON(&mtr.data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	if updateResults, err := mtr.DAO.UpdateByID(objID, mtr.data); err == nil {
		c.JSON(http.StatusOK, gin.H{"ModifiedCount": updateResults.ModifiedCount})
	} else {
		log.Println("ModifiedMtr error: " + err.Error())
		c.JSON(500, err.Error())
	}
}
