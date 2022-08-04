package ui

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/sirupsen/logrus"

	"daily-generator/internal/config"
	"daily-generator/internal/controller"
	"daily-generator/internal/controller/ui/scanner"
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
	fmt.Println("// Welcome to Daily Generator! ")

	dailyData, err := c.scanner.Scan()
	if err != nil {
		logrus.Fatal("failed to scan data", dailyData)
	}

	dailyMessage, err := c.jiraService.GetJiraDailyMessage(dailyData)
	if err != nil {
		logrus.Fatal("failed to map daily data to message")
	}

	fmt.Println("// Generation in progress ...")
	message := c.formatService.FormatDailyMessage(dailyMessage)
	if err := clipboard.WriteAll(message); err != nil {
		logrus.Warn(fmt.Sprintf("failed to copy to clipboard: %s", err))
		fmt.Println("// Generation done!")
	} else {
		fmt.Println("// Generation done! (copied to clipboard)")
	}
	fmt.Printf("%s\n", message)
}
