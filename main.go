package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0226zy/polarisctl/pkg/cmd"
)

func main() {
	start := time.Now()
	root := cmd.NewDefaultPolarisCommand()

	if err := root.Execute(); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("cost:%s\n\n", time.Since(start))
}
