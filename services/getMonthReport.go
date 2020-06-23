package services

import (
	"viper/models"
)

func (ps *InventoryService) GetMonthReport(date string, pageNo, totalPerPage int) []*models.Report {
	report, err := models.GetMonthReport(ps.db, date, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return report
}
