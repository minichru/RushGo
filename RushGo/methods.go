package rushgo

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func Get(url string, headers ...map[string]string) (*http.Response, error) {
	if len(headers) > 0 {
		return SendRequest("GET", url, nil, headers[0])
	}
	return SendRequest("GET", url, nil, nil)
}

func Post(url string, body []byte, headers ...map[string]string) (*http.Response, error) {

	var mergedHeaders map[string]string
	if len(headers) > 0 {
		mergedHeaders = headers[0]
	}
	return SendRequest("POST", url, body, mergedHeaders)
}

func Patch(url string, body []byte, headers ...map[string]string) (*http.Response, error) {
	var mergedHeaders map[string]string
	if len(headers) > 0 {
		mergedHeaders = make(map[string]string)
		for _, header := range headers {
			for key, value := range header {
				mergedHeaders[key] = value
			}
		}
	}
	return SendRequest("PATCH", url, body, mergedHeaders)
}

func Put(url string, body []byte, headers ...map[string]string) (*http.Response, error) {
	var mergedHeaders map[string]string
	if len(headers) > 0 {
		mergedHeaders = make(map[string]string)
		for _, header := range headers {
			for key, value := range header {
				mergedHeaders[key] = value
			}
		}
	}

	return SendRequest("PUT", url, body, mergedHeaders)
}

func Delete(url string, headers ...map[string]string) (*http.Response, error) {
	var mergedHeaders map[string]string
	if len(headers) > 0 {
		mergedHeaders = make(map[string]string)
		for _, header := range headers {
			for key, value := range header {
				mergedHeaders[key] = value
			}
		}
	}
	return SendRequest("DELETE", url, nil, mergedHeaders)
}

var client = &http.Client{
	Timeout: 30 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		Proxy:                 http.ProxyFromEnvironment,
		TLSHandshakeTimeout:   10 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		ExpectContinueTimeout: 1 * time.Second,
		IdleConnTimeout:       60 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		DisableCompression:    true,
		DisableKeepAlives:     false,
		ForceAttemptHTTP2:     true,
	},
}

func GetCookieByName(resp *http.Response, name string) *http.Cookie {
	cookies := resp.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie
		}
	}

	return nil
}

func SendRequest(method, url string, body []byte, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	respChan := make(chan *http.Response)
	errChan := make(chan error)

	go func() {
		resp, err := client.Do(req)
		if err != nil {
			errChan <- err
			return
		}

		respChan <- resp
	}()

	// Wait for either the response or error to be received on the channels
	select {
	case resp := <-respChan:
		return resp, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(30 * time.Second):
		return nil, fmt.Errorf("Request timed out")
	}
}
