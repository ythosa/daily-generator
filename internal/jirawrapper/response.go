package jirawrapper

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

func FormatJiraErrorResponse(response *jira.Response, receivedError error) (string, error) {
	if receivedError == nil {
		return "", nil
	}

	jiraResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrap(receivedError, "failed to read jira response")
	}

	var result strings.Builder

	result.WriteString(fmt.Sprintf("message: %s\n", receivedError))
	result.WriteString(fmt.Sprintf("response: %s", jiraResponse))

	return result.String(), nil
}
