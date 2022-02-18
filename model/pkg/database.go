package pkg

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//mysqlに接続する
func Connect() *gorm.DB {
	d := os.Getenv("DATABASE_DOMAIN")
	n := os.Getenv("DATABASE_NAME")
	un := os.Getenv("DATABASE_USER_NAME")
	up := os.Getenv("DATABASE_USER_PASSWORD")
	//ex) "root:root@tcp(db:3306)/emotion_sns?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", un, up, d, n)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return db
}
