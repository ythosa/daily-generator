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

	t.printMessage("🍉 Input yesterday issues")
	yesterdayIssues, err := t.scanIssues()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan yesterday issues")
	}
	result.Yesterday = yesterdayIssues

	t.printMessage("🍒 Input today issues")
	todayIssues, err := t.scanIssues()
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan today issues")
	}
	result.Today = todayIssues

	t.printMessage("🍑 Input problems")
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

	if len(strings.TrimSpace(line)) == 0 {
		return nil, nil
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

const messagePrefix = ">>"

func (t *terminalScanner) printMessage(message string) {
	fmt.Printf("%s %s: ", messagePrefix, message)
}
