package goslacksender

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

var slackHookUrl = os.Getenv("SLACK_HOOK_URL")
var slackChannel = os.Getenv("SLACK_CHANNEL")

type sendMessageReqBody struct {
	Username string `json:"username"`
	Text     string `json:"text"`
	Channel  string `json:"channel"`
}

// SendMessage via Telegram
func SendMessage(message string, username string) error {
	if slackHookUrl == "" {
		return errors.New("SLACK_HOOK_URL is not set")
	}
	if slackChannel == "" {
		return errors.New("SLACK_CHANNEL is not set")
	}

	// Creates an instance of our custom sendMessageReqBody Type
	reqBody := &sendMessageReqBody{
		Text:     message,
		Username: username,
		Channel:  slackChannel,
	}

	// Convert our custom type into json format
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Make a request to send our message using the POST method to the telegram bot API
	resp, err := http.Post(
		slackHookUrl,
		"application/json",
		bytes.NewBuffer(reqBytes),
	)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		errClose := Body.Close()
		if errClose != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected reponse status" + resp.Status)
	}

	return err
}
