package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Levi-1103/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {

	userConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := &state{
		config: &userConfig,
	}

	cmds := commands{
		commandsMap: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	commandName := args[1]
	commandArgs := args[2:]

	err = cmds.run(appState, command{name: commandName, args: commandArgs})
	if err != nil {
		log.Fatal(err)
	}

}
