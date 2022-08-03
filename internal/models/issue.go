package models

import "fmt"

type JiraIssueID string

func (j JiraIssueID) GetURL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, j)
}

type Issue struct {
	ID      JiraIssueID
	Summary string
}

func NewIssue(ID string, summary string) *Issue {
	return &Issue{ID: JiraIssueID(ID), Summary: summary}
}
