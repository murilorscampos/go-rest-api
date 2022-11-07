package main

import (
	"fmt"
	"log"

	"github.com/murilorscampos/go-rest-api/database"
	"github.com/murilorscampos/go-rest-api/routes"
)

func main() {

	log.Println("Conectando ao banco de dados...")
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor com Go...")
	routes.HandleRequest()
}
