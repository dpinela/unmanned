package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/browser"
)

func main() {
	var serveAt string
	flag.StringVar(&serveAt, "addr", "localhost:", "Serve at `host:port`")
	flag.Parse()
	srv, err := net.Listen("tcp", serveAt)
	if err != nil {
		log.Fatal(err)
	}
	if args := flag.Args(); len(args) > 0 {
		go func() {
			// The browser's connection is guaranteed to succeed, since we're already listening for connections at
			// this point.
			if err := browser.OpenURL("http://" + srv.Addr().String() + "/" + strings.Join(args, "/")); err != nil {
				log.Println(err)
			}
		}()
	}
	r := mux.NewRouter()
	r.HandleFunc("/mandoc.css", handleStylesheet)
	r.Handle(`/{page:\w+}`, handleWithErrors(handleSearch))
	r.Handle(`/{section:\d+}/{page:\w+}`, handleWithErrors(handleManpage))
	log.Printf("Serving manpages at %s. Press Ctrl-C to stop", srv.Addr())
	log.Println(http.Serve(srv, r))
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

func getManpageLocation(ctx context.Context, section, name string) (string, error) {
	var args []string
	if section != "" {
		args = []string{"-w", section, name}
	} else {
		args = []string{"-w", name}
	}
	path, err := exec.CommandContext(ctx, "man", args...).Output()
	// We assume that if `man -w` reports a failure, it's because it can't find the
	// manpage. We don't assume a specific exit status, because they might vary between
	// platforms.
	if _, ok := err.(*exec.ExitError); ok {
		return "", os.ErrNotExist
	}
	return strings.TrimSpace(string(path)), err
}

func handleManpage(w http.ResponseWriter, req *http.Request) error {
	vars := mux.Vars(req)
	return serveManpage(w, req, vars["section"], vars["page"])
}

func serveManpage(w http.ResponseWriter, req *http.Request, section, page string) error {
	fname, err := getManpageLocation(req.Context(), section, page)
	if err != nil {
		return err
	}
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// We pipe the mandoc output directly to the ResponseWriter, so after it returns we can't
	// return an error and change the status code because we've already sent a 200.
	if err := renderMandoc(req.Context(), filepath.Join(fname, "..", ".."), w, f); err != nil {
		log.Println(err)
	}
	return nil
}

func handleSearch(w http.ResponseWriter, req *http.Request) error {
	return serveManpage(w, req, "", mux.Vars(req)["page"])
}
