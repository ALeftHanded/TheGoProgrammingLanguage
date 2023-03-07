package script

import (
	"gorm.io/gorm"

	"AlgorithmGolang/Utils/random"
)

func InsertRandomUserData(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		name := random.GenRandStrSeq(10)
		email := name + "@example.com"
		user := User{Name: name, Email: email}
		db.Create(&user)
	}
}

func BatchInsertRandomUserData(db *gorm.DB, count, batchSize int) {
	for i := 0; i < count/batchSize; i++ {
		users := make([]User, 0, batchSize)
		for j := 0; j < batchSize; j++ {
			name := random.GenRandStrSeq(10)
			email := name + "@example.com"
			users = append(users, User{Name: name, Email: email})
		}
		db.CreateInBatches(&users, batchSize)
	}
}

func ConcurrentInsertRandomUserData(db *gorm.DB, count, concurrency int) {

}
