package main

import (
	"be-ep/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	r.Run(":8080")
}
