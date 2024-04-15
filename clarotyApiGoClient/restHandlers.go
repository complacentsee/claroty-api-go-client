// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func (api *ClarotyAPI) restGet(path string) ([]byte, error) {
	baseURL := api.configuration.URI
	client := api.configuration.Client
	ignoreSSL := api.configuration.IgnoreSSL
	url := baseURL + path

	token := api.configuration.Apikey

	if ignoreSSL != nil && *ignoreSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", *token))

	var resp *http.Response
	initialBackoff := 500 * time.Millisecond
	maxBackoff := 16 * time.Second
	backoff := initialBackoff
	maxRetry := 3
	if api.MaxRetry != nil {
		maxRetry = *api.MaxRetry
	}

	for i := 0; i <= maxRetry; i++ {
		resp, err = client.Do(req)
		if err == nil && !shouldRetry(resp.StatusCode) {
			break
		}

		if i < maxRetry {
			sleepDuration := time.Duration(float64(backoff) * (1 + rand.Float64()))
			if sleepDuration > maxBackoff {
				sleepDuration = maxBackoff
			}

			time.Sleep(sleepDuration)
			backoff *= 2
		}
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("restGet: Non-OK HTTP status: " + resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (api *ClarotyAPI) restPost(path string, body []byte) ([]byte, error) {
	baseURL := api.configuration.URI
	client := api.configuration.Client
	ignoreSSL := api.configuration.IgnoreSSL
	url := baseURL + path

	// If ignoreSSL is true, create a custom HTTP client that ignores SSL certificate errors
	if ignoreSSL != nil && *ignoreSSL {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	token := api.configuration.Apikey

	req.Header.Set("Authorization", fmt.Sprintf("Bearer: %s", *token))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("restPost: Non-OK HTTP status: " + resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func shouldRetry(statusCode int) bool {
	return statusCode == http.StatusTooManyRequests || statusCode == http.StatusRequestTimeout ||
		statusCode == http.StatusServiceUnavailable || statusCode == http.StatusGatewayTimeout
}
