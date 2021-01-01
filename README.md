# OneTimeSecret (OTS) - Go

This serves as a Golang API client for [onetimesecret](https://onetimesecret.com/).

## Installation

As this is a Go package, it is easiest to install via `go get`

As such, you can retrieve this package by running the command

    go get -u "github.com/jdockerty/onetimesecret-go/ots"

## Usage

Prior to using the package, you must have an account with OTS as this gives you a username, which is your email, to use and an API token. Various functions are exposed which match to the naming conventions of the OTS API.

A simple use case for checking the status of the OTS API is shown below
```go
package main

import (
    "fmt"
	"github.com/jdockerty/onetimesecret-go/ots"
)

func main() {

	var c ots.Client
    client := c.New("YOUR_EMAIL", "API_TOKEN")
    
    err := client.Status() // Get the status of the OTS API
    if err != nil {
        fmt.Println(err) // If the API is offline, an error is returned.
    }
}
```

The other exported functions can return a `Secret` type, which is struct that contains the expected responses from the API, such as a list of recipients for the secret or it's time-to-live value.