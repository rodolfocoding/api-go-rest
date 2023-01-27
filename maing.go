package main

import (
	"fmt"

	"github.com/rodolfocoding/api-go-rest/models"
	"github.com/rodolfocoding/api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Ininciando o servidor Rest com Go")
	routes.HandleRequest()
}
