package main

import "fmt"

func main() {
	dim := 100
	s := make([]int, 0, dim)

	for i := 0; i < dim; i++ {
		s = append(s, i+1)
	}

	s = append(s[:10], s[dim-10:]...)
	dim = len(s)

	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}
	fmt.Println(s)
}
