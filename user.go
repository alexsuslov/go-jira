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
)

/**
 _   _
| | | |___  ___ _ __ ___
| | | / __|/ _ \ '__/ __|
| |_| \__ \  __/ |  \__ \
 \___/|___/\___|_|  |___/

*/

const USER_ASSIGNABLE_MULTIPROJECTSEARCH = "/rest/api/2/user/assignable/multiProjectSearch"
const USER_ASSIGNABLE_SEARCH = "/rest/api/2/user/assignable/search"
const USER_PERMISSION_SEARCH = "/rest/api/2/user/permission/search"
const USER_PICKER = "/rest/api/2/user/picker"
const USER_SEARCH = "/rest/api/2/user/search"
const USER_SEARCH_QUERY = "/rest/api/2/user/search/query"
const USER_SEARCH_QUERY_KEY = "/rest/api/2/user/search/query/key"
const USER_VIEWISSUE_SEARCH = "/rest/api/2/user_/viewissue/search"

type UserService struct {
	Service
}

func (SD *SD) UserService() *UserService {
	IS := Service{
		ctx: context.Background(), sd: SD}
	IS.Operation = map[string]ContextReq{
		"Search": SD.CReq(POST, USER_SEARCH),
	}
	return &UserService{IS}
}
