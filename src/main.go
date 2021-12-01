package main

import (
	"net/http"

	"mechanic/src/file"

	"github.com/lxn/walk"
	dec "github.com/lxn/walk/declarative"
)

func main() {
	var wv *walk.WebView
	var mw *walk.MainWindow
	ic, _ := walk.NewIconFromFile("./assets/favicon.ico")
	var html, err = assets.Asset("assets/index.html")
	if err == nil {
		go func () {
			mux := http.NewServeMux()
			mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
				w.Write(html)
			})
			http.ListenAndServe(":8889", mux)
		}()
	}
	dec.MainWindow{
		Icon: ic,
		AssignTo: &mw,
		Title: "Mechanic",
		MinSize: dec.Size{ Width: 600, Height: 400 },
		Size: dec.Size{ Width: 600, Height: 400 },
		Layout:  dec.VBox{MarginsZero: true},
		Children: []dec.Widget{
			dec.WebView{
				AssignTo: &wv,
				Name: "wv",
				URL: "http://127.0.0.1:8889",
			},
		},
	}.Run()
}
