package models

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
