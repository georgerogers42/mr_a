package app

import (
	"github.com/georgerogers42/mr_a/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var App = mux.NewRouter()

func init() {
	App.HandleFunc("/", Index).Methods("GET")
}

var baseTpl = template.Must(template.ParseFiles("tpl/base.tpl"))

var indexTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("tpl/index.tpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	indexTpl.Execute(w, models.ArticleList)
}
