// MIT License

// Copyright (c) 2017 FLYING

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package settings

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

var (
	Listen           = ":8080"
	HmacSampleSecret = []byte("whatever")
	LogFile          = "/var/log/beauty/beauty.log"
	Domain           = "xxxx.com"
	DefaultOrigin    = "http://origin.com"
	Local            = map[string]string{}
)

func InitLocal(jsonPath string) {
	if strings.TrimSpace(jsonPath) == "" {
		jsonPath = "/srv/filestore/settings/latest.json"
	}

	bytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		// log.Println("ReadFile: ", err.Error())
		log.Println("ReadFile:", "read the setting json file failed.")
		return
	}

	if err := json.Unmarshal(bytes, &Local); err != nil {
		log.Println("Unmarshal:", err.Error())
		return
	}

	log.Println("ReadFile:", "read the setting json file done or updated.")
}

func init() {
	InitLocal("")
}
