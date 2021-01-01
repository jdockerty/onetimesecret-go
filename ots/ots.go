package ots

import (
	// "bytes"
	// "encoding/json"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	base = "https://onetimesecret.com/api/v1"
)

// Client is used to set the user's 'Username' and 'Token' for interaction with the OneTimeSecret API.
type Client struct {
	Username string
	Token    string
	hc http.Client
}

type Secret struct {
	CustomerID string `json:"custid"`
	MetadataKey string `json:"metadata_key"`
	SecretKey string `json:"secret_key"`
	TTL int `json:"ttl"`
	MetadataTTL int `json:"metadata_ttl"`
	SecretTTL int `json:"secret_ttl"`
	Recipient []string `json:"recipient`
	State string `json:"state"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
	PassphraseRequired bool `json:"passphrase_required"`
}

// New returns a populated client to OneTimeSecret, this uses your provided username (email) and token (API token in your account)
// in order to authenticate to the API server with OTS.
func (c *Client) New(user, token string) *Client {
	return &Client{Username: user, Token: token}
}

// Status will check the current status of the OTS system.
// This returns an error if the servers are not online or there are other problems with the request.
func (c *Client) Status() error {
	endpoint := createURI("status")

	req, err := http.NewRequest("GET", endpoint, strings.NewReader(""))
	if err != nil {
		return err
	}
	log.Println(c.Token, c.Username)
	req.SetBasicAuth(c.Username, c.Token)

	resp, err := c.hc.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if s := string(body); s == "offline" {
		return errors.New("server is offline, try again later")
	}

	return nil
}

// Create will POST a secret to be stored within OTS, this is shared with the individual you specify via email.
// Secret is the value that you wish to store.
// Passphrase is the string with which the recipient is allowed to view the secret.
// Recipient is who you wish to send the secret to, using their email address.
// TTL is the time-to-live of the secret, in seconds. Once this expires, the secret is deleted.
// POST https://onetimesecret.com/api/v1/share
func (c *Client) Create(secret, passphrase, recipient string, ttl int) error {
	
	endpoint := createURI("share")

	v := url.Values{}
	v.Set("secret", secret)
	v.Set("passphrase", passphrase)
	v.Set("ttl", strconv.Itoa(ttl))
	v.Set("recipient", recipient)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.Username, c.Token)
	
	resp, err := c.hc.Do(req)
	if err != nil {
		return err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var otsResponse Secret

	err = json.Unmarshal(bodyText, &otsResponse)
	if err != nil {
		return err
	}

	// log.Printf("API RESP: %+v\n", otsResponse)

	return nil
}

// Generate will return a short is useful for temporary passwords, one-time pads, salts etc.
func (c *Client) Generate() {

}


func createURI(s string) string {
	URI := fmt.Sprintf("%s/%s", base, s)
	return URI
}
