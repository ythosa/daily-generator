package ui

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"daily-generator/internal/config"
	"daily-generator/internal/controller"
	"daily-generator/internal/controller/ui/scanner"
	"daily-generator/internal/models"
	"daily-generator/internal/service"
)

var _ controller.Controller = (*controllerImpl)(nil)

type controllerImpl struct {
	cfg           *config.Config
	scanner       scanner.Scanner
	jiraService   service.JiraService
	formatService service.FormatService
}

func NewController(
	cfg *config.Config,
	scanner scanner.Scanner,
	jiraService service.JiraService,
	formatService service.FormatService,
) *controllerImpl {
	return &controllerImpl{
		cfg:           cfg,
		scanner:       scanner,
		jiraService:   jiraService,
		formatService: formatService,
	}
}

func (c *controllerImpl) Start() {
	fmt.Println("Daily Generator!!!")

	dailyData, err := c.scanner.Scan()
	if err != nil {
		logrus.Fatal("failed to scan data", dailyData)
	}

	dailyMessage, err := c.mapDailyDataToDailyMessage(dailyData)
	if err != nil {
		logrus.Fatal("failed to map daily data to message")
	}

	message := c.formatService.FormatDailyMessage(dailyMessage)
	fmt.Printf("Generation done:\n%s\n ", message)
}

func (c *controllerImpl) mapDailyDataToDailyMessage(data *models.DailyData) (*models.DailyMessage, error) {
	yesterdayIssues, err := c.jiraIDsToIssues(data.Yesterday)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map yesterday issues")
	}

	todayIssues, err := c.jiraIDsToIssues(data.Today)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map today issues")
	}

	return models.NewDailyMessage(yesterdayIssues, todayIssues, data.Problems), nil
}

func (c *controllerImpl) jiraIDsToIssues(ids []string) ([]*models.Issue, error) {
	var result []*models.Issue

	for _, id := range ids {
		fullJiraID := c.fullJiraID(id)
		summary, err := c.jiraService.GetJiraSummaryByIssueID(fullJiraID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get jira summary")
		}

		result = append(result, models.NewIssue(fullJiraID, summary))
	}

	return result, nil
}

func (c *controllerImpl) fullJiraID(id string) string {
	return fmt.Sprintf("%s-%s", c.cfg.JiraProject, id)
}
