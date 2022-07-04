package main

import (
	router "github.com/victormanduca/personal-finances/src/routes"
)

func main() {
	router.Router().Run("localhost:8000")
}
