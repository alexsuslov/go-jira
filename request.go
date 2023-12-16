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
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

func CloseErrLog(closer io.Closer) {
	if err := closer.Close(); err != nil {
		logrus.Error(err)
	}
}

func (SD *SD) ContextRequest(ctx context.Context,
	Method string, URL url.URL, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, Method, URL.String(), io.NopCloser(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(SD.JiraUser(), SD.JiraPass())

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 && res.StatusCode >= 400 {
		data, _ := io.ReadAll(res.Body)
		defer CloseErrLog(res.Body)

		return nil, fmt.Errorf("%d:%s", res.StatusCode, string(data))
	}

	return res, nil
}
