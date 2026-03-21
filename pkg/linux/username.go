package linux

import (
	"fmt"
	"os/user"
)

func Username() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user.Username
}
