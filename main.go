package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	addr      = flag.String("addr", ":8080", "http service address")
	assets    = flag.String("assets", defaultAssetPath(), "path to assets")
	game      = Manager
	homeTempl *template.Template
)

func defaultAssetPath() string {
	p, err := build.Default.Import("/", "", build.FindOnly)
	if err != nil {
		return "."
	}
	return p.Dir
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func main() {
	flag.Parse()

	// Setup our template html
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "test.html")))

	// Setup our hub's routine
	go h.run()

	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", wsHandler)

	// Log any fatal serve errors
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
