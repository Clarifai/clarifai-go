# Clarifai Golang Library

Unofficial library written for the [Clarifai](http://www.clarifai.com) API. It's still a work in progress.

## Usage
```go
package main

import clarifai "github.com/samuelcouch/clarifai-go"

func main() {
	// Create a new clarifai client
  client := clarifai.NewClient("<client_id>", "<client_secret>")

  // Request a new token
  token, err := client.requestAccessToken()
  if err != nil {
    fmt.Println("whoops")
  }

  // Print the token from the request
  fmt.Printf("%v\n", token.AccessToken)
  // Show that the token is now saved to the client
  fmt.Printf("%v\n", client.accessToken)
}
```
