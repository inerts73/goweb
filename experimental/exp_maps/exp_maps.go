package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	products := []map[int]string{{1: "mobies"}, {2: "tv"}, {3: "laptops"}}
	fmt.Println(products)

	var info0 map[string]string
	info0 = make(map[string]string)
	info0["name"] = "ADAM"
	info0["age"] = "47"
	fmt.Println(info0)

	info1 := make(map[string]string)
	info1["name"] = "LANDY"
	info1["age"] = "37"
	fmt.Println(info1)

	info2 := map[string]string{"name": "Derek", "age": "5"}
	fmt.Println(info2)

	infos := []map[string]string{
		{"name": "Derek", "age": "5"},
		{"name": "Cindy", "age": "5"},
	}
	infos = append(infos, map[string]string{"name": "Landy", "age": "37"})
	fmt.Println(infos)

	id := "2"
	var product map[int]string
	for _, p := range products {
		for k := range p {
			pID, err := strconv.Atoi(id)
			if err != nil {
				os.Exit(1)
			}
			if pID == k {
				product = p
			}
		}

	}
	fmt.Println(product)

	var dummy map[string]string
	fmt.Print(dummy == nil)
}