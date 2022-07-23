package excel

import (
	"encoding/json"
	"excel2jira/client"
	"excel2jira/utils"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"strconv"
)

const (
	colJiraId = "jira_id"
	colResult = "result"
)

var (
	File         *excelize.File
	SheetName    string
	FileName     string
	RelationFile string
	colNameList  = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC"}
)

func Init() {
	var err error
	File, err = excelize.OpenFile(FileName)
	if err != nil {
		panic(err.Error())
	}
}

func Close() {
	// Close the spreadsheet.
	if err := File.Close(); err != nil {
		fmt.Println(err)
	}
}

func getSheetName() string {
	if SheetName != "" {
		return SheetName
	} else {
		return File.GetSheetList()[0]
	}
}

type ColRelation map[string]*ColRelationDetails
type ColRelationDetails struct {
	Name    string `json:"name,omitempty"`
	Id      int    `json:"id,omitempty"`
	Default string `json:"default,omitempty"`
}

func SetDefaultColRelation(relation ColRelation) {
	for _, col := range []string{
		"project", "summary", "issue_type",
		colJiraId, colResult} {
		if _, ok := relation[col]; !ok {
			relation[col] = &ColRelationDetails{Name: col, Id: -1}
		}
	}
}

func loadRelation() (relation ColRelation) {
	relation = make(ColRelation)
	if utils.IsFileExist(RelationFile) {
		file, _ := ioutil.ReadFile(RelationFile)
		err := json.Unmarshal([]byte(file), &relation)
		if err != nil {
			panic(fmt.Sprintf("load relation（%s）failed, err：%s.\n", RelationFile, err.Error()))
		}
		fmt.Printf("load relation（%s）success!\n", RelationFile)
	} else {
		fmt.Printf("relation file（%s）not exists，using default relation.\n", RelationFile)
	}
	SetDefaultColRelation(relation)
	return
}

func Start() {
	// get excel data
	rows, err := File.GetRows(getSheetName())
	if err != nil {
		fmt.Println(err)
		return
	}

	// load relation
	relation := loadRelation()
	maxColIndex := 0
	for colNum, colName := range rows[0] {
		maxColIndex += 1
		updateColNum(relation, colName, colNum+1)
	}

	// cols for update
	jiraCol, resultCol := getWriteCol(relation, maxColIndex)

	// create one by one
	for rowNum, rowContent := range rows[1:] {
		lineNumStr := strconv.FormatInt(int64(rowNum)+2, 10)
		row := RowService{rowContent, relation}
		resLoc := resultCol + lineNumStr

		if isCreated(row.Get(colJiraId)) {
			res := "jira id exists"
			fmt.Printf("Creating for line A%s jumped, reason: %s\n", lineNumStr, res)
			_ = File.SetCellValue(getSheetName(), resLoc, "jumped: "+res)
			_ = File.Save()
			continue
		}

		createParams := client.CreateIssueParams{
			ProjectRaw:      row.Get("project"),
			Summary:         row.Get("summary"),
			Description:     row.Get("description"),
			TypeRaw:         row.Get("issue_type"),
			PriorityRaw:     row.Get("priority"),
			TimeTrackingRaw: row.Get("time_tracking"),
			AssigneeRaw:     row.Get("assignee"),
			FixVersionRaw:   row.Get("fix_version"),
			ComponentRaw:    row.Get("component"),
		}

		if jiraId, err := client.Create(createParams); err != nil {
			// create with error
			fmt.Printf("Creating for line A%s failed, err: %s\n", lineNumStr, err.Error())
			_ = File.SetCellValue(getSheetName(), resLoc, "failed: "+err.Error())
			_ = File.Save()
		} else {
			// create succeed
			fmt.Printf("creating for line A%s success\n", lineNumStr)
			jiraIdLoc := jiraCol + lineNumStr
			_ = File.SetCellValue(getSheetName(), jiraIdLoc, jiraId)
			_ = File.SetCellValue(getSheetName(), resLoc, "success。")
			_ = File.Save()
		}
	}
}

func getWriteCol(relation ColRelation, maxColIndex int) (jiraCol string, resultCol string) {
	jiraCol, maxColIndex = getColOrNew(relation, colJiraId, maxColIndex)
	resultCol, _ = getColOrNew(relation, colResult, maxColIndex)
	return jiraCol, resultCol
}

func getColOrNew(relation ColRelation, key string, maxColIndex int) (col string, newMaxColIndex int) {
	if val, ok := relation[key]; ok && (val.Id != -1) {
		return colNameList[val.Id-1], maxColIndex
	} else {
		col := colNameList[maxColIndex]
		_ = File.SetCellValue(getSheetName(), col+"1", key)
		_ = File.Save()
		return col, maxColIndex + 1
	}
}

func updateColNum(relation ColRelation, name string, num int) {
	for k, v := range relation {
		if v.Name == name {
			relation[k].Id = num
			break
		}
	}
}

func isCreated(jiraId string) bool {
	if jiraId != "" {
		return true
	}
	return false
}
