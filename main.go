package main

import (
	"fmt"
	"log"

	"github.com/Levi-1103/gator/internal/config"
)

func main() {

	userConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	err = userConfig.SetUser("levi")
	if err != nil {
		log.Fatalf("error setting config: %v", err)
	}

	userConfig, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println(userConfig)

}
