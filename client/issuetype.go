package client

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/google/go-querystring/query"
)

func GetIssueTypeList() (*IssueTypeList, *jira.Response, error) {
	issueTypeService := IssueTypeService{
		client: JiraClient,
	}
	return issueTypeService.GetList()
}

type IssueTypeService struct {
	client *jira.Client
}

type IssueTypeList []jira.IssueType

func (s *IssueTypeService) ListWithOptionsWithContext(ctx context.Context, options *jira.GetQueryOptions) (*IssueTypeList, *jira.Response, error) {
	apiEndpoint := "rest/api/2/issuetype"
	req, err := s.client.NewRequestWithContext(ctx, "GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	projectList := new(IssueTypeList)
	resp, err := s.client.Do(req, projectList)
	if err != nil {
		jiraErr := jira.NewJiraError(resp, err)
		return nil, resp, jiraErr
	}

	return projectList, resp, nil
}

func (s *IssueTypeService) GetListWithContext(ctx context.Context) (*IssueTypeList, *jira.Response, error) {
	return s.ListWithOptionsWithContext(ctx, &jira.GetQueryOptions{})
}

func (s *IssueTypeService) GetList() (*IssueTypeList, *jira.Response, error) {
	return s.GetListWithContext(context.Background())
}
