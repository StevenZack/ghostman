package util

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func DoReq(method, url, body string, headers map[string]string) (string, map[string]string, string, error) {
	reader := strings.NewReader(body)
	r, e := http.NewRequest(method, url, reader)
	if e != nil {
		return "", nil, "", e
	}

	for k, v := range headers {
		r.Header.Set(k, v)
	}

	c := http.Client{}
	rp, e := c.Do(r)
	if e != nil {
		return "", nil, "", e
	}
	defer rp.Body.Close()

	b, e := ioutil.ReadAll(rp.Body)
	if e != nil {
		return "", nil, "", e
	}

	status := rp.Status
	rpheader := make(map[string]string)
	for k, v := range rp.Header {
		if len(v) == 0 {
			continue
		}
		rpheader[k] = v[0]
	}
	rpbody := string(b)

	return status, rpheader, rpbody, nil
}
