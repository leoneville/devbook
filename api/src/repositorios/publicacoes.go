package repositorios

import (
	"api/src/models"
	"database/sql"
)

type publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *publicacoes {
	return &publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma única publicação do banco de dados
func (repositorio publicacoes) BuscarPorID(publicacaoID uint64) (models.Publicacao, error) {
	linha, err := repositorio.db.Query(`
		SELECT p.*, u.nick from
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	)
	if err != nil {
		return models.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao models.Publicacao

	if linha.Next() {
		if err = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); err != nil {
			return models.Publicacao{}, err
		}
	}

	return publicacao, nil
}
