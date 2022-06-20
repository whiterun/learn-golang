package main

import _ "fmt"

import (
	"log"
	"bagogo/routes"
)

func main() {
	r := routes.SetupRoutes()

    log.Fatal(r.Run("localhost:3000"))
	// fmt.Println("lorem")
}
