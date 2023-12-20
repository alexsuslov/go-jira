# Simple go-jira


Get issue
```
export JIRA_HOST=https://host.com
export JIRA_USER=user
export JIRA_PASS=pass

import (
    gojira "github.com/andygrunwald/go-jira"
)

SD:=jira.SD{}

SD.Debug(true)

IssueService:=SD.IssueService()

issue:=gojira.Issue{}
if err := IssueService.Issue("issueKey", &issue); err!=nil{
    panic(err)
}

SD.Debug(false)

```
