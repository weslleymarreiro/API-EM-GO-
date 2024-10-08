package modelos

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"api.em.GO/src/seguranca"
	"github.com/badoux/checkmail"
)

type Usuario struct {
    ID        uint64    `json:"id,omitempty"`
    Nome      string    `json:"nome,omitempty"`
    Nick      string    `json:"nick,omitempty"`
    Email     string    `json:"email,omitempty"`
    Senha     string    `json:"senha,omitempty"`
    CriadoEm  time.Time `json:"criado_em,omitempty"`
}

func (usuario *Usuario) Preparar(etapa string) error {
    if erro := usuario.validar(etapa); erro != nil {
        return erro
    }
    if erro := usuario.formata(etapa); erro != nil {
        return erro
    }
    return nil
}

func (usuario *Usuario) validar(etapa string) error {
    if usuario.Nome == "" {
        return errors.New("o nome é obrigatório e não pode estar em branco")
    }
    fmt.Println(usuario)
    if usuario.Nick == "" {
        return errors.New("o nick é obrigatório e não pode estar em branco")
    }
    if usuario.Email == "" {
        return errors.New("o e-mail é obrigatório e não pode estar em branco")
    }

    if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
        return errors.New("o e-mail inserido é invalido")
    }

    if etapa == "cadastro" && usuario.Senha == "" {
        return errors.New("a senha é obrigatória e não pode estar em branco")
    } 
    return nil
}

func (usuario *Usuario) formata(etapa string) error {
    usuario.Nome = strings.TrimSpace(usuario.Nome)
    usuario.Nick = strings.TrimSpace(usuario.Nick)
    usuario.Email = strings.TrimSpace(usuario.Email)

    if etapa == "cadastro" {
        senhaComHash, erro := seguranca.Hash(usuario.Senha)
        if erro != nil { 
            return erro
        }
        usuario.Senha = string(senhaComHash)
    }
    return nil
}