package routes

import (
	"radapp/src/api/v1/controllers"

	"github.com/gin-gonic/gin"
)

func CreateProdutoRoutes(routes *gin.Engine, produto controllers.IControllerProduto) {
	routes.POST("/produtos", produto.Save)
}
