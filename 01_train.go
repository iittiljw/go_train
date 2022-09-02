package main

import (
	"fmt"
)

/*
const (

	asa = iota
	b   = 1 << (iota * 10)
	c   = 1 << (iota * 10)
	d   = 1 << (iota * 10)
	e   = 1 << (iota * 10)
	f   = 1 << (iota * 10)
	g   = 1 << (iota * 10)

)

	func Map2String(m map[string]string) (ret string) {
		for k, v := range m {
			ret += fmt.Sprintf("%s=%s\n", k, v)
		}
		return
	}
*/
func main() {
	//	a := [...]int{5, 3, 2, 5, 12}
	//	fmt.Println(a)
	m := map[int]string{1: "OK", 2: "AA", 3: "BB", 4: "CC", 5: "DD"}

	//m["OK"] = 11
	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println(m)

	/*
		num := len(a)
		for i := 0; i < num; i++ {
			for j := i + 1; j < num; j++ {
				if a[i] < a[j] {
					temp := a[i]
					a[i] = a[j]
					a[j] = temp
				}

			}
		}
		fmt.Println(a)
	*/
	//fmt.Println(b)
	//	fmt.Println(c)
	//	fmt.Println(d)
	//	fmt.Println(e)
	//	fmt.Println(f)
	//	fmt.Println(g)

}
