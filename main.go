package main

import (
	"github.com/taufiqnasrullah195/task-5-vix-btpns-TaufiqNashrullah/tree/main/database"
	"github.com/taufiqnasrullah195/task-5-vix-btpns-TaufiqNashrullah/tree/main/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
