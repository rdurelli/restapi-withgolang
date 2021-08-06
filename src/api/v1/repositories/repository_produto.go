package repositories

import (
	"database/sql"
	"radapp/src/api/v1/models"

	"github.com/google/uuid"
)

type repositoryProduto struct {
	db *sql.DB
}

func NewRepositoryProduto(db *sql.DB) IRepositoryProduto {
	return &repositoryProduto{
		db: db,
	}
}

func (rP repositoryProduto) Save(produto *models.Produto) (*models.Produto, error) {
	stmt, err := rP.db.Prepare("INSERT INTO produto VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic("Unable to get uuid")
	}
	result, err := stmt.Exec(uuid.ID(), produto.Nome, produto.Preco)
	if err != nil {
		panic("Unable to execute statement")
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic("Unable to get ID")
	}
	produto.Id = string(id)
	return produto, nil
}

func (rP repositoryProduto) Update(id string, produto *models.Produto) (*models.Produto, error) {
	panic("implement me")
}

func (rP repositoryProduto) FindById(id string) (*models.Produto, error) {
	panic("implement me")
}

func (rP repositoryProduto) FindAll() (*[]models.Produto, error) {
	panic("implement me")
}

func (rP repositoryProduto) FindByCategoria(categoria string) (*[]models.Produto, error) {
	panic("implement me")
}
