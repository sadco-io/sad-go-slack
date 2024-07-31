// Package slack provides functionality to send messages to Slack using webhooks.
package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sadco-io/sad-go-logger/logger"

	"go.uber.org/zap"
)

// GlobalSlackService is the global instance of SlackService.
var GlobalSlackService *SlackService

var hostname string
var serviceName string

// SlackService represents a service for sending messages to Slack.
type SlackService struct {
	webhookURL string
}

// SlackPayload defines the structure of the payload to be sent to Slack.
type SlackPayload struct {
	Text        string            `json:"text"`
	Attachments []SlackAttachment `json:"attachments"`
}

// SlackAttachment defines the structure of a Slack message attachment.
type SlackAttachment struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

// init initializes the global SlackService instance and sets up necessary variables.
func init() {
	var err error

	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		logger.Log.Info("SLACK_WEBHOOK_URL is not set, SlackService will be disabled")
		GlobalSlackService = nil
	} else {
		GlobalSlackService = &SlackService{webhookURL: webhookURL}
	}

	hostname, err = os.Hostname()
	if err != nil {
		logger.Log.Warn("Error retrieving hostname", zap.Error(err))
		logger.Log.Warn("Setting hostname to unkw")
		hostname = "unkw"
	}

	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		logger.Log.Info("SERVICE_NAME is not set, using Melina as default")
		serviceName = "Melina"
	}
}

// PostMessage sends a message to the configured Slack webhook.
// It takes a text message and optional attachments as parameters.
// Returns an error if the message couldn't be sent.
func (s *SlackService) PostMessage(text string, attachments []SlackAttachment) error {
	newText := serviceName + "_" + hostname + "_" + text
	slackBody, _ := json.Marshal(SlackPayload{
		Text:        newText,
		Attachments: attachments,
	})

	req, err := http.NewRequest(http.MethodPost, s.webhookURL, bytes.NewBuffer(slackBody))
	if err != nil {
		logger.Log.Error("Error creating Slack request", zap.Error(err))
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		logger.Log.Error("Error sending Slack payload", zap.Error(err))
		return err
	}

	return nil
}
