package main

import (
	"flag"
	"fmt"
	"strconv"
	"reflect"
)

func callRecovery(funcname string, fargs []string) {
	if r := recover(); r != nil {
		fmt.Printf("Calling function %s with args %v failed.\n", funcname, fargs)
		fmt.Println("Please ensure the function exists and takes the supplied arguments.")
	}
}


func call(funcname string, fargs []string) {
	defer callRecovery(funcname, fargs)

	callable := reflect.ValueOf(funcname)
	var cargs []reflect.Value

	for idx, val := range fargs {
		cargs[idx] = reflect.ValueOf(val)
	}

	callable.Call(cargs)
}

func main() {
	var problem int
	flag.IntVar(&problem, "problem", 1, "number of problem to start")
	flag.Parse()

	funcname := "Problem" + strconv.Itoa(problem)
	call(funcname, flag.Args())
}