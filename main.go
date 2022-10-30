package main

import (
	"fmt"

	"github.com/murilorscampos/go-rest-api/database"
	"github.com/murilorscampos/go-rest-api/routes"
)

func main() {

	fmt.Println("Conetando ao banco de dados...")
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Restcom Go...")
	routes.HandleRequest()
}
