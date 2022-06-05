package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gobuffalo/packr/v2"
	"github.com/va1da5/simple-emoji-picker/webview"
)

func getRandomPort() string {
	min := 10000
	max := 65535
	return fmt.Sprintf(":%d", rand.Intn(max-min)+min)
}

func main() {
	port := getRandomPort()
	debug := false

	if debug == true {
		fmt.Printf("Listening on  http://localhost%s/\n", port)
	}

	// Bind folder path for packaging with Packr
	folder := packr.New("static", "./gui/build")

	// Handle to ./static/build folder on root path
	http.Handle("/", http.FileServer(folder))

	// Run server at port 8000 as goroutine
	// for non-block working
	go http.ListenAndServe(port, nil)

	// Create new webview

	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("🍭 Emoji Picker")
	w.SetSize(344, 400, webview.HintFixed)
	w.Navigate(fmt.Sprintf("http://localhost%s", port))
	w.Run()
}
