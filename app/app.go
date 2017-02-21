package app

import (
	"fmt"
	"github.com/georgerogers42/mr_a/models"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var App = mux.NewRouter()

func init() {
	App.HandleFunc("/", Index).Methods("GET")
	App.HandleFunc("/article/{i}", Article).Methods("GET")
}

var baseTpl = template.Must(template.ParseFiles("tpl/base.tpl"))

var indexTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("tpl/index.tpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	indexTpl.Execute(w, models.ArticleMap)
}

var articleTpl = template.Must(template.Must(baseTpl.Clone()).ParseFiles("tpl/article.tpl"))

func Article(w http.ResponseWriter, r *http.Request) {
	i := 0
	vars := mux.Vars(r)
	_, err := fmt.Sscan(vars["i"], &i)
	if err != nil || i > len(models.ArticleMap) {
		http.NotFound(w, r)
		return
	}
	articleTpl.Execute(w, models.ArticleMap[i])
}
