package me

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	ContentParagraph string = "p"
	ContentHeading   string = "h"
	ContentImage     string = "i"
	ContentVideo     string = "v"
	ContentTable     string = "t"
	ContentLink      string = "l"
	ContentBold      string = "b"
	ContentListItem  string = "li"
)

type Content struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Project struct {
	Title        string    `json:"title"`
	Links        []string  `json:"links"`
	Summary      string    `json:"summary"`
	StartDate    Date      `json:"startDate"`
	Contents     []Content `json:"contents"`
	Context      string    `json:"context"`
	Active       bool      `json:"active"`
	Tags         []string  `json:"tags"`
	Languages    []string  `json:"languages"`
	PreviewImage string    `json:"previewImage"`
	EndPoint     string
}

type ProjectData struct {
	Projects []*Project `json:"projects"`
}

func (p *ProjectData) Load() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, "data", "projects.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return err
	}
	for _, proj := range p.Projects {
		proj.EndPoint = GenerateEndPoint(proj.Title)
	}
	return nil
}

func (p *ProjectData) Save() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := filepath.Join(cwd, "data", "projects.json")
	data, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	return nil
}

func (projects *ProjectData) Add(project *Project) error {
	for _, proj := range projects.Projects {
		if proj.Title == project.Title {
			return errors.New("Project already exists")
		}
	}
	projects.Projects = append(projects.Projects, project)
	return nil
}

func (projects *ProjectData) Remove(index int) error {
	if index >= len(projects.Projects) {
		return errors.New("Index out of range")
	}
	projects.Projects = append(projects.Projects[:index], projects.Projects[index+1:]...)
	return nil
}

var (
	nonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)
	edgeDashes      = regexp.MustCompile(`^-|-$`)
)

func GenerateEndPoint(title string) string {
	endPoint := strings.ToLower(strings.TrimSpace(title))
	endPoint = nonAlphanumeric.ReplaceAllString(endPoint, "-")
	endPoint = edgeDashes.ReplaceAllString(endPoint, "")
	return endPoint
}
