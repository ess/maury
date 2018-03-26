package accounts

import (
  "encoding/json"
  "net/url"
)

type Updater interface {
  Put(string, url.Values, []byte) ([]byte, error)
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

  response, err := driver.Put("accounts/" + account.ID, nil, data)
  if err != nil {
    return nil, err
  }

  wrapped := struct{Account *Entity `json:"account,omitempty"`}{}

  err = json.Unmarshal(response, &wrapped)
  if err != nil {
    return nil, err
  }

  return wrapped.Account, nil
}
