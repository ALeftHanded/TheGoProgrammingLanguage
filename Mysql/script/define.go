package script

import "gorm.io/gorm"

// DSN mysql password config
// * docker run --name mysql57 -p 3306:3306 -v ~/mysql_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=1q1q1q -d mysql:5.7
// ! if get error: ERROR 1045 (28000): Access denied for user 'root'@'172.17.0.1' (using password: YES)
// * CREATE USER 'test'@'172.17.0.1' IDENTIFIED BY '1q1q1q';
// * GRANT ALL PRIVILEGES ON *.* TO 'test'@'172.17.0.1' WITH GRANT OPTION;
// * flush privileges;
const DSN = "test:1q1q1q@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

func (User) TableName() string {
	return "user_tab"
}
