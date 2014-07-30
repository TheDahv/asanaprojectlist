package web
import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  m "github.com/thedahv/asanaprojectlist/models"
  "github.com/spf13/viper"
)

type projectResponse struct {
  Data []m.Project
}

// GetProjects - gets all projects for the current user
func GetProjects() []m.Project {
  client := &http.Client{}

  apikey := viper.GetString("asanakey")
  req, _ := http.NewRequest("GET", "https://app.asana.com/api/1.0/projects", nil)
  req.SetBasicAuth(apikey, "")

  resp, err := client.Do(req)
  if err != nil {
    panic(err.Error())
  }

  if err != nil {
    panic(err.Error())
  } else {
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    var responseData projectResponse
    err = json.Unmarshal(body, &responseData)
    if err != nil {
      panic(err.Error())
    }

    return responseData.Data
  }
}
