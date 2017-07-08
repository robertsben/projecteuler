package main

import (
	"flag"
	"fmt"
	"strconv"
	"reflect"
)

type Solution struct {}
var sol Solution

func callRecovery(funcname string, fargs []interface{}) {
	if r := recover(); r != nil {
		fmt.Printf("Calling function %s with args %v failed because: ", funcname, fargs)
		fmt.Println(r)
		fmt.Println("Please ensure the function exists and takes the supplied arguments.")
	}
}

func argConvertor(argValue reflect.Value, cargType reflect.Type) interface{} {
	var conversion interface{}
	var err error

	switch argValue.Kind(); cargType.Kind() {
	case reflect.String, reflect.Int:
		conversion, err = strconv.Atoi(argValue.String())
	default:
		conversion, err = argValue.Convert(cargType), nil
	}

	if err != nil {
		panic(err)
	}

	return conversion
}


func call(funcname string, fargs []interface{}) []reflect.Value {
	defer callRecovery(funcname, fargs)

	callable := reflect.ValueOf(sol).MethodByName(funcname)
	cargs := make([]reflect.Value, len(fargs))

	for idx := range fargs {
		cargs[idx] = reflect.ValueOf(
			argConvertor(
				reflect.ValueOf(fargs[idx]),
				callable.Type().In(idx),
			),
		)
	}

	return callable.Call(cargs)
}

func parseArgs(args []string) []interface{} {
	fargs := make([]interface{}, len(args))
	for i, v := range args {
		fargs[i] = v
	}
	return fargs
}

func main() {
	var problem int
	flag.IntVar(&problem, "problem", 1, "number of problem to start")
	flag.Parse()

	funcname := "Problem" + strconv.Itoa(problem)
	solution := call(funcname, parseArgs(flag.Args()))

	for i := 0; i<len(solution); i++ {
		fmt.Printf("Solution part %d: %s\n", i, solution[i])
	}
}