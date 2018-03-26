package accounts

import (
	"errors"
	"fmt"
	"net/url"
	"testing"
)

type updater struct {
	responses map[string]string
}

func (r *updater) Put(path string, params url.Values, data []byte) ([]byte, error) {
	key := r.key(path)

	if r.responses == nil {
		r.reset()
	}

	response, ok := r.responses[key]
	if ok {
		return []byte(response), nil
	}

	return nil, errors.New("Updater didn't update")
}

func (r *updater) set(path string, response string) {
	key := r.key(path)

	if r.responses == nil {
		r.reset()
	}

	r.responses[key] = response
}

func (r *updater) key(path string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "api.engineyard.com",
		Path:   path,
	}

	return u.String()
}

func (r *updater) reset() {
	r.responses = make(map[string]string)
}

func TestUpdate(t *testing.T) {
	id := "12345"
	path := "accounts/" + id
	name := "George"
	emergency := "Larry"

	generate := func(id, name, contact string) string {
		return fmt.Sprintf(
			`{"account" : {"id" : "%s", "name" : "%s", "emergency_contact" : "%s"}}`,
			id,
			name,
			contact,
		)
	}

	original := &Entity{ID: id, Name: name, EmergencyContact: emergency}

	t.Run("when updating the name", func(t *testing.T) {
		change := &Changes{Name: emergency}

		t.Run("and the call succeeds", func(t *testing.T) {
			driver := &updater{}
			driver.set(path, generate(id, emergency, emergency))

			updated, err := Update(driver, original, change)

			t.Run("the entity has a new name", func(t *testing.T) {
				if updated.Name != emergency {
					t.Errorf("Expected for the name to be updated")
				}
			})

			t.Run("has no error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected no error")
				}
			})
		})

		t.Run("and the call fails", func(t *testing.T) {
			driver := &updater{}

			updated, err := Update(driver, original, change)

			t.Run("the entity is nil", func(t *testing.T) {
				if updated != nil {
					t.Errorf("Expected a nil entity")
				}
			})

			t.Run("the error is not nil", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error")
				}
			})
		})

		t.Run("and the API returns bad data", func(t *testing.T) {
			driver := &updater{}
			driver.set(path, "Just a string here.")

			updated, err := Update(driver, original, change)

			t.Run("the entity is nil", func(t *testing.T) {
				if updated != nil {
					t.Errorf("Expected a nil entity")
				}
			})

			t.Run("the error is not nil", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error")
				}
			})
		})
	})

	t.Run("when updating the emergency contact", func(t *testing.T) {
		change := &Changes{EmergencyContact: name}

		t.Run("and the call succeeds", func(t *testing.T) {
			driver := &updater{}
			driver.set(path, generate(id, name, name))

			updated, err := Update(driver, original, change)

			t.Run("the entity has a emergency contact", func(t *testing.T) {
				if updated.EmergencyContact != name {
					t.Errorf("Expected for the emergency contact to be updated")
				}
			})

			t.Run("has no error", func(t *testing.T) {
				if err != nil {
					t.Errorf("Expected no error")
				}
			})
		})

		t.Run("and the call fails", func(t *testing.T) {
			driver := &updater{}

			updated, err := Update(driver, original, change)

			t.Run("the entity is nil", func(t *testing.T) {
				if updated != nil {
					t.Errorf("Expected a nil entity")
				}
			})

			t.Run("the error is not nil", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error")
				}
			})
		})

		t.Run("and the API returns bad data", func(t *testing.T) {
			driver := &updater{}
			driver.set(path, "Just a string here.")

			updated, err := Update(driver, original, change)

			t.Run("the entity is nil", func(t *testing.T) {
				if updated != nil {
					t.Errorf("Expected a nil entity")
				}
			})

			t.Run("the error is not nil", func(t *testing.T) {
				if err == nil {
					t.Errorf("Expected an error")
				}
			})
		})
	})
}
