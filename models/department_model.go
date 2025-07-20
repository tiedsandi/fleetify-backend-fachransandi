package models

import "time"

type Department struct {
	ID              uint      `gorm:"primaryKey"`
	DepartmentName  string    `gorm:"type:varchar(255);not null" json:"department_name"`
	MaxClockInTime  string    `gorm:"column:max_clock_in_time;type:time(0);not null" json:"max_clock_in_time"`
	MaxClockOutTime string    `gorm:"column:max_clock_out_time;type:time(0);not null" json:"max_clock_out_time"`
	CreatedAt       time.Time `gorm:"type:dateTime(0)" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:dateTime(0)" json:"updated_at"`
}
