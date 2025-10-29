package db

/*
 import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
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

// Documentação GORM
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

// Adicionar um estudante
func AddStudent() {
	db := Init()
	student := Student{
		Name:   "Daniel",
		CPF:    12345678910,
		Email:  "daniel.jy94@gmail.com",
		Age:    31,
		Active: true,
	}

	// Resultado da criação do estudante
	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Error to create a student")
	}

	fmt.Println("Student created!")
}
*/

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

func AddStudent() {
	db := Init()
	student := Student{
		Name:   "Daniel",
		CPF:    12345678910,
		Email:  "daniel.jy94@gmail.com",
		Age:    31,
		Active: true,
	}

	// Resultado da criação do estudante
	if result := db.Create(&student); result.Error != nil {
		fmt.Println("Error to create a student")
	}

	fmt.Println("Student created!")
}
