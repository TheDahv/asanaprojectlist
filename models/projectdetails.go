package models

// ProjectMember - The name of a member on the project team
type ProjectMember struct {
  Name string
}

// ProjectDetails - Returns the detailed information of a project
type ProjectDetails struct {
  Notes string
  Members []ProjectMember
  Tasks []ProjectTask
}
