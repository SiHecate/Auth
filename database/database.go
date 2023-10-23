package database

import (
	"fmt"

	model "Auth/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() {
	Database()
	Migrate()

}

func Database() {
	dsn := "host=postgres user=postgres password=393406 dbname=Auth port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Migrate() {
	fmt.Println("Migrating...")
	Conn.AutoMigrate(
		&model.User{},
	)
}
