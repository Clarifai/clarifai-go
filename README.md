# Clarifai Golang Library

Library for our [API](https://developer.clarifai.com/docs/).

[![GoDoc](https://godoc.org/github.com/Clarifai/clarifai-go?status.svg)](https://godoc.org/github.com/Clarifai/clarifai-go)

## Usage
`go get github.com/clarifai/clarifai-go`


```go
package main

import (
	"fmt"

	"github.com/clarifai/clarifai-go"
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
	tag_data, err := client.Tag(clarifai.TagRequest{URLs: urls})

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

##Special Thanks

Thanks so much to [Sam Couch](https://github.com/samuelcouch) for his amazing contribution to the Clarifai client libraries by starting this one and graciously handing it off to us. Follow him on Twitter [@SamuelCouch](http://twitter.com/SamuelCouch).
