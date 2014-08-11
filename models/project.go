package models

import (
  r "regexp"
)

// Project - Models an Asana project summary
type Project struct {
  ID int
  Name string
  ProjectName string
  Status string
}

func getNthMatch(pattern string, searchable string, position int) string {
  rx, err := r.Compile(pattern)

  if err != nil {
    panic(err.Error())
  }

  matches := rx.FindStringSubmatch(searchable)

  // If we have multiple matches and are asking for a
  // match within the length of addressable matches
  if len(matches) > 1 && (position + 1) < len(matches) {
    // The first array element is the entire match itself
    // So we increment the client's position to get the desired match
    return matches[position + 1]
  }
  return ""
}

func getFirstMatch(pattern string, searchable string) string {
  return getNthMatch(pattern, searchable, 0)
}

// IsAProject -- Checks the project name to see if it first the
// official Asana project format
func (p Project) IsAProject() bool {
  // Projects match the format: "[STATUS] PROJECT NAME GOES HERE"
  match, _ := r.MatchString("^\\[.\\] .*$", p.Name)
  return match
}

// GetName - Project#GetName parses the project's name from the data
func (p Project) GetName() string {
  if p.IsAProject() {
    return getFirstMatch("^\\[.\\] (.*)$", p.Name)
  }
  return p.Name
}

// GetStatus - Project#GetStatus parses the project's status from the data
func (p Project) GetStatus() string {
  status := getFirstMatch("^\\[([R|Y|G|!])\\] .*$", p.Name)

  if status == "" {
    return "Unknown"
  }

  if status == "!" {
    return "Done"
  }

  return status
}

// Prepare - Project#Prepare populates virtual attributes for JSON output
// and returns the reference to the project
func (p* Project) Prepare() {
  p.ProjectName = p.GetName()
  p.Status = p.GetStatus()
}
