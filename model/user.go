package model

import "gopkg.in/mgo.v2"

type User struct {
	User_ID string `json:"user_id" example:"18717992222" format:"string"`
	Phone   string `json:"phone" example:"18717992222"`
	Name    string `json:"name" example:"Ryan"`
	Birth   string `json:"birth" example:"2009-08-23"`
	Gender  string `json:"gender" example:"男"`
	ABO     string `json:"abo" example:"B"`
	Rh      bool   `json:"rh" example:"true"`
	Height  int    `json:"height" example:"170"`
	Weight  int    `json:"weight" example:"65"`
}

var MongoClient *mgo.Session
