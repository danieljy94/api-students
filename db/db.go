package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    int
	Email  string
	Age    int
	Active bool
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=student port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	// Abre a conexão com o banco
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent(student Student) error {
	db := Init()

	// Resultado da criação do estudante
	if result := db.Create(&student); result.Error != nil {
		return result.Error
	}

	fmt.Println("Student created!")
	return nil
}

func GetStudents() ([]Student, error) {
	students := []Student{}

	db := Init()
	err := db.Find(&students).Error
	return students, err
}
