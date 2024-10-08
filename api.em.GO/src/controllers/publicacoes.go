package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"api.em.GO/src/autenticacao"
	"api.em.GO/src/banco"
	"api.em.GO/src/modelos"
	"api.em.GO/src/repositorios"
	"api.em.GO/src/respostas"
	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r * http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusCreated, publicacao)
}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacoes, erro := repositorio.Buscar(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	fmt.Printf("Publicações encontradas: %+v\n", publicacoes)
	respostas.JSON(w, http.StatusOK, publicacoes)
}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	publicacao, erro := repositorio.Buscar(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	fmt.Printf("Publicação encontrada: %+v\n", publicacao) 

	respostas.JSON(w, http.StatusOK, publicacao)
}
func AtualizarPublicacao(w http.ResponseWriter, r * http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	ppublicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if ppublicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao é possivel atualizar uma publicacao que nao seja sua"))
		return
	}
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var publicacao modelos.Publicacao
	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = publicacao.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = repositorio.Atualizar(publicacaoID, publicacao); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusNoContent, nil)
}
func DeletarPublicacao(w http.ResponseWriter, r * http.Request){
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)
	ppublicacaoSalvaNoBanco, erro := repositorio.BuscarPorID(publicacaoID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if ppublicacaoSalvaNoBanco.AutorID != usuarioID {
		respostas.Erro(w, http.StatusForbidden, errors.New("nao é possivel deletar uma publicacao que nao seja sua"))
		return
	}
	if erro = repositorio.DeletarPublicacao(publicacaoID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	} 
}
func BuscarPublicacoesPorUsuario(w http.ResponseWriter, r *http.Request) {
    parametros := mux.Vars(r)
    usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
    if erro != nil {
        respostas.Erro(w, http.StatusBadRequest, erro)
        return
    }

    db, erro := banco.Conectar()
    if erro != nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }
    defer db.Close()

    repositorio := repositorios.NovoRepositorioDePublicacoes(db)  
    publicacoes, erro := repositorio.BuscarPublicacoesPorUsuario(usuarioID)
    if erro != nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }

    respostas.JSON(w, http.StatusOK, publicacoes)
}
func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
    publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
    if erro != nil {
        respostas.Erro(w, http.StatusBadRequest, erro)
        return
    }

    db, erro := banco.Conectar()
    if erro != nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }
    defer db.Close()

    repositorio := repositorios.NovoRepositorioDePublicacoes(db)  
    if erro := repositorio.Curtir(publicacaoID); erro != nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }

    respostas.JSON(w, http.StatusNoContent, nil)
}
func DescurtiPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
    publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
    if erro != nil {
        respostas.Erro(w, http.StatusBadRequest, erro)
        return
    }

    db, erro := banco.Conectar()
    if erro != nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }
    defer db.Close()

    repositorio := repositorios.NovoRepositorioDePublicacoes(db)  
    if erro := repositorio.Descurtir(publicacaoID); erro != nil {
        respostas.Erro(w, http.StatusInternalServerError, erro)
        return
    }

    respostas.JSON(w, http.StatusNoContent, nil)
}