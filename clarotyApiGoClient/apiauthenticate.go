// Copyright (c) 2023, Adam Traeger
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package clarotyApiGoClient

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (api *ClarotyAPI) Authenticate() error {
	creds := APICredentials{
		Username: api.configuration.Username,
		Password: api.configuration.Password,
	}
	postBody, err := json.Marshal(creds)
	if err != nil {
		fmt.Println("Error marshaling credentials:", err)
		return fmt.Errorf("Error marshaling credentials: %v", err)
	}

	response, err := api.restPost("/auth/authenticate", postBody)

	var authResponse APIAuthenticationResponse

	err = json.Unmarshal(response, &authResponse)
	if err != nil {
		fmt.Println("Error unmarshaling authentication response:", err)
		return fmt.Errorf("Error unmarshaling authentication response: %v", err)
	}

	if authResponse.PasswordExpired {
		fmt.Println("Account password has expired")
		return fmt.Errorf("Account password has expired")
	}

	api.authentication = &authResponse
	return nil
}

func (api *ClarotyAPI) GetAuthenticationToken() (string, error) {
	if api.authentication == nil {
		err := api.Authenticate()
		if err != nil {
			fmt.Println("Error authenticateing:", err)
			return "", err
		}
	}

	expired, err := isTokenExpired(api.authentication.Token)
	if err != nil {
		fmt.Println("Error checking token expiration:", err)
		return "", err
	}

	if expired {
		err := api.Authenticate()
		if err != nil {
			fmt.Println("Error authenticateing:", err)
			return "", err
		}
	}

	return api.authentication.Token, nil
}

func isTokenExpired(tokenString string) (bool, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return false, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("error parsing token claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, fmt.Errorf("token missing 'exp' claim")
	}

	expirationTime := time.Unix(int64(exp), 0)
	return expirationTime.Before(time.Now()), nil
}
