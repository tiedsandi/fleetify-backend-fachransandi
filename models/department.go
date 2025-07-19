package models

import "time"

type Department struct {
	ID              uint   `gorm:"primaryKey"`
	DepartmentName  string `gorm:"type:varchar(255);not null"`
	MaxClockInTime  string `gorm:"type:time;not null"`
	MaxClockOutTime string `gorm:"type:time;not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
