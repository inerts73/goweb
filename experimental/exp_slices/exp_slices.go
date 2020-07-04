package main

import "fmt"

func main(){
	primes := [6]int{2,3,5,7,11,13}

	var s []int = primes[1:4]
	fmt.Println(s)

	names1 := []string{"ADAM", "DEREK"}
	fmt.Println(names1)

	names2 := make([]string, 2)
	names2[0] = "CINDY"
	names2[1] = "LANDY"
	fmt.Println(names2)

	test := map[string]string{"name": "ADAM", "gedner": "MALE"}
	fmt.Println("********************")
	for k, v := range test {
		fmt.Println(k, v)
	}
}