package result

import (
  "errors"
  "testing"
)

func TestResult_Ok(t *testing.T) {
  t.Run("when the result has an error", func(t *testing.T) {
    result := New(1, errors.New("onoes"))

    t.Run("it is false", func(t *testing.T) {
      if result.Ok() {
        t.Errorf("Expected Ok to be false")
      }
    })
  })

  t.Run("when the result has no error", func(t *testing.T) {
    result := New(1, nil)

    t.Run("it is true", func(t *testing.T) {
      if !result.Ok() {
        t.Errorf("Expected Ok to be true")
      }
    })
  })
}

func TestResult_Value(t *testing.T) {
  t.Run("when the result has an error", func(t *testing.T) {
    result := New(1, errors.New("onoes"))

    t.Run("it is nil", func(t *testing.T) {
      if result.Value() != nil {
        t.Errorf("Expected value to be nil, got an actual value")
      }
    })
  })

  t.Run("when the result has no error", func(t *testing.T) {
    result := New(1, nil)

    t.Run("it is the wrapped value", func(t *testing.T) {
      if result.Value() == nil {
        t.Errorf("Expected value not to be nil, but it was")
      }
    })
  })
}

func TestResult_Error(t *testing.T) {
  err := errors.New("onoes")
  good := New(1, nil)
  bad := New(1, err)

  t.Run("it is the wrapped error", func(t *testing.T) {
    if good.Error() != nil {
      t.Errorf("Expected nil without an error")
    }

    if bad.Error() != err {
      t.Errorf("Expected the wrapped error when present")
    }
  })
}
