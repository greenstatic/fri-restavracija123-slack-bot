package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/palantir/stacktrace"
	"net/http"
)

type Config struct {
	WebhookUrl string
}

type MessagePayload struct {
	Mrkdwn bool   `json:"mrkdwn"`
	Text   string `json:"text"`
}

func (c Config) SendMessage(msg string) error {
	p := MessagePayload{true, msg}

	body, err := json.Marshal(p)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to marshal payload into JSON")
	}

	b := bytes.NewBuffer(body)
	resp, err := http.Post(c.WebhookUrl, "application/json", b)
	if err != nil {
		return stacktrace.Propagate(err, "Failed to send HTTP POST request")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("unsuccessful status code")
	}

	return nil
}
