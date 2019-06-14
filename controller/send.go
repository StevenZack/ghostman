package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/pretty"

	"github.com/StevenZack/ghostman/model"
	"github.com/StevenZack/ghostman/util"
	"github.com/StevenZack/ghostman/vars"
	"github.com/StevenZack/livedata"
)

func init() {
	livedata.Observe(vars.StateSend, stateSend)
}
func stateSend(v interface{}) {
	ss := v.([]string)
	method := ss[0]
	url := ss[1]
	header := ss[2]
	body := ss[3]
	bodyr := strings.NewReader(body)
	r, e := http.NewRequest(method, url, bodyr)
	if e != nil {
		model.SetStateSent(e)
		return
	}
	for _, h := range util.ParseHeader(header) {
		r.Header.Set(h.Key, h.Value)
	}
	c := http.Client{}
	rp, e := c.Do(r)
	if e != nil {
		model.SetStateSent(e)
		return
	}
	back := "Status:\n"
	back += rp.Status + "\n\nHeaders:\n"
	b, e := json.Marshal(rp.Header)
	if e != nil {
		model.SetStateSent(e)
		return
	}
	pb := pretty.Pretty(b)
	back += string(pb) + "\n\nBody:\n"

	defer rp.Body.Close()
	rpBody, e := ioutil.ReadAll(rp.Body)
	if e != nil {
		model.SetStateSent(e)
		return
	}
	if rp.Header.Get("Content-Type") == "application/json" {
		rpBody = pretty.Pretty(rpBody)
	}
	back += string(rpBody) + "\n\n"
	model.SetStateSent(back)
}
