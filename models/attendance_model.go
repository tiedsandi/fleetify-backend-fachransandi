package models

import "time"

type Attendance struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	EmployeeID   string    `gorm:"type:varchar(50);not null;index"`
	AttendanceID string    `gorm:"type:varchar(100);not null;unique" json:"attendance_id"`
	ClockIn      time.Time `gorm:"type:dateTime(0)" json:"clock_in"`
	ClockOut     time.Time `gorm:"type:dateTime(0)" json:"clock_out"`
	CreatedAt    time.Time `gorm:"type:dateTime(0)" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:dateTime(0)" json:"updated_at"`
}

func (Attendance) TableName() string {
	return "attendance"
}
