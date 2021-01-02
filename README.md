# OneTimeSecret (OTS) - Go Client Library

This serves as a Golang API client for [onetimesecret](https://onetimesecret.com/).

## Installation

As this is a Go package, it is easiest to install via `go get`

As such, you can retrieve this package by running the command

    go get -u "github.com/jdockerty/onetimesecret-go/ots"

## Usage

Prior to using the package, you must have an account with OTS as this gives you a username, which is your email, to use and an API token. Various functions are exposed which match to the naming conventions of the OTS API.

A simple use case for checking the status of the OTS API and creating a secret is shown below
```go
package main

import (
    "log"
    "github.com/jdockerty/onetimesecret-go/ots"
)

func main() {

    var c ots.Client
    client := c.New("YOUR_EMAIL", "API_TOKEN")
    
    err := client.Status() // Get the status of the OTS API
    if err != nil {
        return // If the API is offline, an error is returned.
    }

    // Send a secret to your friend's email address, it is destroyed within 60 seconds. 
    // They must enter the passphrase to view it.
    secretResp, err := client.Create(
	"my super secret value", 
	"very secret passphrase", 
	"bestfriend@gmail.com",
	60)
    if err != nil {
	// Handling errors from sending requests and dealing with JSON go here
	log.Fatal(err) 
	}

    // Log the Secret struct response with various information to stdout in an easy to read format.
    secretResp.PrettyPrint()

}
```


The other exported functions can return a `Secret` or `Secrets` type, which is struct that contains the expected responses from the API, such as a list of recipients for the secret or it's time-to-live value. Which fields are used is left to the user and further details on the various functions are available via the [godoc](https://godoc.org/github.com/jdockerty/onetimesecret-go/ots) page.
