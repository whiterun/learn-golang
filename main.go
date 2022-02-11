package main

import _ "fmt"

import "bagogo/routes"

func main() {
	r := routes.SetupRoutes()
    r.Run("localhost:8080")
	// fmt.Println("lorem")
}
