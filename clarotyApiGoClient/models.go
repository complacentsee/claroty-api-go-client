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
	Apikey                *string      `json:"Apikey"`
	Client                *http.Client `json:"-"`
	IgnoreSSL             *bool        `json:"ignoreSSL"`
	MaxConcurrentRequests *int         `json:"maxConcurrentRequests"`
	MaxRetry              *int         `json:"maxRetry"`
}

type ClarotyAPI struct {
	configuration APIConfiguration
	MaxRetry      *int
	semaphore     *chan struct{}
}
