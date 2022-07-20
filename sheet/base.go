package sheet

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io/ioutil"
	"jira/client"
	"jira/utils"
	"strconv"
)

var (
	File     *excelize.File
	fileName = "sample.xlsx"
	a2z      = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
)

func init() {
	var err error
	File, err = excelize.OpenFile(fileName)
	if err != nil {
		panic(err.Error())
	}

	// Get value from cell by given worksheet name and axis.
	//cell, err := f.GetCellValue("Sheet1", "B2")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(cell)

}

func Close() {
	// Close the spreadsheet.
	if err := File.Close(); err != nil {
		fmt.Println(err)
	}
}

func getDefaultSheet() string {
	return File.GetSheetList()[0]
}

//type ColRelation struct {
//	Project ColRelationDetails `json:"project"`
//	Summary ColRelationDetails `json:"summary"`
//}

type ColRelation map[string]*ColRelationDetails
type ColRelationDetails struct {
	Name string `json:"name"`
	Id   int    `json:"id,omitempty"`
}

func SetDefaultColRelation(relation ColRelation) {
	// 只填必填字段
	for _, col := range []string{
		"project", "summary", "description",
		"priority", "jira_id"} {
		if _, ok := relation[col]; !ok {
			relation[col] = &ColRelationDetails{col, -1}
		}
	}
}

func loadRelation() (relation ColRelation) {
	relation = make(ColRelation)
	// TODO: lin, read from arg
	fileName := "relation.json"
	if utils.IsFileExist(fileName) {
		file, _ := ioutil.ReadFile(fileName)
		err := json.Unmarshal([]byte(file), &relation)
		if err != nil {
			panic(fmt.Sprintf("load relation（%s）failed, err：%s.\n", fileName, err.Error()))
		}
		fmt.Printf("load relation（%s）success!\n", fileName)
	} else {
		fmt.Printf("relation file（%s）not exists，using default relation.\n", fileName)
	}
	SetDefaultColRelation(relation)
	return
}

func Start() {
	rows, err := File.GetRows(getDefaultSheet())
	if err != nil {
		fmt.Println(err)
		return
	}
	relation := loadRelation()
	maxColIndex := 0
	for colNum, colName := range rows[0] {
		maxColIndex += 1
		updateColNum(relation, colName, colNum)
	}
	//for k, v := range relation {
	//	fmt.Printf("k: %v, v.name: %v, v.id: %v\n", k, v.Name, v.Id)
	//}
	// cols for update
	// TODO: lin, check if exists
	jiraCol := a2z[relation["jira_id"].Id]
	resultCol := a2z[maxColIndex]

	// create one by one
	for rowNum, rowContent := range rows[1:] {
		lineNumStr := strconv.FormatInt(int64(rowNum)+2, 10)
		createParams := client.CreateIssueParams{
			ProjectRaw:  getColValue(rowContent, relation, "project"),
			Summary:     getColValue(rowContent, relation, "summary"),
			Description: getColValue(rowContent, relation, "description"),
		}
		resLoc := resultCol + lineNumStr

		if jiraId, err := client.Create(createParams); err != nil {
			// create with error
			fmt.Printf("creating for line A%s failed, err: %s\n", lineNumStr, err.Error())
			_ = File.SetCellValue(getDefaultSheet(), resLoc, "failed.\\n"+err.Error())
			//fmt.Println(err)
			File.Save()
		} else {
			// create succeed
			jiraIdLoc := jiraCol + lineNumStr
			_ = File.SetCellValue(getDefaultSheet(), jiraIdLoc, jiraId)
			_ = File.SetCellValue(getDefaultSheet(), resLoc, "success。")
			//fmt.Println(err)
			File.Save()
		}
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

func getColValue(rowContent []string, relation ColRelation, key string) string {
	if val, ok := relation[key]; ok {
		if val.Id == -1 {
			// not in excel
			return ""
		} else {
			// in excel
			// TODO: lin, index out of range [1] with length 1
			return rowContent[val.Id]
		}
	} else {
		// relation not found
		return ""
	}
}
