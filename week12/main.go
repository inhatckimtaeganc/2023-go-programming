package main

import "fmt"

func main() {
	a := make([]string, 4, 5)
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2]
	//c := append(a, "y") // 메모리공간이 늘어나지 않는다
	c := append(a, "y", "x") // 메모리 공간을 늘린다
	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
	as[1] = "z"
	fmt.Printf("%x %x %x\n", &a[0], &as[0], &c[0])
	c[0] = "Q"
	fmt.Println(a, len(a), cap(a))
	fmt.Println(c, len(c), cap(c))
}
