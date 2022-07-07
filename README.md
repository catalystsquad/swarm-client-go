# swarm-client-go

A Go client library for https://swarmiolabs.com/


## Getting Started

Add the swarm client to your project with:
```sh
go get github.com/catalystsquad/swarm-client-go
```

Import the project and initialize a client by supplying your customer ID and an
API token to the client constructor:
```go
import (
	swarm "github.com/catalystsquad/swarm-client-go"
)

func main() {
	client := swarm.NewClient("MYCUSTOMERID", "MYAPITOKEN")
}
```

## Usage

Services exist for all API paths. Each service will have methods for REST API
actions for listing, creating, deleting, etc.

### Publishing Messages

The publish service has one Publish method which can be used to send messages
to Pipelines. It will marshal structs into JSON for you. Simply pass a pipeline
ID and your data as parameters:
``` go
type SomeData struct {
	SomeField string `json:"someField"`
}

myPipeline := "02J6XFSVVAT8FNRGXP4BX0A1PM"

ctx := context.TODO()
_, err := client.Publish.Publish(ctx, myPipeline, &SomeData{SomeField: "myData"})
if err != nil {
	return fmt.ErrorF("error publishing: %s", err)
}
```
