package accounts

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"

	"github.com/ess/maury/users"
)

// Reader provides an interface for the finder functions to talk to the API
type Reader interface {
	Get(string, url.Values) ([]byte, error)
}

// All returns an array of account entities from the API. If params are
// provided, they are passed along to the API for consideration.
func All(driver Reader, params url.Values) []*Entity {
	return allPages(driver, "accounts", params)
}

// ForUser returns an array of account entities from the API scoped to the
// given user. If params are provided, they are passed along to the API for
// consideration.
func ForUser(driver Reader, user *users.Entity, params url.Values) []*Entity {
	pathParts := []string{"users", user.ID, "accounts"}

	return allPages(driver, strings.Join(pathParts, "/"), params)
}

// Find queries the API for a single account entity by account ID. If there
// are problems along the way, a non-nil error is returned. Otherwise, the
// error is nil and the entity is populated.
func Find(driver Reader, id string) (*Entity, error) {
	response, err := driver.Get("accounts/"+id, nil)
	if err != nil {
		return nil, err
	}

	wrapper := struct {
		Account *Entity `json:"account,omitempty"`
	}{}

	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Account, nil
}

func allPages(driver Reader, path string, params url.Values) []*Entity {
	var accounts []*Entity

	maxResults := 100
	page := 1

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", strconv.Itoa(maxResults))

	wrapper := struct {
		Accounts []*Entity `json:"accounts,omitempty"`
	}{
		Accounts: []*Entity{{}},
	}

	for len(wrapper.Accounts) > 0 {
		params.Set("page", strconv.Itoa(page))

		if response, err := driver.Get(path, params); err == nil {
			if jerr := json.Unmarshal(response, &wrapper); jerr == nil {
				accounts = append(accounts, wrapper.Accounts...)

				if len(wrapper.Accounts) < maxResults {
					break
				}
			}
		}

		page = page + 1
	}

	return accounts
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
