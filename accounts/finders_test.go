package accounts

import (
  "errors"
  "fmt"
  "net/url"
  "strings"
  "testing"

  "github.com/ess/maury/users"
)

type reader struct {
  responses map[string]string
}

func (r *reader) Get(path string, params url.Values) ([]byte, error) {
  key := r.key(path, params)

  if r.responses == nil {
    r.reset()
  }

  response, ok := r.responses[key]
  if ok {
    return []byte(response), nil
  }

  return nil, errors.New("Getter didn't get")
}

func (r *reader) set(path string, params url.Values, response string) {
  key := r.key(path, params)

  if r.responses == nil {
    r.reset()
  }

  r.responses[key] = response
}

func (r *reader) key(path string, params url.Values) string {
  if params == nil {
    params = url.Values{}
  }

  u := url.URL{
    Scheme: "https",
    Host: "api.engineyard.com",
    Path: path,
    RawQuery: params.Encode(),
  }

  return u.String()
}

func (r *reader) reset() {
  r.responses = make(map[string]string)
}

func TestAll(t *testing.T) {
  generate := func(start, finish int) string {
    var accounts []string

    for x := start; x <= finish; x++ {
      accounts = append(accounts, fmt.Sprintf(`{"id" : "%d"}`, x))
    }

    return fmt.Sprintf(`{"accounts" : [%s]}`, strings.Join(accounts, ","))
  }

  t.Run("when there are no accounts visible", func(t *testing.T) {
    driver := &reader{}
    params := url.Values{}
    params.Set("page", "1")
    params.Set("per_page", "100")

    driver.set("accounts", params, `{"accounts" : []}`)

    all := All(driver, nil)

    t.Run("it is empty", func(t *testing.T) {
      if len(all) > 0 {
        t.Errorf("Expected an empty array, got one with %d members", len(all))
      }
    })
  })

  t.Run("when there are accounts visible", func(t *testing.T) {
    t.Run("and there are fewer than 100 results", func (t *testing.T) {
      driver := &reader{}
      params := url.Values{}
      params.Set("page", "1")
      params.Set("per_page", "100")

      driver.set("accounts", params, generate(1, 10))

      all := All(driver, nil)

      t.Run("it contains the entities the API returned", func(t *testing.T) {
        if len(all) != 10 {
          t.Errorf("Expected 10 entities in the collection, got %d", len(all))
        }
      })
    })

    t.Run("and there are more than 100 results", func(t *testing.T) {
      driver := &reader{}
      params := url.Values{}
      params.Set("page", "1")
      params.Set("per_page", "100")

      driver.set("accounts", params, generate(1, 100))

      params.Set("page", "2")

      driver.set("accounts", params, generate(101, 110))

      all := All(driver, nil)

      t.Run("it contains all of the entities the API returned", func(t *testing.T) {
        if len(all) != 110 {
          t.Errorf("Expected 110 entities, got %d", len(all))
        }
      })
    })
  })
}

func TestForUser(t *testing.T) {
  generate := func(start, finish int) string {
    var accounts []string

    for x := start; x <= finish; x++ {
      accounts = append(accounts, fmt.Sprintf(`{"id" : "%d"}`, x))
    }

    return fmt.Sprintf(`{"accounts" : [%s]}`, strings.Join(accounts, ","))
  }

  user := &users.Entity{
    ID: "12345",
  }

  path := "users/12345/accounts"

  t.Run("when there are no accounts visible", func(t *testing.T) {
    driver := &reader{}
    params := url.Values{}
    params.Set("page", "1")
    params.Set("per_page", "100")

    driver.set(path, params, `{"accounts" : []}`)

    all := ForUser(driver, user, nil)

    t.Run("it is empty", func(t *testing.T) {
      if len(all) > 0 {
        t.Errorf("Expected an empty array, got one with %d members", len(all))
      }
    })
  })

  t.Run("when there are accounts visible", func(t *testing.T) {
    t.Run("and there are fewer than 100 results", func (t *testing.T) {
      driver := &reader{}
      params := url.Values{}
      params.Set("page", "1")
      params.Set("per_page", "100")

      driver.set(path, params, generate(1, 10))

      all := ForUser(driver, user, nil)

      t.Run("it contains the entities the API returned", func(t *testing.T) {
        if len(all) != 10 {
          t.Errorf("Expected 10 entities in the collection, got %d", len(all))
        }
      })
    })

    t.Run("and there are more than 100 results", func(t *testing.T) {
      driver := &reader{}
      params := url.Values{}
      params.Set("page", "1")
      params.Set("per_page", "100")

      driver.set(path, params, generate(1, 100))

      params.Set("page", "2")

      driver.set(path, params, generate(101, 110))

      all := ForUser(driver, user, nil)

      t.Run("it contains all of the entities the API returned", func(t *testing.T) {
        if len(all) != 110 {
          t.Errorf("Expected 110 entities, got %d", len(all))
        }
      })
    })
  })
}

func TestFind(t *testing.T) {
  id := "8675309"
  path := "accounts/" + id

  t.Run("when the account does not exist", func(t *testing.T) {
    driver := &reader{}

    account, err := Find(driver, id)

    t.Run("the entity is nil", func(t *testing.T) {
      if account != nil {
        t.Errorf("Expected no value")
      }
    })

    t.Run("the error is not nil", func(t *testing.T) {
      if err == nil {
        t.Errorf("Expected an error")
      }
    })
  })

  t.Run("when the account exists", func(t *testing.T) {
    driver := &reader{}

    driver.set(path, nil, fmt.Sprintf(`{"account" : {"id" : "%s"}}`, id))

    account, err := Find(driver, id)

    t.Run("the entity is not nil", func(t *testing.T) {
      if account == nil {
        t.Errorf("Expected an account entity")
      }
    })

    t.Run("the error is nil", func(t *testing.T) {
      if err != nil {
        t.Errorf("Expected no error")
      }
    })

  })

  t.Run("when the API sends bad data", func(t *testing.T) {
    driver := &reader{}

    driver.set(path, nil, "This is a string.")

    account, err := Find(driver, id)

    t.Run("the entity is nil", func(t *testing.T) {
      if account != nil {
        t.Errorf("Expected a nil entity")
      }
    })

    t.Run("the error is not nil", func(t *testing.T) {
      if err == nil {
        t.Errorf("Expected an error")
      }
    })
  })
}
