# FRI Restavracija123 Slack Bot
A simple Slack bot that posts the daily FRI Restavracija123.si cafeteria menu.

Data is extracted from: [www.restavracija123.si/](https://www.restavracija123.si/).

## Environment variables
* `SLACK_WEBHOOK` (required) - Slack Webhook URL
* `MESSAGE_TRIGGER` (required) - HH:MM in UTC (e.g. "10:00")
* `DEBUG` - true/false

## Dev
Do not forget to change version in `main.go`

To build docker image
```bash
make docker
```