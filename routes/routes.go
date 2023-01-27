package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfocoding/api-go-rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonaldiade).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))
}
