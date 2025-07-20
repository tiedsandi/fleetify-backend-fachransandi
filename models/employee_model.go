package models

import "time"

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	EmployeeID   string     `gorm:"type:varchar(50);not null;unique" json:"employee_id"`
	DepartmentID uint       `gorm:"not null" json:"department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID" json:"department"`
	Name         string     `gorm:"type:varchar(255);not null" json:"name"`
	Address      string     `gorm:"type:text;not null" json:"address"`
	CreatedAt    time.Time  `gorm:"type:dateTime(0)" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:dateTime(0)" json:"updated_at"`
}
