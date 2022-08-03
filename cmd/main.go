package main

import (
	"github.com/sirupsen/logrus"

	"daily-generator/internal/config"
	"daily-generator/internal/controller/ui"
	"daily-generator/internal/controller/ui/scanner/terminal"
	"daily-generator/internal/jirawrapper"
	"daily-generator/internal/service/formatservice"
	"daily-generator/internal/service/jiraservice"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		logrus.Fatal("failed to get config: ", err)
	}

	jiraClient, err := jirawrapper.GetJiraClient(cfg)
	if err != nil {
		logrus.Fatal("failed to get jira client: ", err)
	}

	formatService := formatservice.NewFormatService(cfg.JiraURL)
	jiraService := jiraservice.NewJiraService(jiraClient)

	scanner := terminal.NewTerminalScanner()

	ui.NewController(cfg, scanner, jiraService, formatService).Start()
}
