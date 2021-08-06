package controllers

import (
	"radapp/src/api/v1/services"

	"github.com/gin-gonic/gin"
)

type Reader interface {
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	FindByCategoria(c *gin.Context)
}

type Writer interface {
	Save(c *gin.Context)
	Update(c *gin.Context)
}

type IControllerProduto interface {
	Writer
	Reader
}

type controllerProduto struct {
	serviceProduto services.IServiceProduto
}

func NewControllerProduto(serviceProduto services.IServiceProduto) IControllerProduto {
	return &controllerProduto{
		serviceProduto: serviceProduto,
	}
}

func (controllerProduto) Save(c *gin.Context) {
	panic("implement me")
}

func (controllerProduto) Update(c *gin.Context) {
	panic("implement me")
}

func (controllerProduto) FindById(c *gin.Context) {
	panic("implement me")
}

func (controllerProduto) FindAll(c *gin.Context) {
	panic("implement me")
}

func (controllerProduto) FindByCategoria(c *gin.Context) {
	panic("implement me")
}
