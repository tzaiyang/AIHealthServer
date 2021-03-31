package controller

import (
	"AIHealth_Server/dbs"
	"AIHealth_Server/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
	"gopkg.in/mgo.v2/bson"
)

var db = make(map[string]string)

func AuthorizeAccount(c *gin.Context) {
	log.Println(c)
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

// @Description View users detail information
// @Accept  json
// @Produce  json
// @Param name path string false "Account Name"
// @Success 200
// @Header 200 {string} user "User name"
// @Router /accounts/name/{name} [get]
func GetAccountByName(c *gin.Context) {
	collection := model.MongoClient.DB("AIHealth").C("users")

	user_name := c.Params.ByName("name")
	log.Println(user_name)
	var data []interface{}
	if err := collection.Find(bson.M{"name": user_name}).All(&data); err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		log.Print(err)
	}
}

// @Description View users detail information with User ID
// @Accept  json
// @Produce  json
// @Param user_id path string true "Account User ID"
// @Success 200
// @Header 200 {string} user "User name"
// @Router /accounts/id/{user_id} [get]
func GetAccountByID(c *gin.Context) {
	collection := model.MongoClient.DB("AIHealth").C("users")

	user_name := c.Params.ByName("user_id")
	log.Println(user_name)
	var data []interface{}
	if err := collection.Find(bson.M{"user_id": user_name}).All(&data); err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		log.Print(err)
	}
}

// @Description View all users information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} user "User name"
// @Router /accounts [get]
func GetAccount(c *gin.Context) {
	collection := model.MongoClient.DB("AIHealth").C("users")

	var data []interface{}
	if err := collection.Find(bson.M{}).All(&data); err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		log.Print(err)
	}
}

// @Description Add user detail information
// @Accept json
// @Produce json
// @Param account body model.User true "Add account"
// @Success 200
// @Router /accounts [post]
func AddAccount(c *gin.Context) {
	// path c.Param
	// @Produce formData c.PostForm
	//
	log.Println(c)

	var data model.User

	if err := c.ShouldBindJSON(&data); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	// curl -X POST "http://localhost:8080/accounts" -H "accept: application/json" -H "Content-Type: application/json" -d '{"user_id":"13","name":"1","gender":"2"}'
	// curl -X POST "http://localhost:8080/accounts" -d 'name=1&gender=2&user_id=14'
	// data := model.User{
	// 	User_ID: c.PostForm("user_id"),
	// 	Name:    c.PostForm("name"),
	// 	Birth:   c.PostForm("birth"),
	// 	Gender:  c.PostForm("gender"),
	// 	ABO:     c.PostForm("abo"),
	// 	Phone:   c.PostForm("phone"),
	// 	// Rh:     bool(c.PostForm("rh")),
	// 	// Height: c.PostForm("height"),
	// 	// Weight: c.PostForm("weight"),
	// }
	// data := model.User{
	// 	Name:   c.Param("name"),
	// 	Birth:  c.Param("birth"),
	// 	Gender: c.Param("gender"),
	// 	ABO:    c.Param("abo"),
	// 	Phone:  c.Param("phone"),
	// 	// Rh:     bool(c.PostForm("rh")),
	// 	// Height: c.PostForm("height"),
	// 	// Weight: c.PostForm("weight"),
	// }

	err := dbs.InsertMongo(model.MongoClient, "AIHealth", "users", data)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{
		"user_id": data.User_ID,
		"gender":  data.Gender,
		"birth":   data.Birth,
		"abo":     data.ABO,
		"name":    data.Name,
		"phone":   data.Phone,
		"rh":      data.Rh,
		"height":  data.Height,
		"weight":  data.Weight,
	})
}

// @Description Delete user by user_id
// @Accept json
// @Produce json
// @Param  user_id path string true "User Id"
// @Success 200
// @Router /accounts/{user_id} [delete]
func DeleteAccountByUserID(c *gin.Context) {
	collection := model.MongoClient.DB("AIHealth").C("users")
	user_id := c.Params.ByName("user_id")
	log.Println(user_id)
	if err := collection.Remove(bson.M{"user_id": user_id}); err == nil {
		c.JSON(200, gin.H{"status": "deleted succes"})
	} else {
		// log.Fatal(err)
		c.JSON(400, gin.H{"err": err.Error()})
	}
}

// @Description Update user by user_id
// @Accept json
// @Produce json
// @Param  user_id path string true "User Id"
// @Param account body model.User true "Add account"
// @Success 200
// @Router /accounts/{user_id} [patch]
func UpdateAccountByUserID(c *gin.Context) {
	user_id := c.Param("user_id")

	var updateAccount model.User
	if err := c.ShouldBindJSON(&updateAccount); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	collection := model.MongoClient.DB("AIHealth").C("users")
	err := collection.Update(bson.M{"user_id": user_id}, updateAccount)
	if err != nil {
		// log.Fatal(err)
		c.JSON(400, gin.H{"err": err.Error()})
	}

	c.JSON(http.StatusOK, updateAccount)
}
