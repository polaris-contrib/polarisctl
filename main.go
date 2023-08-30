package main

import (
	"fmt"
	"os"

	"github.com/0226zy/polarisctl/pkg/cmd"
)

func main() {
	root := cmd.NewDefaultPolarisCommand()

	if err := root.Execute(); err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}
}
