package main

import "fmt"

func main() {
	var multiple1, multiple2, sum int = 3, 5, 0

	for i := 3; i < 1000; i++ {
		if i%multiple1 == 0 || i%multiple2 == 0 {
			sum += i
		}
	}

	fmt.Println(sum)
}
