package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"xylight.dev/pastebin/db"
	"xylight.dev/pastebin/routes"
)

func main() {
	db.InitDB()

	_ = routes.NewServer()
}
