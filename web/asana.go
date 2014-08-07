package web
import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  m "github.com/thedahv/asanaprojectlist/models"
  "github.com/spf13/viper"
  "strconv"
)

type projectResponse struct {
  Data []m.Project
}

type projectDetailsResponse struct {
  Data m.ProjectDetails
}

type projectTaskIDs struct {
  ID int
}
type projectTaskIdsResponse struct {
  Data []projectTaskIDs
}

type projectTaskDetailResponse struct {
  Data m.ProjectTask
}

func authenticatedGet(path string) ([]byte, error) {
  client := &http.Client{}

  apikey := viper.GetString("asanakey")
  req, _ := http.NewRequest("GET", "https://app.asana.com/api/1.0/" + path, nil)
  req.SetBasicAuth(apikey, "")

  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  body, _ := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()

  return body, nil
}

// GetProjects - gets all projects for the current user
func GetProjects() []m.Project {
  body, err := authenticatedGet("projects")
  if err != nil {
    panic(err.Error())
  } else {
    var responseData projectResponse
    err = json.Unmarshal(body, &responseData)
    if err != nil {
      panic(err.Error())
    }

    return responseData.Data
  }
}

func getProjectTaskIDS(projectID string) []projectTaskIDs {
  // Get list of task IDS
  idsData, err := authenticatedGet("projects/" + projectID + "/tasks")
  if err != nil {
    panic(err.Error())
  }

  var idsList projectTaskIdsResponse
  err = json.Unmarshal(idsData, &idsList)

  if err != nil {
    panic(err.Error())
  }

  return idsList.Data
}

func getProjectTaskDetail(taskID int) m.ProjectTask {
  taskDetailData, err := authenticatedGet("tasks/" + strconv.Itoa(taskID))
  if err != nil {
    panic(err.Error())
  }

  var taskDetail projectTaskDetailResponse
  err = json.Unmarshal(taskDetailData, &taskDetail)

  if err != nil {
    panic(err.Error())
  }

  return taskDetail.Data
}

// GetProjectDetails - Returns the project details for the given project ID
func GetProjectDetails(projectID string) m.ProjectDetails {
  body, err := authenticatedGet("projects/" + projectID)
  if err != nil {
    panic(err.Error())
  }

  // Get the Project body
  var responseData projectDetailsResponse
  err = json.Unmarshal(body, &responseData)
  if err != nil {
    panic(err.Error())
  }

  // Get list of IDs
  idsList := getProjectTaskIDS(projectID)

  // Expand the list of project tasks
  var projectTaskDetails []m.ProjectTask

  for _, task := range idsList {
    taskDetail := getProjectTaskDetail(task.ID)
    projectTaskDetails = append(projectTaskDetails, taskDetail)
  }

  responseData.Data.Tasks = projectTaskDetails

  return responseData.Data
}
