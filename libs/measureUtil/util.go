package measureUtil

import (
	"reflect"
	"runtime"
	"strings"
)

// getFunctionName Get the name of the function
func getFunctionName(fn interface{}) string {
	pc := reflect.ValueOf(fn).Pointer()
	return strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]
}

// convertRet Convert the result slice to a slice of interfaces{}
func convertRet(funcRes []reflect.Value) (ret []interface{}) {
	for _, r := range funcRes {
		ret = append(ret, r.Interface())
	}
	return ret
}

func callPrepare(fn interface{}, args interface{}) (callFunc reflect.Value, inputs []reflect.Value) {
	// Convert the function to a reflect.Value
	callFunc = reflect.ValueOf(fn)
	// Convert the arguments to a slice of reflect.Values
	for _, arg := range args.([]interface{}) {
		inputs = append(inputs, reflect.ValueOf(arg))
	}
	return
}
