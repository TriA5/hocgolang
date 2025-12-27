package main

import (
	"be-ep/db"
	"be-ep/migrations"
	"be-ep/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	migrations.RunMigration()

	router.SetupRouter(r)

	r.Run(":8080")
}
