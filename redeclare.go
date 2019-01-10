package main

import (
	"errors"
	"fmt"
)

func getWorld(s string) (string, error) {
	if s != "hello" {
		return "", errors.New("say hello")
	}
	return "world", nil
}

func redeclareGood() {
	var err error
	s, err := getWorld("hi")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func redeclareBad() {
	var s string
	if s == "" {
		s, err := getWorld("hello")
		if err != nil {
			fmt.Println(err)
			return
		}
		s += " peace"
	}
	fmt.Println(s)
}
