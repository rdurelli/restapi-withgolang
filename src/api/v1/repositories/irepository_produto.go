package repositories

import "radapp/src/api/v1/models"

type Reader interface {
	FindById(id string) (*models.Produto, error)
	FindAll() (*[]models.Produto, error)
	FindByCategoria(categoria string) (*[]models.Produto, error)
}

type Writer interface {
	Save(produto *models.Produto) (*models.Produto, error)
	Update(id string, produto *models.Produto) (*models.Produto, error)
}

type IRepositoryProduto interface {
	Writer
	Reader
}
