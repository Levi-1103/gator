package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Levi-1103/gator/internal/config"
	"github.com/Levi-1103/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	config *config.Config
	db     *database.Queries
}

func main() {

	userConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appState := &state{
		config: &userConfig,
	}

	db, err := sql.Open("postgres", userConfig.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	appState.db = dbQueries

	cmds := commands{
		commandsMap: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)

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
