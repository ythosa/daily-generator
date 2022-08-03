package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"

	"daily-generator/internal/controller/ui/scanner"
	"daily-generator/internal/models"
	"daily-generator/pkg/collections"
)

var _ scanner.Scanner = (*terminalScanner)(nil)

type terminalScanner struct {
	reader *bufio.Reader
}

func NewTerminalScanner() *terminalScanner {
	return &terminalScanner{reader: bufio.NewReader(os.Stdin)}
}

func (t *terminalScanner) Scan() (*models.DailyData, error) {
	var result = new(models.DailyData)

	fmt.Printf("Input yesterday issues: ")
	yesterdayIssues, err := t.scanIssues()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan yesterday issues")
	}
	result.Yesterday = yesterdayIssues

	fmt.Printf("Input today issues: ")
	todayIssues, err := t.scanIssues()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan today issues")
	}
	result.Today = todayIssues

	fmt.Printf("Input problems: ")
	problems, err := t.scanProblems()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan problems")
	}
	result.Problems = problems

	return result, nil
}

func (t *terminalScanner) scanIssues() ([]string, error) {
	line, err := t.reader.ReadString('\n')
	if err != nil {
		return nil, errors.Wrap(err, "failed to read string")
	}

	return collections.MapSlice(strings.Split(line, ","), func(id string) string {
		return strings.TrimSpace(id)
	}), nil
}

func (t *terminalScanner) scanProblems() (string, error) {
	line, err := t.reader.ReadString('\n')
	if err != nil {
		return "", errors.Wrap(err, "failed to read string")
	}

	return line, nil
}
