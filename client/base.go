package client

import (
	"github.com/andygrunwald/go-jira"
)

var (
	JiraClient *jira.Client
	// TODO: lin, read from arg
	jiraUrl  = "http://xxx.xxx.xxx"
	jiraUser = "xxxx"
	jiraPsw  = "xxxx"
)

func init() {
	var err error
	auth := jira.BasicAuthTransport{
		Username: jiraUser,
		Password: jiraPsw,
	}
	JiraClient, err = jira.NewClient(auth.Client(), jiraUrl)
	if err != nil {
		panic(err.Error())
	}

	checkConnection()
}

func checkConnection() {
	list, _, _ := JiraClient.Project.GetList()
	if list == nil {
		panic("no project available, pls check connection or user")
	}
}
