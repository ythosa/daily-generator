package scanner

import "daily-generator/internal/models"

type Scanner interface {
	Scan() (*models.DailyData, error)
}
