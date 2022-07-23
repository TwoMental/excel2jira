package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
)

type CreateIssueParams struct {
	ProjectRaw  string
	projectId   string
	Summary     string
	Description string
	TypeRaw     string
	typeID      string
	PriorityRaw string
	priorityId  string
	// TODO: lin, validate below fields
	TimeTrackingRaw string
	AssigneeRaw     string
	FixVersionRaw   string
	ComponentRaw    string
	customFields    interface{}
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
			params.projectId = projectKey
		}
	}
	// verify - issue type
	if params.TypeRaw != "" {
		typeId, err := validateIssueType(params.TypeRaw)
		if err != nil {
			return "", err
		} else {
			params.typeID = typeId
		}
	}
	// verify - priority
	if params.PriorityRaw != "" {
		priorityId, err := validatePriority(params.PriorityRaw)
		if err != nil {
			return "", err
		} else {
			params.priorityId = priorityId
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
	// fields
	fields := jira.IssueFields{
		Project:     jira.Project{ID: params.projectId},
		Summary:     params.Summary,
		Description: params.Description,
		Type:        jira.IssueType{ID: params.typeID},
		Priority:    &jira.Priority{ID: params.priorityId},
	}

	// issue
	issue := jira.Issue{Fields: &fields}
	createdIssue, res, err := JiraClient.Issue.Create(&issue)
	jiraRes := jiraRes{}
	_ = json.NewDecoder(res.Body).Decode(&jiraRes)
	return createdIssue, fmt.Sprintf("%v", jiraRes.Errs), err
}
