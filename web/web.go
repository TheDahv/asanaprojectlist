package web

import (
	"code.google.com/p/gorilla/mux"
	"net/http"
	m "github.com/thedahv/asanaprojectlist/models"
	"encoding/json"
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
	var projects []m.Project

	for _, project := range GetProjects() {
		// Filter out projects that are not "top-tier" projects
		if project.IsAProject() {
			// Prepare virtual fields for JSON output
			project.Prepare()
			projects = append(projects, project)
		}
	}

	output, _ := json.Marshal(projects)
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
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

	projectTasks := GetProjectTasks(projectID)

	output, _ := json.Marshal(projectTasks)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
