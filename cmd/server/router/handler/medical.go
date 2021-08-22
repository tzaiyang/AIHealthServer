package handler

import (
	"aihealth/internal/pkg/db"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Medical struct {
	PZWH              string  `json:"pzwh" example:"H20046681"`
	GYZZ              string  `json:"gyzz" example:"H20046681"`
	ZCZH              string  `json:"zczh"`
	Name              string  `json:"name" example:"聚乙烯醇滴眼液 (瑞珠)"`
	Dosage_form       string  `json:"dosage_form" example:"眼用制剂(滴眼剂)"`
	Packing_unit      string  `json:"packing_unit" example:"盒"`
	Specification     string  `json:"specification" example:"0.8ml*25支"`
	Single_dose       string  `json:"single_dose" example:"0.2ml"`
	Frequency         string  `json:"frequency" example:"每日4次"`
	Usage             string  `json:"usage" example:"点双眼"`
	Major_Functions   string  `json:"major_functions" example:"异物感  眼疲劳  眼部干涩"`
	Price             float32 `json:"price" example:"58.16"`
	Manufacturer      string  `json:"manufacturer" example:"湖北远大天天明制药有限公司"`
	Bar_code          string  `json:"bar_code" example:"6935899801619"`
	Prescription_Only bool    `json:"prescription_only" example:"true"`
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
	if err := c.ShouldBindJSON(&medical); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMedical)
	_, err := collection.InsertOne(context.Background(), medical)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, medical)
}

// @Description View all medicals information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} medical "medical name"
// @Router /medicals [get]
func GetMedicals(c *gin.Context) {
	var results []bson.M

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMedical)
	if cursor, err := collection.Find(context.Background(), bson.M{}); err == nil {
		err = cursor.All(context.TODO(), &results)
		c.JSON(http.StatusOK, results)
	} else {
		log.Println("getMedical error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func DelMedicals(c *gin.Context) {
	medical_id := c.Params.ByName("medical_id")
	log.Println(medical_id)
	objID, _ := primitive.ObjectIDFromHex(medical_id)

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMedical)

	if results, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID}); err == nil {
		c.JSON(http.StatusOK, gin.H{"deletedCount": results.DeletedCount})
	} else {
		log.Println("DeleteMedical error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func UpdateMedicalByID(c *gin.Context) {
	medical_id := c.Params.ByName("medical_id")
	log.Println(medical_id)
	objID, _ := primitive.ObjectIDFromHex(medical_id)

	var updateMedical Medical

	if err := c.ShouldBindJSON(&updateMedical); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		// log.Error(err)
	}

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMedical)
	if updateResults, err := collection.UpdateByID(context.Background(), objID, updateMedical); err == nil {
		c.JSON(http.StatusOK, gin.H{"ModifiedCount": updateResults.ModifiedCount})
	} else {
		log.Println("ModifiedMedical error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
