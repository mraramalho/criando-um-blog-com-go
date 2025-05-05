package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
)

var posts map[string]*Post

// Post represents a blog post loaded from a YAML file.
//
// Fields:
//   - Title:       The title of the blog post.
//   - Excerpt:     A short summary or preview of the post.
//   - Date:        The creation date of the post (from the YAML field `created`).
//   - MDContent:   The original post content written in Markdown (from the YAML field `content`).
//   - HTMLContent: The HTML-rendered version of the Markdown content. Filled at runtime.
//   - Slug:        The URL-friendly identifier derived from the filename (without extension).
//
// This struct is populated by parsing `.yaml` files in the `posts/` directory
// and is used to render blog posts dynamically in templates.
type Post struct {
	Title       string `yaml:"title"`
	Excerpt     string `yaml:"excerpt"`
	Date        string `yaml:"created"`
	MDContent   string `yaml:"content"`
	HTMLContent template.HTML
	Slug        string
}

// markdownToHTML converts a Markdown string into HTML.
//
// It uses the Goldmark library (a CommonMark-compliant Markdown parser) to parse and convert
// the Markdown content into HTML. The result is returned as a string.
//
// Parameters:
//   - markdown: A string containing Markdown-formatted content.
//
// Returns:
//   - The HTML representation of the Markdown content.
//   - An error, if the conversion fails.
//
// This function is used to transform post content (written in Markdown) into HTML
// before rendering it in templates.
func markdownToHTML(markdown string) (string, error) {
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// loadPosts loads all blog posts from the "posts/" directory.
//
// It searches for all `.yaml` files and parses each one into a `Post` struct.
// For every file, it performs the following steps:
//   1. Reads the YAML content.
//   2. Unmarshals the content into a Post struct.
//   3. Extracts the slug from the filename (removing the ".yaml" extension).
//   4. Converts the Markdown content (`MDContent`) to HTML and stores it in `HTMLContent`.
//   5. Stores the post in the global `posts` map, using the slug as the key.
//
// If any step fails (e.g. reading a file, parsing YAML, converting markdown),
// the function returns an error, halting the loading process.
//
// Note: The global `posts` map is cleared and rebuilt on every call.
func loadPosts() error {
	posts = make(map[string]*Post)
	files, err := filepath.Glob("posts/*.yaml")
	if err != nil {
		return fmt.Errorf("Error reading files: %w", err)
	}

	for _, file := range files {
		yamlFile, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("Error reading file: %w", err)
		}

		postData := &Post{}
		err = yaml.Unmarshal(yamlFile, postData)
		if err != nil {
			return fmt.Errorf("Error unmarshalling file: %w", err)
		}

		slug := strings.TrimSuffix(filepath.Base(file), ".yaml")
		postData.Slug = slug

		htmlContent, err := markdownToHTML(postData.MDContent)
		if err != nil {
			return fmt.Errorf("Error converting markdown to HTML: %w", err)
		}
		postData.HTMLContent = template.HTML(htmlContent)
		posts[slug] = postData
	}

	return nil
}
