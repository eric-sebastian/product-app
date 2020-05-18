package http

import (
	"product-app/src/internal/entity/product"
	"product-app/src/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ProductUsecase usecase.ProductUsecase
}

func ProvideProductHandler(p usecase.ProductUsecase) ProductHandler {
	return ProductHandler{ProductUsecase: p}
}

func (p *ProductHandler) GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products := p.ProductUsecase.GetAll()
		c.JSON(200, gin.H{
			"status":  200,
			"message": "",
			"data":    products,
		})
	}
}

func (p *ProductHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		product := p.ProductUsecase.GetById(id)
		c.JSON(200, gin.H{
			"status":  200,
			"message": "",
			"data":    product,
		})
	}
}

func (p *ProductHandler) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productForm product.ProductForm
		if c.BindJSON(&productForm) == nil {
			newProduct := p.ProductUsecase.CreateProduct(productForm)
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Product created!",
				"data":    newProduct,
			})
		}
	}
}

func (p *ProductHandler) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var productForm product.ProductForm
		if c.BindJSON(&productForm) == nil {
			updatedProduct := p.ProductUsecase.UpdateProduct(id, productForm)
			c.JSON(200, gin.H{
				"status":  200,
				"message": "Product updated!",
				"data":    updatedProduct,
			})
		}
	}
}

func (p *ProductHandler) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		p.ProductUsecase.DeleteProdct(id)
		c.JSON(200, gin.H{
			"status":  200,
			"message": "Product deleted!",
			"data":    "",
		})
	}
}
