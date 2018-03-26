package accounts

import (
	"encoding/json"
	"net/url"
)

// Updater provides an interface for the update functions to talk to the API
type Updater interface {
	Put(string, url.Values, []byte) ([]byte, error)
}

// Changes models the aspects of an Account that we are allowed to change
type Changes struct {
	Name             string `json:"name,omitempty"`
	EmergencyContact string `json:"emergency_contact,omitempty"`
	SupportPlan      string `json:"support_plan,omitempty"`
}

// Update requests that an account be updated on the API to match the provided
// changes. If there are issues along the way, a non-nil error is returned.
// Otherwise, the error is nil and the returned entity contains the requested
// changes.
func Update(driver Updater, account *Entity, changes *Changes) (*Entity, error) {

	wrappedChanges := struct {
		Account *Changes `json:"account,omitempty"`
	}{
		Account: changes,
	}

	data, err := json.Marshal(&wrappedChanges)
	if err != nil {
		return nil, err
	}

	response, err := driver.Put("accounts/"+account.ID, nil, data)
	if err != nil {
		return nil, err
	}

	wrapped := struct {
		Account *Entity `json:"account,omitempty"`
	}{}

	err = json.Unmarshal(response, &wrapped)
	if err != nil {
		return nil, err
	}

	return wrapped.Account, nil
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
