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

func redeclare() {
	var err error

	s, err := getWorld("hi")
	if err != nil {
		fmt.Println("error: ", err)
	}

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
