package blog_service

import (
	"fmt"
)

// [url]/api/v1/posts
// GET posts/id
// { id, html, }
//
// POST posts/id
// PUT posts/id
// DELETE posts/id

// comments
// GET for post, POST add to post

// categories

////
// html
// [url]/ ?page=
// [url]/about
// [url]/posts/[post_id]
// [url]/categories/[category_id]
// [url]/admin
// [url]/admin/posts
// [url]/admin/logs

type AdminConfig struct {
	Password string    `json:"password"`
	Jwt      jwtConfig `json:"jwt"`
}

type jwtConfig struct {
	Secret    string `json:"secret"`
	ExpiresIn int    `json:"expires_in"`
}

func Do_work() {
	fmt.Println("hello blog")
}
