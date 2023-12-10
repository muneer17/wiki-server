package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Article is a struct that holds the name and content of an article
type Article struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// Articles is a slice of Article
type Articles []Article

// Global variable to hold all articles
var articles Articles

func GetArticlesHandler(w http.ResponseWriter, r *http.Request) {

	// Check if name is provided in URL
	name := strings.TrimPrefix(r.URL.Path, "/articles/")
	if name != "" {
		GetArticleHandler(w, r)
		return
	}

	// If no name is provided, return all articles
	articlesJson, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(articlesJson)
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get article name from URL
	name := strings.TrimPrefix(r.URL.Path, "/articles/")

	// Find article by name
	for _, article := range articles {
		if article.Name == name {
			// Encode article to JSON
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	// No article found, return 404
	w.WriteHeader(http.StatusNotFound)
}

// PutArticleHandler handles PUT requests to /articles/:name
func PutArticleHandler(w http.ResponseWriter, r *http.Request) {

	// Get article name from URL
	parts := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	name := parts[len(parts)-1]

	// Read and parse request body
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	content := requestData["content"]

	// Log the name and new content
	fmt.Printf("Updating article %s with content: %s\n", name, content)

	// Find article by name
	for i, article := range articles {
		if article.Name == name {
			// Update existing article
			articles[i].Content = content

			// Log the updated article
			fmt.Printf("Updated article: %+v\n", articles[i])

			// Return 200
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	// Article not found, create new article
	newArticle := Article{Name: name, Content: content}
	articles = append(articles, newArticle)

	// Log the new article
	fmt.Printf("Created new article: %+v\n", newArticle)

	// Return 201
	w.WriteHeader(http.StatusCreated)
}

func main() {

	// Add some articles
	articles = append(articles, Article{Name: "wiki", Content: "A wiki is a knowledge base website"})
	articles = append(articles, Article{Name: "rest api", Content: "REST API is a way to communicate between client and server"})

	// Register handlers
	http.HandleFunc("/articles/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if strings.TrimPrefix(r.URL.Path, "/") == "articles" || strings.TrimPrefix(r.URL.Path, "/") == "articles/" {
				GetArticlesHandler(w, r)
			} else {
				GetArticleHandler(w, r)
			}
		case http.MethodPut:
			PutArticleHandler(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	fmt.Println("Server listening on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
