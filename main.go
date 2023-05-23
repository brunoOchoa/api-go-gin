package main

import (
	"github.com/brunoOchoa/api-rest-gin/database"
	"github.com/brunoOchoa/api-rest-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
