package clarotyApiGoClient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
)

func (api *ClarotyAPI) restGet(path string) ([]byte, error) {
	baseURL := api.configuration.URI
	client := api.configuration.Client
	ignoreSSL := api.configuration.IgnoreSSL
	url := baseURL + path

	token, err := api.GetAuthenticationToken()
	if err != nil {
		return nil, err
	}

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
	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("restGet: Non-OK HTTP status: " + resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
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

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if api.authentication == nil {
		req.Header.Set("Authorization", "Authorization")
	} else {
		token := api.authentication.Token
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("restPost: Non-OK HTTP status: " + resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
