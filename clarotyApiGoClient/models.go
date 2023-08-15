// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import (
	"net/http"
)

type APIConfiguration struct {
	URI                   string       `json:"uri"`
	Username              *string      `json:"username"`
	Password              *string      `json:"password"`
	Apikey                *string      `json:"Apikey"`
	Client                *http.Client `json:"-"`
	IgnoreSSL             *bool        `json:"ignoreSSL"`
	MaxConcurrentRequests *int         `json:"maxConcurrentRequests"`
	MaxRetry              *int         `json:"maxRetry"`
}

type APICredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ClarotyAPI struct {
	configuration  APIConfiguration
	MaxRetry       *int
	authentication APIAuthentication
	semaphore      *chan struct{}
}

type APIAuthentication interface {
	getToken() string
	isPasswordExpired() bool
	isAPIKey() bool
}

type APIAuthenticationResponse struct {
	ID              int    `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Token           string `json:"token"`
	PasswordExpired bool   `json:"password_expired"`
	Mail            string `json:"mail"`
}

func (a *APIAuthenticationResponse) getToken() string {
	return a.Token
}

func (a *APIAuthenticationResponse) isPasswordExpired() bool {
	return a.PasswordExpired
}

func (a *APIAuthenticationResponse) isAPIKey() bool {
	return false
}

type APIAuthenticationKey struct {
	Apikey string `json:"apikey"`
}

func (a *APIAuthenticationKey) getToken() string {
	return a.Apikey
}

func (a *APIAuthenticationKey) isPasswordExpired() bool {
	return false
}

func (a *APIAuthenticationKey) isAPIKey() bool {
	return true
}
