package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	configsFolderPathEnv     = "DG_CONFIGS_FOLDER_PATH"
	defaultConfigsFolderPath = "./configs"

	configNameEnv     = "DG_CONFIG_NAME"
	defaultConfigName = "config"
)

type Config struct {
	JiraURL     string
	JiraProject string
	*JiraCredentials
}

func GetConfig() (*Config, error) {
	if err := initConfigParser(); err != nil {
		return nil, errors.Wrap(err, "failed to init config parser")
	}

	jiraURL, err := getJiraURL()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get jira url")
	}

	jiraCredentials, err := getJiraCredentials()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get jira credentials")
	}

	jiraProject, err := getJiraProject()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get jira project")
	}

	return &Config{
		JiraURL:         jiraURL,
		JiraProject:     jiraProject,
		JiraCredentials: jiraCredentials,
	}, nil
}

func getJiraURL() (string, error) {
	jiraURL := viper.GetString("jira.url")
	if _, err := url.Parse(jiraURL); err != nil {
		return "", errors.Wrap(err, "invalid jira url")
	}

	return jiraURL, nil
}

func getJiraProject() (string, error) {
	jiraURL := viper.GetString("jira.project")
	if _, err := url.Parse(jiraURL); err != nil {
		return "", errors.Wrap(err, "invalid jira project")
	}

	return jiraURL, nil
}

type JiraCredentials struct {
	Username string
	Password string
}

func getJiraCredentials() (*JiraCredentials, error) {
	username := viper.GetString("jira.username")
	if username == "" {
		return nil, errors.New("empty jira username")
	}

	password := viper.GetString("jira.password")
	if password == "" {
		return nil, errors.New("empty jira password")
	}

	return &JiraCredentials{
		Username: username,
		Password: password,
	}, nil
}

func initConfigParser() error {
	configsFolderPath := os.Getenv(configsFolderPathEnv)
	if configsFolderPath == "" {
		logrus.Warn(fmt.Sprintf("empty configs folder path environment variable: %s, using default: %s",
			configsFolderPathEnv, defaultConfigsFolderPath))
		configsFolderPath = defaultConfigsFolderPath
	}

	configName := os.Getenv(configNameEnv)
	if configName == "" {
		logrus.Warn(fmt.Sprintf("empty config name environment variable: %s, using default: %s",
			configNameEnv, defaultConfigName))
		configName = defaultConfigName
	}

	viper.AddConfigPath(configsFolderPath)
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while reading config: %w", err)
	}

	return nil
}
