package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Client struct {
	httpClient *http.Client
	baseUrl    string
	token      string
	userAgent  string
}

func NewClient(baseUrl string, token string) (*Client, error) {
	c := Client{
		httpClient: http.DefaultClient,
		token:      token,
		userAgent:  "go-scm/0.1 (+https://github.com/cloudogu/go-scm)",
		baseUrl:    baseUrl,
	}
	return &c, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) Put(url string, body []byte, headers map[string]string) error {
	return c.handleRequest(http.MethodPut, url, body, nil, headers)
}

func (c *Client) Delete(url string, headers map[string]string) error {
	return c.handleRequest(http.MethodDelete, url, nil, nil, headers)
}

func (c *Client) GetJson(url string, respModel interface{}, headers map[string]string) error {
	return c.handleRequest(http.MethodGet, url, nil, &respModel, headers)
}

func (c *Client) handleRequest(method, url string, body []byte, respModel interface{}, headers map[string]string) error {
	request := &http.Request{}
	var err error
	if method == "GET" {
		request, err = http.NewRequest(method, c.baseUrl+url, nil)
		if err != nil {
			return err
		}
	} else {
		request, err = http.NewRequest(method, c.baseUrl+url, bytes.NewBuffer(body))
		if err != nil {
			return err
		}
	}

	request.Header.Set("User-Agent", c.userAgent)
	//request.Header.Set("Accept", "application/vnd.scmm-me+json;v=2")
	request.Header.Set("Authorization", "Bearer "+c.token)

	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}

	request.Close = true
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(response.Body)

	// handle response error body
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
		errModel := ErrorResponse{}
		// if we can not handle the error response body, create error with status code only
		if json.NewDecoder(response.Body).Decode(&errModel) != nil {
			return errors.New("http status " + strconv.Itoa(response.StatusCode))
		}
		return errors.New("http status " + strconv.Itoa(response.StatusCode) + " - " + errModel.String())
	}

	if respModel == nil {
		return nil
	}

	// read response body
	err = json.NewDecoder(response.Body).Decode(&respModel)
	if err != nil {
		return errors.New("http status " + strconv.Itoa(response.StatusCode) + " - " + err.Error())
	}

	return nil
}

type ErrorResponse struct {
	TransactionId string `json:"transactionId,omitempty"`
	ErrorCode     string `json:"errorCode,omitempty"`
	Context       []struct {
		Type string `json:"type,omitempty"`
		Id   string `json:"id,omitempty"`
	} `json:"context,omitempty"`
	Message            string `json:"message,omitempty"`
	AdditionalMessages []struct {
		Key     string `json:"key,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"additionalMessages,omitempty"`
	Violations []struct {
		Path    string `json:"path,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"violations,omitempty"`
	Url string `json:"url,omitempty"`
}

func (e *ErrorResponse) String() string {
	return fmt.Sprintf("%+q", *e)
}
