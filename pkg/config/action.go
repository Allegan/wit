package config

import (
	"errors"
	"fmt"
	"os"
)

// Action wraps the command slices used for determining actions
type Action struct {
	Verb   string
	Target string
	File   string
}

// ParseCommandLine parses the command line parameters for the application
// and returns an Action struct with the information
func (a *Action) ParseCommandLine() (action *Action, err error) {
	action = &Action{"", "", ""}
	args := os.Args[1:]

	fmt.Println(len(args))

	if len(args) < 1 {
		return nil, errors.New("There must be at least 2 arguments provided")
	}

	action.Verb = args[0]
	action.Target = args[1]

	if len(args) == 3 {
		action.File = args[2]
	}

	return action, nil
}
