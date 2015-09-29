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

  // Print the full response from the token request
  // &{AccessToken:<token> ExpiresIn:<time> Scope:<scope of token> TokenType:Bearer}
  fmt.Printf("%+v\n", token)
  // Show that the token is now saved to the client
  fmt.Printf("%v\n", client.accessToken)
}
```
