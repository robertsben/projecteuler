package main

import (
	"flag"
	"strconv"
	"reflect"
)

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