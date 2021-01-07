// github패키지는 Github 이슈 트래커에 대한 Go API를 제공한다.
package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
