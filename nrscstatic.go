package nrscstatic

import (
	"bitbucket.org/tebeka/nrsc"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

type Handler interface{}

// NrscStatic returns a middleware handler (based on martini's Static middleware)
// that serves static files via nrsc (https://bitbucket.org/tebeka/nrsc)
func NrscStatic(path string) Handler {
	if err := nrsc.Initialize(); err != nil {
		panic("[NrscStatic] Unable to initialize nrsc: " + err.Error())
	}
	return func(res http.ResponseWriter, req *http.Request, log *log.Logger) {
		file := req.URL.Path

		// nrsc expects there not to be a leading slash
		if file[0] == '/' {
			file = file[1:]
		}

		f := nrsc.Get(file)
		if f == nil {
			return
		}

		rdr, err := f.Open()
		if err != nil {
			log.Println("[NrscStatic] Can't open " + f.Name() + ": " + err.Error())
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		}
		defer rdr.Close()

		log.Println("[NrscStatic] Serving " + file)
		mtype := mime.TypeByExtension(filepath.Ext(req.URL.Path))
		if len(mtype) != 0 {
			res.Header().Set("Content-Type", mtype)
		}
		res.Header().Set("Content-Size", fmt.Sprintf("%d", f.Size()))
		res.Header().Set("Last-Modified", f.ModTime().UTC().Format(http.TimeFormat))

		io.Copy(res, rdr)
	}
}
