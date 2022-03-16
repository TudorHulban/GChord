package main

import (
	"fmt"
	"gchord/fingers"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("missing echo node port")
	}

	n, errNew := fingers.NewEchoNode(os.Args[1])
	if errNew != nil {
		fmt.Printf("errNew: %s", errNew)
		os.Exit(1)
	}

	n.Stop()
}
