package models

import (
	"github.com/aldodarel/go_students_crud_mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Kolom-kolom yang ada di tabel student
type Student struct {
	gorm.Model
	NIM string `gorm:"json:nim"`
	Name string `json:"name"`
	IPK string `json:"ipk"`
	Jurusan string `json:"jurusan"`
	Angkatan string `json:"angkatan"`
	StatusAktif string `json:"status_aktif"`
	Username string `json:"username"`
	EmailAkademik string `json:"email_akademik"`
	WaliMahasiswa string `json:"wali_mahasiswa"`
	JalurMasuk string `json:"jalur_USM"`

}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}
func (b *Student) CreateStudent() *Student {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}
func GetStudentbyId(nim int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("NIM=?", nim).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(nim int64) Student {
	var student Student
	db.Where("NIM=?", nim).Delete(student)
	return student
}

