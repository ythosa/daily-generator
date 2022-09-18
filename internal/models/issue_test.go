package models

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestJiraIssueID_GetURL(t *testing.T) {
	t.Parallel()

	jiraURL := "https://jira.ru"
	jiraIssueID := "353"

	assert.Equal(t, JiraIssueID(jiraIssueID).GetURL(jiraURL), "https://jira.ru/browse/353")
}
