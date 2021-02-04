package models

import "time"

//CTI //
type CTI struct {
	ID             int32     `gorm:"primaryKey;column:CTI_CLASS"`
	ParentID       int32     `gorm:"column:PARENT_ID"`
	ClassCode      string    `gorm:"column:CLASS_CODE"`
	FullName       string    `gorm:"column:FULL_NAME"`
	Remark         string    `gorm:"column:REMARK"`
	CreateDate     time.Time `gorm:"column:CREATE_DATE"`
	CreateUser     string    `gorm:"column:CREATE_USER"`
	LastModifyDate time.Time `gorm:"column:LAST_MODIFY_DATE"`
	LastModifyUser string    `gorm:"column:LAST_MODIFY_USER"`
	DelFlag        bool      `gorm:"column:DEL_FLAG"`
}

//TableName ///
func (CTI) TableName() string {
	return "tb_ai_model_ctiinfo"
}
