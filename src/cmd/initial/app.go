package initial

import (
	"database/sql"
	"product-app/src/internal/handler/http"
	"product-app/src/internal/repo/products"
	"product-app/src/internal/usecase"
)

func InitProductAPI(db *sql.DB) http.ProductHandler {

	// wire.Build(products.ProvideProductRepostiory, usecase.ProvideProductUsecase, http.ProvideProductHandler)

	repo := products.ProvideProductRepostiory(db)
	usecase := usecase.ProvideProductUsecase(repo)
	api := http.ProvideProductHandler(usecase)

	return api
}
