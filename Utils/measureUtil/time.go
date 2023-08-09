package measureUtil

import (
	"fmt"
	"time"
)

func ExecutionTime(fn interface{}, args ...interface{}) ([]interface{}, time.Duration) {
	// call Prepare
	name := getFunctionName(fn)
	callFunc, inputs := callPrepare(fn, args)

	// Call the function with the arguments
	start := time.Now()
	result := callFunc.Call(inputs)
	totalTime := time.Since(start)
	fmt.Printf("Function %s took %s\n", name, totalTime)

	return convertRet(result), totalTime
}
