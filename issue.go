// (c) 2023 Alex Suslov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package jira

import (
	"context"
	"fmt"

	gojira "github.com/andygrunwald/go-jira"
)

/**
 ___
|_ _|_______  _ ___
 | |(_-<_-< || / -_)
|___/__/__/\_,_\___|

*/

const ISSUES = "/rest/api/2/issue"
const ISSUE = "/rest/api/2/issue/{issueIdOrKey}"
const ISSUE_ARCHIVE = "/rest/api/2/issue/archive"
const ISSUE_BULK = "/rest/api/2/issue/bulk"
const ISSUE_CREATEMETA = "/rest/api/2/issue/createmeta"
const ISSUE_CREATEMETA_ISSUETYPES = "/rest/api/2/issue/createmeta/{projectIdOrKey}/issuetypes"
const ISSUE_CREATEMETA_ISSUETYPE = "/rest/api/2/issue/createmeta/{projectIdOrKey}/issuetypes/{issueTypeId}"
const ISSUE_UNARCHIVE = "/rest/api/2/issue/unarchive"
const ISSUE_ASSIGNEE = "/rest/api/2/issue/{issueIdOrKey}/assignee"
const ISSUE_CHANGELOG = "/rest/api/2/issue/{issueIdOrKey}/changelog"
const ISSUE_CHANGELOG_LIST = "/rest/api/2/issue/{issueIdOrKey}/changelog/list"
const ISSUE_EDIT_META = "/rest/api/2/issue/{issueIdOrKey}/editmeta"
const ISSUE_NOTIFY = "/rest/api/2/issue/{issueIdOrKey}/notify"
const ISSUE_TRANSITIONS = "/rest/api/2/issue/{issueIdOrKey}/transitions"
const ISSUES_ARCHIVE_EXPORT = "/rest/api/2/issues/archive/export"

var configIssue = map[string][2]string{
	"Issue":            {GET, ISSUE},
	"Create":           {POST, ISSUES},
	"ArchiveByID":      {PUT, ISSUE_ARCHIVE},
	"ArchiveJQL":       {POST, ISSUE_ARCHIVE},
	"CreateBulk":       {POST, ISSUE_BULK},
	"Meta":             {GET, ISSUE_CREATEMETA},
	"ProjectMetaTypes": {GET, ISSUE_CREATEMETA_ISSUETYPES},
	"ProjectMetaType":  {GET, ISSUE_CREATEMETA_ISSUETYPE},
	"UnArchive":        {PUT, ISSUE_UNARCHIVE},
	"Edit":             {PUT, ISSUE},
	"Del":              {DEL, ISSUE},
	"Assign":           {PUT, ISSUE_ASSIGNEE},
	"Changelog":        {GET, ISSUE_CHANGELOG},
	"ChangelogID":      {POST, ISSUE_CHANGELOG_LIST},
	"EditMeta":         {GET, ISSUE_EDIT_META},
	"Notification":     {POST, ISSUE_NOTIFY},
	"Transitions":      {GET, ISSUE_TRANSITIONS},
	"DoTransitions":    {POST, ISSUE_TRANSITIONS},
	"Export":           {PUT, ISSUES_ARCHIVE_EXPORT},
}

type IssueService struct {
	Service
}

func (SD *SD) IssueService() *IssueService {
	IS := Service{
		ctx: context.Background(), sd: SD, Operation: map[string]ContextReq{}}
	for k, v := range configDashboard {
		IS.Operation[k] = SD.CReq(v[0], v[1])
	}

	return &IssueService{IS}
}

func (I *IssueService) ContextCreate(ctx context.Context, NewIssue *gojira.Issue, result interface{}) error {
	if _, ok := I.Operation["Create"]; !ok {
		return fmt.Errorf("no operation")
	}
	return I.Operation["Create"](ctx, nil, NewIssue, result)
}

func (I *IssueService) Create(NewIssue *gojira.Issue, result interface{}) error {
	return I.ContextCreate(I.ctx, NewIssue, result)
}

func (I *IssueService) ContextIssue(ctx context.Context, issueIdOrKey string, result interface{}) error {
	if _, ok := I.Operation["Issue"]; !ok {
		return fmt.Errorf("no operation")
	}
	return I.Operation["Issue"](ctx, Values{"issueIdOrKey": issueIdOrKey}, nil, result)
}

func (I *IssueService) Issue(issueIdOrKey string, result interface{}) error {
	return I.ContextIssue(I.ctx, issueIdOrKey, result)
}

func (I *IssueService) ContextTransitions(ctx context.Context, issueIdOrKey string, result interface{}) error {
	if _, ok := I.Operation["Transitions"]; !ok {
		return fmt.Errorf("no operation")
	}
	return I.Operation["Transitions"](ctx, Values{"issueIdOrKey": issueIdOrKey}, nil, result)
}

func (I *IssueService) Transitions(issueIdOrKey string, result interface{}) error {
	return I.ContextTransitions(I.ctx, issueIdOrKey, result)
}

func (I *IssueService) ContextDoTransitions(ctx context.Context, issueIdOrKey, transitionID string, result interface{}) error {
	if _, ok := I.Operation["DoTransitions"]; !ok {
		return fmt.Errorf("no operation")
	}
	payload := gojira.CreateTransitionPayload{
		Transition: gojira.TransitionPayload{
			ID: transitionID}}

	return I.Operation["DoTransitions"](ctx, Values{"issueIdOrKey": issueIdOrKey}, payload, result)
}
