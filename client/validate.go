package client

import (
	"errors"
	"fmt"
	"strings"
)

var (
	memo map[string]map[string]interface{}
)

func validateProject(projectRaw string) (projectKey string, err error) {
	// 确认缓存
	if val, ok := memo["project"][projectRaw]; ok {
		return val.(string), nil
	}
	// 获取项目列表
	list, _, _ := JiraClient.Project.GetList()
	if list == nil {
		return "", errors.New("no project available")
	}

	// 查找项目
	for _, each := range *list {
		if strings.Contains(each.Name, projectRaw) ||
			strings.Contains(projectRaw, each.Name) ||
			strings.Contains(each.Key, projectRaw) ||
			projectRaw == each.ID {
			return each.Key, nil
		}

	}
	// 没找到
	return "", errors.New(fmt.Sprintf(
		"can't find project named %s", projectRaw))
}
