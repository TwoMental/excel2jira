# intro
- Automatically creating Jira issue using data in Excel.
## support fields
- project
- summary
- description
- issue_type
- priority

# usage
## config
- `relation.json`: Jira field in Excel.
    ```json
    {
      "project": {                // field name
        "name": "project",        // col name in excel
        "defalut": "test issue"   // default value of this field
      },
      ...
    }
    ```
## run 
- run with go 
  - show help
      ```shell
      go run main.go -h
      ```
  - run with full args
      ```shell
      go run main.go \
        --url=http://atlassian.sample.com \
        --username=YourName \
        --password=ChangeMe \
        --relation=relation.json \ 
        --excel=sample.xlsx
      ```

- run with release
  - show help
      ```shell
      ./excel2jira -h
      ```
  - run with full args
      ```shell
      ./excel2jira \
        --url=http://atlassian.sample.com \
        --username=YourName \
        --password=ChangeMe \
        --relation=relation.json \ 
        --excel=sample.xlsx
      ```
## build
- for linux
    ```shell
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o excel2jira
    ```
- for mac
    ```shell
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o excel2jira
    ```
- for windows
    ```shell
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o excel2jira
    ```

# notes
## TODOs
- [ ] custom fields
- [ ] jira auth
- [ ] choose sheet
- [x] parse args

## pkg
### excel
- :star:12.2k [excelize](https://github.com/qax-os/excelize)
- :star:5.2k [xlsx](https://github.com/tealeg/xlsx)
### arg pars
- flag
- [argparse]()
- :star:2.3k [go-flags](https://github.com/jessevdk/go-flags)