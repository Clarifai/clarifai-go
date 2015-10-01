# Clarifai Golang Library

Unofficial library written for the [Clarifai](http://www.clarifai.com) API. It's still a work in progress.

## Usage
```go
package main

import "fmt"

func main() {
  client := NewClient("<client_id>", "<client_secret>")

  info, _ := ClarifaiInfo(client)

  fmt.Printf("%+v\n", info)
  // &{StatusCode:OK StatusMessage:All images in request have completed successfully.  Results:{MaxImageSize:100000 DefaultLanguage:en MaxVideoSize:100000 MaxImageBytes:10485760 DefaultModel:default MaxVideoBytes:104857600 maxVideoDuration:0 MaxVideoBatchSize:1 MinVideoSize:1 MinImageSize:1 MaxBatchSize:128 APIVersion:0.1}}
}

```

## Testing
Run `go test`
