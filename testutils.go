package goscm

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// This starts a Testify server which responds with a JSON if and only if the requested URL matches the expected one.
// It is advised to end each Server instance with a server.Close() statement after finishing a unit test.
//
// If SendsMethodResponse is set to true, the server will respond with a generic template by the pattern of OK-[Method]
// if and only if both the expected URL and the method as a value match.
func setupTestServer(urlToJson map[string]string, sendsMethodResponse bool, t *testing.T) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if sendsMethodResponse {
			urlIncluded, _ := checkMethodResponse(urlToJson, request, writer)
			if !urlIncluded {
				t.Logf("Received url " + request.RequestURI + " with method " + request.Method + " doesn't match any predefined url and method for this server.")
				writer.WriteHeader(404)
			}
		} else {
			urlIncluded, _ := checkJsonResponse(urlToJson, request, writer)
			if !urlIncluded {
				t.Logf("Received url " + request.RequestURI + " doesn't match any predefined url for this server.")
				writer.WriteHeader(404)
			}
		}

	}))
	return server
}

// Checks whether both URL and method are registered and sends back a signal text.
func checkMethodResponse(urlToMethod map[string]string, request *http.Request, writer http.ResponseWriter) (bool, error) {
	for expectedURL, method := range urlToMethod {
		if request.RequestURI == expectedURL && request.Method == method {
			writer.Header().Set("Content-Type", "text/plain")
			_, err := writer.Write([]byte("OK-" + method))
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

// Checks whether the URL is registered and sends back the predefined JSON file.
func checkJsonResponse(urlToJson map[string]string, request *http.Request, writer http.ResponseWriter) (bool, error) {
	for expectedUrl, jsonFile := range urlToJson {
		if request.RequestURI == expectedUrl {
			writer.Header().Set("Content-Type", "application/json")
			content, _ := os.ReadFile(jsonFile)
			_, err := writer.Write(content)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

// Cf. setupTestServer.
func setupSingleTestServer(jsonFile string, expectedUrl string, t *testing.T) *httptest.Server {
	return setupTestServer(map[string]string{expectedUrl: jsonFile}, false, t)
}
