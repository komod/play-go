package main

import "fmt"

func getSwitch(com string) {
	switch com {
	case "sony", "microsoft":
		fmt.Println("you wish")
	case "rayark":
		fmt.Println("magic")
		fallthrough
	case "nintendo":
		fmt.Println("switch get!")
	default:
		fmt.Println("no way")
	}
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func breakSwitch() {
	src := make([]int, 10)
	for i := range src {
		src[i] = i
	}
Loop:
	for i := 0; i < len(src); i++ {
		switch {
		case src[i]%2 != 0:
			fmt.Println("odd")
		case src[i]%5 == 1:
			break Loop
		}
	}
}
