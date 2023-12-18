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

package v0

import "context"

/**
 ___     _
| _ \___| |___ ___
|   / _ \ / -_|_-<
|_|_\___/_\___/__/

*/

const ROLE = "/rest/api/2/applicationrole"
const ROLE_KEY = "/rest/api/2/applicationrole/{key}"

type RoleService struct {
	Service
}

func (SD *SD) RoleService() *RoleService {
	IS := Service{
		ctx: context.Background(), sd: SD}
	IS.Operation = map[string]ContextReq{
		"Get": SD.CReq(GET, ROLE),
		"KEY": SD.CReq(GET, ROLE_KEY),
	}
	return &RoleService{IS}
}
