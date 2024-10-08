package repositorios

import (
	"database/sql"

	"api.em.GO/src/modelos"
) 

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
} 
func (repositorio Publicacoes) Criar(Publicacao modelos.Publicacao) (uint64, error) {
	statemnt, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",)
	if erro != nil { 
		return 0, erro 
	}
	defer statemnt.Close()

	resultado, erro := statemnt.Exec(Publicacao.Titulo, Publicacao.Conteudo, Publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}
	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil
}
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		"SELECT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id = p.autor_id WHERE p.id = ?", publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
		return publicacao, nil
	}

	return modelos.Publicacao{}, nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(
		"SELECT DISTINCT p.*, u.nick FROM publicacoes p INNER JOIN usuarios u ON u.id =  p.autor_id INNER JOIN seguidores s ON p.autor_id = s.usuario_id WHERE u.id = ? OR s.seguidor_id = ? order by 1 desc", usuarioID,  usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linha.Close()

	var publicacoes []modelos.Publicacao

	for linha.Next() {
		var publicacao modelos.Publicacao

		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil 
}
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil { 
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	} 
	return nil
}
func (repositorio Publicacoes) DeletarPublicacao(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil { 
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	} 
	return nil
}  
func (repositorio Publicacoes) BuscarPublicacoesPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT p.*, u.nick FROM publicacoes p JOIN usuarios u ON u.id = p.autor_id WHERE p.autor_id = ?",
		usuarioID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadoEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil 
}
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statent, erro := repositorio.db.Prepare("update publicacoes set curtidas= curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statent.Close()

	if _, erro = statent.Exec(publicacaoID); erro != nil {
		return erro 
	}
	return nil
}
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statent, erro := repositorio.db.Prepare("update publicacoes set curtidas = CASE WHEN curtidas > 0 THEN curtidas - 1 ELSE 0 END where id = ?")
	if erro != nil {
		return erro
	}
	if _, erro = statent.Exec(publicacaoID); erro != nil {
		return erro 
	}
	return nil
}
