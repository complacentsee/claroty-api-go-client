// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

func (request *APISiteRequest) Setrid_exact(rid_exact string) {
	request.Rid__exact = &rid_exact
}

func (request *APISiteRequest) SetSortName() {
	var sort = "name"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortNameDescending() {
	var sort = "-name"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortTrainingMode() {
	var sort = "training_mode"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortTrainingModeDescending() {
	var sort = "-training_mode"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortVersion() {
	var sort = "version"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortVersionDescending() {
	var sort = "-version"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortID() {
	var sort = "id"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortIDDescending() {
	var sort = "-id"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortLocationLatitude() {
	var sort = "location_latitude"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortLocationLatitudeDescending() {
	var sort = "-location_latitude"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortLocationLongitude() {
	var sort = "location_longitude"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortLocationLongitudeDescending() {
	var sort = "-location_longitude"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortDescription() {
	var sort = "description"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortDescriptionDescending() {
	var sort = "-description"
	request.Sort = &sort
}

func (request *APISiteRequest) SetSortScore() {
	var sort = "score"
	request.Sort = &sort
}

func (request *APISiteRequest) SetPage(page int) {
	var reqPage = page
	request.Page = &reqPage
}

func (request *APISiteRequest) SetPerPage(per_page int) {
	var reqPerPage = per_page
	request.Per_page = &reqPerPage
}

func (request *APISiteRequest) GetQueryParameters() string {
	values := url.Values{}

	if request.Page == nil {
		request.Page = new(int)
		*request.Page = 1
	}

	if request.Per_page == nil {
		request.Per_page = new(int)
		*request.Per_page = 20
	}

	value := reflect.ValueOf(request).Elem()
	st := reflect.TypeOf(request).Elem()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldName := st.Field(i).Tag.Get("json")

		if fieldValue.IsNil() {
			continue
		}

		switch fieldValue.Interface().(type) {
		case *string:
			values.Set(fieldName, *fieldValue.Interface().(*string))
		case *int:
			values.Set(fieldName, strconv.Itoa(*fieldValue.Interface().(*int)))
		case *bool:
			values.Set(fieldName, strconv.FormatBool(*fieldValue.Interface().(*bool)))
		default:
			// handle other types if needed
		}
	}

	return values.Encode()

}

func (request *APISiteRequest) GetResponseCount() (int, error) {
	if request.Page != nil && request.Per_page != nil {
		return *request.Page * *request.Per_page, nil
	}
	err := fmt.Errorf("page or per_page is nil")
	return 0, err
}
