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
| | | |___ ___ _ _
| |_| (_-</ -_) '_|
 \___//__/\___|_|

figlet -f small user
*/

const USER_ASSIGNABLE_MULTIPROJECTSEARCH = "/rest/api/2/user/assignable/multiProjectSearch"
const USER_ASSIGNABLE_SEARCH = "/rest/api/2/user/assignable/search"
const USER_PERMISSION_SEARCH = "/rest/api/2/user/permission/search"
const USER_PICKER = "/rest/api/2/user/picker"
const USER_SEARCH = "/rest/api/2/user/search"
const USER_SEARCH_QUERY = "/rest/api/2/user/search/query"
const USER_SEARCH_QUERY_KEY = "/rest/api/2/user/search/query/key"
const USER_VIEWISSUE_SEARCH = "/rest/api/2/user_/viewissue/search"
const USER = "/rest/api/2/user"
const USERS = "/rest/api/2/users"
const USER_BULK = "/rest/api/2/user/bulk"
const USER_MIGRATION = "/rest/api/2/user/bulk"
const USER_COLUMNS = "/rest/api/2/user/columns"
const USER_EMAIL = "/rest/api/2/user/email"
const USER_EMAIL_BULK = "/rest/api/2/user/email/bulk"
const USER_GROUPS = "/rest/api/2/user/groups"

type UserService struct {
	Service
}

func (SD *SD) UserService() *UserService {
	IS := Service{
		ctx: context.Background(), sd: SD}
	IS.Operation = map[string]ContextReq{
		"User":       SD.CReq(GET, USER),
		"Users":      SD.CReq(GET, USERS),
		"Create":     SD.CReq(POST, USER),
		"Del":        SD.CReq(DEL, USER),
		"Bulk":       SD.CReq(GET, USER_BULK),
		"Migration":  SD.CReq(GET, USER_MIGRATION),
		"Columns":    SD.CReq(GET, USER_COLUMNS),
		"ColumnsPut": SD.CReq(PUT, USER_COLUMNS),
		"ColumnsDel": SD.CReq(DEL, USER_COLUMNS),
		"Email":      SD.CReq(GET, USER_EMAIL),
		"EmailBulk":  SD.CReq(GET, USER_EMAIL_BULK),
		"Groups":     SD.CReq(GET, USER_GROUPS),
		"Search":     SD.CReq(POST, USER_SEARCH),
	}
	return &UserService{IS}
}
