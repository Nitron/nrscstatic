nrscstatic
==========

[Martini](https://github.com/codegangsta/martini) middleware for serving static assets (Javascript, CSS, images, etc.) via [nrsc](https://bitbucket.org/tebeka/nrsc). Should be interchangeable with Martini's built-in Static middleware.

Usage
-----

```go
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
```

```bash
$ mkdir public
$ echo "test" > public/test.txt
$ go get
$ go install bitbucket.org/tebeka/nrsc
$ go build nrsc_example.go
$ $GOPATH/bin/nrsc nrsc_example public
$ ./nrsc_example
```

Then visit [http://localhost:8080/test.txt](http://localhost:8080/test.txt).