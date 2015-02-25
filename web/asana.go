package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
	m "github.com/startupweekend/asanaprojectlist/models"
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
	fmt.Printf("authenticatedGet(%s)\n", path)
	client := &http.Client{}

	apikey := viper.GetString("asanakey")
	reqPath := "https://app.asana.com/api/1.0/" + path
	req, err := http.NewRequest("GET", reqPath, nil)
	if err != nil {
		fmt.Printf("Error creating request for %s: %s\n", reqPath, err.Error())
		return nil, err
	}
	req.SetBasicAuth(apikey, "")

	fmt.Printf("Issuing request to %s\n", reqPath)
	if len(apikey) == 0 {
		fmt.Printf("Authentication key is empty")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error on client request %s\n", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading API response: %s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}

// GetProjects - gets all projects for the current user or returns an error
func GetProjects() (projects []m.Project, err error) {
	fmt.Println("GetProjects()")
	body, err := authenticatedGet("projects")
	if err != nil {
		fmt.Printf("Got an error loading projects: %s", err.Error())
		return
	}

	var responseData projectResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Printf("Got an error parsing unmarshalling projects response: %s\n", err.Error())
		return
	}

	projects = responseData.Data

	return
}

func getProjectTaskIDS(projectID string) (ids []projectTaskIDs, err error) {
	fmt.Printf("getProjectTaskIDS(%s)\n", projectID)
	// Get list of task IDS
	idsData, err := authenticatedGet("projects/" + projectID + "/tasks")
	if err != nil {
		fmt.Printf("Error in project IDS API call: %s\n", err.Error())
		return nil, err
	}

	var idsList projectTaskIdsResponse
	err = json.Unmarshal(idsData, &idsList)

	if err != nil {
		fmt.Printf("Error unmarshalling project task IDS: %s\n", err.Error())
		return nil, err
	}

	ids = idsList.Data

	return
}

func getProjectTaskDetail(taskID int) (task m.ProjectTask, err error) {
	taskDetailData, err := authenticatedGet("tasks/" + strconv.Itoa(taskID))
	if err != nil {
		return
	}

	var taskDetail projectTaskDetailResponse
	err = json.Unmarshal(taskDetailData, &taskDetail)

	if err != nil {
		return
	}

	task = taskDetail.Data
	return
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

	return responseData.Data
}

// GetProjectTasks - Given a project, download the list of tasks and then
// load the task details for each
func GetProjectTasks(projectID string) (tasks []m.ProjectTask, err error) {
	// Get list of IDs
	idsList, err := getProjectTaskIDS(projectID)
	if err != nil {
		return
	}

	// Concurrently expand the list of project tasks
	tasksLength := len(idsList)

	type empty struct{}                  // Semaphore for timing and sequencing
	sem := make(chan empty, tasksLength) // as we are loading tasks
	taskErrorChan := make(chan error)

	// Empty slice to hold our tasks details
	tasks = make([]m.ProjectTask, tasksLength)

	for i, task := range idsList {
		// Spin up a goroutine as a closure over the
		// results slice and loop through each task
		go func(i int, taskID int) {
			tasks[i], err = getProjectTaskDetail(taskID)
			if err != nil {
				taskErrorChan <- err
			}
			// Ping back on the channel when it is done
			sem <- empty{}
		}(i, task.ID)
	}
	// Wait for each goroutine on the channel to ping back
	for i := 0; i < tasksLength; i++ {
		select {
		case _ = <-sem: // Do nothing
		case err := <-taskErrorChan:
			fmt.Printf("Error when loading task details: %s\n", err.Error())
			// Bail early from further processing
			break
		}
	}

	close(taskErrorChan)
	close(sem)

	return
}
