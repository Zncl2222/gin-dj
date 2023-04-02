package cmd

import (
	"fmt"
	"os"
)

func isGoModule() bool {
	_, err := os.Stat("go.mod")
	if err != nil {
		fmt.Println("go.mod does not exist in the current directory.")
		return false
	}
	return true
}
