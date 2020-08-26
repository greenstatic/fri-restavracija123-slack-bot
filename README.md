# Slack Food Bot
A simple Slack bot that posts the daily FRI & BF cafeteria menu.

Data is scraped from: 
* [www.restavracija123.si](https://www.restavracija123.si)
* [www.studentska-prehrana.si](https://www.studentska-prehrana.si/sl/restaurant/Details/2023)

Docker Image: [greenstatic/slack-food-bot](https://hub.docker.com/r/greenstatic/slack-food-bot)

## Environment variables
* `SLACK_WEBHOOK` (required) - Slack Webhook URL
* `MESSAGE_TRIGGER` (required) - HH:MM in UTC (e.g. "10:30")
* `DEBUG` - true/false

## Dev
Do not forget to change version in `main.go`

To build docker image
```bash
make docker
```

### Insider Knowledge
The FRI cafeteria finalizes the menu by 10:30 (local time).
