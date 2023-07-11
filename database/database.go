package database

import (
	"fmt"

	"github.com/Pedro-Cecilio/projeto-livro-go/env"
	"github.com/Pedro-Cecilio/projeto-livro-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func CreateConnection(){
	dns := fmt.Sprintf("%s:%s@tcp(localhost)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.Env.DB_USER, env.Env.DB_PASSWORD, env.Env.DB_NAME)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil{
		panic(err.Error())
	}
	err = db.Transaction(func(tx *gorm.DB) error { // Executa as migrations
		if err := tx.AutoMigrate(&models.User{}); err != nil{
			return err
		}
		if err := tx.AutoMigrate(&models.ValidToken{}); err != nil{
			return err
		}
		if err := tx.AutoMigrate(&models.BookRead{}); err != nil{
			return err
		}
		if err := tx.AutoMigrate(&models.BooksBeingRead{}); err != nil{
			return err
		}
		if err := tx.AutoMigrate(&models.BooksToRead{}); err != nil{
			return err
		}

		return nil
	})
	if err != nil {
		panic(err.Error())
	}
	DB = db

}