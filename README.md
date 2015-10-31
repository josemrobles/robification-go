# robification-go [![Build Status](https://travis-ci.org/josemrobles/robification-go.svg?branch=master)](https://travis-ci.org/josemrobles/robification-go)
A golang library which uses the RobiFication API to broadcast messages to Flowdock, Slack, HipChat, SMS.

Usage:
```go
package main

import (
    "fmt"
    "github.com/josemrobles/robification-go"
)

func main() {
  yourMessage := robification.NewFdChat("YOUR_CHAT_API_TOKEN", "YOUR_MESSAGE")
  err := robification.Send(yourMessage)
  if err != nil {
    fmt.Println(err)
  }
}
```

Todo:
- Concurrent messages and message types to multiple providers.
- Better error handling
- Tests
