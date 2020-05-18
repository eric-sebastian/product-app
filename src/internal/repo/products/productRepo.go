package products

import (
	"database/sql"
	"log"
	"product-app/src/internal/entity/product"
)

type ProductRepository struct {
	DB *sql.DB
}

func ProvideProductRepostiory(DB *sql.DB) ProductRepository {
	return ProductRepository{DB: DB}
}

func (p *ProductRepository) GetAll() []product.Product {
	var singleProduct product.Product
	var multiProduct []product.Product

	rows, err := p.DB.Query("SELECT * FROM products")
	handleError(err)
	for rows.Next() {
		err = rows.Scan(&singleProduct.ID, &singleProduct.Name, &singleProduct.Price)
		handleError(err)

		multiProduct = append(multiProduct, singleProduct)
	}

	return multiProduct
}

func (p *ProductRepository) GetById(id int) product.Product {
	var singleProduct product.Product

	rows, err := p.DB.Query("SELECT * FROM products WHERE id = ?", id)
	handleError(err)
	for rows.Next() {
		err = rows.Scan(&singleProduct.ID, &singleProduct.Name, &singleProduct.Price)
		handleError(err)
	}

	return singleProduct
}

func (p *ProductRepository) CreateProduct(newProduct product.ProductForm) product.Product {
	var CreatedProduct product.Product

	tx, err := p.DB.Begin()
	handleError(err)
	res, err := tx.Exec("INSERT INTO products (name, price) VALUES (?, ?)", newProduct.Name, newProduct.Price)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	handleError(err)

	CreatedProduct.ID = int(id)
	CreatedProduct.Name = newProduct.Name
	CreatedProduct.Price = newProduct.Price
	handleError(tx.Commit())

	return CreatedProduct
}

func (p *ProductRepository) UpdateProduct(id int, modifyProduct product.ProductForm) product.Product {
	var UpdatedProduct product.Product

	tx, err := p.DB.Begin()
	handleError(err)

	res, err := tx.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", modifyProduct.Name, modifyProduct.Price, id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	if res == nil {

	}

	UpdatedProduct.ID = id
	UpdatedProduct.Name = modifyProduct.Name
	UpdatedProduct.Price = modifyProduct.Price
	handleError(tx.Commit())

	return UpdatedProduct
}

func (p *ProductRepository) DeleteProduct(id int) {
	tx, err := p.DB.Begin()
	handleError(err)

	res, err := tx.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	if res == nil {

	}
	handleError(tx.Commit())
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
