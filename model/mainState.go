package model

import (
	"github.com/StevenZack/ghostman/vars"
	"github.com/StevenZack/livedata"
	. "github.com/gofaith/faithtop.libui"
)

func init() {
	livedata.Observe(vars.StateSend, stateSend)
	livedata.Observe(vars.StateSent, stateSent)
}

// SetStateSend send
func SetStateSend(method, url, header, body string) {
	livedata.Set(vars.StateSend, []string{method, url, header, body})
}

func stateSend(v interface{}) {
	if mainWin == nil {
		return
	}
	RunOnUIThread(func() {
		mainModel.SendButton.Enabled(false)
	})
}

// SetStateSent sent
func SetStateSent(v interface{}) {
	livedata.Set(vars.StateSent, v)
}
func stateSent(v interface{}) {
	if mainWin == nil {
		return
	}

	e, ok := v.(error)
	if ok {
		RunOnUIThread(func() {
			mainModel.SendButton.Enabled(true)
			ShowErr(mainWin, "错误", e.Error())
		})
		return
	}
	s := v.(string)
	RunOnUIThread(func() {
		mainModel.SendButton.Enabled(true)
		mainModel.RpText.Text(s)
	})
}
