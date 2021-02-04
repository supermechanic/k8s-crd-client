package models

import (
	"model-service/client"

	"gorm.io/gorm"
)

//Predict running predictors
type Predict struct {
	gorm.Model
	ModelID    int32
	Version    string
	Status     string
	ServerType string
	Flow       int
	Creator    string
	StopTime   string
}

//Add create a predict record in database
func (p *Predict) Add() error {
	db := client.GetDB()
	if err := db.Create(p).Error; err != nil {
		return err
	}
	return nil
}
