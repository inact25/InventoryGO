package services

import (
	"database/sql"
	"viper/models"
)

type ProductService struct {
	db *sql.DB
}


func ViewAllProductService(db *sql.DB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) GetDataProduct(pageNo, totalPerPage int) []*models.Product {
	products, err := models.GetDataProduct(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return products
}

func Call(){

}