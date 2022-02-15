package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 4; j <= 6; j++ {
		fmt.Println(j)
	}

	p := 7
	for {
		if p <= 11 {
			fmt.Println(p)
			p += 1
		} else {
			break
		}
	}

	for n := 12; n < 25; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
