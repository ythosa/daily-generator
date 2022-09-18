package jiraservice

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	cmap "github.com/orcaman/concurrent-map/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

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

	issueSummaryByKeyCache cmap.ConcurrentMap[string]
}

func NewJiraService(cfg *config.Config, jiraClient *jira.Client) *jiraService {
	return &jiraService{
		logger: logrus.WithField("service", "jira-service").Logger,

		cfg:        cfg,
		jiraClient: jiraClient,

		issueSummaryByKeyCache: cmap.New[string](),
	}
}

func (j *jiraService) GetJiraDailyMessage(data *models.DailyData) (*models.DailyMessage, error) {
	allIssues, err := j.jiraIDsToIssues(append(data.Yesterday, data.Today...))
	if err != nil {
		return nil, errors.Wrap(err, "failed to ")
	}

	yesterdayIssues := allIssues[:len(data.Yesterday)]
	todayIssues := allIssues[len(data.Yesterday):]

	return models.NewDailyMessage(yesterdayIssues, todayIssues, data.Problems), nil
}

func (j *jiraService) jiraIDsToIssues(ids []string) ([]*models.Issue, error) {
	var (
		result   = make([]*models.Issue, len(ids))
		errGroup errgroup.Group
	)

	for i, id := range ids {
		position := i
		jiraID := id

		errGroup.Go(func() error {
			fullJiraID := j.fullJiraID(jiraID)

			summary, err := j.getJiraSummaryByIssueID(fullJiraID)
			if err != nil {
				return errors.Wrap(err, "failed to get jira summary")
			}

			result[position] = models.NewIssue(fullJiraID, summary)

			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}

func (j *jiraService) getJiraSummaryByIssueID(id string) (string, error) {
	if summary, ok := j.issueSummaryByKeyCache.Get(id); ok {
		return summary, nil
	}

	j.logger.Infof("getting issue by id=%s", id)

	issue, resp, err := j.jiraClient.Issue.Get(id, nil)
	if err != nil {
		jiraError, err := jirawrapper.FormatJiraErrorResponse(resp, err)
		if err != nil {
			return "", errors.Wrap(err, "failed to format jira error response")
		}

		j.logger.Errorf("failed to get issue by id=%s: %s", id, jiraError)

		return "", errors.Wrapf(err, "failed to get jira issue by id=%s", id)
	}

	j.issueSummaryByKeyCache.Set(id, issue.Fields.Summary)

	return issue.Fields.Summary, nil
}

func (j *jiraService) fullJiraID(id string) string {
	return fmt.Sprintf("%s-%s", j.cfg.JiraProject, id)
}
