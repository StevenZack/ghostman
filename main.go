package main

import (
	"strings"

	"github.com/StevenZack/ghostman/logx"
	"github.com/StevenZack/ghostman/util"
	"github.com/StevenZack/ghostman/views"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	run()
}
func run() {
	w, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, sciter.NewRect(100, 150, 400, 600))
	if e != nil {
		logx.Error(e)
		return
	}
	bindFuncs(w)
	w.SetTitle("Ghost Man")
	w.LoadHtml(views.Str_index, "")
	w.Show()
	w.Run()
}

func bindFuncs(w *window.Window) {
	w.DefineFunction("doReq", func(args ...*sciter.Value) *sciter.Value {
		if len(args) != 4 {
			panic("args.len!=4")
		}
		method := args[0].String()
		url := args[1].String()
		header, e := util.UnmarshalMap(args[2].String())
		if e != nil {
			logx.Error(e)
			showErr(e)
			return sciter.NullValue()
		}
		body := args[3].String()
		go func() {
			status, rpheader, rpbody, e := util.DoReq(method, url, body, header)
			if e != nil {
				logx.Error(e)
				showErr(e)
				return
			}

			w.Call("response", sciter.NewValue(status), sciter.NewValue(util.MarshalMap(rpheader)), sciter.NewValue(rpbody))
		}()
		return sciter.NullValue()
	})
}

func showErr(e error) {
	w, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS, sciter.NewRect(120, 150, 100, 100))
	if e != nil {
		logx.Error(e)
		return
	}
	w.SetTitle("错误")

	str := strings.Replace(views.Str_error, `{{.}}`, e.Error(), -1)
	w.LoadHtml(str, "")
	w.Show()
}
