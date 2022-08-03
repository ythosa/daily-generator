package jirawrapper

import (
	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"

	"daily-generator/internal/config"
)

func GetJiraClient(config *config.Config) (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: config.JiraCredentials.Username,
		Password: config.JiraCredentials.Password,
	}

	client, err := jira.NewClient(tp.Client(), string(config.JiraURL))
	if err != nil {
		return nil, errors.Wrap(err, "failed to init jira client")
	}

	return client, nil
}
