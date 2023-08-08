package script

import (
	"testing"

	"AlgorithmGolang/Utils/measureUtil"
)

func TestInsertRandomUserDataSpeed(t *testing.T) {
	db, err := InitMySqlConnection()
	if err != nil {
		panic(err)
	}
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insert 10 random user data",
			args: args{
				count: 10,
			},
		},
		{
			name: "insert 1000 random user data",
			args: args{
				count: 1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := generateRandomUserDataList(tt.args.count)
			const batchSize = 50
			const goroutineCount = 20
			measureUtil.ExecutionTime(InsertRandomUserData, db, users)
			truncateUserTable(db)
			measureUtil.ExecutionTime(BatchInsertRandomUserData, db, users, batchSize)
			truncateUserTable(db)
			measureUtil.ExecutionTime(ConcurrentInsertRandomUserData, db, users, goroutineCount)
			truncateUserTable(db)
		})
	}
}
