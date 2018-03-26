package accounts

import (
  "encoding/json"
  "errors"
  "net/url"
  "strconv"
  "strings"

  "github.com/ess/maury/result"
  "github.com/ess/maury/users"
)

type Reader interface{
  Get(path string, params url.Values) *result.Result
}

func All(driver Reader, params url.Values) []*Entity {
  return allPages(driver, "accounts", params)
}

func ForUser(driver Reader, user *users.Entity, params url.Values) []*Entity {
  pathParts := []string{"users", user.ID, "accounts"}

  return allPages(driver, strings.Join(pathParts, "/"), params)
}

func Find(driver Reader, id string) (*Entity, error) {
  response := driver.Get("accounts/" + id, nil)

  wrapper := struct{Account *Entity `json:"account,omitempty"`}{}
  if response.Ok() {
    data := response.Value().([]byte)
    err := json.Unmarshal(data, &wrapper)
    if err != nil {
      return nil, err
    }

    return wrapper.Account, nil
  }

  return nil, errors.New("Could not find account with ID")
}

func allPages(driver Reader, path string, params url.Values) []*Entity {
  var accounts []*Entity

  maxResults := 100
  page := 1

  if params == nil {
    params = url.Values{}
  }

  params.Set("per_page", strconv.Itoa(maxResults))

  wrapper := struct{Accounts []*Entity `json:"accounts,omitempty"`}{
    Accounts: []*Entity{&Entity{}},
  }

  for len(wrapper.Accounts) > 0 {
    params.Set("page", strconv.Itoa(page))

    response := driver.Get(path, params)
    if response.Ok() {
      data := response.Value().([]byte)
      if err := json.Unmarshal(data, &wrapper); err == nil {
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
