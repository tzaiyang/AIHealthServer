package model

type Prescription struct {
	User_ID               string `json:"user_id"`
	MTR_ID                string `json:"mtr_id"`
	Prescription_ID       string `json:"prescription_id"`
	Name                  string
	Detail                string
	Prescription_Medicals []Medical
}
