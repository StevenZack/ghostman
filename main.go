package main

import (
	"github.com/StevenZack/ghostman/logx"
	"github.com/StevenZack/ghostman/views"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	run()
}
func run() {
	w, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, nil)
	if e != nil {
		logx.Error(e)
		return
	}
	w.SetTitle("title")
	w.DefineFunction("showWin", func(args ...*sciter.Value) *sciter.Value {
		win, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, 
			sciter.NewRect(150, 50, 100, 100))
		if e != nil {
			logx.Error(e)
			return sciter.NullValue()
		}

		win.LoadHtml(views.Str_index, "")
		win.Show()
		return sciter.NullValue()
	})
	w.LoadHtml(views.Str_index, "")
	w.Show()
	w.Run()
}
