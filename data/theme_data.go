package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MyData struct {
	Title           string     `json:"Title"`
	SiteName        string     `json:"SiteName"`
	Query           string     `json:"Query"`
	MainPages       []PageItem `json:"MainPages"`
	CurrentMainPage string     `json:"CurrentMainPage"`
	Pages           []PageItem `json:"Pages"`
	Pagination      struct {
		Total   int   `json:"Total"`
		Current int   `json:"Current"`
		Pages   []int `json:"Pages"`
	} `json:"Pagination"`
	Posts    []PostItem    `json:"Posts"`
	Comments []CommentItem `json:"Comments"`
	Page     PageItem      `json:"Page,omitempty"`
	Post     PostItem      `json:"Post,omitempty"`
}

type PageItem struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Slug    string `json:"Slug"`
	Content string `json:"Content"`
	Updated int64    `json:"Updated"`
}

type PostItem struct {
	ID           string `json:"ID"`
	Title        string `json:"Title"`
	Slug         string `json:"Slug"`
	LastIndex    int    `json:"LastIndex"`
	CurrentIndex int    `json:"CurrentIndex"`
	Score        int    `json:"Score"`
	URL          string `json:"Url"`
	URI          string `json:"Uri"`
	Author       struct {
		Username string `json:"Username"`
		URL      string `json:"Url"`
	} `json:"Author"`
	Updated int64 `json:"Updated"`
}

type CommentItem struct {
	ID           string `json:"ID"`
	Content      string `json:"Content"`
	LastIndex    int    `json:"LastIndex"`
	CurrentIndex int    `json:"CurrentIndex"`
	User         struct {
		Username string `json:"Username"`
		URL      string `json:"Url"`
	} `json:"User"`
	Updated int64 `json:"Updated"`
}

func (md *MyData) GetDataFromFile(path string) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	// Sorry about that
	err = json.Unmarshal(body, &md)
	if err != nil {
		log.Fatalf("not valid json file: %v", err)
	}

	// Add PageItem
	md.Page = md.Pages[0]

	// Add PostItem
	md.Post = md.Posts[0]
}
