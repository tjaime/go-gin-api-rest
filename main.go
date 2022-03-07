package main

import (
	"github.com/tjaime/go-gin-api-rest/db"
	"github.com/tjaime/go-gin-api-rest/routes"
)

func main() {
	db.ConnectDb()
	routes.HandleRequest()
}
