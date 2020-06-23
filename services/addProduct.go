package services

import (
	"viper/models"
)

func (ps *InventoryService) AddNewProduct(productID, productDesc, categoryID, productPrice, discountID, brandDesc string) (*models.Product, error) {
	product := &models.Product{
		ProductID:    productID,
		ProductDesc:  productDesc,
		CategoryID:   categoryID,
		ProductPrice: productPrice,
		DiscountID:   discountID,
		BrandDesc:    brandDesc,
	}
	id, err := models.AddNewProduct(ps.db, product)
	if err != nil {
		return nil, err
	}
	product.ProductID = id
	return product, nil
}
