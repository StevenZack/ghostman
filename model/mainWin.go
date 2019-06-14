package model

import (
	"github.com/StevenZack/ghostman/view"
	. "github.com/gofaith/faithtop.libui"
)

var (
	mainWin   *FWin
	mainModel view.MainModel
)

func ShowMainWin() {
	if mainWin != nil {
		RunOnUIThread(func() {
			mainWin.Show()
		})
		return
	}

	RunOnUIThread(func() {
		mainWin = view.CreateMainWin(&mainModel)
		mainWin.OnClose(func() bool {
			mainWin = nil
			return true
		}).QuitOnCloseClicked()

		mainModel.SendButton.OnClick(func() {
			method := mainModel.MethodCombo.GetSelectedString()
			url := mainModel.URLEdit.GetText()
			header := mainModel.HeaderEdit.GetText()
			body := mainModel.BodyEdit.GetText()
			SetStateSend(method, url, header, body)
		})

		mainWin.Show()
	})
}
