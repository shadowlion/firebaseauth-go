package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	ApiKey     string
}

func New(apiKey string) *Client {
	return &Client{
		HttpClient: &http.Client{},
		ApiKey:     apiKey,
	}
}

type ErrorItem struct {
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type ErrorPayload struct {
	Errors  []ErrorItem `json:"errors"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
}

type ErrorResponse struct {
	Error ErrorPayload `json:"error"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HttpClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes ErrorResponse

		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return errors.New(errRes.Error.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return err
}
