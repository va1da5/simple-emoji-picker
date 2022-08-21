package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/va1da5/simple-emoji-picker/webview"
)

//go:embed gui/build
var content embed.FS

func fsHandler() http.Handler {
	sub, err := fs.Sub(content, "gui/build")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(sub))
}

func main() {
	addr := "localhost:44665"
	debug := false

	if debug == true {
		fmt.Printf("Listening on  http://%s/\n", addr)
	}

	go http.ListenAndServe(addr, fsHandler())

	// Create new webview

	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("🍭 Emoji Picker")
	w.SetSize(344, 400, webview.HintFixed)
	w.Navigate(fmt.Sprintf("http://%s", addr))
	w.Run()
}
