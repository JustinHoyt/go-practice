package main

import "fmt"

func main() {
	s1 := []int{2, 3, 4}
	fmt.Println(s1)
	s2 := s1[:1]
	s2 = append(s2, 10)

	fmt.Println(s1)
	fmt.Println(s2)
}
