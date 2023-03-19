// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import "net/http"

type APIConfiguration struct {
	URI                   string       `json:"uri"`
	Username              string       `json:"username"`
	Password              string       `json:"password"`
	Client                *http.Client `json:"-"`
	IgnoreSSL             *bool        `json:"ignoreSSL"`
	MaxConcurrentRequests *int         `json:"maxConcurrentRequests"`
}

type APICredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ClarotyAPI struct {
	configuration  APIConfiguration
	authentication *APIAuthenticationResponse
	semaphore      *chan struct{}
}

type APIAuthenticationResponse struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Token           string `json:"token"`
	PasswordExpired bool   `json:"password_expired"`
	Mail            string `json:"mail"`
}
