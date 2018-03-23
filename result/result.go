// Package result provides basic invariant result capabilities
package result

// Result wraps both a value and an error
type Result struct {
  value interface{}
  err error
}

// New takes a value and an error, and returns a Result that wraps them
func New(value interface{}, err error) *Result {
  return &Result{value, err}
}

// Ok is true if the Result contains an error, and false otherwise.
func (result *Result) Ok() bool {
  if result.err == nil {
    return true
  }

  return false
}

// Value is the value wrapped by the Result. If the result wraps an error
// as well as a value, then this is nil. Otherwise, it is the wrapped value.
func (result *Result) Value() interface{} {
  if result.Ok() {
    return result.value
  }

  return nil
}

// Error is the error wrapped by the Result.
func (result *Result) Error() error {
  return result.err
}
