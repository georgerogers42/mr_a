package main

import (
	"flag"
	"fmt"
	"github.com/georgerogers42/mr_a/app"
	"net/http"
)

func main() {
	on := flag.String("o", "0.0.0.0:8080", "Address To Listen On")
	flag.Parse()
	fmt.Println("Listening on:", *on)
	http.Handle("/", app.App)
	http.ListenAndServe(*on, nil)
}
