package rotas

import (
	"net/http"

	"api.em.GO/src/controllers"
)

var rotaLogin = Rota{
	Uri:               "/login",
	Metodo:            http.MethodPost,
	Funcao:            controllers.Login,
	RequerAutenticacao: false,
}