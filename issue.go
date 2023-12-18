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

package v2

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

	"CommentsByID": {GET, "/rest/api/2/issue/{issueIdOrKey}/comment"},
	"CommentAdd":   {POST, "/rest/api/2/issue/{issueIdOrKey}/comment"},
}

type IssueService struct {
	Service
}

func (SD *SD) IssueService() *IssueService {
	IS := Service{
		ctx: context.Background(), sd: SD, Operation: map[string]ContextReq{}}
	for k, v := range configIssue {
		IS.Operation[k] = SD.CReq(v[0], v[1])
	}

	return &IssueService{IS}
}

// СommentsCtx Returns all comments for an issue.
//
// https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-issue-issueidorkey-comment-get
func (I *IssueService) CommentsCtx(ctx context.Context, issueIdOrKey string,
	result interface{}) error {
	fn, ok := I.Operation["CommentsByID"]
	if !ok {
		return fmt.Errorf("no operation")
	}
	res, err := fn(ctx, Values{"issueIdOrKey": issueIdOrKey}, nil)
	return I.sd.JsonDecode(res, err, result)
}

// Comments Returns all comments for an issue.
func (I *IssueService) Comments(ctx context.Context, issueIdOrKey string,
	result interface{}) error {
	return I.CommentsCtx(ctx, issueIdOrKey, result)
}

// CommentAddCtx adds a new comment to issueID.
//
// Jira API docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-issue-issueidorkey-comment-post
func (I *IssueService) CommentAddCtx(ctx context.Context, issueIdOrKey string,
	comment *gojira.Comment, result interface{}) error {

	fn, ok := I.Operation["Issue"]
	if !ok {
		return fmt.Errorf("no operation")
	}
	res, err := fn(ctx, Values{"issueIdOrKey": issueIdOrKey}, comment)
	return I.sd.JsonDecode(res, err, result)
}

func (I *IssueService) CommentAdd(issueIdOrKey string,
	comment *gojira.Comment, result interface{}) error {
	return I.CommentAddCtx(I.ctx, issueIdOrKey, comment, result)
}

func (I *IssueService) CreateCtx(ctx context.Context, NewIssue *gojira.Issue, result interface{}) error {
	if _, ok := I.Operation["Create"]; !ok {
		return fmt.Errorf("no operation")
	}
	res, err := I.Operation["Create"](ctx, nil, NewIssue)
	return I.sd.JsonDecode(res, err, result)
}

func (I *IssueService) Create(NewIssue *gojira.Issue, result interface{}) error {
	return I.CreateCtx(I.ctx, NewIssue, result)
}

func (I *IssueService) IssueCtx(ctx context.Context, issueIdOrKey string, result interface{}) error {
	fn, ok := I.Operation["Issue"]
	if !ok {
		return fmt.Errorf("no operation")
	}
	res, err := fn(ctx, Values{"issueIdOrKey": issueIdOrKey}, nil)
	return I.sd.JsonDecode(res, err, result)
}

func (I *IssueService) Issue(issueIdOrKey string, result interface{}) error {
	return I.IssueCtx(I.ctx, issueIdOrKey, result)
}

func (I *IssueService) TransitionsCtx(ctx context.Context, issueIdOrKey string, result interface{}) error {
	if _, ok := I.Operation["Transitions"]; !ok {
		return fmt.Errorf("no operation")
	}
	res, err := I.Operation["Transitions"](ctx, Values{"issueIdOrKey": issueIdOrKey}, nil)
	return I.sd.JsonDecode(res, err, result)
}

func (I *IssueService) Transitions(issueIdOrKey string, result interface{}) error {
	return I.TransitionsCtx(I.ctx, issueIdOrKey, result)
}

func (I *IssueService) DoTransitionCtx(ctx context.Context, issueIdOrKey, transitionID string, result interface{}) error {
	if _, ok := I.Operation["DoTransitions"]; !ok {
		return fmt.Errorf("no operation")
	}
	payload := gojira.CreateTransitionPayload{
		Transition: gojira.TransitionPayload{
			ID: transitionID}}

	res, err := I.Operation["DoTransitions"](ctx, Values{"issueIdOrKey": issueIdOrKey}, payload)
	return I.sd.JsonDecode(res, err, result)
}

func (I *IssueService) DoTransition(issueIdOrKey, transitionID string, result interface{}) error {
	return I.DoTransitionCtx(I.ctx, issueIdOrKey, transitionID, result)
}
