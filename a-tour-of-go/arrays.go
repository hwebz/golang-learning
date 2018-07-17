package main

import "fmt"

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	g := names[0:2]
	h := names[1:3]
	fmt.Println(g, h)

	h[0] = "XXX"
	fmt.Println(g, h)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, false}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{7, false},
		{15, true},
	}
	fmt.Println(s2)

	s3 := s2[1:4]
	fmt.Println(s3)

	s4 := s2[:2]
	fmt.Println(s4)

	s5 := s2[1:]
	fmt.Println(s5)

	printSlices(s)
	printSlices(q)

	var s6 []int
	fmt.Println(s6, len(s6), cap(s6))

	if s == nil {
		fmt.Println("nil!")
	}
}
