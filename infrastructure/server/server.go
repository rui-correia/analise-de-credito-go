package server

import (
	"log"
)

func Start() {

	if err := run("8080"); err != nil {
		log.Printf("Error: %v", err)
	}
}

func run(port string) error {
	router := mapRoutes()
	return router.Run(":" + port)
}
