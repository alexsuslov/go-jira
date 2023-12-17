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

import "context"

/**
   _           _
  /_\__ ____ _| |_ __ _ _ _
 / _ \ V / _` |  _/ _` | '_|
/_/ \_\_/\__,_|\__\__,_|_|

*/

const AVATAR = "/rest/api/2/applicationrole/{key}"
const AVATARS = "/rest/api/2/universal_avatar/type/{type}/owner/{entityId}"
const LOAD_AVATAR = "/rest/api/2/universal_avatar/type/{type}/owner/{entityId}"
const DEL_AVATAR = "/rest/api/2/universal_avatar/type/{type}/owner/{owningObjectId}/avatar/{id}"
const IMAGE = "/rest/api/2/universal_avatar/view/type/{type}"
const IMAGE_ID = "/rest/api/2/universal_avatar/view/type/{type}/avatar/{id}"
const IMAGE_OWNER = "/rest/api/2/universal_avatar/view/type/{type}"

type AvatarService struct {
	Service
}

func (SD *SD) AvatarService() *AvatarService {
	IS := Service{
		ctx: context.Background(), sd: SD}
	IS.Operation = map[string]ContextReq{
		"Get":          SD.CReq(GET, AVATAR),
		"Avatars":      SD.CReq(GET, AVATARS),
		"Load":         SD.CReq(POST, LOAD_AVATAR),
		"Del":          SD.CReq(DEL, DEL_AVATAR),
		"Image":        SD.CReq(GET, IMAGE),
		"ImageByID":    SD.CReq(GET, IMAGE_ID),
		"ImageByOwner": SD.CReq(GET, IMAGE_OWNER),
	}
	return &AvatarService{IS}
}
