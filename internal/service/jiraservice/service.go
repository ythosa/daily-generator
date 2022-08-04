package jiraservice

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"daily-generator/internal/config"
	"daily-generator/internal/jirawrapper"
	"daily-generator/internal/models"
	"daily-generator/internal/service"
)

var _ service.JiraService = (*jiraService)(nil)

type jiraService struct {
	logger *logrus.Logger

	cfg        *config.Config
	jiraClient *jira.Client
}

func NewJiraService(cfg *config.Config, jiraClient *jira.Client) *jiraService {
	return &jiraService{
		logger: logrus.WithField("service", "jira-service").Logger,

		cfg:        cfg,
		jiraClient: jiraClient,
	}
}

func (j *jiraService) GetJiraDailyMessage(data *models.DailyData) (*models.DailyMessage, error) {
	yesterdayIssues, err := j.jiraIDsToIssues(data.Yesterday)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map yesterday issues")
	}

	todayIssues, err := j.jiraIDsToIssues(data.Today)
	if err != nil {
		return nil, errors.Wrap(err, "failed to map today issues")
	}

	return models.NewDailyMessage(yesterdayIssues, todayIssues, data.Problems), nil
}

func (j *jiraService) jiraIDsToIssues(ids []string) ([]*models.Issue, error) {
	var result []*models.Issue

	for _, id := range ids {
		fullJiraID := j.fullJiraID(id)
		summary, err := j.getJiraSummaryByIssueID(fullJiraID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get jira summary")
		}

		result = append(result, models.NewIssue(fullJiraID, summary))
	}

	return result, nil
}

func (j *jiraService) getJiraSummaryByIssueID(id string) (string, error) {
	issue, resp, err := j.jiraClient.Issue.Get(id, nil)
	if err != nil {
		jiraError, err := jirawrapper.FormatJiraErrorResponse(resp, err)
		if err != nil {
			return "", errors.Wrap(err, "failed to format jira error response")
		}

		j.logger.Errorf("failed to get issue by id: %s", jiraError)

		return "", errors.Wrap(err, "failed to get jira issue by id")
	}

	return issue.Fields.Summary, nil
}

func (j *jiraService) fullJiraID(id string) string {
	return fmt.Sprintf("%s-%s", j.cfg.JiraProject, id)
}
