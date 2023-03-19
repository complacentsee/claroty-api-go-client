// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/complacentsee/clarotyApiGoClient"
)

type Configuration struct {
	ClarotyAPI clarotyApiGoClient.APIConfiguration `json:"clarotyAPI"`
}

func loadConfig(filename string) (*Configuration, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Configuration
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	httpClient := http.Client{}
	config.ClarotyAPI.Client = &httpClient

	clarotyClient := clarotyApiGoClient.NewClarotyAPI(&config.ClarotyAPI)

	siteRequest := clarotyApiGoClient.APISiteRequest{}

	siteChan, err := clarotyClient.StreamSites(&siteRequest)
	if err != nil {
		fmt.Println(err)
	}

	var siteCount int
	var assetTotal int
	var wg sync.WaitGroup

	for site := range siteChan {
		siteCount++
		fmt.Println("Site Name:", site.Name)

		wg.Add(1)

		go func(site clarotyApiGoClient.Site) {
			defer func() {
				wg.Done()
			}()

			assetRequest := clarotyApiGoClient.APIAssetRequest{}
			assetRequest.SetFormatAssetList()
			assetRequest.SetPage(1)
			assetRequest.SetPerPage(50)
			assetRequest.SetSiteIdExact(site.ID)
			assetRequest.SetVendorIContains("Rockwell")

			assetChan, err := clarotyClient.StreamAssets(&assetRequest)
			if err != nil {
				fmt.Println(err)
			}

			var assetCount int
			for range assetChan {
				assetTotal++
				assetCount++
				//fmt.Println("Asset Name:", asset.Name)
			}

			fmt.Printf("Assets for Site %s: %d\n", site.Name, assetCount)
		}(site)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("Site count:", siteCount)
	fmt.Println("Asset count:", assetTotal)
}
