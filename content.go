package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

// GetContent Get file content for revision and path
func (c *Client) GetContent(namespace string, name string, revision string, path string) (string, *http.Response, error) {
	r, err := c.httpClient.Get(c.baseUrl + "/api/v2/repositories/" + namespace + "/" + name + "/content/" + url.PathEscape(revision) + "/" + url.PathEscape(path))
	if err != nil {
		return "", nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(r.Body)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(b), r, nil
}

// GetContentForDefaultBranch Get file content for latest revision on default branch and path
func (c *Client) GetContentForDefaultBranch(namespace string, name string, path string) (string, *http.Response, error) {
	defaultBranch, err := c.GetDefaultBranch(namespace, name)
	if err != nil {
		return "", nil, err
	}
	content, r, err := c.GetContent(namespace, name, defaultBranch.Revision, path)
	if err != nil {
		return "", r, err
	}
	return content, r, nil
}
