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

/**
 __  __     _   _            _
|  \/  |___| |_| |_  ___  __| |
| |\/| / -_)  _| ' \/ _ \/ _` |
|_|  |_\___|\__|_||_\___/\__,_|

*/

const GET = "GET"
const POST = "POST"
const PUT = "PUT"
const DEL = "DEL"

type Head map[string]string

var JSON = Head{"Content-Type": "application/json"}

//var NO_CHECK = Head{"X-Atlassian-Token": "no-check"}
//var headers = map[string]Head{"json": JSON}
