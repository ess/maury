// Package maury is a repository-inspired client for the Engine Yard API
package maury

import (
  "github.com/ess/maury/client"
)

// NewClient returns a low-level HTTP driver configurd for the Engine Yard API
// for the given base URL and token. If there are problems initializing the
// client, then an error is returned.
func NewClient(baseURL string, token string) (*client.Driver, error) {
  return client.New(baseURL, token)
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
