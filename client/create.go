package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
)

type CreateIssueParams struct {
	ProjectRaw  string
	ProjectKey  string
	Summary     string
	Description string
	TypeRaw     string
	TypeID      string
}

type jiraRes struct {
	Msg  interface{} `json:"errorMessages"`
	Errs interface{} `json:"errors"`
}

func Create(params CreateIssueParams) (string, error) {
	// verify - project
	if params.ProjectRaw != "" {
		projectKey, err := validateProject(params.ProjectRaw)
		if err != nil {
			return "", err
		} else {
			params.ProjectKey = projectKey
		}
	}
	// verify - issue type
	if params.TypeRaw != "" {
		typeId, err := validateIssueType(params.TypeRaw)
		if err != nil {
			return "", err
		} else {
			params.TypeID = typeId
		}
	}

	// create
	createdIssue, errDetail, err := createIssue(params)

	// error
	if err != nil {
		return "", errors.New(errDetail)
	} else {
		return createdIssue.Key, nil
	}
}

func createIssue(params CreateIssueParams) (*jira.Issue, string, error) {
	//utils.PrintStruct(params)
	// fields
	fields := jira.IssueFields{
		Project:     jira.Project{Key: params.ProjectKey},
		Summary:     params.Summary,
		Description: params.Description,
		Type:        jira.IssueType{ID: params.TypeID},
	}

	// issue
	issue := jira.Issue{Fields: &fields}
	createdIssue, res, err := JiraClient.Issue.Create(&issue)
	jiraRes := jiraRes{}
	json.NewDecoder(res.Body).Decode(&jiraRes)
	return createdIssue, fmt.Sprintf("%v", jiraRes.Errs), err
}
