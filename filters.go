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

import "context"

var configFilter = map[string][2]string{
	"Create":       {POST, "/rest/api/3/filter"},
	"Favorite":     {GET, "/rest/api/3/filter/favourite"},
	"My":           {GET, "/rest/api/3/filter/my"},
	"Search":       {GET, "/rest/api/3/filter/search"},
	"Get":          {GET, "/rest/api/3/filter/{id}"},
	"Put":          {PUT, "/rest/api/3/filter/{id}"},
	"Del":          {DEL, "/rest/api/3/filter/{id}"},
	"Columns":      {GET, "/rest/api/3/filter/{id}/columns"},
	"ColumnSet":    {PUT, "/rest/api/3/filter/{id}/columns"},
	"ColumnReset":  {DEL, "/rest/api/3/filter/{id}/columns"},
	"FilterAdd":    {PUT, "/rest/api/3/filter/{id}/favourite"},
	"FilterRemove": {DEL, "/rest/api/3/filter/{id}/favourite"},
	"FilterOwner":  {PUT, "/rest/api/3/filter/{id}/owner"},
}

type FilterService struct {
	Service
}

func (SD *SD) FilterService() *FilterService {
	IS := Service{
		ctx: context.Background(), sd: SD, Operation: map[string]ContextReq{}}
	for k, v := range configFilter {
		IS.Operation[k] = SD.CReq(v[0], v[1])
	}

	return &FilterService{IS}
}
