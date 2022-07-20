package main

import "jira/sheet"

func main() {
	defer sheet.Close()
	sheet.Start()
}
