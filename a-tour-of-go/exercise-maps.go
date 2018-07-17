package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)

	b := strings.Split(s, " ")
	for i := 0; i < len(b); i++ {
		if a[b[i]] > 0 {
			a[b[i]]++
		} else {
			a[b[i]] = 1
		}
	}

	return a
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog."))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	fmt.Println(WordCount("A man a plan a canal panama."))
}
