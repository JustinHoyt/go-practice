package main

import (
	"context"
	"log"
	"net/http"
)

var MockArticlesData = []*Article{
	{Title: "Hello", Desc: "Article Description1", Content: "Article Content1"},
	{Title: "Hello 2", Desc: "Article Description2", Content: "Article Content2"},
}

type ArticleServiceRPC struct{}

func (s *ArticleServiceRPC) Ping(ctx context.Context) (bool, error) {
	return true, nil
}

func (s *ArticleServiceRPC) GetAllArticles(ctx context.Context) ([]*Article, error) {
	return MockArticlesData, nil
}

func main() {
	webrpcHandler := NewArticleServiceServer(&ArticleServiceRPC{})
	log.Fatal(http.ListenAndServe(":3000", webrpcHandler))
}
