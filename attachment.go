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
	"io"

	req "github.com/imroc/req/v3"
)

type AttachmentService struct {
	Service
}

const ATTACHMENT_POST = "/rest/api/2/issue/{issueIdOrKey}/attachments"
const ATTACHMENT = "/secure/attachment/{id}/"

var configAttachment = map[string][2]string{
	//"AttachmentPost": {POST, "/rest/api/2/issue/{issueIdOrKey}/attachments", ""},
	"Get": {GET, "/rest/api/2/attachment/content/{id}"},
	"Del": {DEL, "/rest/api/2/attachment/{id}"},

	"Meta":      {GET, "/rest/api/2/attachment/meta"},
	"Thumbnail": {GET, "/rest/api/2/attachment/thumbnail/{id}"},
	"Metadata":  {GET, "/rest/api/2/attachment/{id}"},
	"Human":     {GET, "/rest/api/2/attachment/{id}/expand/human"},
	"Raw":       {GET, "/rest/api/2/attachment/{id}/expand/raw"},
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
func (AS *AttachmentService) AttachmentPostCtx(ctx context.Context, issueIdOrKey string,
	r io.Reader, attachmentName string, result interface{}) error {

	Path := Replace(ATTACHMENT_POST, Values{"issueIdOrKey": issueIdOrKey})
	u, err := AS.sd.Parse(Path)
	if err != nil {
		return err
	}

	client := req.C()
	res, err := client.R().
		SetHeader("X-Atlassian-Token", "no-check").
		SetBasicAuth(AS.sd.JiraUser(), AS.sd.JiraPass()).
		SetFileReader("file", attachmentName, r).
		SetContext(ctx).
		Post(u.String())

	return AS.sd.JsonDecode(res.Body, err, result)
}

func (AS *AttachmentService) AttachmentPost(issueIdOrKey string,
	r io.Reader, attachmentName string, result interface{}) error {
	return AS.AttachmentPostCtx(AS.ctx, issueIdOrKey, r, attachmentName, result)
}

func (AS *AttachmentService) DownloadAttachmentCtx(ctx context.Context, attachmentID string) (io.ReadCloser, error) {

	u, err := AS.sd.Parse(
		Replace(ATTACHMENT, Values{"attachmentID": attachmentID}))
	if err != nil {
		return nil, err
	}

	return AS.sd.ContextRequest(ctx, GET, u, nil)
}

func (AS *AttachmentService) DownloadAttachment(attachmentID string) (io.ReadCloser, error) {
	return AS.DownloadAttachmentCtx(AS.ctx, attachmentID)
}
