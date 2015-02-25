package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"code.google.com/p/gorilla/mux"
	m "github.com/startupweekend/asanaprojectlist/models"
)

// SetupRoutes returns a router with all routes defined
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/projects", projectsHandler)
	r.HandleFunc("/projects/{projectID}/tasks", projectTasksHandler)
	r.HandleFunc("/projects/{projectID}", projectDetailsHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./webapp")))

	return r
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("projectsHandler")
	projects := []m.Project{}

	response, err := GetProjects()
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		for _, project := range response {
			// Filter out projects that are not "top-tier" projects
			if project.IsAProject() {
				// Prepare virtual fields for JSON output
				project.Prepare()
				projects = append(projects, project)
			}
		}

		fmt.Printf("Replying with %d projects\n", len(projects))
		output, err := json.Marshal(projects)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error marshalling projects: " + err.Error()))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(output)
		}
	}
}

func projectDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID := vars["projectID"]

	projectDetails := GetProjectDetails(projectID)

	output, _ := json.Marshal(projectDetails)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func projectTasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID := vars["projectID"]

	projectTasks, err := GetProjectTasks(projectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		output, _ := json.Marshal(projectTasks)

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	}
}
