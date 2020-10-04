package main

import (
	"fmt"

	"github.com/Allegan/wit/pkg/config"
)

func main() {
	var err error

	defer func() {
		if err != nil {
			errMsg := fmt.Errorf("[FATAL]%w", err)
			fmt.Println(errMsg)
		}
	}()

	flags, err := config.New()
	if err != nil {
		return
	}

	fmt.Printf("%v\n", flags.Host)
}
