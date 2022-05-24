# go-slack-sender
Send message to your Slack channel

## How to use it

```shell
export SLACK_HOOK_URL=https://hooks.slack.com/services...
export SLACK_CHANNEL="#channel"
```

and in your go code:

```go
package main

import (
    sc "github.com/scor2k/go-slack-sender"
)

func main() {
	if err := sc.SendMessage("Hello, world!", "yourBotName"); err != nil {
        panic(err)
    }
}
```
