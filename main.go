package main

import (
	"fmt"
	"log"

	"github.com/JustSomeHack/one-oauth2-server/controllers"
)

var version string

func main() {
	fmt.Printf("Starting one-oauth2-server %s\n\n", version)

	router, err := controllers.SetupRouter()
	if err != nil {
		log.Fatal(err)
	}

	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
