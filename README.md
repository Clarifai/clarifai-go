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
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", info)
	}
	// Let's get some context about these images
	urls := []string{"http://www.clarifai.com/img/metro-north.jpg", "http://www.clarifai.com/img/metro-north.jpg"}
	// Give it to Clarifai to run their magic
	tag_data, err := client.Tag(urls, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", tag_data) // See what we got!
	}

	feedback_data, err := client.Feedback(clarifai.FeedbackForm{
		URLs:    urls,
		AddTags: []string{"cat", "animal"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", feedback_data)
	}
}
```

## Testing
Run `go test`
