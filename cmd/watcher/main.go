package main

import (
	"fmt"
	"log"

	"github.com/casnerano/web-watcher/internal/config"
)

func main() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
