package client

import (
	"excel2jira/utils"
)

func GetIssueById(id string) {
	issue, _, _ := JiraClient.Issue.Get(id, nil)
	utils.PrintStruct(issue)
}
