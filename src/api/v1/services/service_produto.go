package services

import (
	"radapp/src/api/v1/models"
	"radapp/src/api/v1/repositories"
)

type serviceProduto struct {
	repository repositories.IRepositoryProduto
}

func NewServiceProduto(repository repositories.IRepositoryProduto) IServiceProduto {
	return &serviceProduto{
		repository: repository,
	}
}

func (sP serviceProduto) Save(produto *models.Produto) (*models.Produto, error) {
	produto, err := sP.repository.Save(produto)
	if err != nil {
		panic(err)
	}
	return produto, err
}

func (serviceProduto) Update(id string, produto *models.Produto) (*models.Produto, error) {
	panic("implement me")
}

func (serviceProduto) FindById(id string) (*models.Produto, error) {
	panic("implement me")
}

func (serviceProduto) FindAll() (*[]models.Produto, error) {
	panic("implement me")
}

func (serviceProduto) FindByCategoria(categoria string) (*[]models.Produto, error) {
	panic("implement me")
}
