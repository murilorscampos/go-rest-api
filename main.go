package main

import (
	"fmt"

	"github.com/murilorscampos/go-rest-api/database"
	"github.com/murilorscampos/go-rest-api/models"
	"github.com/murilorscampos/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{{Id: 1, Nome: "Nome 1", Historia: "Historia 1"}, {Id: 2, Nome: "Nome 2", Historia: "Historia 2"}}

	fmt.Println("Conetando ao banco de dados...")
	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Restcom Go...")
	routes.HandleRequest()
}
