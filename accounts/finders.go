package accounts

import (
  "encoding/json"
  "net/url"
  "strconv"
  "strings"

  "github.com/ess/maury/users"
)

type Reader interface{
  Get(string, url.Values) ([]byte, error)
}

func All(driver Reader, params url.Values) []*Entity {
  return allPages(driver, "accounts", params)
}

func ForUser(driver Reader, user *users.Entity, params url.Values) []*Entity {
  pathParts := []string{"users", user.ID, "accounts"}

  return allPages(driver, strings.Join(pathParts, "/"), params)
}

func Find(driver Reader, id string) (*Entity, error) {
  response, err := driver.Get("accounts/" + id, nil)
  if err != nil {
    return nil, err
  }

  wrapper := struct{Account *Entity `json:"account,omitempty"`}{}

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

  wrapper := struct{Accounts []*Entity `json:"accounts,omitempty"`}{
    Accounts: []*Entity{&Entity{}},
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
