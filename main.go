package main

import (
	"fmt"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Not supported!")
	} else {
		fmt.Println(currentUser.Username)
	}
}
