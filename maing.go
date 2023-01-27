package main

import (
	"fmt"

	"github.com/rodolfocoding/api-go-rest/routes"
)

func main() {
	fmt.Println("Ininciando o servidor Rest com Go")
	routes.HandleRequest()
}
