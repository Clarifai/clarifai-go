# Clarifai Golang Library

Unofficial library written for the [Clarifai](http://www.clarifai.com) API. It's still a work in progress.

## Usage
```go
package main

import (
  "fmt"

  "github.com/samuelcouch/clarifai"
)

func main() {
  client := clarifai.NewClient("<client_id>", "<client_secret>")
  // Get the current status of things
  info, err := clarifai.Info(client)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%+v\n", info)
  }
  // Let's get some context about these images
  urls := []string{"http://www.clarifai.com/img/metro-north.jpg", "http://www.clarifai.com/img/metro-north.jpg"}
  // Give it to Clarifai to run their magic
  tag_data, err := clarifai.Tag(client, urls, nil)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("%+v\n", tag_data) // See what we got!
  }
}

```

## Testing
Run `go test`
