package main

import "fmt"

func main() {
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for k := 10; k <= 20; k++ {
		if k%2 == 0 {
			continue
		}

		if k == 17 {
			break
		}

		fmt.Println(k)
	}
}
