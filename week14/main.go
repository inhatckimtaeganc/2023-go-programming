package main

import "fmt"

func status(name string) {
	balls := map[string]int{"성기훈": 20, "오일남": 0}
	//ball := balls[name]
	ball, exists := balls[name]
	fmt.Println(ball, exists)
}
func main() {
	status("오일남")
	status("강철")
	status("성기훈")
}
