package rotas

import (
	"api.em.GO/src/controllers" 
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/usuario", controllers.CriarUsuario).Methods("POST")
	return Configurar(r) 
}
