package services

import (
	"viper/models"
)

func (ps *InventoryService) GetAllReport(pageNo, totalPerPage int) []*models.Report {
	report, err := models.GetAllReport(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return report
}
