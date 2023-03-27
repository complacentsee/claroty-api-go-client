// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import (
	"encoding/json"
	"log"
	"math"
	"sync"
)

func (api *ClarotyAPI) GetSites(query *APISiteRequest) ([]Site, error) {
	var sites []Site

	siteChan, err := api.StreamSites(query)
	if err != nil {
		return nil, err
	}

	for site := range siteChan {
		sites = append(sites, site)
	}

	return sites, nil
}

func (api *ClarotyAPI) StreamSites(query *APISiteRequest) (<-chan Site, error) {
	siteChan := make(chan Site)

	go func() {
		defer close(siteChan)
		if err := api.fetchSites(query, siteChan); err != nil {
			log.Printf("Error fetching sites: %v", err)
		}
	}()

	return siteChan, nil
}

func (api *ClarotyAPI) fetchSites(query *APISiteRequest, siteChan chan<- Site) error {
	uri := "/ranger/sites?" + query.GetQueryParameters()

	// Acquire a semaphore slot
	*api.semaphore <- struct{}{}

	response, err := api.restGet(uri)
	if err != nil {
		return err
	}

	// Release the semaphore slot when the rest call has returned
	<-*api.semaphore

	var siteResponse APISiteResponse
	err = json.Unmarshal(response, &siteResponse)
	if err != nil {
		return err
	}

	for _, site := range siteResponse.Sites {
		siteChan <- site
	}

	responseCount, err := query.GetResponseCount()
	if err != nil {
		return err
	}

	if *query.Page > 1 {
		return nil
	}

	if responseCount < siteResponse.CountTotal {
		pages := int(math.Ceil(float64(siteResponse.CountTotal) / float64(responseCount)))
		wg := &sync.WaitGroup{}

		for i := 2; i <= pages; i++ {
			wg.Add(1)

			go func(page int) {
				defer wg.Done()
				newQuery := *query
				newQuery.Page = &page
				err := api.fetchSites(&newQuery, siteChan)
				if err != nil {
					log.Printf("Error fetching sites for page %d: %v", page, err)
				}
			}(i)
		}

		wg.Wait()
	}

	return nil
}
