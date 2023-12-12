package manager

import (
	"fmt"
	"os"
)

func ReadPublicKey() {
	file, err := os.Open("~/.ssh/id_rsa.pub")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Println(file)
}
