package main

import (
	"github.com/deyweson/go-gin-api-rest/db"
	"github.com/deyweson/go-gin-api-rest/routes"
)

func main() {
	db.ConectarDB()
	routes.HandleRequest()

}
