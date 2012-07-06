package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Doer interface {
	Do(request *http.Request) (*http.Response, error)
}

type Client struct {
	HttpClient *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{HttpClient: client}
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	if token, err := ReadToken(); err == nil {
		request.Header.Set("Authorization", token)
	}
	response, err := c.HttpClient.Do(request)
	if err != nil {
		msg := fmt.Sprintln("A problem occurred while trying to do the request.", "Original error message is:", err.Error(), "")
		return nil, errors.New(msg)
	}
	if response.StatusCode > 399 {
		defer response.Body.Close()
		result, _ := ioutil.ReadAll(response.Body)
		return nil, errors.New(string(result))
	}
	return response, nil
}
