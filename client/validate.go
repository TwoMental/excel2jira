package client

import (
	"errors"
	"fmt"
	"strings"
)

var (
	memo map[string]map[string]interface{}
)

func validateProject(rawValue string) (id string, err error) {
	fieldName := "project"
	// find in memo
	if val, ok := memo[fieldName][rawValue]; ok {
		return val.(string), nil
	}
	// get available value list
	list, _, _ := JiraClient.Project.GetList()
	if list == nil {
		return "", errors.New(fmt.Sprintf(
			"no %s available", fieldName))
	}
	// find match value
	for _, each := range *list {
		if strings.Contains(each.Name, rawValue) ||
			strings.Contains(rawValue, each.Name) ||
			strings.Contains(each.Key, rawValue) ||
			rawValue == each.ID {
			return each.ID, nil
		}

	}
	// no match value
	return "", errors.New(fmt.Sprintf(
		"can't find %s named %s", fieldName, rawValue))
}

func validateIssueType(rawValue string) (id string, err error) {
	fieldName := "type"
	// find in memo
	if val, ok := memo[fieldName][rawValue]; ok {
		return val.(string), nil
	}
	// 获取类型列表
	list, _, _ := GetIssueTypeList()
	if list == nil {
		return "", errors.New(fmt.Sprintf(
			"no %s available", fieldName))
	}
	// get available value list
	for _, each := range *list {
		if strings.Contains(each.Name, rawValue) ||
			strings.Contains(rawValue, each.Name) ||
			rawValue == each.ID {
			return each.ID, nil
		}

	}
	// find match value
	return "", errors.New(fmt.Sprintf(
		"can't find %s named %s", fieldName, rawValue))
}

func validatePriority(rawValue string) (id string, err error) {
	fieldName := "priority"
	// find in memo
	if val, ok := memo[fieldName][rawValue]; ok {
		return val.(string), nil
	}
	// get available value list
	list, _, _ := JiraClient.Priority.GetList()
	if len(list) == 0 {
		return "", errors.New(fmt.Sprintf(
			"no %s available", fieldName))
	}
	// find match value
	for _, each := range list {
		if strings.Contains(each.Name, rawValue) ||
			strings.Contains(rawValue, each.Name) ||
			rawValue == each.ID {
			return each.ID, nil
		}

	}
	// no match value
	return "", errors.New(fmt.Sprintf(
		"can't find %s named %s", fieldName, rawValue))
}
