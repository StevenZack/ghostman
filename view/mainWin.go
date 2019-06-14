package view

import (
	. "github.com/gofaith/faithtop.libui"
)

type MainModel struct {
	MethodCombo          *FCombo
	URLEdit              *FEdit
	HeaderEdit, BodyEdit *FTextArea
	SendButton           *FButton
	RpText               *FTextArea
}

func CreateMainWin(model *MainModel) *FWin {
	model.MethodCombo = Combo().Append("GET").Append("POST").Append("DELETE").Append("PATCH").Append("PUT").Select(0)
	model.URLEdit = Edit().Expand().Text("http://")
	model.SendButton = Button().Text("SEND")
	model.HeaderEdit = TextArea().Text("Content-Type:application/json\nToken:123")
	model.BodyEdit = TextArea()
	model.RpText = TextArea()
	return Win(600, 600).Title("Ghostman").DeferShow().VBox(
		HBox().Append(
			model.MethodCombo,
			Text("URL:"),
			model.URLEdit,
			model.SendButton,
		),
		Tab().Append("RequestHeader", VBox().Append(
			model.HeaderEdit,
		)).Append("RequestBody", VBox().Append(
			model.BodyEdit,
		)).Append("Response", VBox().Append(
			model.RpText,
		)).Expand(),
	)
}
