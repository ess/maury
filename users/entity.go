// Package users provides the data structures and functions for modeling
// the Users endpoint on the Engine Yard API
package users

// Entity is a flat data structure that maps to an upstream User
type Entity struct {
	ID string `json:"id,omitempty"`
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