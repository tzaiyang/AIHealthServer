package controller

import (
	"AIHealth_Server/dbs"
	"AIHealth_Server/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
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
// @Param name path string true "Account Name"
// @Success 200
// @Header 200 {string} user "User name"
// @Header 200 {string} status "User status"
// @Router /accounts/{name} [get]
func ShowAccount(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}

// @Description Add user detail information
// @Param name formData string true "Account Name"
// @Param birth formData string false "Birth"
// @Param gender formData string true "Gender"
// @Param abo formData string false "ABO"
// @Success 200
// @Router /accounts/ [post]
func AddAccount(c *gin.Context) {
	log.Println(c)
	mongo, err := mgo.Dial("127.0.0.1")

	defer mongo.Close()
	if err != nil {
		fmt.Print("connect error")
	}

	data := model.User{
		Name:   c.PostForm("name"),
		Birth:  c.PostForm("birth"),
		Gender: c.PostForm("gender"),
		ABO:    c.PostForm("abo"),
	}

	// log.Println(data.Name)
	dbs.InsertMongo(mongo, "AIHealth", "users", &data)
	c.JSON(200, gin.H{
		"status": "posted",
		"gender": data.Gender,
		"birth":  data.Birth,
		"abo":    data.ABO,
		"name":   data.Name,
	})
}
