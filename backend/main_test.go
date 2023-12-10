package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setup() {
	articles = append(articles, Article{Name: "wiki", Content: "A wiki is a knowledge base website"})
	articles = append(articles, Article{Name: "rest api", Content: "REST API is a way to communicate between client and server"})
}

func TestGetArticlesHandler(t *testing.T) {
	setup()
	req, err := http.NewRequest("GET", "/articles/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetArticlesHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetArticleHandler(t *testing.T) {
	setup()
	req, err := http.NewRequest("GET", "/articles/wiki", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetArticleHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPutArticleHandler(t *testing.T) {
	setup()
	var jsonStr = []byte(`{"content":"updated content"}`)
	req, err := http.NewRequest("PUT", "/articles/wiki", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PutArticleHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
