package models

import "time"

type Department struct {
	ID              uint      `gorm:"primaryKey"`
	DepartmentName  string    `gorm:"type:varchar(255);not null" json:"department_name"`
	MaxClockInTime  string    `gorm:"column:max_clock_in_time;type:time(3);not null" json:"max_clock_in_time"`
	MaxClockOutTime string    `gorm:"column:max_clock_out_time;type:time(3);not null" json:"max_clock_out_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
