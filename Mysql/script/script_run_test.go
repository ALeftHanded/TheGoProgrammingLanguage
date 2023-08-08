package script

import (
	"testing"

	"AlgorithmGolang/Utils/measureUtil"
)

func TestInsertRandomUserData(t *testing.T) {
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
			measureUtil.ExecutionTime(BatchInsertRandomUserData, db, tt.args.count, 100)
			measureUtil.ExecutionTime(InsertRandomUserData, db, tt.args.count)
			measureUtil.ExecutionTime(ConcurrentInsertRandomUserData, db, tt.args.count, 100)
		})
	}
}
