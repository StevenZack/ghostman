package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/StevenZack/tools/cryptoToolkit"

	"github.com/StevenZack/db"
	"github.com/StevenZack/ghostman/util"
	"github.com/StevenZack/ghostman/views"
	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	log.SetFlags(log.Lshortfile)
	db.InitDB(strToolkit.Getrpath(fileToolkit.GetHomeDir())+".ghostman", "pw")
	run(loadOldData())
}

func loadOldData() []string {
	oldData := db.Get("history.txt").Val()
	if oldData != "" {
		strs := []string{}
		e := json.Unmarshal([]byte(oldData), &strs)
		if e != nil {
			fmt.Println("unmarshal old data error :", e)
			return nil
		}

		if len(strs) < 3 {
			fmt.Println("old data len<3")
			return nil
		}

		return strs
	}

	return nil
}

func setOldData(method, url, cypher, header, body string) {
	db.Set("history.txt", []string{method, url, cypher, header, body})
}

func run(old []string) {
	w, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, sciter.NewRect(100, 150, 400, 600))
	if e != nil {
		log.Println(e)
		return
	}
	bindFuncs(w)
	w.SetTitle("Ghost Man")
	w.LoadHtml(views.Str_index, "")
	w.Show()

	if len(old) == 5 {
		method := "GET"
		if old[0] != "" {
			method = old[0]
		}
		url := "http://localhost:8080/"
		if old[1] != "" {
			url = old[1]
		}
		cypher := old[2]
		header := "{}"
		if old[3] != "" {
			header = old[3]
		}
		body := old[4]
		w.Call("setupData", sciter.NewValue(method), sciter.NewValue(url), sciter.NewValue(cypher), sciter.NewValue(header), sciter.NewValue(body))
	}
	w.Run()
}

func bindFuncs(w *window.Window) {
	w.DefineFunction("showErr", func(args ...*sciter.Value) *sciter.Value {
		w2, e := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_ENABLE_DEBUG, sciter.NewRect(200, 200, 300, 200))
		if e != nil {
			log.Println(e)
			return sciter.NullValue()
		}
		w2.SetTitle("错误")

		w2.LoadHtml(views.Str_error, "")
		w2.Show()
		w2.Call("showInfo", args[0])
		return sciter.NullValue()
	})

	w.DefineFunction("doReq", func(args ...*sciter.Value) *sciter.Value {
		if len(args) != 5 {
			log.Fatal("args != 5")
		}
		method := args[0].String()
		url := args[1].String()
		cypher := args[2].String()
		header, e := util.UnmarshalMap(args[3].String())
		if e != nil {
			log.Println(e)
			showErr(w, e.Error())
			return sciter.NullValue()
		}
		body := args[4].String()
		encBody := ""
		if cypher != "" {
			encBodyB := cryptoToolkit.Encrypt([]byte(body), cypher)
			encBody = string(encBodyB)
		}
		go func() {
			bodyToSend := body
			if cypher != "" {
				bodyToSend = encBody
			}
			status, rpheader, rpbody, e := util.DoReq(method, url, bodyToSend, header)
			if e != nil {
				log.Println(e)
				showErr(w, e.Error())
				return
			}

			if status != "200 OK" {
				showErr(w, rpbody)
				return
			}

			headerB, e := json.Marshal(header)
			if e != nil {
				showErr(w, e.Error())
				return
			}

			go setOldData(method, url, cypher, string(headerB), body)
			w.Call("response", sciter.NewValue(status), sciter.NewValue(util.MarshalMap(rpheader)), sciter.NewValue(rpbody))
		}()
		return sciter.NullValue()
	})
}

func showErr(w *window.Window, e string) {
	w.Call("showErr", sciter.NewValue(e))
}
