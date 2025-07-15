package main

import "log"

type GormLogger struct{}

func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":

		log.Println("module:gorm","type:sql","rows:",v[5],"src_ref:",v[1],"values:",v[4],"debug:",v[3])
		//log.WithFields(
		//	map[string]interface{}{
		//		"module":  "gorm",
		//		"type":    "sql",
		//		"rows":    v[5],
		//		"src_ref": v[1],
		//		"values":  v[4],
		//	},
		//).Debug(v[3])
	case "log":
		//log.WithFields(map[string]interface{}{"module": "gorm", "type": "log"}).Info(v[2])
		log.Println(v[2])
	}
}