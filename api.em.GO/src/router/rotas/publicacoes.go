package rotas

import (
	"net/http"
	"api.em.GO/src/controllers"
)

const publicacaoIDPath = "/publicacoes/{publicacaoId}"

var rotasPublicacoes = []Rota{
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true, 
	},
	{
		Uri:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacoes,
		RequerAutenticacao: true, 
	},
	{
		Uri:                publicacaoIDPath,
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacao,
		RequerAutenticacao: true, 
	},
	{
		Uri:                publicacaoIDPath,
		Metodo:             http.MethodPut, 
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true, 
	},
	{
		Uri:                publicacaoIDPath,
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublicacao,
		RequerAutenticacao: true, 
	},
	{
		Uri:                "/usuarios/{usuarioID}/publicacoes",
		Metodo:             http.MethodGet, 
		Funcao:             controllers.BuscarPublicacoesPorUsuario,
		RequerAutenticacao: true, 
	},
	{
		Uri:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true, 
	},
	{
		Uri:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtiPublicacao,
		RequerAutenticacao: true, 
	},
}
