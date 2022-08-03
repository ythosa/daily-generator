package service

import "daily-generator/internal/models"

type JiraService interface {
	GetJiraSummaryByIssueID(id string) (string, error)
}

type FormatService interface {
	FormatDailyMessage(data *models.DailyMessage) string
}
