package models

import (
	"bufio"
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Metadata struct {
	Serial              uint64
	Author, Title, Slug string
	Posted              time.Time
}

type Article struct {
	Metadata
	Content string
}

func (a *Article) HtmlContent() template.HTML {
	return template.HTML(a.Content)
}

type Articles []*Article

func (a Articles) Len() int {
	return len(a)
}
func (a Articles) Less(i, j int) bool {
	return a[i].Serial < a[j].Serial
}
func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Articles) ArticleMap() ArticlesMap {
	articlesMap := ArticlesMap{}
	for _, article := range a {
		articlesMap[article.Slug] = article
	}
	return articlesMap
}

func LoadArticles(pat string) (Articles, error) {
	fnames, err := filepath.Glob(pat)
	if err != nil {
		return nil, err
	}
	articles := make(Articles, 0, len(fnames))
	for _, fname := range fnames {
		article, err := LoadArticle(fname)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	sort.Sort(articles)
	return articles, nil
}

func LoadArticle(fname string) (*Article, error) {
	fh, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	s := bufio.NewScanner(fh)
	mdata := ""
	for s.Scan() {
		t := s.Text()
		if t == "" {
			break
		}
		mdata = mdata + t + "\n"
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	article := &Article{}
	err = json.Unmarshal([]byte(mdata), &article.Metadata)
	if err != nil {
		return nil, err
	}
	contents := ""
	for s.Scan() {
		t := s.Text()
		contents = contents + t + "\n"
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	article.Content = contents
	return article, nil
}

type ArticlesMap map[string]*Article

var ArticleList Articles

var ArticleMap = ArticleList.ArticleMap()

func init() {
	articleList, err := LoadArticles("articles/*.html")
	if err != nil {
		panic(err)
	}
	ArticleList = articleList
}
