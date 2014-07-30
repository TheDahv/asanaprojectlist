package models

import (
  "testing"
)

func makeGoodProject() Project {
  return Project{ ID: 1234, Name: "[G] Test Name" }
}

func makeBadProject() Project {
  return Project{ ID: 134, Name: "Not a real project" }
}

func TestGetName(t *testing.T) {
  name := makeGoodProject().GetName()

  if name != "Test Name" {
    t.Error("Expected 'Test Name', got ", name)
  }
}

func TestGetStatus(t *testing.T) {
  status := makeGoodProject().GetStatus()

  if status != "G" {
    t.Error("Expected 'G', got ", status)
  }
}

func TestDetectsProjectFormat(t *testing.T) {
  if makeGoodProject().IsAProject() == false {
    t.Error("Expected this to be a project")
  }
}

func TestRejectsNonProjects(t *testing.T) {
  if makeBadProject().IsAProject() {
    t.Error("Expected this to not be a project")
  }
}
