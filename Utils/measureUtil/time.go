package measureUtil

import (
	"fmt"
	"time"
)

func ExecutionTime(fn interface{}, args ...interface{}) ([]interface{}, time.Duration) {
	start := time.Now()

	name := getFunctionName(fn)

	callFunc, inputs := callPrepare(fn, args)

	// Call the function with the arguments
	result := callFunc.Call(inputs)
	totalTime := time.Since(start)
	fmt.Printf("Function %s took %s\n", name, totalTime)

	return convertRet(result), totalTime
}
