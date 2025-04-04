package main

import (
	"wine-shop/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
