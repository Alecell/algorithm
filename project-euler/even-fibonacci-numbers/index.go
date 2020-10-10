package main

import "fmt"

func main() {
	var prev, cur, sum int32 = 1, 1, 0

	for cur < 4000000 {
		var tmp = prev + cur
		prev = cur
		cur = tmp

		if tmp%2 == 0 {
			sum += tmp
		}
	}

	fmt.Println(sum)
}
