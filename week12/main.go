package main

import "fmt"

func main() {
	s := []int{0, 0, 0, 0, 0} // slice literal

	s[4] = 99
	s[2] = 11

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

	test := [3]string{"inha", "go", "student"}
	//testS := test[0:4] // error invalid argument: index 4 out of bounds [0:4]
	testS := test[0:2]
	fmt.Println(len(testS))
}
