package ots

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	base = "https://onetimesecret.com/api/v1"
)

// Client is used to set the user's 'Username' and 'Token' for interaction with the OneTimeSecret API.
type Client struct {
	Username string
	Token    string
}

func createURI(s string) string {
	URI := fmt.Sprintf("%s/%s", base, s)
	return URI
}

// New returns a populated client to OneTimeSecret, this uses your provided username (email) and token (API token in your account)
// in order to authenticate to the API server with OTS.
func (c *Client) New(user, token string) *Client {
	return &Client{Username: user, Token: token}
}

// Status will check the current status of the OTS system.
// This returns either 'nominal' or 'offline'
func (c *Client) Status() error {
	endpoint := createURI("status")
	m := http.Client{}
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		panic(err)
	}
	log.Println("AUTH:", c.Username, c.Token)

	req.SetBasicAuth(c.Username, c.Token)
	resp, err := m.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	s := string(body)
	log.Println(s)

	if s == "offline" {
		return errors.New("server is unavailable, please try again later")
	}

	return nil
}
