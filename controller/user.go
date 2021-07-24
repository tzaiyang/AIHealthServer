package controller

import (
	"aihealth/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	user_name := c.Params.ByName("name")
	log.Println(user_name)

	var user model.User

	if data, err := user.Find(bson.M{"name": user_name}); err == nil {
		c.JSON(200, data)
	} else {
		log.Print(err)
		c.JSON(500, "server error")
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
	user_name := c.Params.ByName("user_id")
	log.Println(user_name)
	var user model.User

	if data, err := user.Find(bson.M{"user_id": user_name}); err == nil {
		c.JSON(200, data)
	} else {
		log.Print(err)
		c.JSON(500, "server error")
	}
}

// @Description View all users information
// @Accept  json
// @Produce  json
// @Success 200
// @Header 200 {string} user "User name"
// @Router /accounts [get]
func GetAccount(c *gin.Context) {
	var user model.User
	if data, err := user.Find(bson.M{}); err == nil {
		c.JSON(200, data)
	} else {
		log.Print(err)
		c.JSON(500, "server error")
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
	// @Accept formData c.PostForm
	log.Println(c)
	var data model.User
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	err := data.InsertOne()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}
	c.JSON(200, data)
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

	count := model.UserDeleteByID(user_id)

	c.JSON(200, gin.H{"status": "deleted succes", "count": count})
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
		c.JSON(http.StatusBadRequest, err.Error())
		log.Fatal(err)
	}
	updateAccount.UpdateByID(user_id)
	c.JSON(http.StatusOK, updateAccount)
}
