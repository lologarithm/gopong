package main

import (
	"flag"
	"github.com/lologarithm/gopong/server"
	"go/build"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var (
	addr      = flag.String("addr", "localhost:8080", "http service address")
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
	homeTempl = template.Must(template.ParseFiles(filepath.Join(*assets, "index.html")))

	var h = server.NewHub()

	// Setup our hub's routine
	go h.Run()

	// Register handlers
	http.HandleFunc("/", defaultFileServe)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { server.HandleWebSocket(w, r, h) })

	http.ListenAndServe(*addr, http.DefaultServeMux)
	//http.ListenAndServe("localhost:8000", http.FileServer(http.Dir("client/build/")))
}

func defaultFileServe(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/index.html"
	}
	data, err := ioutil.ReadFile(filepath.Join("./	client/build", r.URL.Path))
	if err != nil {
		log.Printf("Err reading file: %s", err.Error())
		return
	}
	w.Header().Add("Content-Type", http.DetectContentType(data))
	w.Write(data)
}
