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
	"encoding/json"
	"net/url"
)

type UserService struct {
	ctx context.Context
	sd  *SD
}

func (SD *SD) UserService() *IssueService {
	return &IssueService{
		context.Background(), SD}
}

func (IS *IssueService) Search(value url.Values, result interface{}) error {
	return IS.ContextSearch(IS.ctx, value, result)
}

func (IS *IssueService) ContextSearch(ctx context.Context, values url.Values, result interface{}) error {
	u, err := IS.sd.Parse(USER_SEARCH)
	if err != nil {
		return err
	}

	u.RawQuery = values.Encode()

	res, err := IS.sd.ContextRequest(ctx, POST, *u, nil)
	if err != nil {
		return err
	}
	defer CloseErrLog(res.Body)
	return json.NewDecoder(res.Body).Decode(result)
}
