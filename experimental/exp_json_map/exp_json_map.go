package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
//	"reflect"
)

func main() {
	content, err := ioutil.ReadFile("single_person.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	var singlePerson map[string]interface{}
	json.Unmarshal(content, &singlePerson)
	fmt.Println(singlePerson)
	// fmt.Println(singlePerson["friends"])
	// fmt.Println(reflect.TypeOf(singlePerson["friends"]).Kind())

	fmt.Println("========================")

	content1, err1 := ioutil.ReadFile("persons.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}

	var persons []map[string]interface{}
	json.Unmarshal(content1, &persons)
	// for _, v := range persons {
	// 	fmt.Println(v)
	// }
	fmt.Println(persons)
}
