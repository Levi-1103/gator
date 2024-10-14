package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {

	userConfig, _ := config.Read()
	userConfig.SetUser("levi")
	userConfig, _ = config.Read()

	fmt.Println(userConfig)

}
