package main

import "fmt"

func main() {
	var S1, S2 string
	fmt.Scanln(&S1, &S2)

	if S1 == "sick" && S1 == S2 {
		fmt.Println(1)
	} else if S1 == "sick" {
		fmt.Println(2)
	} else if S2 == "sick" {
		fmt.Println(3)
	} else {
		fmt.Println(4)
	}
}
