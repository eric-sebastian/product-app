package usecase

import (
	"log"
	"product-app/src/internal/entity/product"
	"product-app/src/internal/repo/products"
	"strconv"
)

type ProductUsecase struct {
	ProductRepository products.ProductRepository
}

func ProvideProductUsecase(p products.ProductRepository) ProductUsecase {
	return ProductUsecase{ProductRepository: p}
}

func (p *ProductUsecase) GetAll() []product.Product {
	return p.ProductRepository.GetAll()
}

func (p *ProductUsecase) GetById(id string) product.Product {
	convId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	return p.ProductRepository.GetById(convId)
}

func (p *ProductUsecase) CreateProduct(productForm product.ProductForm) product.Product {
	newProduct := p.ProductRepository.CreateProduct(productForm)

	return newProduct
}

func (p *ProductUsecase) UpdateProduct(id string, productForm product.ProductForm) product.Product {
	convId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	newProduct := p.ProductRepository.UpdateProduct(convId, productForm)

	return newProduct
}

func (p *ProductUsecase) DeleteProdct(id string) {
	convId, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	p.ProductRepository.DeleteProduct(convId)
}
