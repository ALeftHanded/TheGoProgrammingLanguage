package measureUtil

import (
	"reflect"
	"runtime"
	"strings"
)

func getFunctionName(fn interface{}) string {
	// Get the name of the function
	pc := reflect.ValueOf(fn).Pointer()
	ptr := runtime.FuncForPC(pc)
	name := strings.Split(ptr.Name(), ".")[1]
	return name
}

func convertRet(funcRes []reflect.Value) (ret []interface{}) {
	// Convert the result slice to a slice of interfaces{}
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
