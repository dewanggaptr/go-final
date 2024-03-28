package main

import (
	"final-project/databases"
	"final-project/routers"

	// "github.com/gin-gonic/gin"
)

func main() {
	database.DBInit()
	r := routes.StartApp()
	r.Run(":8080")
}
