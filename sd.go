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
	"net/url"
	"os"
	"strings"
)

type SD struct{}

type Values map[string]string

func Replace(src string, values Values) string {
	for k, v := range values {
		src = strings.ReplaceAll(src, k, v)
	}
	return src
}
func (SD *SD) Parse(s string) (*url.URL, error) {
	return url.Parse(SD.JiraHost() + s)
}

func (SD *SD) JiraHost() string {

	return os.Getenv("JIRA_HOST")
}
func (SD *SD) JiraUser() string {

	return os.Getenv("JIRA_USER")
}
func (SD *SD) JiraPass() string {
	return os.Getenv("JIRA_PASS")
}
