package main

import (
	"log"
	"net/http"

	"github.com/pnlinh/goreddit/stores"
	"github.com/pnlinh/goreddit/web"
)

func main() {
	s, err := stores.NewStore("postgres://postgres:secret@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(*s)
	http.ListenAndServe(":8081", h)
}
