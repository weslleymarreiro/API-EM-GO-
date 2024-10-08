package rotas

import (
	"net/http"
	"api.em.GO/src/controllers"
)

const UsuarioIDPath = "/usuarios/{usuarioId}"

type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

var rotasUsuarios = []Rota{
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		Uri:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuario}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioid}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri:                "/usuarios/{usuarioid}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},
	{
		Uri: 				"/usuarios/{usuarioId}/seguir",
		Metodo: 			http.MethodPost,
		Funcao: 			controllers.SeguirUsuario,
		RequerAutenticacao: true,		
	},
	{
		Uri: 				"/usuarios/{usuarioId}/parar-de-seguir",
		Metodo: 			http.MethodPost,
		Funcao: 			controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,		
	},
	{
		Uri: 				"/usuarios/{usuarioId}/seguidores",
		Metodo: 			http.MethodGet,
		Funcao: 			controllers.BuscarSeguidores,
		RequerAutenticacao: true,		
	},
	{
		Uri: 				"/usuarios/{usuarioId}/seguindo",
		Metodo: 			http.MethodGet,
		Funcao: 			controllers.BuscarSeguindo,
		RequerAutenticacao: true,		
	},
	{
		Uri: 				"/usuarios/{usuarioId}/atualizar-senha",
		Metodo: 			http.MethodPost,
		Funcao: 			controllers.AtualizarSenha,
		RequerAutenticacao: true,		
	},
}
