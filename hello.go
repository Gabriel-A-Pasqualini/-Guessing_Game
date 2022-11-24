package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {

	fmt.Println(">---|Go|---<")

	num_rand := randnum()

	var num int
	fmt.Println("Try one number:")
	fmt.Scanln(&num)

	fmt.Println(reflect.TypeOf(num_rand))

	if num_rand == num {
		fmt.Println("✨CONGRATULATIONS!✨")
		fmt.Println("You hit the number!")
		fmt.Println("The number was ", num_rand)
	} else {
		fmt.Println("Sorry...")
		fmt.Println("The number was ", num_rand)
	}
}

func randnum() int {
	num_rand := rand.Intn(100)
	return num_rand
}
