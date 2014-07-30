package web

import (
  "net/http"
  "io/ioutil"
)

// GetEvents does a thing
func GetEvents() string {
  resp, err := http.Get("https://swoop.up.co/events?event_type=Startup%20Weekend&city=Seattle")

  output := ""

  if err == nil {
    body, _ := ioutil.ReadAll(resp.Body)

    output = string(body[:])

    resp.Body.Close()

  } else {
    output = "failure!"
  }

  return output
}
