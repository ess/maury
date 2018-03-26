package accounts

import (
  "encoding/json"
  "net/url"

  "github.com/ess/maury/result"
)

type Updater interface {
  Put(path string, params url.Values, data []byte) *result.Result
}

type Changes struct {
  Name string `json:"name,omitempty"`
  EmergencyContact string `json:"emergency_contact,omitempty"`
  SupportPlan string `json:"support_plan,omitempty"`
}

func Update(driver Updater, account *Entity, changes *Changes) (*Entity, error) {

  wrappedChanges := struct{Account *Changes `json:"account,omitempty"`}{
    Account: changes,
  }

  data, err := json.Marshal(&wrappedChanges)
  if err != nil {
    return nil, err
  }

  response := driver.Put("accounts/" + account.ID, nil, data)
  if response.Ok() {
    changed := response.Value().([]byte)

    wrapped := struct{Account *Entity `json:"account,omitempty"`}{}

    err = json.Unmarshal(changed, &wrapped)
    if err != nil {
      return nil, err
    }

    return wrapped.Account, nil
  }

  return nil, response.Error()
}
