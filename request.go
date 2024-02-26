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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

func CloseErrLog(closer io.Closer) {
	if err := closer.Close(); err != nil {
		logrus.Error(err)
	}
}

func (SD *SD) ContextRequest(ctx context.Context,
	Method string, URL *url.URL, body io.Reader) (r io.ReadCloser, err error) {

	if SD.IsDebug() {
		data, err := io.ReadAll(body)
		if err != nil {
			return nil, err
		}
		logrus.
			WithField("Method", Method).
			WithField("URL", URL.String()).
			Debug(string(data))
		buf := bytes.NewBuffer(data)
		body = io.NopCloser(buf)
	}

	client := req.C()
	Request := client.R().
		SetBasicAuth(SD.JiraUser(), SD.JiraPass())
	if body != nil {
		Request.
			SetHeader("Content-Type", "application/json").
			SetBody(body)
	}
	u := URL.String()
	var res *req.Response
	switch Method {
	case GET:
		res, err = Request.Get(u)
	case POST:
		res, err = Request.Post(u)
	case PUT:
		res, err = Request.Put(u)
	case DEL:
		res, err = Request.Delete(u)
	default:
		return nil, fmt.Errorf("unknown method: %v", Method)
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		data, _ := io.ReadAll(res.Body)
		defer CloseErrLog(res.Body)

		return nil, fmt.Errorf("%d:%s", res.StatusCode, string(data))
	}

	return res.Body, nil
}

func (SD *SD) ContextRequest1(ctx context.Context,
	Method string, URL *url.URL, body io.Reader) (io.ReadCloser, error) {

	u := URL.String()
	req, err := http.NewRequestWithContext(ctx, Method, u, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(SD.JiraUser(), SD.JiraPass())

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		data, _ := io.ReadAll(res.Body)
		defer CloseErrLog(res.Body)

		return nil, fmt.Errorf("%d:%s", res.StatusCode, string(data))
	}

	return res.Body, nil
}

type ContextReq func(ctx context.Context,
	values Values,
	req interface{}) (io.ReadCloser, error)

func (SD *SD) CReq(Method, Path string) ContextReq {
	return func(ctx context.Context, values Values,
		req interface{}) (io.ReadCloser, error) {

		p := Path
		if values != nil {
			p = Replace(p, values)
		}

		u, err := SD.Parse(p)
		if err != nil {
			return nil, err
		}

		if req != nil {
			data, err := json.Marshal(req)
			if err != nil {
				return nil, err
			}
			buf := bytes.NewBuffer(data)
			return SD.ContextRequest(ctx, Method, u, buf)
		}

		return SD.ContextRequest(ctx, Method, u, nil)
	}
}
