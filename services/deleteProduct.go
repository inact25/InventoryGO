package services

import (
	"viper/models"
)

func (ps *InventoryService) DeleteDataProduct(productID string) (*models.Product, error) {
	product := &models.Product{
		ProductID: productID,
	}
	delproID, err := models.DeleteProduct(ps.db, product)
	if err != nil {
		return nil, err
	}
	product.ProductID = delproID
	return product, nil
}
