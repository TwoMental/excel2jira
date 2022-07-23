package main

import (
	"excel2jira/client"
	"excel2jira/excel"
	"flag"
	"fmt"
)

func main() {
	// args
	parseArgs()

	// check jira connection
	client.Init()

	// excel work
	excel.Init()
	defer excel.Close()
	excel.Start()
}

func parseArgs() {
	flag.StringVar(&client.JiraUrl, "url", "http://atlassian.sample.com", "Address of Jira")
	flag.StringVar(&client.JiraUser, "username", "YourName", "Username of Jira")
	flag.StringVar(&client.JiraPsw, "password", "ChangeMe", "Password of Jira")
	flag.StringVar(&excel.RelationFile, "relation", "relation.json", "Relation file of fields")
	flag.StringVar(&excel.FileName, "excel", "sample.xlsx", "Excel file")

	flag.Parse()
	// version
	versionArgs := []string{"v", "V", "version"}
	if stringInSlice(versionArgs, flag.Arg(0)) {
		fmt.Println("v0.1.0")
	}

}
func stringInSlice(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//func Contains[T comparable](arr []T, x T) bool {
//	for _, v := range arr {
//		if v == x {
//			return true
//		}
//	}
//	return false
//}
