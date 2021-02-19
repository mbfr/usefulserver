package usefulserver

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/pkg/errors"
)

func GetServer(port string, dir string) *http.Server {
	t := templateServer{dir: dir}

	server := http.NewServeMux()
	server.HandleFunc("/", t.getPageWrap)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}

	return srv
}

type templateServer struct {
	dir string
}

func (t templateServer) getPageWrap(w http.ResponseWriter, r *http.Request) {
	content, err := getPage(r, t.dir)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	} else {
		_, _ = w.Write(content)
	}
}

func getPage(r *http.Request, dir string) ([]byte, error) {
	var content []byte
	var err error

	url := r.URL
	name := url.Query().Get("name")

	if name == "" {
		path := filepath.Join(dir, "who.html")
		content, err = ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.Errorf("unable to load file '%s'", path)
		}
	} else {
		tpl := template.New("html-template")
		path := filepath.Join(dir, "hello.html")

		var err error
		tpl, err = template.ParseFiles(path)
		if err != nil {
			return nil, errors.Errorf("could not load template '%s'", path)
		}

		var buf bytes.Buffer
		err = tpl.Execute(&buf, map[string]string{"Name": name})
		if err != nil {
			return nil, errors.New("error parsing template")
		}

		content = buf.Bytes()
	}

	return content, nil
}
