package models

import (
	"model-service/client"
	"time"
)

//Model ///
type Model struct {
	ModelID    int32 `gorm:"primaryKey"`
	ModelName  string
	ModelPath  string
	ModelType  string
	ModelVer   string
	ModelDesc  string
	UserID     int32
	PublicFlag string
	CreateDate string
	ModelFrom  string
	SystemTime time.Time
	FileKey    string
	FileName   string
	FilePath   string
	FileSize   string
	Cti        CTI   `gorm:"foreignKey:CTIID"`
	CTIID      int32 `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;column:CTI_CLASS"`
}

//TableName ///
func (Model) TableName() string {
	return "kb_model"
}

//GetModelByID ///
func (m *Model) GetModelByID() error {
	db := client.GetDB()
	if err := db.First(m).Error; err != nil {
		return err
	}
	return nil
}
