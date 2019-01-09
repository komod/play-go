package main

import "fmt"

func rangeOver() {
	arr := [8]int{}
	for idx, i := range arr {
		fmt.Println(idx, i)
	}
	for idx, i := range arr[:] {
		fmt.Println(idx, i)
	}
	for idx, i := range "string" {
		fmt.Println(idx, i)
	}

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	ch := make(chan int)
	go func() {
		for _, i := range m {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}

}
