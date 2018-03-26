package client

import (
  "net/url"
  "testing"

  httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestDriver_Get(t *testing.T) {
  driver, _ := New("https://api.engineyard.com", "faketoken")
  data := []byte(`{"sausages" : "gold"}`)

  t.Run(
    "when on the happy path",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "GET",
        "https://api.engineyard.com/sausages",
        httpmock.NewStringResponder(200, string(data)),
      )

      result, err := driver.Get("sausages", nil)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })

    })

  t.Run(
    "when on the happy path (with params)",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      data = []byte(`{"sausages" : "yup"}`)
      httpmock.RegisterResponder(
        "GET",
        "https://api.engineyard.com/sausages?color=gold",
        httpmock.NewStringResponder(200, string(data)),
      )

      params := url.Values{}
      params.Set("color", "gold")
      result, err := driver.Get("sausages", params)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })
    })

  t.Run(
    "when a wild API error appears",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "GET",
        "https://api.engineyard.com/ed209",
        httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
      )

      result, err := driver.Get("ed209", nil)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })

      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })

  t.Run(
    "when a problem exists betwen the chair and the keyboard",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "GET",
        "https://api.engineyard.com/404",
        httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
      )

      result, err := driver.Get("404", nil)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })

      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })
}

func TestDriver_Post(t *testing.T) {
  driver, _ := New("https://api.engineyard.com", "faketoken")

  data := []byte(`{"sausages" : "gold"}`)

  t.Run(
    "when on the happy path",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "POST",
        "https://api.engineyard.com/sausages",
        httpmock.NewStringResponder(200, string(data)),
      )

      result, err := driver.Post("sausages", nil, data)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })

    })

  t.Run(
    "when on the happy path (with params)",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "POST",
        "https://api.engineyard.com/sausages?color=gold",
        httpmock.NewStringResponder(200, string(data)),
      )

      params := url.Values{}
      params.Set("color", "gold")
      result, err := driver.Post("sausages", params, data)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })
    })

  t.Run(
    "when a wild API error appears",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "POST",
        "https://api.engineyard.com/ed209",
        httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
      )

      result, err := driver.Post("ed209", nil, data)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })

  t.Run(
    "when a problem exists betwen the chair and the keyboard",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "POST",
        "https://api.engineyard.com/404",
        httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
      )

      result, err := driver.Post("404", nil, data)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })
}

func TestDriver_Put(t *testing.T) {
  driver, _ := New("https://api.engineyard.com", "faketoken")

  data := []byte(`{"sausages" : "gold"}`)

  t.Run(
    "when on the happy path",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PUT",
        "https://api.engineyard.com/sausages",
        httpmock.NewStringResponder(200, string(data)),
      )

      result, err := driver.Put("sausages", nil, data)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })

    })

  t.Run(
    "when on the happy path (with params)",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PUT",
        "https://api.engineyard.com/sausages?color=gold",
        httpmock.NewStringResponder(200, string(data)),
      )

      params := url.Values{}
      params.Set("color", "gold")
      result, err := driver.Put("sausages", params, data)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })
    })

  t.Run(
    "when a wild API error appears",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PUT",
        "https://api.engineyard.com/ed209",
        httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
      )

      result, err := driver.Put("ed209", nil, data)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })

  t.Run(
    "when a problem exists betwen the chair and the keyboard",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PUT",
        "https://api.engineyard.com/404",
        httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
      )

      result, err := driver.Put("404", nil, data)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })
}

func TestDriver_Patch(t *testing.T) {
  driver, _ := New("https://api.engineyard.com", "faketoken")

  data := []byte(`{"sausages" : "gold"}`)

  t.Run(
    "when on the happy path",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PATCH",
        "https://api.engineyard.com/sausages",
        httpmock.NewStringResponder(200, string(data)),
      )

      result, err := driver.Patch("sausages", nil, data)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })

    })

  t.Run(
    "when on the happy path (with params)",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PATCH",
        "https://api.engineyard.com/sausages?color=gold",
        httpmock.NewStringResponder(200, string(data)),
      )

      params := url.Values{}
      params.Set("color", "gold")
      result, err := driver.Patch("sausages", params, data)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })
    })

  t.Run(
    "when a wild API error appears",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PATCH",
        "https://api.engineyard.com/ed209",
        httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
      )

      result, err := driver.Patch("ed209", nil, data)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run("it has an error", func(t *testing.T) {
        if err == nil {
          t.Errorf("Expected a non-nil error")
        }
      })

    })

  t.Run(
    "when a problem exists betwen the chair and the keyboard",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "PATCH",
        "https://api.engineyard.com/404",
        httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
      )

      result, err := driver.Patch("404", nil, data)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run("it has an error", func(t *testing.T) {
        if err == nil {
          t.Errorf("Expected a non-nil error")
        }
      })

    })
}

func TestDriver_Delete(t *testing.T) {
  driver, _ := New("https://api.engineyard.com", "faketoken")
  data := []byte(`{"sausages" : "gold"}`)

  t.Run(
    "when on the happy path",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "DELETE",
        "https://api.engineyard.com/sausages",
        httpmock.NewStringResponder(200, string(data)),
      )

      result, err := driver.Delete("sausages", nil)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })

    })

  t.Run(
    "when on the happy path (with params)",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      data = []byte(`{"sausages" : "yup"}`)

      httpmock.RegisterResponder(
        "DELETE",
        "https://api.engineyard.com/sausages?color=gold",
        httpmock.NewStringResponder(200, string(data)),
      )

      params := url.Values{}
      params.Set("color", "gold")
      result, err := driver.Delete("sausages", params)

      t.Run(
        "it is a success",
        func(t *testing.T) {
          if err != nil {
            t.Errorf("Call was not successful!")
          }
        })

      t.Run(
        "it has the expected value",
        func(t *testing.T) {
          if string(result) != string(data) {
            t.Errorf("Expected '%s', got '%s'", string(data), string(result))
          }
        })
    })

  t.Run(
    "when a wild API error appears",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "DELETE",
        "https://api.engineyard.com/ed209",
        httpmock.NewStringResponder(500, "Drop your weapon. You have 20 seconds to comply."),
      )

      result, err := driver.Delete("ed209", nil)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })

  t.Run(
    "when a problem exists betwen the chair and the keyboard",
    func(t *testing.T) {
      httpmock.Activate()
      defer httpmock.DeactivateAndReset()

      httpmock.RegisterResponder(
        "DELETE",
        "https://api.engineyard.com/404",
        httpmock.NewStringResponder(404, "You are now staring into the void. It is staring back."),
      )

      result, err := driver.Delete("404", nil)

      t.Run("it has no result", func(t *testing.T) {
        if result != nil {
          t.Errorf("Expected a nil result")
        }
      })


      t.Run(
        "it has an error",
        func(t *testing.T) {
          if err == nil {
            t.Errorf("Expected a non-nil error")
          }
        })

    })
}
