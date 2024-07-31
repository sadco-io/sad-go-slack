# Slack Service Package

This package provides functionality to send messages to Slack using webhooks. It's designed to be easily integrated into Go applications that require Slack notifications.

## Features

- Send text messages to Slack
- Support for message attachments
- Automatic inclusion of service name and hostname in messages
- Environment-based configuration

## Installation

To use this package in your Go project, you can import it as follows:

```go
import "github.com/sadco-io/sad-go-slack/slack"
```

## Configuration

The Slack service uses environment variables for configuration:

- `SLACK_WEBHOOK_URL`: The webhook URL for your Slack workspace (required)
- `SERVICE_NAME`: The name of your service (optional, defaults to "Melina")

If `SLACK_WEBHOOK_URL` is not set, the Slack service will be disabled.

## Usage

### Initialization

The package automatically initializes a global `SlackService` instance. You can use this instance to send messages:

```go
if slack.GlobalSlackService != nil {
    err := slack.GlobalSlackService.PostMessage("Hello, Slack!", nil)
    if err != nil {
        // Handle error
    }
}
```

### Sending Messages

To send a simple text message:

```go
err := slack.GlobalSlackService.PostMessage("Your message here", nil)
```

To send a message with attachments:

```go
attachments := []slack.SlackAttachment{
    {
        Text:  "Attachment text",
        Color: "#36a64f", // Optional color
    },
}
err := slack.GlobalSlackService.PostMessage("Your message here", attachments)
```

### Message Format

Messages sent through this service will automatically include the service name and hostname in the following format:

```
ServiceName_Hostname_YourMessage
```

## Error Handling

The `PostMessage` function returns an error if the message couldn't be sent. Make sure to handle these errors appropriately in your application.

## Logging

This package uses the `github.com/sadco-io/sad-go-logger/logger` for logging. Ensure this dependency is properly set up in your project.

## Dependencies

- `github.com/sadco-io/sad-go-logger/logger`
- `go.uber.org/zap`

## Contributing

Contributions to improve this package are welcome. Please submit issues and pull requests on the GitHub repository.

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This project is licensed under the MIT License - see the [LICENSE] file for details.

Copyright (c) 2024 SAD co.
