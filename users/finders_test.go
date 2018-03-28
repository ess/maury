package users

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"testing"
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
		Scheme:   "https",
		Host:     "api.engineyard.com",
		Path:     path,
		RawQuery: params.Encode(),
	}

	return u.String()
}

func (r *reader) reset() {
	r.responses = make(map[string]string)
}

func TestAll(t *testing.T) {
	generate := func(start, finish int) string {
		var users []string

		for x := start; x <= finish; x++ {
			users = append(users, fmt.Sprintf(`{"id" : "%d"}`, x))
		}

		return fmt.Sprintf(`{"users" : [%s]}`, strings.Join(users, ","))
	}

	t.Run("when there are no users visible", func(t *testing.T) {
		driver := &reader{}
		params := url.Values{}
		params.Set("page", "1")
		params.Set("per_page", "100")

		driver.set("users", params, `{"users" : []}`)

		all := All(driver, nil)

		t.Run("it is empty", func(t *testing.T) {
			if len(all) > 0 {
				t.Errorf("Expected an empty array, got one with %d members", len(all))
			}
		})
	})

	t.Run("when there are users visible", func(t *testing.T) {
		t.Run("and there are fewer than 100 results", func(t *testing.T) {
			driver := &reader{}
			params := url.Values{}
			params.Set("page", "1")
			params.Set("per_page", "100")

			driver.set("users", params, generate(1, 10))

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

			driver.set("users", params, generate(1, 100))

			params.Set("page", "2")

			driver.set("users", params, generate(101, 110))

			all := All(driver, nil)

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
	path := "users/" + id

	t.Run("when the user does not exist", func(t *testing.T) {
		driver := &reader{}

		user, err := Find(driver, id)

		t.Run("the entity is nil", func(t *testing.T) {
			if user != nil {
				t.Errorf("Expected no value")
			}
		})

		t.Run("the error is not nil", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})

	t.Run("when the user exists", func(t *testing.T) {
		driver := &reader{}

		driver.set(path, nil, fmt.Sprintf(`{"user" : {"id" : "%s"}}`, id))

		user, err := Find(driver, id)

		t.Run("the entity is not nil", func(t *testing.T) {
			if user == nil {
				t.Errorf("Expected a user entity")
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

		user, err := Find(driver, id)

		t.Run("the entity is nil", func(t *testing.T) {
			if user != nil {
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

func TestCurrent(t *testing.T) {
	id := "current"
	path := "users/" + id

	t.Run("when the user does not exist", func(t *testing.T) {
		driver := &reader{}

		user, err := Current(driver)

		t.Run("the entity is nil", func(t *testing.T) {
			if user != nil {
				t.Errorf("Expected no value")
			}
		})

		t.Run("the error is not nil", func(t *testing.T) {
			if err == nil {
				t.Errorf("Expected an error")
			}
		})
	})

	t.Run("when the user exists", func(t *testing.T) {
		driver := &reader{}

		driver.set(path, nil, fmt.Sprintf(`{"user" : {"id" : "%s"}}`, id))

		user, err := Current(driver)

		t.Run("the entity is not nil", func(t *testing.T) {
			if user == nil {
				t.Errorf("Expected a user entity")
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

		user, err := Current(driver)

		t.Run("the entity is nil", func(t *testing.T) {
			if user != nil {
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
