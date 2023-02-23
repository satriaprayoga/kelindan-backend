package models

import "time"

type KUser struct {
	UserID      int       `json:"user_id" gorm:"PRIMARY_KEY"`
	UserName    string    `json:"user_name" gorm:"type:varchar(20)"`
	Name        string    `json:"name" gorm:"type:varchar(60);not null"`
	Telp        string    `json:"telp" gorm:"type:varchar(20)"`
	Email       string    `json:"email" gorm:"type:varchar(60)"`
	IsActive    bool      `json:"is_active" gorm:"type:boolean"`
	JoinDate    time.Time `json:"join_date" gorm:"type:timestamp(0)"`
	BirthOfDate time.Time `json:"birth_of_date" gorm:"type:timestamp(0)"`
	Password    string    `json:"password" gorm:"type:varchar(150)"`
	FileID      int       `json:"file_id" gorm:"type:integer"`
	UserType    string    `json:"user_type" gorm:"type:varchar(10)"`
	Model
}
