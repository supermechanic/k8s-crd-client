package models

import "model-service/client"

func init() {
	db := client.GetDB()
	db.AutoMigrate(&Model{})
	db.AutoMigrate(&Predict{})
}
