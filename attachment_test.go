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
	"os"
	"testing"

	"github.com/andygrunwald/go-jira"
)

func TestAttachmentService_AttachmentPostCtx(t *testing.T) {
	sd := SD{}
	sd.SetDebug(true)
	AS := sd.AttachmentService()

	f, err := os.Open("data/1.png")
	if err != nil {
		panic(err)
	}

	type args struct {
		ctx            context.Context
		issueIdOrKey   string
		r              io.Reader
		attachmentName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"upload file to 149429",

			args{
				context.Background(),
				"149429",
				f,
				"1.png",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := []jira.Attachment{}
			if err := AS.AttachmentPostCtx(tt.args.ctx, tt.args.issueIdOrKey, tt.args.r, tt.args.attachmentName, result); (err != nil) != tt.wantErr {
				t.Errorf("AttachmentPostCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
