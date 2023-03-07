package script

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySqlConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	_ = db.AutoMigrate(&User{})
	return
}
