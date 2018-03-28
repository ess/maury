package users

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// Reader provides an interface for the finder functions to talk to the API
type Reader interface {
	Get(string, url.Values) ([]byte, error)
}

// All returns an array of user entities from the API. If params are
// provided, they are passed along to the API for consideration.
func All(driver Reader, params url.Values) []*Entity {
	return allPages(driver, "users", params)
}

// Find queries the API for a single account entity by account ID. If there
// are problems along the way, a non-nil error is returned. Otherwise, the
// error is nil and the entity is populated.
func Find(driver Reader, id string) (*Entity, error) {
	response, err := driver.Get("users/"+id, nil)
	if err != nil {
		return nil, err
	}

	wrapper := struct {
		User *Entity `json:"user,omitempty"`
	}{}

	err = json.Unmarshal(response, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.User, nil
}

// Current queries the API for the user entity that initiated the request.
// If there are problems along the way, a non-nil error is returned. Otherwise,
// the error is nil and the entity is populated.
func Current(driver Reader) (*Entity, error) {
  response, err := driver.Get("users/current", nil)
  if err != nil {
    return nil, err
  }

  wrapper := struct {
    User *Entity `json:"user,omitempty"`
  }{}

  err = json.Unmarshal(response, &wrapper)
  if err != nil {
    return nil, err
  }

  return wrapper.User, nil
}

func allPages(driver Reader, path string, params url.Values) []*Entity {
	var users []*Entity

	maxResults := 100
	page := 1

	if params == nil {
		params = url.Values{}
	}

	params.Set("per_page", strconv.Itoa(maxResults))

	wrapper := struct {
		Users []*Entity `json:"users,omitempty"`
	}{
		Users: []*Entity{{}},
	}

	for len(wrapper.Users) > 0 {
		params.Set("page", strconv.Itoa(page))

		if response, err := driver.Get(path, params); err == nil {
			if jerr := json.Unmarshal(response, &wrapper); jerr == nil {
				users = append(users, wrapper.Users...)

				if len(wrapper.Users) < maxResults {
					break
				}
			}
		}

		page = page + 1
	}

	return users
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
