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
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

/**
 ___ ___
/ __|   \
\__ \ |) |
|___/___/

*/

type SD struct {
	host, user, pass *string
	debug            bool
}

type Service struct {
	ctx       context.Context
	sd        *SD
	Operation map[string]ContextReq
}

type Values map[string]string

func Replace(src string, values Values) string {
	for k, v := range values {
		src = strings.ReplaceAll(src, fmt.Sprintf("{%s}", k), v)
	}
	return src
}
func (SD *SD) Parse(s string) (*url.URL, error) {

	return url.Parse(SD.JiraHost() + s)
}

func (SD *SD) Debug(v bool) {
	SD.debug = v
}

func (SD *SD) JiraHost() string {
	if SD.host != nil {
		return *SD.host
	}
	return os.Getenv("JIRA_HOST")
}

func (SD *SD) JiraUser() string {
	if SD.user != nil {
		return *SD.user
	}
	return os.Getenv("JIRA_USER")
}
func (SD *SD) JiraPass() string {
	if SD.pass != nil {
		return *SD.pass
	}
	return os.Getenv("JIRA_PASS")
}

func (SD *SD) SetJiraHost(s string) *SD {
	SD.host = &s
	return SD
}

func (SD *SD) SetJiraUser(s string) *SD {
	SD.user = &s
	return SD
}

func (SD *SD) SetJiraPass(s string) *SD {
	SD.pass = &s
	return SD
}

func (SD *SD) JsonDecode(Body io.ReadCloser, err error, result interface{}) error {
	if err != nil {
		return err
	}
	defer CloseErrLog(Body)
	if SD.debug {
		data, err := io.ReadAll(Body)
		if err != nil {
			return err
		}
		body := string(data)
		logrus.
			WithField("body", body).
			Info("body")
		return json.Unmarshal(data, result)
	}
	return json.NewDecoder(Body).Decode(result)
}
