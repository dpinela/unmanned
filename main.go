package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mandoc.css", handleStylesheet)
	r.Handle(`/{section:\d+}`, handleWithErrors(handleSection))
	r.Handle(`/{section:\d+}/{page:\w+}`, handleWithErrors(handleManpage))
	log.Println(http.ListenAndServe("localhost:6006", r))
}

var manpath = []string{"/usr/share/man", "/usr/local/share/man"}

func handleStylesheet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	w.Write(defaultStylesheet)
}

func handleWithErrors(handler func(http.ResponseWriter, *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := handler(w, req)
		if err == nil {
			return
		}
		log.Println(err)
		if os.IsNotExist(err) {
			http.NotFound(w, req)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
	})
}

func handleSection(w http.ResponseWriter, req *http.Request) error {
	section := mux.Vars(req)["section"]
	f, err := os.Open(filepath.Join("/usr/share/man", "man"+section))
	if err != nil {
		return err
	}
	suffix := "." + section
	defer f.Close()
	entries, err := f.Readdirnames(-1)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<!doctype html><html><head><title>")
	io.WriteString(w, "Man section ")
	io.WriteString(w, section)
	io.WriteString(w, "</title></head><body><ul>")
	for _, name := range entries {
		if trimmed := strings.TrimSuffix(name, suffix); len(trimmed) < len(name) {
			io.WriteString(w, `<li><a href="/`+section+"/"+trimmed+`">`+trimmed+"</a></li>")
		}
	}
	io.WriteString(w, "</ul></body></html>")
	return nil
}

func handleManpage(w http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)
	f, err := os.Open(filepath.Join("/usr/share/man", "man"+vars["section"], vars["page"]+"."+vars["section"]))
	if err != nil {
		return err
	}
	defer f.Close()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return renderMandoc(req.Context(), w, f)
}
