package models

import (
	"bufio"
	"encoding/json"
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
	json.Unmarshal([]byte(mdata), &article.Metadata)
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
