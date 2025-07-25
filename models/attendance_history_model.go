package models

import "time"

type AttendanceHistory struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	EmployeeID     string    `gorm:"type:varchar(50);not null;index"`
	AttendanceID   string    `gorm:"type:varchar(100);not null;index"`
	DateAttendance time.Time `gorm:"type:dateTime(0)" json:"date_attendance"`
	AttendanceType int       `gorm:"type:tinyint(1)" json:"attendance_type"`
	Description    string    `gorm:"type:text" json:"description"`
	CreatedAt      time.Time `gorm:"type:dateTime(0)" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:dateTime(0)" json:"updated_at"`
}

func (AttendanceHistory) TableName() string {
	return "attendance_history"
}
