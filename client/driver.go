// Package client provides a low-level driver for interacting with the Engine
// Yard API
package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Driver is an object that knows specifically how to interact with the
// Engine Yard API at the HTTP level
type Driver struct {
	raw     *http.Client
	baseURL url.URL
	token   string
}

// New takes a base URL for an Engine Yard API and a token, returning a Driver
// that can be used to interact with the API in question.
func New(baseURL string, token string) (*Driver, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	d := &Driver{
		&http.Client{Timeout: 20 * time.Second},
		*url,
		token,
	}

	return d, nil
}

// Get performs a GET operation for the given path and params against the
// upstream API. it returns a byte array and an error.
func (driver *Driver) Get(path string, params url.Values) ([]byte, error) {
	return driver.makeRequest("GET", path, params, nil)
}

// Post performs a POST operation for the given path, params, and data against
// the upstream API. it returns a byte array and an error.
func (driver *Driver) Post(path string, params url.Values, data []byte) ([]byte, error) {
	return driver.makeRequest("POST", path, params, data)
}

// Put performs a PUT operation for the given path, params, and data against
// the upstream API. It returns a byte array and an error.
func (driver *Driver) Put(path string, params url.Values, data []byte) ([]byte, error) {
	return driver.makeRequest("PUT", path, params, data)
}

// Patch performs a PATCH operation for the given path, params, and data against
// the upstream API. it returns a byte array and an error.
func (driver *Driver) Patch(path string, params url.Values, data []byte) ([]byte, error) {
	return driver.makeRequest("PATCH", path, params, data)
}

// Delete performs a DELETE operation for the given path and params against the
// upstream API. it returns a byte array and an error.
func (driver *Driver) Delete(path string, params url.Values) ([]byte, error) {
	return driver.makeRequest("DELETE", path, params, nil)
}

func (driver *Driver) makeRequest(verb string, path string, params url.Values, data []byte) ([]byte, error) {

	request, err := http.NewRequest(
		verb,
		driver.constructRequestURL(path, params),
		bytes.NewReader(data),
	)

	if err != nil {
		return nil, err
	}

	request.Header.Add("X-EY-TOKEN", driver.token)
	request.Header.Add("Accept", "application/vnd.engineyard.v3+json")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("User-Agent", "maury-go/0.1.0 (https://github.com/ess/maury)")

	response, err := driver.raw.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		return nil, fmt.Errorf(
			"The upstream API returned the following status: %d",
			response.StatusCode,
		)
	}

	return body, nil
}

func (driver *Driver) constructRequestURL(path string, params url.Values) string {

	pathParts := []string{driver.baseURL.Path, path}

	requestURL := url.URL{
		Scheme:   driver.baseURL.Scheme,
		Host:     driver.baseURL.Host,
		Path:     strings.Join(pathParts, "/"),
		RawQuery: params.Encode(),
	}

	return requestURL.String()
}

// Copyright 2018 Dennis Walters
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
