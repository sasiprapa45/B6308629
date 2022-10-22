package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}
type Doctor struct {
	gorm.Model
	DoctorName       string
	DoctorEmail      string `gorm:"uniqueIndex"`
	DoctorPassword   string
	DoctorDepartment string
	// 1 doctor ดูแล patient ได้หลายคน
	Patient []Patient `gorm:"foreignKey:DoctorID"`
}

type PatientType struct {
	gorm.Model
	Type string
	// 1 patient type เป็น patient ได้หลายคน
	Patient []Patient `gorm:"foreignKey:PatientTypeID"`
}

type Symptoms struct {
	gorm.Model
	SymptomsName string
	// 1 symptoms ให้ patient เป็นได้หลายคน
	Patient []Patient `gorm:"foreignKey:SymptomsID"`
}

type Patient struct {
	gorm.Model
	PatientsName string
	DateAdmit    time.Time
	Age         int
	// DoctorID ทำหน้าที่เป็น FK
	DoctorID *uint
	Doctor   Doctor

	// PatientTypeID ทำหน้าที่เป็น FK
	PatientTypeID *uint
	PatientType   PatientType

	// SymptomsID ทำหน้าที่เป็น FK
	SymptomsID *uint
	Symptoms   Symptoms
}
