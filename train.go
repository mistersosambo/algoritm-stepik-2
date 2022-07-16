package main

import "fmt"

func main() {
	var n, z int
	m := make([]int, n, n)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&z)
		m = append(m, z)
	}
	fmt.Println(m[3])
}
