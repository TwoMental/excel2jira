package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"io"
)

type CreateIssueParams struct {
	ProjectRaw  string
	ProjectKey  string
	Summary     string
	Description string
	TypeRaw     string
	TypeName    string
}

func Create(params CreateIssueParams) (int, error) {
	// project
	if params.ProjectRaw != "" {
		projectKey, err := validateProject(params.ProjectRaw)
		if err != nil {
			return 0, err
		} else {
			params.ProjectKey = projectKey
		}
	}
	fmt.Printf("projectKey: %s", params.ProjectKey)

	return 0, nil
}

func TestSample() {

	//自定义参数
	//customeFields := tcontainer.MarshalMap{
	//	// story points:1
	//	"customfield_10102": 1,
	//}

	issue := jira.Issue{
		Fields: &jira.IssueFields{
			Project:     jira.Project{Key: "TIYAN"},
			Summary:     "test_auto_01",
			Description: "Description",
			// 默认类型: Task（任务）
			Type: jira.IssueType{Name: "任务"},
			// 默认优先级：Minor/major（中/高）
			Priority: &jira.Priority{ID: "3"},
			//// 必填参数：预估时间
			//TimeTracking: &jira.TimeTracking{OriginalEstimate: "3d"},
			// 必填参数：经办人
			Assignee: &jira.User{Name: "linbei"},
			//// 修复版本
			//FixVersions: []*jira.FixVersion{&jira.FixVersion{Name: "无"}},
			//Components:  []*jira.Component{&jira.Component{Name: "无"}},
			//
			//// 其他自定义参数
			//Unknowns: customeFields,
		},
	}
	issueJson, _ := json.Marshal(issue)
	fmt.Printf("%+v\n", string(issueJson))

	createdIssue, res, err := JiraClient.Issue.Create(&issue)
	fmt.Println(createdIssue)
	printBody(res.Body)
	fmt.Println(err)
}

func printBody(body io.ReadCloser) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	fmt.Println(buf.String())
}
