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
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/misc/", http.FileServer(http.Dir("public")))
	http.Handle("/", app.App)
	http.ListenAndServe(*on, nil)
}
