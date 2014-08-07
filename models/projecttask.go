package models

// AssigneeInfo - The information about the owner of the task
type AssigneeInfo struct {
  Name string
}

// ProjectTask - The tasks listed in a given project
type ProjectTask struct {
  ID int
  Assignee AssigneeInfo
  Completed bool
  Name string
  Notes string
}
