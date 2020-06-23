package services

import "database/sql"

type InventoryService struct {
	db *sql.DB
}

func NewService(db *sql.DB) *InventoryService {
	return &InventoryService{db}
}
