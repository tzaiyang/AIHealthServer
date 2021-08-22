package handler

/*
 * @File: handler.mtr.go
 * @Description: Defines medical treatment record information will be returned to the clients
 * @Author: ryan (zaiyangt@163.com)
 */

import (
	"aihealth/internal/pkg/db"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MedicalTreatmentRecord manages
type MedicalTreatmentRecord struct {
	User_ID                    string `json:"user_id" example:"18717992222" format:"string"`
	MTR_ID                     string `json:"mtr_id" example:"mtr0001"`
	Chief_Complaint            string `json:"chief_complaint" example:"双眼疲劳，干涩，一月有余"`
	History_of_present_illness string `json:"history_of_present_illness" example:"双眼干涩，畏光，异物感，一月有余"`
	Past_History               string `json:"past_history" example:"否认其他眼部疾病史，否认外伤史，否认手术史"`
	Physical_Exam              string `json:"physical_exam" example:"普通视力检查OU; 裂隙灯检查OU；小瞳验光(检影，云雾试验，试镜，主导眼检查)；眼底检查（直接眼底镜法）OU；非接触眼压计法（综合门诊）OU；"`
	Assessment                 string `json:"assessment" example:"双眼视疲劳；双眼屈光不正；双眼干眼症"`
	Treatment_Plan             string `json:"treatment_plan"`
	Prescription_ID            string `json:"prescription_id"`
	Date                       string `json:"date" example:"2021-01-24 15:59"`
	Hospital                   string `json:"hospital" example:"中山大学眼科中心"`
	Depart                     string `json:"depart" example:"综合门诊"`
	Attending_Physician        string `json:"attending_physician" example:"孔炳华"`
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

	if err := c.ShouldBindJSON(&mtr); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		// log.Fatal(err)
		return
	}

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMtrs)
	if InsertOneResult, err := collection.InsertOne(context.Background(), mtr); err != nil {
		c.JSON(http.StatusOK, InsertOneResult.InsertedID)
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

// @Description View all MTR information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} mtrs "mtrs name"
// @Router /mtrs [get]
func GetMTR(c *gin.Context) {
	var results []bson.M

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMtrs)
	if cursor, err := collection.Find(context.Background(), bson.M{}); err == nil {
		err = cursor.All(context.TODO(), &results)
		c.JSON(http.StatusOK, results)
	} else {
		log.Print("getMTR error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func DelMtr(c *gin.Context) {
	mtr_id := c.Params.ByName("mtr_id")
	log.Println(mtr_id)
	objID, _ := primitive.ObjectIDFromHex(mtr_id)

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMtrs)

	if results, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID}); err == nil {
		c.JSON(http.StatusOK, gin.H{"deletedCount": results.DeletedCount})
	} else {
		log.Println("Delete error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func UpdateMtrByID(c *gin.Context) {
	mtr_id := c.Params.ByName("mtr_id")
	log.Println(mtr_id)
	objID, _ := primitive.ObjectIDFromHex(mtr_id)

	var mtr MedicalTreatmentRecord
	if err := c.ShouldBindJSON(&mtr); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		// log.Fatal(err)
	}

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColMtrs)
	if updateResults, err := collection.UpdateByID(context.Background(), objID, mtr); err == nil {
		c.JSON(http.StatusOK, gin.H{"ModifiedCount": updateResults.ModifiedCount})
	} else {
		log.Println("Modified error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
