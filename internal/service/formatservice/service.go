package formatservice

import (
	"fmt"
	"strings"

	"daily-generator/internal/models"
	"daily-generator/internal/service"
	"daily-generator/pkg/collections"
)

var _ service.FormatService = (*formatService)(nil)

type formatService struct {
	jiraBaseURL string
}

func NewFormatService(jiraBaseURL string) *formatService {
	return &formatService{jiraBaseURL: jiraBaseURL}
}

func (f *formatService) FormatDailyMessage(data *models.DailyMessage) string {
	var messageBuilder strings.Builder

	messageBuilder.WriteString("**Что вы делали вчера?**\n")
	messageBuilder.WriteString(f.formatIssues(data.Yesterday))
	messageBuilder.WriteString("\n")

	messageBuilder.WriteString("**Что вы будете делать сегодня?**\n")
	messageBuilder.WriteString(f.formatIssues(data.Today))
	messageBuilder.WriteString("\n")

	messageBuilder.WriteString("**Отлично, есть ли какие-то препятствия?**\n")
	messageBuilder.WriteString(data.Problems)

	return messageBuilder.String()
}

func (f *formatService) formatIssues(issues []*models.Issue) string {
	return strings.Join(
		collections.MapSlice(issues, func(issue *models.Issue) string {
			return f.formatIssue(issue)
		}),
		"\n",
	)
}

func (f *formatService) formatIssue(issue *models.Issue) string {
	return fmt.Sprintf("* [%s](%s) - %s", issue.ID, issue.ID.GetURL(f.jiraBaseURL), issue.Summary)
}
