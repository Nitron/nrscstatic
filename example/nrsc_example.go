package main

import (
    "github.com/Nitron/nrscstatic"
    "github.com/codegangsta/martini"
    "net/http"
)

func main() {
    m := martini.New()
    m.Use(martini.Recovery())
    m.Use(martini.Logger())
    m.Use(nrscstatic.NrscStatic("public", true))
    http.ListenAndServe(":8080", m)
}