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
	w.LoadHtml(views.Str_index, "")
	w.Show()
	w.Run()
}
