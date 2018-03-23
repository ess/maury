// Package maury is a repository-inspired client for the Engine Yard API
package maury

import (
  "github.com/ess/maury/client"
)

// NewClient returns a low-level HTTP driver configurd for the Engine Yard API
// for the given base URL and token. If there are problems initializing the
// client, then an error is returned.
func NewClient(baseUrl string, token string) (*client.Driver, error) {
  return client.New(baseUrl, token)
}
