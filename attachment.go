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
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
)

type AttachmentService struct {
	Service
}

var configAttachment = map[string][2]string{
	"AttachmentPost": {POST, "/rest/api/3/issue/{issueIdOrKey}/attachments"},
	"Get":            {GET, "/rest/api/3/attachment/content/{id}"},
	"Del":            {DEL, "/rest/api/3/attachment/{id}"},

	"Meta":      {GET, "/rest/api/3/attachment/meta"},
	"Thumbnail": {GET, "/rest/api/3/attachment/thumbnail/{id}"},
	"Metadata":  {GET, "/rest/api/3/attachment/{id}"},
	"Human":     {GET, "/rest/api/3/attachment/{id}/expand/human"},
	"Raw":       {GET, "/rest/api/3/attachment/{id}/expand/raw"},
}

func (SD *SD) AttachmentService() *AttachmentService {
	IS := Service{
		ctx: context.Background(), sd: SD, Operation: map[string]ContextReq{}}
	for k, v := range configAttachment {
		IS.Operation[k] = SD.CReq(v[0], v[1])
	}
	return &AttachmentService{IS}
}

// AttachmentPostContext uploads r (io.Reader) as an attachment to a given issueIdOrKey
func (AS *AttachmentService) AttachmentPostContext(ctx context.Context, issueIdOrKey string,
	r io.Reader, attachmentName string, result interface{}) error {

	fn, ok := AS.Operation["AttachmentPost"]
	if !ok {
		return fmt.Errorf("no operation")
	}

	b := new(bytes.Buffer)
	writer := multipart.NewWriter(b)

	fw, err := writer.CreateFormFile("file", attachmentName)
	if err != nil {
		return err

	}
	if r != nil {
		if _, err = io.Copy(fw, r); err != nil {

			return err
		}
	}

	err = writer.Close()
	if err != nil {
		return err

	}

	return fn(ctx, Values{"issueIdOrKey": issueIdOrKey}, b, result)
}
