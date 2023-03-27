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

func (api *ClarotyAPI) GetAssets(query *APIAssetRequest) ([]Asset, error) {
	var assets []Asset

	assetChan, err := api.StreamAssets(query)
	if err != nil {
		return nil, err
	}

	for asset := range assetChan {
		assets = append(assets, asset)
	}

	return assets, nil
}

func (api *ClarotyAPI) StreamAssets(query *APIAssetRequest) (<-chan Asset, error) {
	assetChan := make(chan Asset)

	go func() {
		defer close(assetChan)
		if err := api.fetchAssets(query, assetChan); err != nil {
			log.Printf("Error fetching assets: %v", err)
		}
	}()

	return assetChan, nil
}

func (api *ClarotyAPI) fetchAssets(query *APIAssetRequest, assetChan chan<- Asset) error {
	uri := "/ranger/assets?" + query.GetQueryParameters()

	// Acquire a semaphore slot
	*api.semaphore <- struct{}{}

	response, err := api.restGet(uri)
	if err != nil {
		return err
	}

	// Release the semaphore slot when the rest call has returned
	<-*api.semaphore

	var assetResponse APIAssetResponse
	err = json.Unmarshal(response, &assetResponse)
	if err != nil {
		return err
	}

	for _, asset := range assetResponse.Assets {
		assetChan <- asset
	}

	responseCount, err := query.GetResponseCount()
	if err != nil {
		return err
	}

	if *query.Page > 1 {
		return nil
	}

	if responseCount < assetResponse.CountTotal {
		pages := int(math.Ceil(float64(assetResponse.CountTotal) / float64(responseCount)))
		wg := &sync.WaitGroup{}

		for i := 2; i <= pages; i++ {
			wg.Add(1)
			go func(page int) {
				defer wg.Done()
				newQuery := *query
				newQuery.Page = &page
				err := api.fetchAssets(&newQuery, assetChan)
				if err != nil {
					log.Printf("Error fetching assets for page %d: %v", page, err)
				}
			}(i)
		}

		wg.Wait()
	}

	return nil
}
