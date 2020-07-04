package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

// Person ...
type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	content, err := ioutil.ReadFile("persons.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var persons []Person
	json.Unmarshal(content, &persons)
	fmt.Println(persons)
    fmt.Println(persons[1].Age)
    
	content1, err1 := ioutil.ReadFile("single_person.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
    var singlePerson Person
	json.Unmarshal(content1, &singlePerson)
    fmt.Println(singlePerson) 
    fmt.Println(singlePerson.Age)
	fmt.Println("==================")

	var testsPtr *[]int = new([]int)
	*testsPtr = append(*testsPtr, 29)
	*testsPtr = append(*testsPtr, 30)
	fmt.Println(*testsPtr)
	fmt.Println(reflect.TypeOf(testsPtr).Kind())

	var tests []int = make([]int, 0)
	tests = append(tests, 129)
	tests = append(tests, 130)
	fmt.Println(tests)
	fmt.Println(reflect.TypeOf(tests).Kind())
}
