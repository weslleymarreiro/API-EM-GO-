package repositorios

import (
	"database/sql"
	"fmt"

	"api.em.GO/src/modelos"
	
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}
func (repositorio usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
    statement, erro := repositorio.db.Prepare(
        "INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
    )
    if erro != nil {
        return 0, erro
    }
    defer statement.Close()

    resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
    if erro != nil {
        return 0, erro
    }

    ultimoIDInserido, erro := resultado.LastInsertId()
    if erro != nil {
        return 0, erro
    }

    return uint64(ultimoIDInserido), nil
}
func (repositorio usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
    nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
    linhas, erro := repositorio.db.Query(
        "SELECT id, nome, nick, email, criado_em FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
        nomeOuNick, nomeOuNick,
    )
    if erro != nil {
        return nil, erro
    }
    defer linhas.Close()
    
    var usuarios []modelos.Usuario
    for linhas.Next() {
        var usuario modelos.Usuario

        if erro = linhas.Scan(
            &usuario.ID,
            &usuario.Nome,
            &usuario.Nick,
            &usuario.Email,
            &usuario.CriadoEm,
        ); erro != nil {
            return nil, erro
        }

        usuarios = append(usuarios, usuario)
    }
    return usuarios, nil
}
func (repositorio usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
    linhas, erro := repositorio.db.Query(
        "SELECT id, nome, nick, email, criado_em from usuarios where id = ?",
        ID,
    )
    if erro != nil {
        return modelos.Usuario{}, erro
    }
    defer linhas.Close()
    var usuario modelos.Usuario
    if linhas.Next() {
        if erro = linhas.Scan(
            &usuario.ID,
            &usuario.Nome,
            &usuario.Nick,
            &usuario.Email,
            &usuario.CriadoEm,
        ); erro != nil {
            return modelos.Usuario{}, erro
        }
    }
    return usuario, nil
}
func (repositorio usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
    statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")

    if erro != nil {
        return erro
    }
    defer statement.Close()

    if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); erro != nil {
        return erro
    }
    return nil
}
func (repositorio usuarios) Deletar(ID uint64) error {
    statement, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
    if erro != nil { 
        return erro
    }
    defer statement.Close()
    if _,erro = statement.Exec(ID); erro != nil { 
        return erro
    } 
    return nil
}
func (respositorio usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
    linha, erro := respositorio.db.Query("select id, senha from usuarios where email = ?", email)
    if erro != nil{
        return modelos.Usuario{}, erro
    }
    defer linha.Close()

    var usuario modelos.Usuario
    if linha.Next() {
        if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
            return modelos.Usuario{}, erro
        }
    }else {
        return modelos.Usuario{}, fmt.Errorf("usuario não encontrado")
    }
    return usuario, nil
}
func (repositorio usuarios) Seguir(usuarioID, seguidorID uint64) error {
    statement, erro := repositorio.db.Prepare(
        "insert into seguidores (usuario_id, seguidor_id) values (?, ?)",
    )
    if erro != nil {
        return erro
    }
    defer statement.Close()

    if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
        return erro
    }
    return nil 
}
func (repositorio usuarios) PararDeSeguir(usuarioID, seguidorID uint64) error {
    statement, erro := repositorio.db.Prepare(
        "delete from seguidores where usuario_id = ? and seguidor_id = ?",
    )
    if erro != nil {
        return erro
    }
    defer statement.Close()

    if _, erro = statement.Exec(usuarioID, seguidorID); erro != nil {
        return erro
    }
    return nil 
}
func (repositorio usuarios) BuscarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
    linhas, erro := repositorio.db.Query(`
        SELECT u.id, u.nome, u.nick, u.email, u.criado_em 
        FROM usuarios u 
        INNER JOIN seguidores s ON u.id = s.seguidor_id 
        WHERE s.usuario_id = ?
    `, usuarioID)
    
    if erro != nil {
        return nil, erro
    }
    defer linhas.Close()

    var usuarios []modelos.Usuario
    for linhas.Next() { 
        var usuario modelos.Usuario
        if erro = linhas.Scan(
            &usuario.ID,
            &usuario.Nome,
            &usuario.Nick,
            &usuario.Email,
            &usuario.CriadoEm,
        ); erro != nil {
            return nil, erro
        }
        usuarios = append(usuarios, usuario)
    }
    return usuarios, nil
}
func (repositorio usuarios) BuscarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
    linhas, erro := repositorio.db.Query(`
        SELECT u.id, u.nome, u.nick, u.email, u.criado_em 
        FROM usuarios u 
        INNER JOIN seguidores s ON u.id = s.usuario_id 
        WHERE s.seguidor_id = ?
    `, usuarioID)
    
    if erro != nil {
        return nil, erro
    }
    defer linhas.Close()

    var usuarios []modelos.Usuario
    for linhas.Next() { 
        var usuario modelos.Usuario
        if erro = linhas.Scan(
            &usuario.ID,
            &usuario.Nome,
            &usuario.Nick,
            &usuario.Email,
            &usuario.CriadoEm,
        ); erro != nil {
            return nil, erro
        }
        usuarios = append(usuarios, usuario)
    }
    return usuarios, nil
}
func (repositorio usuarios) BuscarSenha(usuarioID uint64) (string, error) {
    linha, erro := repositorio.db.Query("select senha from usuarios where id = ?", usuarioID)
    if erro != nil {
        return "", erro
    }
    defer linha.Close()

    var usuario modelos.Usuario
    if linha.Next() {
        if erro = linha.Scan(&usuario.Senha); erro != nil {
            return "", erro
        }
    }
    return usuario.Senha, nil
}
func (repositorio usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
    statement, erro := repositorio.db.Prepare("upadate usuarios set senha = ? where id = ?")
    if erro != nil {
        return erro 
    }
    defer statement.Close()

    if _, erro = statement.Exec(senha, usuarioID); erro != nil {
        return erro
    }
    return nil

}
