package jiraservice

import (
	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"daily-generator/internal/jirawrapper"
	"daily-generator/internal/service"
)

var _ service.JiraService = (*jiraService)(nil)

type jiraService struct {
	logger     *logrus.Logger
	jiraClient *jira.Client
}

func NewJiraService(jiraClient *jira.Client) *jiraService {
	return &jiraService{
		logger:     logrus.WithField("service", "jira-service").Logger,
		jiraClient: jiraClient,
	}
}

func (j *jiraService) GetJiraSummaryByIssueID(id string) (string, error) {
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
