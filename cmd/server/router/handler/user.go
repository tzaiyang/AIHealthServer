package handler

import (
	"aihealth/internal/pkg/db"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// User manages
type User struct {
	User_ID    string `json:"user_id" example:"18717992222" format:"string"`
	Phone      string `json:"phone" example:"18717992222"`
	Name       string `json:"name" example:"Ryan"`
	Birth      string `json:"birth" example:"2009-08-23"`
	Sex        string `json:"sex" example:"男"`
	ABO        string `json:"abo" example:"B"`
	Rh         bool   `json:"rh" example:"true"`
	Height     int    `json:"height" example:"170"`
	Weight     int    `json:"weight" example:"65"`
	Occupation string `json:"occupation" example:"程序员"`
	Updated    string `json:"updated" example:"2021-03-30 15:59"`
}

// var db = make(map[string]string)

// func AuthorizeAccount(c *gin.Context) {
// 	log.Println(c)
// 	user := c.MustGet(gin.AuthUserKey).(string)

// 	// Parse JSON
// 	var json struct {
// 		Value string `json:"value" binding:"required"`
// 	}

// 	if c.Bind(&json) == nil {
// 		db[user] = json.Value
// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 	}
// }

// @Description View users detail information
// @Accept  json
// @Produce  json
// @Param name path string false "Account Name"
// @Success 200
// @Header 200 {string} user "User name"
// @Router /accounts [get]
func GetAccount(c *gin.Context) {
	user_name := c.Params.ByName("name")
	user_id := c.Params.ByName("user_id")
	log.Println(user_id)
	log.Println(user_name)
	var results []bson.M

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColUsers)
	if cursor, err := collection.Find(context.Background(), bson.M{"name": user_name, "user_id": user_id}); err == nil {
		if err := cursor.All(context.TODO(), &results); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, results)
		}
	} else {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

// @Description Add user detail information
// @Accept json
// @Produce json
// @Param account body models.User true "Add account"
// @Success 200
// @Router /accounts [post]
func AddAccount(c *gin.Context) {
	// path c.Param
	// @Accept formData c.PostForm
	log.Println(c)
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColUsers)
	if InsertOneResult, err := collection.InsertOne(context.Background(), user); err != nil {
		c.JSON(http.StatusOK, InsertOneResult.InsertedID)
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

// @Description Delete user by user_id
// @Accept json
// @Produce json
// @Param  user_id path string true "User Id"
// @Success 200
// @Router /accounts/{user_id} [delete]
func DeleteAccountByUserID(c *gin.Context) {
	user_id := c.Params.ByName("user_id")
	log.Println(user_id)

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColUsers)

	if results, err := collection.DeleteOne(context.Background(), bson.M{"user_id": user_id}); err == nil {
		c.JSON(http.StatusOK, gin.H{"deletedCount": results.DeletedCount})
	} else {
		log.Println("DeleteMedical error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

// @Description Update user by user_id
// @Accept json
// @Produce json
// @Param  user_id path string true "User Id"
// @Param account body models.User true "Add account"
// @Success 200
// @Router /accounts/{user_id} [patch]
func UpdateAccountByUserID(c *gin.Context) {
	user_id := c.Param("user_id")
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}

	collection := db.Database.MgDbClient.Database(db.Database.DbName).Collection(ColUsers)
	if updateResults, err := collection.UpdateOne(context.Background(), bson.M{"user_id": user_id}, user); err == nil {
		c.JSON(http.StatusOK, gin.H{"ModifiedCount": updateResults.ModifiedCount})
	} else {
		log.Println("ModifiedMedical error: " + err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
