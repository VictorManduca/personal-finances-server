package main

import (
	"os"

	router "github.com/victormanduca/personal-finances/src/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router.Router().Run(":" + port)
}
