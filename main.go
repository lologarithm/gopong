package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/lologarithm/gopong/server"
)

var (
	addr      = flag.String("addr", ":8080", "http service address")
	assets    = flag.String("assets", defaultAssetPath(), "path to assets")
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
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "client/build/test.html")))

	var h = server.NewHub()

	// Setup our hub's routine
	go h.Run()

	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { server.HandleWebSocket(w, r, h) })

	// Log any fatal serve errors
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
