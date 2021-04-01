Database design

# MongoDB
`users`
```json
{
  "user_id": "18707122919",
  "phone": "18707122919",
  "name": "Ryan Tang",
  "birth": "1995-06-26",
  "rh": false,
  "abo": "B",
  "gender": "男",
  "height": 165,
  "weight": 65,
  "occupation": "程序员",
  "updated": "2021-03-20"
}
```

`eyes.refraction_report`
```json
{
  "user_id":"18707122919",
  "family_history":null,
  "type":"OU",      
  "dominant_eye": "OD",
  "OD":-3.00,
  "OS":-2.50,
  "PD":60,
  "astigmatism":false,
  "position":"正位眼",
  "optometry_status":"Normal pupil",
  "updated_date":"2021-01-24",
  "updated_hospital":"中山大学眼科中心附属眼科医院验光配镜中心",
  "Updated_depart":"验光配镜中心",
  "updated_doctor":"王子轩"
}
```

Medical Treatment Record
`eyes.mtr`
```json
{
  "user_id":"18707122919",
  "mtr_id":"mtr0001",
  "主诉":"双眼疲劳，干涩，一月有余",
  "现病史":"双眼干涩，畏光，异物感，一月有余",
  "既往史":"否认其他眼部疾病史，否认外伤史，否认手术史",
  "检查":"普通视力检查OU; 裂隙灯检查OU；小瞳验光(检影，云雾试验，试镜，主导眼检查)；眼底检查（直接眼底镜法）OU；非接触眼压计法（综合门诊）OU；",
  "诊断":"双眼视疲劳；双眼屈光不正；双眼干眼症",
  "updated_date":"2021-01-24",
  "updated_hospital":"中山大学眼科中心",
  "updated_depart":"综合门诊",
  "updated_doctor":"孔炳华"
}
{
  "user_id":"18707122919",
  "mtr_id":"mtr0002",
  "主诉":"双眼前黑影飘动数年",
  "现病史":"近视",
  "既往史":"否认其他眼部疾病史，否认外伤史，否认手术史",
  "体征":"右眼视力0.1，左眼0.08，矫正1.0，双眼角膜透明，前房清，瞳孔圆，晶体透明，玻璃体絮状混浊，眼底视盘界限清，网膜平伏，未见出血及渗出，黄斑区中心凹反光可见。",
  "检查":"扫描激光眼底检查（SLO）；眼压检查（非接触眼压记法）；试镜；验光；普通视力检查；裂隙灯检查",
  "诊断":"双眼近视，视疲劳，玻璃体混浊",
  "Updated_date":"2020-05-31",
  "Updated_hospital":"东莞康华医院",
  "Updated_depart":"眼科门诊",
  "Updated_doctor":"金涛"
}
```

`eyes.prescription`
```json
{
"user_id": "18707122919",
"mtr_id":"mtr0002",
"prescription_id":"pres0001",
"name":"",
"detail":"",
"prescription_medicals":[
  {
    "name":"聚乙烯醇滴眼液",
    "single_dose":"0.2ml",
    "frequency":"每日4次",
    "天数":14,
    "usage":"点双眼"
  },
  {
    "name":"明目地黄胶囊",
    "single_dose":"3",
    "frequency":"每日3次",
    "天数":14,
    "usage":"口服"
  },
  {
    "name":"脉血康胶囊",
    "single_dose":"3",
    "frequency":"每日3次",
    "天数":14,
    "usage":"口服"
  }
]
}
{
  "user_id": "18707122919",
  "mtr_id":"mtr0001",
  "prescription_id":"pres0002",
  "name": "Prescription for spectacles",
  "prescription":{
    "Distance":{
      "OD": {
        "Sph":-2.75,
        "CVA":1.0
      },
      "OS": {
        "Sph":-2.25,
        "CVA":1.0
      }
    },
    "Near":{

    },
    "PD": 60,
    "RPD":null,
    "LPD":null,
    "PH":null,
    "RPH":null,
    "LPH":null
  },
}
```

`medical`
```json
{
  "gyzz": "H20046681",
  "name": "聚乙烯醇滴眼液 (瑞珠)",
  "dosage_form": "眼用制剂(滴眼剂)",
  "packing_unit": "盒",
  "specification": "0.8ml*25支",
  "single_dose": "0.2ml",
  "frequency": "每日4次",
  "usage": "点双眼",
  "major_functions": "异物感  眼疲劳  眼部干涩",
  "price": 58,
  "major_functions":"湖北远大天天明制药有限公司"
}
{
  "gyzz":"Z20050408",
  "name":"明目地黄胶囊(众生)",
  "dosage_form":"胶囊剂",
  "specification":"0.5g*30粒",
  "packing_unit":"盒",
  "single_dose":"3",
  "frequency":"每日3次",
  "usage":"口服",
  "manufacturer":"广东众生药业股份有限公司",
  "bar_code":"6902170000443",
  "major_functions":"迎风流泪  肝肾阴虚  视物模糊  目涩畏光",
  "price":24.15
},
{
  "gyzz":"Z10970056",
  "name":"脉血康胶囊(多泰)",
  "dosage_form":"胶囊剂",
  "packing_unit":"盒",
  "single_dose":"3",
  "frequency":"每日3次",
  "usage":"口服",
  "major_functions":"症瘕痞块  血瘀经闭  跌打损伤",
  "price":26.16
}
{
  "name":"右旋糖酐70滴眼液 (润齐)",
  "specification":"0.4ml:0.4mg*10支",
  "dosage_form":"滴眼剂",
  "packing_unit":"盒",
  "gyzz":"H20133327",
  "prescription_only":true,
  "manufacturer":"齐鲁制药有限公司",
  "bar_code":"6915798003802",
  "major_functions":"眼部刺激感  眼部灼热  眼部干燥  眼部不适"
}
{
  "gyzz":"S20020016",
  "name":"重组人表皮生长因子滴眼液(酵母) (易贝)",
  "packing_unit":"盒",
  "specification":"4ml:4万IU",
  "major_functions":"化学烧伤  点状角膜病变  角膜上皮缺损  干眼症  角膜损伤",
  "prescription_only":true,
  "price":41.21
}
{
  "gyzz":"J20150129",
  "zczh":"H20150390",
  "name":"维生素B12滴眼液(散克巴)",
  "major_functions":"眼疲劳",
  "price":22.37
}
{
  "gyzz":"H35020244",
  "name":"维生素AD软胶囊(星鲨)",
  "specification":"3000U:300U*100粒",
  "dosage_form":"胶囊剂(软胶囊)",
  "packing_unit":"瓶",
  "manufacturer":"国药控股星鲨制药(厦门)有限公司",
  "bar_code":"6912283508125",
  "major_functions":"维生素D缺乏  维生素A缺乏  小儿手足抽搐症  佝偻病  夜盲症",
  "price":12.02
}
{
  "gyzz":"H20051825",
  "name":"加替沙星滴眼液(祝宁)",
  "specification":"5ml:15mg",
  "packing_unit":"盒",
  "manufacturer":"安徽省双科药业有限公司",
  "bar_code":"6935392400104",
  "major_functions":"假单孢菌属  布兰氏卡他菌  棒状杆菌属  细菌菌属  葡萄球菌属  链球菌属  嗜血"
}
{
  "zczh":"H20150150",
  "name":"玻璃酸钠滴眼液(海露)",
  "usage用量":"每日三次，一次一滴",
  "specification":"10ml*1支",
  "packing_unit":"盒",
  "manufacturer":"德国URSAPHARM Arzneimittel GmbH(德国)",
  "bar_code":"4020799602205",
  "major_functions":"斯-约二氏综合症  角结膜上皮损伤  眼干燥症  干燥综合症",
  "price":69.52
}
{
  "zczh":"H20140811",
  "name":"妥布霉素滴眼液(托百士)",
  "specification":"5ml*15mg",
  "packing_unit":"支",
  "major_functions":"适用于外眼及附属器敏感菌株感染的局部抗感染治疗",
  "price":19.51
}
{
  "zczh":"H20040234",
  "name":"盐酸左氧氟沙星眼用凝胶(杰奇)",
  "specification":"5g:0.015g",
  "dosage_form":"眼用制剂(眼用凝胶剂)",
  "packing_unit":"支",
  "prescription_only":true,
  "manufacturer":"湖北远大天天明制药有限公司",
  "bar_code":"6935899800117",
  "major_functions":"外眼感染  细菌性结膜炎  角膜溃疡  角膜炎  泪囊炎  术后感染",
  "price":19.19
}
```

# MySQL
`book` 
```json
{
    "isbn":"9781617291784",
    "author":"Brian Ketelsen & Erik St. Martin & William Kennedy",
    "title":"Go in Action",
    "publisher":"Manning Publications",
    "year":"2015",
    "tags":"Web & Programming",
    "language":"english",
    "rating":4
}

{
  "title":"眼、耳鼻喉科学笔记"
}

{
  "title":"眼.耳鼻咽喉与口腔科常见疾病诊治"
}

{
  "title":"眼耳鼻咽喉科护理学"
}
```


