package script

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"AlgorithmGolang/libs/random"
)

func genRandomUserData(nameLen int) User {
	name := random.GenRandStrSeq(nameLen)
	return User{
		Name:  name,
		Email: name + "@example.com",
	}
}

func generateRandomUserDataList(count int) []User {
	users := make([]User, 0, count)
	for i := 0; i < count; i++ {
		users = append(users, genRandomUserData(10))
	}
	return users
}

func InsertRandomUserData(db *gorm.DB, users []User) {
	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			log.Println("Error inserting error, ", err.Error())
		}
	}
}

func BatchInsertRandomUserData(db *gorm.DB, users []User, batchSize int) {
	for i := 0; i < len(users)/batchSize; i++ {
		tmpUsers := users[i*batchSize : (i+1)*batchSize]
		err := db.CreateInBatches(&tmpUsers, batchSize).Error
		if err != nil {
			log.Println("Error inserting error, ", err.Error())
		}
	}
}

func ConcurrentInsertRandomUserData(db *gorm.DB, users []User, goroutineCount int) {
	var wg sync.WaitGroup
	wg.Add(goroutineCount)

	chunkSize := len(users) / goroutineCount
	for i := 0; i < goroutineCount; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == goroutineCount-1 {
			end = len(users) // Handle the remainder
		}

		go func(start, end int) {
			defer wg.Done()
			BatchInsertRandomUserData(db, users[start:end], end-start)
		}(start, end)
	}

	wg.Wait()
}

func truncateUserTable(db *gorm.DB) {
	db.Exec("truncate table user_tab")
}
