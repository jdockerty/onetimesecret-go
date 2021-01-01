package ots

import (
	"encoding/json"
	"errors"
	"fmt"
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
	hc       http.Client
}

// Secret is a struct which contains the expected fields from the /share API endpoint.
type Secret struct {
	CustomerID         string   `json:"custid,omitempty"`
	MetadataKey        string   `json:"metadata_key,omitempty"`
	SecretKey          string   `json:"secret_key,omitempty"`
	Value              string   `json:"value,omitempty"`
	TTL                int      `json:"ttl,omitempty"`
	MetadataTTL        int      `json:"metadata_ttl,omitempty"`
	SecretTTL          int      `json:"secret_ttl,omitempty"`
	Recipient          []string `json:"recipient,omitempty`
	State              string   `json:"state,omitempty"`
	Created            int64    `json:"created,omitempty"`
	Updated            int64    `json:"updated,omitempty"`
	PassphraseRequired bool     `json:"passphrase_required,omitempty"`
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
// This request is sent via POST https://onetimesecret.com/api/v1/share
func (c *Client) Create(secret, passphrase, recipient string, ttl int) (*Secret, error) {

	endpoint := createURI("share")

	v := url.Values{}
	v.Set("secret", secret)
	v.Set("passphrase", passphrase)
	v.Set("ttl", strconv.Itoa(ttl))
	v.Set("recipient", recipient)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var otsResponse *Secret

	err = json.Unmarshal(bodyText, &otsResponse)
	if err != nil {
		return nil, err
	}

	return otsResponse, nil
}

// Generate will return a short, unique secret which is useful for temporary passwords, one-time pads, salts etc.
// The response value is the same format as Create(), but the Value field is populated.
// This request is sent via POST https://onetimesecret.com/api/v1/generate
func (c *Client) Generate(passphrase, recipient string, ttl int) (*Secret, error) {

	endpoint := createURI("generate")

	v := url.Values{}
	v.Set("passphrase", passphrase)
	v.Set("ttl", strconv.Itoa(ttl))
	v.Set("recipient", recipient)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var otsResponse *Secret

	err = json.Unmarshal(bodyText, &otsResponse)
	if err != nil {
		return nil, err
	}

	return otsResponse, nil
}

// Retrieve is used to get the value of a secret which was previously stored. Once you retrieve the secret, it is no longer available.
// The secretKey parameter is gained from the response when initially creating a secret that is to be shared and the passphrase is what was
// specified upon creation of the said secret.
func (c *Client) Retrieve(secretKey, passphrase string) (*Secret, error) {

	route := fmt.Sprintf("secret/%s", secretKey)
	endpoint := createURI(route)

	v := url.Values{}

	v.Set("secret_key", secretKey)
	v.Set("passphrase", passphrase)

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var otsResponse *Secret

	err = json.Unmarshal(bodyText, &otsResponse)
	if err != nil {
		return nil, err
	}

	log.Println(otsResponse)

	return otsResponse, nil

}


func (c *Client) RetrieveMetadata(metadataKey string) (*Secret, error) {

	route := fmt.Sprintf("private/%s", metadataKey)

	endpoint := createURI(route)

	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Token)

	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var otsResponse *Secret

	err = json.Unmarshal(bodyText, &otsResponse)
	if err != nil {
		return nil, err
	}

	log.Println(otsResponse)

	return otsResponse, nil
}

func createURI(s string) string {
	URI := fmt.Sprintf("%s/%s", base, s)
	return URI
}
