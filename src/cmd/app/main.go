package main

import (
	"database/sql"
	"product-app/src/cmd/initial"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root:1234567890@tcp(127.0.0.1:3306)/Phonebook")
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := initial.InitProductAPI(db)

	r := gin.Default()
	r.GET("/products", productAPI.GetProducts())
	r.GET("/product/:id", productAPI.GetUser())
	r.POST("/products", productAPI.CreateProduct())
	r.PUT("/product/:id", productAPI.UpdateProduct())
	r.DELETE("/product/:id", productAPI.DeleteProduct())

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
