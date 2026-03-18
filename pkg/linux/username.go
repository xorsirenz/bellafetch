package linux

import (
	"fmt"
	"os/user"
)

func Username() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return currentUser.Username
}
