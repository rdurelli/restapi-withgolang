package main

import (
	"database/sql"
	"log"
	"os"
	"radapp/src/api/v1/controllers"
	"radapp/src/api/v1/middleware"
	"radapp/src/api/v1/repositories"
	"radapp/src/api/v1/routes"
	"radapp/src/api/v1/services"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db          *sql.DB
	mysqlUser   string
	mysqlPass   string
	mysqlAddr   string
	mysqlPort   string
	mysqlSchema string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("not possible to load env variables")
	}
	mysqlUser = os.Getenv("MYSQL_USER")
	mysqlPass = os.Getenv("MYSQL_PASS")
	mysqlAddr = os.Getenv("MYSQL_ADDRESS")
	mysqlPort = os.Getenv("MYSQL_PORT")
	mysqlSchema = os.Getenv("MYSQL_SCHEMA")

}

func init() {

	DB, err := sql.Open("mysql", mysqlUser+":"+mysqlPass+"@tcp("+mysqlAddr+":"+mysqlPort+")/"+mysqlSchema+"?parseTime=true")
	if err != nil {
		log.Fatal("Something went wrong trying to connect to the database ")
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Something went wrong trying to ping to the database ")
	}
	db = DB
}

var (
	produtoController controllers.IControllerProduto
)

func init() {
	produtoController = controllers.NewControllerProduto(services.NewServiceProduto(repositories.NewRepositoryProduto(db)))

}

func main() {
	r := gin.Default()
	r.Use(gin.CustomRecovery(middleware.RecoverOfPanic))
	serverPort := os.Getenv("SERVER_PORT")
	routes.CreateProdutoRoutes(r, produtoController)
	r.Run(serverPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
