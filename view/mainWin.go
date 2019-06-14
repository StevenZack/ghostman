package view

import (
	. "github.com/gofaith/faithtop.libui"
)

type MainModel struct {
	MethodCombo                   *FCombo
	URLEdit, HeaderEdit, BodyEdit *FEdit
	SendButton                    *FButton
	RpText                        *FText
}

func CreateMainWin(model *MainModel) *FWin {
	model.MethodCombo = Combo().Append("GET").Append("POST").Append("DELETE").Append("PATCH").Append("PUT").Select(0)
	model.URLEdit = Edit().Expand().Text("http://")
	model.SendButton = Button().Text("SEND")
	model.HeaderEdit = Edit()
	model.BodyEdit = Edit()
	model.RpText = Text("")
	return Win(600, 600).Title("Ghostman").DeferShow().VBox(
		HBox().Append(
			model.MethodCombo,
			Text("URL:"),
			model.URLEdit,
			model.SendButton,
		),
		Text("Headers:"),
		model.HeaderEdit,
		Text("Body:"),
		model.BodyEdit,
		Text("Response:"),
		model.RpText,
	)
}
