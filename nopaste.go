package nopaste

import (
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Nopaste struct {
	config *Config
}

func Run(configPath string) {
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("cannot load config file: %v\n", err)
		return
	}
	np := New(config)
	http.Handle("/", np)
	http.ListenAndServe(config.Listen, nil)
}

func New(config *Config) *Nopaste {
	return &Nopaste{
		config: config,
	}
}

func isAlnum(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9')
}

func (np *Nopaste) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	upath := req.URL.Path

	if upath == np.config.Root || upath == np.config.Root+"/" {
		np.rootHandler(w, req)
		return
	}

	if p := strings.TrimPrefix(upath, np.config.Root+"/"); len(p) < len(upath) {
		req.URL.Path = p
		if strings.IndexFunc(upath, isAlnum) != -1 {
			np.dataHandler(w, req)
			return
		}
	}

	http.NotFound(w, req)
}

func (np *Nopaste) rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		np.saveContent(w, req)
	} else if err := tmplRoot.ExecuteTemplate(w, "index", nil); err != nil {
		log.Println(err)
		serverError(w)
	}
}

func (np *Nopaste) saveContent(w http.ResponseWriter, req *http.Request) {
	data := []byte(req.FormValue("text"))
	if len(data) == 0 {
		http.Redirect(w, req, np.config.Root, http.StatusFound)
		return
	}

	hex := fmt.Sprintf("%x", sha1.Sum(data))
	id := hex[0:10]
	log.Printf("saving: %s\n", id)
	err := ioutil.WriteFile(np.dataFilePath(id), data, 0644)
	if err != nil {
		log.Println(err)
		serverError(w)
		return
	}
	http.Redirect(w, req, np.config.Root+"/"+id, http.StatusFound)
}

func (np *Nopaste) dataHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Path
	if len(id) >= 1 && id[0] == '/' {
		id = id[1:]
	}

	f, err := os.Open(np.dataFilePath(id))
	if err != nil {
		log.Println(err)
		http.NotFound(w, req)
		return
	}
	w.Header().Add("Content-Type", "text/plain")
	io.Copy(w, f)
}

func (np *Nopaste) dataFilePath(id string) string {
	return fmt.Sprintf("%s/%s.txt", np.config.DataDir, id)
}

func serverError(w http.ResponseWriter) {
	code := http.StatusInternalServerError
	http.Error(w, http.StatusText(code), code)
}
