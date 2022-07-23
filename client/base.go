package client

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

var (
	JiraClient *jira.Client
	JiraUrl    string
	JiraUser   string
	JiraPsw    string
)

func Init() {
	var err error
	auth := jira.BasicAuthTransport{
		Username: JiraUser,
		Password: JiraPsw,
	}
	JiraClient, err = jira.NewClient(auth.Client(), JiraUrl)
	if err != nil {
		panic(err.Error())
	}

	checkConnection()
}

func checkConnection() {
	list, _, _ := JiraClient.Project.GetList()
	if list == nil {
		panic("no project available, pls check connection or user")
	} else {
		fmt.Printf("Connecting to Jira %s success!\n", JiraUrl)
	}
}
