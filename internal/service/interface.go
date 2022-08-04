package service

import "daily-generator/internal/models"

type JiraService interface {
	GetJiraDailyMessage(data *models.DailyData) (*models.DailyMessage, error)
}

type FormatService interface {
	FormatDailyMessage(data *models.DailyMessage) string
}
