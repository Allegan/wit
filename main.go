package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Allegan/wit/pkg/config"
)

func main() {
	conf := &config.Config{}
	action := &config.Action{}

	// read config from file
	conf, err := conf.ReadFromFile("./.witrc")
	if err != nil {
		log.Fatalln(fmt.Errorf("Quitting wit: \n\t%w", err))
		os.Exit(1)
	}

	// parse action from command line
	action, err = action.ParseCommandLine()
	if err != nil {
		log.Fatalln(fmt.Errorf("Quitting wit: \n\t%w", err))
		os.Exit(1)
	}
}
