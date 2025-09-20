package db

import (
	"github.com/BellaMez/api-students/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {

	}

	db.AutoMigrate(&schemas.Student{})

	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student schemas.Student) error {
	// result := db.Create(&student)
	// if result.Error != nil {
	// 	fmt.Println("Error to create student")
	// }
	if result := s.DB.Create(&student); result.Error != nil {
		log.Error().Msg("Failed to crete student")
		return result.Error
	}

	log.Info().Msg("Create sudents!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]schemas.Student, error) {
	students := []schemas.Student{}
	err := s.DB.Find(&students).Error
	return students, err
}
func (s *StudentHandler) GetFilteredStudents(active bool) ([]schemas.Student, error) {
	filteredsStudents := []schemas.Student{}
	err := s.DB.Where("active = ? ", active).Find(&filteredsStudents)
	return filteredsStudents, err.Error
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	var student schemas.Student
	err := s.DB.First(&student, id)

	return student, err.Error
}

func (s *StudentHandler) UpdateStudent(updatestudent schemas.Student) error {
	return s.DB.Save(&updatestudent).Error
}

func (s *StudentHandler) DeleteStudent(student schemas.Student) error {
	return s.DB.Delete(&student).Error
}
