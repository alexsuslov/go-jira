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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type IssueService struct {
	ctx context.Context
	sd  *SD
}

func (SD *SD) IssueService() *IssueService {

	return &IssueService{
		context.Background(), SD}
}

func (IS *IssueService) Create(Issue interface{}, result interface{}) error {
	return IS.ContextCreate(IS.ctx, Issue, result)
}

func (IS *IssueService) ContextCreate(ctx context.Context, Issue interface{}, result interface{}) error {
	u, err := IS.sd.Parse(ISSUE)
	if err != nil {
		return err
	}

	data, err := json.Marshal(Issue)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)
	res, err := IS.sd.ContextRequest(ctx, POST, *u, buf)
	if err != nil {
		return err
	}
	defer CloseErrLog(res.Body)
	return json.NewDecoder(res.Body).Decode(result)
}

func (IS *IssueService) Issue(IssueQ string, result interface{}) error {
	return IS.ContextIssue(IS.ctx, IssueQ, result)
}

func (IS *IssueService) ContextIssue(ctx context.Context,
	IssueQ string, result interface{}) error {
	u, err := IS.sd.Parse(
		fmt.Sprintf("%s/%s", ISSUE, IssueQ))
	if err != nil {
		return err
	}
	res, err := IS.sd.ContextRequest(ctx, GET, *u, nil)
	if err != nil {
		return err
	}
	defer CloseErrLog(res.Body)
	return json.NewDecoder(res.Body).Decode(result)
}

func (IS *IssueService) Transitions(issueIdOrKey string, result interface{}) error {
	return IS.ContextTransitions(IS.ctx, issueIdOrKey, result)
}

func (IS *IssueService) ContextTransitions(ctx context.Context,
	issueIdOrKey string, result interface{}) error {

	u, err := IS.sd.Parse(
		Replace(ISSUE_TRANSITIONS, Values{"issueIdOrKey": issueIdOrKey}))
	if err != nil {
		return err
	}

	res, err := IS.sd.ContextRequest(ctx, GET, *u, nil)
	if err != nil {
		return err
	}
	defer CloseErrLog(res.Body)
	return json.NewDecoder(res.Body).Decode(result)
}

func (IS *IssueService) DoTransitions(issueIdOrKey, transitionID string, result interface{}) error {
	return IS.ContextDoTransitions(IS.ctx, issueIdOrKey, transitionID, result)
}

// TransitionPayloadCommentBody represents body of comment in payload
type TransitionPayloadCommentBody struct {
	Body string `json:"body,omitempty"`
}

// TransitionPayloadCommentBody represents body of comment in payload
type TransitionPayloadComment struct {
	Add TransitionPayloadCommentBody `json:"add,omitempty" structs:"add,omitempty"`
}

// TransitionPayloadUpdate represents the updates of Transition calls like DoTransition
type TransitionPayloadUpdate struct {
	Comment []TransitionPayloadComment `json:"comment,omitempty" structs:"comment,omitempty"`
}

// TransitionPayload represents the request payload of Transition calls like DoTransition
type TransitionPayload struct {
	ID string `json:"id" structs:"id"`
}

// Resolution represents a resolution of a Jira issue.
// Typical types are "Fixed", "Suspended", "Won't Fix", ...
type Resolution struct {
	Self        string `json:"self" structs:"self"`
	ID          string `json:"id" structs:"id"`
	Description string `json:"description" structs:"description"`
	Name        string `json:"name" structs:"name"`
}

// TransitionPayloadFields represents the fields that can be set when executing a transition
type TransitionPayloadFields struct {
	Resolution *Resolution `json:"resolution,omitempty" structs:"resolution,omitempty"`
}

type CreateTransitionPayload struct {
	Update     TransitionPayloadUpdate `json:"update,omitempty" structs:"update,omitempty"`
	Transition TransitionPayload       `json:"transition" structs:"transition"`
	Fields     TransitionPayloadFields `json:"fields" structs:"fields"`
}

func (IS *IssueService) ContextDoTransitions(ctx context.Context,
	issueIdOrKey string, transitionID string, result interface{}) error {

	u, err := IS.sd.Parse(
		Replace(ISSUE_TRANSITIONS, Values{"issueIdOrKey": issueIdOrKey}))
	if err != nil {
		return err
	}

	payload := CreateTransitionPayload{
		Transition: TransitionPayload{
			ID: transitionID,
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)

	res, err := IS.sd.ContextRequest(ctx, POST, *u, buf)
	if err != nil {
		return err
	}
	defer CloseErrLog(res.Body)
	return json.NewDecoder(res.Body).Decode(result)

}
