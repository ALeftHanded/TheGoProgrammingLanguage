package measureUtil

import (
	"fmt"
	"runtime"
)

func MemoryCost(fn interface{}, args ...interface{}) ([]interface{}, uint64) {
	name := getFunctionName(fn)

	callFunc, inputs := callPrepare(name, args)

	// Get the current memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	memStart := memStats.TotalAlloc

	// Call the function with the arguments
	result := callFunc.Call(inputs)

	// Get the new memory usage
	runtime.ReadMemStats(&memStats)
	memEnd := memStats.TotalAlloc
	totalAlloc := memEnd - memStart
	fmt.Printf("Function: %s Total Alloc: %d", name, totalAlloc)

	return convertRet(result), totalAlloc
}
