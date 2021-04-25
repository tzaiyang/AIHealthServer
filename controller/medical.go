package controller

import (
	"log"
	"net/http"

	"AIHealthServer/model"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"gopkg.in/mgo.v2/bson"
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
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	err := data.Insert()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
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
	var data []interface{}
	collection := model.MongoSession.DB("AIHealth").C("medicals")
	if err := collection.Find(bson.M{}).All(&data); err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		log.Print(err)
	}
}
