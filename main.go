package main

import (
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
	w.DefineFunction("showErr", func(args ...*sciter.Value) *sciter.Value {
		w2, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, sciter.NewRect(200, 200, 300, 200))
		if e != nil {
			logx.Error(e)
			return sciter.NullValue()
		}
		w2.SetTitle("错误")

		w2.LoadHtml(views.Str_error, "")
		w2.Show()
		w2.Call("showInfo", args[0])
		return sciter.NullValue()
	})

	w.DefineFunction("doReq", func(args ...*sciter.Value) *sciter.Value {
		if len(args) != 4 {
			panic("args.len!=4")
		}
		method := args[0].String()
		url := args[1].String()
		header, e := util.UnmarshalMap(args[2].String())
		if e != nil {
			logx.Error(e)
			showErr(w, e)
			return sciter.NullValue()
		}
		body := args[3].String()
		go func() {
			status, rpheader, rpbody, e := util.DoReq(method, url, body, header)
			if e != nil {
				logx.Error(e)
				showErr(w, e)
				return
			}

			w.Call("response", sciter.NewValue(status), sciter.NewValue(util.MarshalMap(rpheader)), sciter.NewValue(rpbody))
		}()
		return sciter.NullValue()
	})
}

func showErr(w *window.Window, e error) {
	w.Eval(`view.showErr('` + e.Error() + `')`)
}
