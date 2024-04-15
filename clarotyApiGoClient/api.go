// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

func NewClarotyAPI(config *APIConfiguration) ClarotyAPI {

	if config.MaxConcurrentRequests == nil {
		maxConcurrentRequests := 20
		config.MaxConcurrentRequests = &maxConcurrentRequests
	}

	semaphore := make(chan struct{}, *config.MaxConcurrentRequests)

	return ClarotyAPI{
		configuration: *config,
		semaphore:     &semaphore,
	}
}
