package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// blogHandler handles GET requests to the root ("/") route.
// If a different method is used, it responds with a 405 status code.
// If any path other than "/" is requested, it responds with 404.
//
// It dynamically loads all posts from the "posts" directory
// on every GET request.
//
// Note: This handler does not implement caching. Posts are reloaded
// from disk on every request. While this is inefficient for high-traffic
// sites, the choice was intentional — for a personal blog, the simplicity
// and real-time updates outweigh the performance cost.
func blogHandler(w http.ResponseWriter, r *http.Request) {
	// Checks if path isn't / and return a 404 status code
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Checks if method isn't GET and return a 405 status code
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	if err := loadPosts(); err != nil {
		log.Println("Error loading posts:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error loading posts: %v", err)))
		return
	}

	renderTemplate(w, "blog", posts)
}

// postHandler handles GET requests for individual blog posts using a URL path in the format /post/{slug}.
//
// It first validates the HTTP method and returns a 405 Method Not Allowed
// if it's anything other than GET. Then, it loads the posts from disk.
//
// If the slug is missing (i.e., the path is just "/post/"), it redirects
// the user back to the home page with a 303 See Other status.
//
// If no post is found with the given slug, it returns a 404 Not Found.
//
// The post content is rendered using the "posts" template.
//
// Note: As with blogHandler, posts are reloaded on every request and not cached.
// This ensures that any changes in the post files are reflected immediately,
// but might impact performance on high-traffic blogs.
func postHandler(w http.ResponseWriter, r *http.Request) {
	// Checks if method isn't GET and return a 405 status code
	if r.Method != http.MethodGet {
		log.Println("Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	if err := loadPosts(); err != nil {
		log.Println("Error loading posts:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error loading posts: %v", err)))
		return
	}

	slug := strings.TrimPrefix(r.URL.Path, "/post/")
	// Se o slug estiver vazio, redireciona para a página principal
	if slug == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	post, ok := posts[slug]
	if !ok {
		log.Println("Post not found")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Post not found"))
		return
	}
	renderTemplate(w, "posts", post)

}
