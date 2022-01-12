package main

import (
	"fmt"
	"strconv"
)

type employee struct {
	name   string
	age    int
	batch  string
	salary int
}

func (e employee) details() (bool, string) {
	if e.age < 22 {
		return false, ""
	} else {
		result := "Name: " + e.name + "," + " age: " + strconv.Itoa(e.age) + "," + " batch: " + e.batch + "," + " salary: " + strconv.Itoa(e.salary)
		return true, result
	}
}
func main() {
	e1 := employee{"sanjay", 25, "go", 30000}
	e2 := employee{"akash", 21, "ios", 30000}

	fmt.Println(e1.details())
	fmt.Println(e2.details())
}
