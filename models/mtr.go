/*
 * @File: models.mtr.go
 * @Description: Defines medical treatment record information will be returned to the clients
 * @Author: ryan (zaiyangt@163.com)
 */
package models

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
