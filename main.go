package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/mandoc.css", handleStylesheet)
	r.Handle(`/{page:\w+}`, handleWithErrors(handleSearch))
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
	sort.Strings(entries)
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
	fname := filepath.Join("/usr/share/man", "man"+vars["section"], vars["page"]+"."+vars["section"])
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// We pipe the mandoc output directly to the ResponseWriter, so after it returns we can't
	// return an error and change the status code because we've already sent a 200.
	if err := renderMandoc(req.Context(), "/usr/share/man", w, f); err != nil {
		log.Println(err)
	}
	return nil
}

func handleSearch(w http.ResponseWriter, req *http.Request) error {
	query := mux.Vars(req)["page"]
	mandir, err := os.Open("/usr/share/man")
	if err != nil {
		return err
	}
	defer mandir.Close()
	dirs, err := mandir.Readdir(-1)
	if err != nil {
		return err
	}
	for _, d := range dirs {
		if d.IsDir() && strings.HasPrefix(d.Name(), "man") {
			section := strings.TrimPrefix(d.Name(), "man")
			_, err := os.Stat(filepath.Join("/usr/share/man", d.Name(), query+"."+section))
			if err == nil {
				http.Redirect(w, req, "/"+section+"/"+query, http.StatusFound)
				return nil
			}
			if !os.IsNotExist(err) {
				return err
			}
		}
	}
	return os.ErrNotExist
}
