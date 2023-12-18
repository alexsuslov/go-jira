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

import (
	"context"
)

/**
 ___          _    _                      _
|   \ __ _ __| |_ | |__  ___  __ _ _ _ __| |
| |) / _` (_-< ' \| '_ \/ _ \/ _` | '_/ _` |
|___/\__,_/__/_||_|_.__/\___/\__,_|_| \__,_|

*/

var configDashboard = map[string][2]string{
	"Get":           {GET, "/rest/api/2/dashboard"},
	"Create":        {POST, "/rest/api/2/dashboard"},
	"PUT":           {PUT, "/rest/api/2/dashboard/bulk/edit"},
	"Gadgets":       {GET, "/rest/api/2/dashboard/gadgets"},
	"Search":        {GET, "/rest/api/2/dashboard/search"},
	"GadgetsByID":   {GET, "/rest/api/2/dashboard/{dashboardId}/gadget"},
	"GadgetAdd":     {POST, "/rest/api/2/dashboard/{dashboardId}/gadget"},
	"GadgetPut":     {PUT, "/rest/api/2/dashboard/{dashboardId}/gadget/{gadgetId}"},
	"GadgetDel":     {DEL, "/rest/api/2/dashboard/{dashboardId}/gadget/{gadgetId}"},
	"PropertyKeys":  {GET, "/rest/api/2/dashboard/{dashboardId}/items/{itemId}/properties"},
	"Property":      {GET, "/rest/api/2/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}"},
	"PropertyPut":   {PUT, "/rest/api/2/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}"},
	"PropertyDel":   {DEL, "/rest/api/2/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey}"},
	"Dashboard":     {GET, "/rest/api/2/dashboard/{id}"},
	"DashboardPut":  {PUT, "/rest/api/2/dashboard/{id}"},
	"DashboardDel":  {DEL, "/rest/api/2/dashboard/{id}"},
	"DashboardCopy": {POST, "/rest/api/2/dashboard/{id}/copy"},
}

type DashboardService struct {
	Service
}

func (SD *SD) DashboardService() *DashboardService {
	IS := Service{
		ctx: context.Background(), sd: SD, Operation: map[string]ContextReq{}}

	for k, v := range configDashboard {
		IS.Operation[k] = SD.CReq(v[0], v[1])
	}

	return &DashboardService{IS}
}
