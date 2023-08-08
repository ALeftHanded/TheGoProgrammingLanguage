package script

import (
	"sync"

	"gorm.io/gorm"

	"AlgorithmGolang/Utils/random"
)

func genRandomUserData(nameLen int) User {
	name := random.GenRandStrSeq(nameLen)
	return User{
		Name:  name,
		Email: name + "@example.com",
	}
}

func InsertRandomUserData(db *gorm.DB, count int) {
	for i := 0; i < count; i++ {
		user := genRandomUserData(10)
		db.Create(&user)
	}
}

func BatchInsertRandomUserData(db *gorm.DB, count, batchSize int) {
	for i := 0; i < count/batchSize; i++ {
		users := make([]User, 0, batchSize)
		for j := 0; j < batchSize; j++ {
			users = append(users, genRandomUserData(10))
		}
		db.CreateInBatches(&users, batchSize)
	}
}

func ConcurrentInsertRandomUserData(db *gorm.DB, count, concurrency int) {
	var wg sync.WaitGroup
	wg.Add(concurrency)

	chunkSize := count / concurrency

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < chunkSize; j++ {
				user := genRandomUserData(10)
				db.Create(&user)
			}
		}()
	}

	wg.Wait()
}
