package utils

import (
	"fmt"
	"log"
	"os/user"
)

func Home() string {
	current, err := user.Current()
	Check(err)
	return current.HomeDir
}

func Check(err error) {
	if err != nil {
		log.Fatalf("fatal: %v\n", err)
	}
}

func EndlessSelect(max int, index int) int {
	print("\nChoose one out of them: ")
	fmt.Scanf("%d", &index)
	if index <= max {
		return index
	}
	return EndlessSelect(max, index)
}
